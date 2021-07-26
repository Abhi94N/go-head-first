package magazine

//When exporting struct the fields must be capitalized
//just type for addresss not name
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

// type Employee struct {
// 	Name        string
// 	Salary      float64
// 	HomeAddress Address
// }
type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
