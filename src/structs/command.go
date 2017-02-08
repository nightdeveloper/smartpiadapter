package structs

type Command struct {
	Name	string
	Action	string

	Enable	bool
	R	int
	G	int
	B	int
	Data	[]byte
	Result	string
}