package object

import (
	"context"
)

type TagRepository interface {
	ReadTag(context.Context, *[]Tag) error
	CreateTag(context.Context, ...Tag) error
	UpdateTag(context.Context, ...Tag) error
	DeleteTag(context.Context, ...TagID) error
}

type Tag interface {
	Node
	GetTagID() TagID
	GetState() State
}

type TagID interface {
	// TODO need methods
}
