package models

type Port struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float32
	Province    string
	Timezone    string
	Code        string
	Unlocs      []string
}
