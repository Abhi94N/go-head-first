package gadget

type Switch string

func (s *Switch) Toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
}

type Toggleable interface {
	Toggle()
}
