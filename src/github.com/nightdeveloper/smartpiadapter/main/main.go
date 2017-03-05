package main

import ( 
	"github.com/nightdeveloper/smartpiadapter/logger"
	"github.com/nightdeveloper/smartpiadapter/devices"
	"github.com/nightdeveloper/smartpiadapter/settings"
	"github.com/nightdeveloper/smartpiadapter/web"
	"github.com/nightdeveloper/smartpiadapter/interfaces"
	"strconv"
	"strings"
	"github.com/nightdeveloper/smartpiadapter/chats"
)

func addTemperatureHamidityDevice(dm *devices.DeviceManager, name string, pin int) {

	var td devices.Dht22Device = devices.Dht22Device{}
	td.SetProps(name, pin);

	var tdi interfaces.IStatusDevice = &td;
	dm.AddStatusDevice(&tdi);
}

func addRgbLedDevice(dm *devices.DeviceManager, c *settings.Config) {

	var ld devices.RgbLedDevice = devices.RgbLedDevice{}
	ld.SetProps(c.LedRPin, c.LedGPin, c.LedBPin);

	var ldi interfaces.IRgbLedDevice = &ld;
	dm.AddRgbLedDevice(&ldi);
}

func addSystemPropertyDevice(dm *devices.DeviceManager, name string, command string, isCommand bool,
	rc interfaces.ResultConverter) {

	var spd devices.SystemPropertyDevice = devices.SystemPropertyDevice{}
	spd.SetProps(name, command, isCommand, rc);

	var spdi interfaces.ISystemPropertyDevice = &spd;
	dm.AddSystemProperyDevice(&spdi);
}

func addIpInfoDevice(dm *devices.DeviceManager, name string) {

	var spd devices.IpInfoDevice = devices.IpInfoDevice{}
	spd.SetProps(name, "", false, func(r string) (int, string) { return 0, ""});

	var spdi interfaces.ISystemPropertyDevice = &spd;
	dm.AddSystemProperyDevice(&spdi);
}

func main() {
	logger.Info("application started");

	c := settings.Config{};
	c.Load();

	dm := devices.DeviceManager{};
	dm.Init(&c);

	addTemperatureHamidityDevice(&dm, "Temperature", c.TemperatureSensorPin);

	addSystemPropertyDevice(&dm, "CpuTemperature", "/sys/class/thermal/thermal_zone0/temp", false,
		func(r string) (int, string) {
			r = strings.Replace(r, "\n","", -1);

			i, err := strconv.ParseFloat(r, 64);
			if err != nil {
				return 0, "error parsing " + r + " " + err.Error();
			}

			var res = i / 1000;

			return int(res), strconv.FormatFloat(res, 'f', 2, 64);
		});

	addIpInfoDevice(&dm, "IpInfo")

	dm.Start();

	cm := chats.ChatManager{}
	cm.Init(&c, &dm);
	cm.Start();

	web := web.Server{}
	web.Start(&dm, &c);
}