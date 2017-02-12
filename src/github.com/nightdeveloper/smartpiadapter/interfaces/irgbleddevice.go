package interfaces

import "github.com/nightdeveloper/smartpiadapter/structs"

type IRgbLedDevice interface {
	GetName() string
	Init()
	GetStatus() structs.DeviceStatus
	SetProps(rPin, gPin, bPin int)
	SetLight(r, g, b int)
}
