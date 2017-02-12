package rpio

type Direction uint8
type Pin uint8
type State uint8
type Pull uint8

func Open() {
}

func (pin Pin) Input() {
}

func (pin Pin) Output() {
}

func (pin Pin) High() {
}

func (pin Pin) Low() {
}

func (pin Pin) Toggle() {
}

func (pin Pin) Mode(dir Direction) {
}

func (pin Pin) Write(state State) {
}

func (pin Pin) Read() State {
	return 25;
}

func (pin Pin) Pull(pull Pull) {
}

func (pin Pin) PullUp() {
}

func (pin Pin) PullDown() {
}

func (pin Pin) PullOff() {
}
