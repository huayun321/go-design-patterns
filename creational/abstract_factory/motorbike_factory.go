package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType = iota
	CruiseMotorbikeType
)

type MotorbikeFactory struct{}

func (*MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recongnized\n", v))
	}
}
