package devices

import (
	"github.com/nightdeveloper/smartpiadapter/structs"
	"github.com/nightdeveloper/smartpiadapter/logger"
)

type RgbLedDevice struct {
	name string
	r, g, b int
	rPinNum int
	gPinNum int
	bPinNum int
}

func (ld *RgbLedDevice) GetName() string {
	return "RGB Led";
}

func (ld *RgbLedDevice) SetProps(rPin, gPin, bPin int) {
	ld.rPinNum = rPin;
	ld.gPinNum = gPin;
	ld.bPinNum = bPin;
}

func (ld *RgbLedDevice) Init() {
}

func (ld *RgbLedDevice) SetLight(r, g, b int) {

	ld.r = r;
	ld.g = g;
	ld.b = b;

	logger.Info("rgb " + ld.GetStatus().Status);
}

func (ld *RgbLedDevice) GetStatus() structs.DeviceStatus {

	var a structs.DeviceStatus;
	a.Status = "windows is not supported";
	a.Value = 0;
	return a;
}




