package object

import (
	"context"
)

type ProtocolRepository interface {
	ReadProtocol(context.Context, *[]Protocol) error
	CreateProtocol(context.Context, ...Protocol) error
	UpdateProtocol(context.Context, ...Protocol) error
	DeleteProtocol(context.Context, ...ProtocolID) error
}

type Protocol interface {
	GetProtocolID() ProtocolID
	GetName() string
}

type ProtocolID interface{}
