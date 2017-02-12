package interfaces

import "github.com/nightdeveloper/smartpiadapter/structs"

type IStatusDevice interface {
	GetName() string
	Init()
	GetStatus() structs.DeviceStatus
	SetProps(name string, pin int)
}
