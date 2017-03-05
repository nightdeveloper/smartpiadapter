package devices

import (
	"github.com/nightdeveloper/smartpiadapter/structs"
)

type Dht22Device struct {
	name string
	pinNum int
	temperature float32
	humidity float32
}


func (td *Dht22Device) GetName() string {
	return td.name;
}

func (td *Dht22Device) SetProps(name string, pin int) {
}

func (td *Dht22Device) Init() {
}

func (td *Dht22Device) GetStatus() structs.DeviceStatus {

	var a structs.DeviceStatus;
	a.Status = "windows is not supported";
	a.Value = 0;
	return a;
}




