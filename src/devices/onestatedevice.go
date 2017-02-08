package devices

import (
	"logger"
	"strconv"
	"rpio"
	"structs"
)

type OneStateDevice struct {
	name string
	pinNum int
	pin rpio.Pin
	value int
}


func (td *OneStateDevice) GetName() string {
	return td.name;
}

func (td *OneStateDevice) SetProps(name string, pin int) {
	td.name = name;
	td.pinNum = pin;
}

func (td *OneStateDevice) Init() {

	td.pin = rpio.Pin(td.pinNum)
	td.pin.Input()

	td.value = 0;

	logger.Info(td.name + " Sensor: using pin " +  strconv.Itoa(td.pinNum));
}

func (td *OneStateDevice) GetStatus() structs.DeviceStatus {

	td.value = int(td.pin.Read());

	var ds structs.DeviceStatus
	ds.DeviceType = "OneStateSensor"
	ds.Status = "value " + strconv.Itoa(int(td.value));
	ds.Value = td.value;

	return ds;
}




