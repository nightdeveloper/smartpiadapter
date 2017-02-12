package interfaces;

import "github.com/nightdeveloper/smartpiadapter/structs"

type ResultConverter func(string) (int, string)

type ISystemPropertyDevice interface {
	GetName() string
	GetStatus() structs.DeviceStatus
	SetProps(name string, command string, isCommand bool, rc ResultConverter)
}
