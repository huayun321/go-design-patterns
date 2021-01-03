package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType = iota
	MotorbikeFactoryType
)

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}
