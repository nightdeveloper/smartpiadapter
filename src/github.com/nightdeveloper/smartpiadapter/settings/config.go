package settings

import (
	"encoding/json"
	"github.com/nightdeveloper/smartpiadapter/logger"
	"path/filepath"
	"io/ioutil"
)

type Config struct {
	PoolTimeoutSecs int
	HttpPort int
	TelegramKey string
	TemperatureSensorPin int
	HumiditySensorPin int
	AirSensorPin int
	LedRPin int
	LedGPin int
	LedBPin int
}

func (c *Config) Load() {

	absPath, _ := filepath.Abs("../");

	filename := absPath + "/config.json";

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		logger.Error("Config reading error ", err);
		panic("config reading error");
	}

	err = json.Unmarshal(file, c);

	if err != nil || c == nil {
		logger.Error("Config decoding error ", err);
		panic("config decoding error");
	}

	out, _ := json.Marshal(c);
	logger.Info("config read " + filename+ ": " + string(out));
}