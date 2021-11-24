package geo

import (
	"errors"
)

type Landmark struct {
	name string
	//can call struct in same directory/package
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("Invalid name")
	}
	l.name = name
	return nil
}
