package devices

import (
	"github.com/nightdeveloper/smartpiadapter/logger"
	"github.com/nightdeveloper/smartpiadapter/settings"
	"time"
	"github.com/nightdeveloper/smartpiadapter/interfaces"
	"github.com/nightdeveloper/smartpiadapter/structs"
)

type DeviceManager struct {
	c *settings.Config

	statusDevices []*interfaces.IStatusDevice
	rgbLedDevices map[string]*interfaces.IRgbLedDevice
	systemPropertyDevices []*interfaces.ISystemPropertyDevice;
}

func (d DeviceManager) Test() {
	logger.Info("devicemanager test")
}

func (d *DeviceManager) Init(c *settings.Config) {
	d.c = c;

	d.rgbLedDevices = make(map[string]*interfaces.IRgbLedDevice)
}

func (d *DeviceManager) AddStatusDevice(device *interfaces.IStatusDevice) {
	(*device).Init();
	d.statusDevices = append(d.statusDevices, device);
}

func (d *DeviceManager) AddRgbLedDevice(device *interfaces.IRgbLedDevice) {
	(*device).Init();

	d.rgbLedDevices[(*device).GetName()] = device;
}

func (d *DeviceManager) SetRgbLedDeviceValue(name string, r, g, b int) {

	if d.rgbLedDevices[name] == nil {
		logger.Error("Led with name [" + name + "] not found", nil);
		return;
	}

	(*d.rgbLedDevices[name]).SetLight(r, g, b);
}

func (d *DeviceManager) AddSystemProperyDevice(device *interfaces.ISystemPropertyDevice) {

	d.systemPropertyDevices = append(d.systemPropertyDevices, device);
}

func (d *DeviceManager) Start() {

	go func() {
		logger.Info("device manager loop start")
		for {
			time.Sleep(time.Duration(d.c.PoolTimeoutSecs) * time.Second);
			for k, v := range (*d).GetStatus() {
				logger.Info(k + " " + v.Status)
			}
		}
	}()

}

func (d *DeviceManager) GetStatus() structs.StatusMap {

	results := make(map[string]structs.DeviceStatus)

	var ds structs.DeviceStatus
	ds.DeviceType = "DeviceManager"
	ds.Status = "alive at " + logger.CurrentTime()

	results["Device Manager"] = ds;

	for _, v  := range d.statusDevices {
		results[ (*v).GetName() ] = (*v).GetStatus();
	}

	for k, v  := range d.rgbLedDevices {
		results[ k ] = (*v).GetStatus();
	}

	for _, v  := range d.systemPropertyDevices {
		results[ (*v).GetName() ] = (*v).GetStatus();
	}

	return results;
}

func (d *DeviceManager) SetRgbLedState(r, g, b int) {

	for k := range d.rgbLedDevices {

		var dev *interfaces.IRgbLedDevice = d.rgbLedDevices[k];
		(*dev).SetLight(r, g, b);
		return;
	}
}
















