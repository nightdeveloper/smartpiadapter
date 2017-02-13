package main

import ( 
	"github.com/nightdeveloper/smartpiadapter/logger"
	"github.com/nightdeveloper/smartpiadapter/devices"
	"github.com/nightdeveloper/smartpiadapter/settings"
	"github.com/nightdeveloper/smartpiadapter/web"
	"github.com/nightdeveloper/smartpiadapter/interfaces"
	"github.com/stianeikeland/go-rpio"
	"strconv"
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

	addTemperatureHamidityDevice(&dm, "Temperature", c.TemperatureSensorPin);

	addSystemPropertyDevice(&dm, "CpuTemperature", "/sys/class/thermal/thermal_zone0/temp", false,
		func(r string) (int, string) {

			i, err := strconv.Atoi(r);
			if err != nil {
				return 0, r
			}

			return i / 1000, r;
		});

	dm.Start();

	web := web.Server{}
	web.Start(&dm, &c);
}