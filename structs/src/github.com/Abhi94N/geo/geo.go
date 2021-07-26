package geo

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

type Landmark struct {
	Name string
	Coordinates
}
