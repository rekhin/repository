package object

import (
	"context"
)

type DeviceRepository interface {
	ReadDevice(context.Context, *[]Device) error
	CreateDevice(context.Context, ...Device) error
	UpdateDevice(context.Context, ...Device) error
	DeleteDevice(context.Context, ...DeviceID) error
}

type Device interface {
	Node
	GetDeviceID() DeviceID
	GetObjectID() ObjectID
	GetSourceIDs() []SourceID
	GetState() State
}

type DeviceID interface{}

type State int

const (
	StateInactive State = 0
	StateActive   State = 1
)
