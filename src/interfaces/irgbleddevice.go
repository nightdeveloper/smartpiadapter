package interfaces

import "structs"

type IRgbLedDevice interface {
	GetName() string
	Init()
	GetStatus() structs.DeviceStatus
	SetProps(rPin, gPin, bPin int)
	SetLight(r, g, b int)
}
