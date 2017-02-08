package web

import (
	"io/ioutil"
	"net/http"
	"logger"
	"path/filepath"
	"strings"
	"devices"
	"settings"
	"html/template"
	"strconv"
	"encoding/json"
	"structs"
)

type Server struct {
	dm *devices.DeviceManager
	c *settings.Config
}

func (d *Server) response(w http.ResponseWriter, body string, httpCode int, contentType string) {

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(httpCode)
	w.Write([]byte(body));
}

func (d *Server) renderIndexPage(w http.ResponseWriter, body []byte, mainPageMessage string) {

	var funcs = template.FuncMap{"join": strings.Join}
	var statuses = d.dm.GetStatus();

	t, err := template.New("index").Funcs(funcs).Parse(string(body));
	if err != nil {
		logger.Error("can not generate index", err);
		d.response(w, "Can not parse page " + mainPageMessage, 500, "text/html");
	}

	err = t.ExecuteTemplate(w, "T", statuses)
	if err != nil {
		logger.Error("can not generate index", err);
		d.response(w, "Can not render page " + mainPageMessage, 500, "text/html");
	}
}

func (d *Server) responseJson(w http.ResponseWriter, v interface{}) {

	responseCode := 200;

	var responseTemplate = make(map[string]interface{});
	responseTemplate["data"] = v;
	responseTemplate["error"] = "";

	var response []byte
	var err error

	response, err = json.Marshal(responseTemplate);
	if err != nil {
		logger.Error("can not marshal response", err)
		responseTemplate["error"] = "Can not marshal response";
		responseCode = 500

		response, err = json.Marshal(responseTemplate);
	}

	logger.Info("action response: " + string(response));

	d.response(w, string(response), responseCode, "text/json");
}

func (d *Server) doAction(w http.ResponseWriter, body []byte) {

	logger.Info("-------");
	logger.Info("incoming action: " + string(body));

	var c structs.Command

	var err = json.Unmarshal(body, &c)

	if err != nil {
		logger.Error("can not unmarshal", err)
		return;
	}

	if "rgbLedState" == c.Action {
		d.dm.SetRgbLedState(c.R, c.G, c.B)
	}

	d.responseJson(w, d.dm.GetStatus());

	logger.Info("-------");
}

func (d *Server) loadPage(w http.ResponseWriter, path string, requestBody []byte) {

	loweredPath := strings.ToLower(path);

	if loweredPath == "action" {
		d.doAction(w, requestBody)
		return
	}

	if loweredPath == "" {
		loweredPath = "index.html"
	}

	mimeType := "";

	if strings.Contains(loweredPath, ".html") {
		mimeType = "text/html";

	} else if strings.Contains(loweredPath, ".png") {
		mimeType = "image/png";

	} else if strings.Contains(loweredPath, ".js") {
		mimeType = "application/javascript";

	} else if strings.Contains(loweredPath, ".css") {
		mimeType = "text/css";

	} else {
		mimeType = "application/octet-stream"
	}

	secureAbsPath, _ := filepath.Abs("../web")

	absPath, err := filepath.Abs("../web/" + loweredPath)

	mainPageMessage := "<a href='/'>go to main page</a>";

	if err != nil {
		logger.Error("error while looking for file: ", err)
		d.response(w, "Invalid path " + mainPageMessage, 400, "text/html");
		return
	}

	if strings.Index(absPath, secureAbsPath) != 0 {

		logger.Error("url denied: " + absPath + " (out of " + secureAbsPath + ")", nil)
		d.response(w, "Access denied " + mainPageMessage, 403, "text/html");
		return
	}

	body, err := ioutil.ReadFile(absPath)

	if err != nil {
		logger.Error("error while reading page: ", err);
		d.response(w, "File not found " + mainPageMessage, 404, "text/html");
		return
	}

	w.Header().Set("Content-Type", mimeType)

	w.Write(body)

	return
}

func (d *Server) handler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:];

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("error while reading request body", err)
	}

	d.loadPage(w, path, body);
}

func (d *Server) Start(dm *devices.DeviceManager, c *settings.Config) {
	d.dm = dm;
	d.c = c;

	port := strconv.Itoa(c.HttpPort);

	logger.Info("web server start at port " + port);

	http.HandleFunc("/", d.handler)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		logger.Error( "webserver listener error: ", err)
	}
}


