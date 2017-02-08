package main

import ( 
	"logger"
	"devices"
	"settings"
	"web"
	"rpio"
	"interfaces"
	"strconv"
)

func addOneStateDevice(dm *devices.DeviceManager, name string, pin int) {

	var td devices.OneStateDevice = devices.OneStateDevice{}
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

func addSystemPropertyDevice(dm *devices.DeviceManager, name string, command string, isCommand bool, rc interfaces.ResultConverter) {

	var spd devices.SystemPropertyDevice = devices.SystemPropertyDevice{}
	spd.SetProps(name, command, isCommand, rc);

	var spdi interfaces.ISystemPropertyDevice = &spd;
	dm.AddSystemProperyDevice(&spdi);
}

func main() {
	logger.Info("application started");

	c := settings.Config{};
	c.Load();

	rpio.Open();

	dm := devices.DeviceManager{};
	dm.Init(&c);

	addOneStateDevice(&dm, "Temperature", c.TemperatureSensorPin);
	addOneStateDevice(&dm, "Humidity", c.HumiditySensorPin);
	addOneStateDevice(&dm, "AirSensorPin", c.AirSensorPin);

	addSystemPropertyDevice(&dm, "CpuTemperature", "/sys/class/thermal/thermal_zone0/temp", false,
		func(r string) (int, string) {

			i, err := strconv.Atoi(r);
			if err != nil {
				return 0, r
			}

			return i / 1000, r;
		});

	addRgbLedDevice(&dm, &c);

	dm.Start();

	web := web.Server{}
	web.Start(&dm, &c);
}