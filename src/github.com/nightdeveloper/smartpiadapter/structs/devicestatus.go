package structs

type DeviceStatus struct {
	DeviceType string
	Status string
	Value int
	R, G, B int
	Data []byte
}