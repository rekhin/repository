package object

import (
	"context"
)

type SourceRepository interface {
	ReadSource(context.Context, *[]Source) error
	CreateSource(context.Context, ...Source) error
	UpdateSource(context.Context, ...Source) error
	DeleteSource(context.Context, ...SourceID) error
}

type Source interface {
	GetSourceID() SourceID
	GetProtocolID() ProtocolID
	GetDeviceIDs() []DeviceID
	GetState() State
}

type SourceID interface{}
