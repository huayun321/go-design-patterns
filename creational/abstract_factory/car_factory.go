package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = iota
	FamiliarCarType
)

type CarFactory struct{}

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamiliarCarType:
		return new(FamiliarCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recongnized\n", v))
	}
}
