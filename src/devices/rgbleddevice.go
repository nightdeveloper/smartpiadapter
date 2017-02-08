package devices

import (
	"logger"
	"strconv"
	"rpio"
	"structs"
)

type RgbLedDevice struct {
	name string
	r, g, b int
	rPinNum int
	gPinNum int
	bPinNum int
	rPin rpio.Pin
	gPin rpio.Pin
	bOin rpio.Pin
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

	logger.Info("RgbLedDevice: using " +
		"red pin "   + strconv.Itoa(ld.rPinNum) + " " +
		"green pin " + strconv.Itoa(ld.gPinNum) + " " +
		"blue pin "  + strconv.Itoa(ld.bPinNum)	)

	ld.rPin = rpio.Pin(ld.rPinNum)
	ld.gPin = rpio.Pin(ld.gPinNum)
	ld.bOin = rpio.Pin(ld.bPinNum)
	ld.rPin.Output()
	ld.gPin.Output()
	ld.bOin.Output()

	ld.SetLight(0, 0, 0);
}

func (ld *RgbLedDevice) SetLight(r, g, b int) {

	ld.r = r;
	ld.g = g;
	ld.b = b;

	logger.Info("rgb " + ld.GetStatus().Status);

	if r < 127 { ld.rPin.Low() } else { ld.rPin.High() }
	if g < 127 { ld.gPin.Low() } else { ld.gPin.High() }
	if b < 127 { ld.bOin.Low() } else { ld.bOin.High() }
}

func (ld *RgbLedDevice) GetStatus() structs.DeviceStatus {

	var ds structs.DeviceStatus

	ds.DeviceType = "RgbLedDevice"
	ds.Status = "color (" +
		strconv.Itoa(ld.r) + ", " +
		strconv.Itoa(ld.g) + ", " +
		strconv.Itoa(ld.b) + ")"

	ds.R = ld.r
	ds.G = ld.g
	ds.B = ld.b

	return ds
}




