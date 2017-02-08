package interfaces

import "structs"

type IStatusDevice interface {
	GetName() string
	Init()
	GetStatus() structs.DeviceStatus
	SetProps(name string, pin int)
}
