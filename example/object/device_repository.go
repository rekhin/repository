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
	DeviceID() DeviceID
	GetState() State
}

type DeviceID int

type State int

const (
	StateInactive State = 0
	StateActive   State = 1
)
