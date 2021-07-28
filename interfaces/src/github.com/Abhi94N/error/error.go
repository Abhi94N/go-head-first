package error

import (
	"fmt"
)

type Error interface {
	Error() string
}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

//return function satisfies interface
func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}
