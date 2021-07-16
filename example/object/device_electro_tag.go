package object

import (
	"context"
)

type DeviceElectroTagRepository interface {
	ReadDeviceElectroTag(context.Context, *[]DeviceElectroTag) error
	CreateDeviceElectroTag(context.Context, ...DeviceElectroTag) error
	UpdateDeviceElectroTag(context.Context, ...DeviceElectroTag) error
	DeleteDeviceElectroTag(context.Context, ...DeviceElectroTagID) error
}

type DeviceElectroTag interface {
	ElectroTag
	GetDeviceElectroTagID() DeviceElectroTagID
	GetDeviceID() DeviceID
}

type DeviceElectroTagID struct {
	DeviceID
	ElectroTagType
}

// type DeviceElectroTagID interface {
// 	ElectroTagID
// 	GetDeviceID() DeviceID
// 	GetElectroTagType() ElectroTagType
// }
