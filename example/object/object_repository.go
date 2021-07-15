package object

import (
	"context"
)

type ObjectRepository interface {
	ReadObject(context.Context, *[]Object) error
	CreateObject(context.Context, ...Object) error
	UpdateObject(context.Context, ...Object) error
	DeleteObject(context.Context, ...ObjectID) error
}

type Object interface {
	Node
	ObjectID() ObjectID
}

type ObjectID int
