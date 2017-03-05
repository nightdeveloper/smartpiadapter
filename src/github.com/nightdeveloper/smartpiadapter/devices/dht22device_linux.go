package devices

import (
	"github.com/nightdeveloper/smartpiadapter/logger"
	"github.com/nightdeveloper/smartpiadapter/structs"
	"github.com/d2r2/go-dht"
	"strconv"
	"fmt"
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
	td.name = name;
	td.pinNum = pin;
}

func (td *Dht22Device) Init() {

	rpio.Open();

	td.temperature = 0;
	td.humidity = 0;

	logger.Info(td.name + " Sensor: using pin " +  strconv.Itoa(td.pinNum));
}

func (td *Dht22Device) GetStatus() structs.DeviceStatus {

	var ds structs.DeviceStatus
	ds.DeviceType = "TemperatureSensor"

	sensorType := dht.DHT22
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(sensorType, td.pinNum, false, 2)
	if err != nil {
		ds.Status = "error while reading"
		logger.Error("error while read dht22", err)
		return ds
	}
	td.temperature = temperature;
	td.humidity = humidity;
	ds.Status =
		fmt.Sprintf("temperature = %.2f, humidity = %.2f (retried %d times)",
		temperature, humidity, retried)
	ds.Value = int(td.temperature);
	// todo add humanity

	return ds;
}




