package repository

import "context"

type SomethingRepository interface {
	ReadSomething(context.Context, *[]Something, ...Filter) error
	CreateSomething(context.Context, ...Something) error
	UpdateSomething(context.Context, ...Something) error
	DeleteSomething(context.Context, ...ID) error
}

type Something interface {
	GetID() ID
}

type ID interface {
	// TODO need methods
}

type Filter struct{}
