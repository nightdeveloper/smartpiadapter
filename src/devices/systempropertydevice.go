package devices

import (
	"structs"
	"logger"
	"io/ioutil"
	"interfaces"
)

type SystemPropertyDevice struct {
	name string
	command string
	isCommand bool
	value string
	rc interfaces.ResultConverter
}


func (spd *SystemPropertyDevice) GetName() string {
	return spd.name;
}

func (spd *SystemPropertyDevice) SetProps(name string, command string, isCommand bool, rc interfaces.ResultConverter) {
	spd.name = name
	spd.command = command
	spd.isCommand = isCommand
	spd.rc = rc
}

func (spd *SystemPropertyDevice) Init() { }

func (spd *SystemPropertyDevice) GetStatus() structs.DeviceStatus {

	var ds structs.DeviceStatus
	ds.DeviceType = "SystemPropertyDevice"

	if !spd.isCommand {
		bytes, err := ioutil.ReadFile(spd.command);
		if err != nil {
			logger.Error("Error while fetching data from " + spd.command, err);
			ds.Value = 0
			ds.Status = "error"
		} else {
			ds.Value, ds.Status = spd.rc(string(bytes));
		}
	} else {
		logger.Error("command requests not implemented yet", nil)
	}

	return ds;
}




