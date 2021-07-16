package object

import (
	"context"
)

type NodeRepository interface {
	ReadNode(context.Context, *[]Node) error
	CreateNode(context.Context, ...Node) error
	UpdateNode(context.Context, ...Node) error
	DeleteNode(context.Context, ...NodeID) error
}

type Node interface {
	GetNodeID() NodeID
	GetParentID() NodeID
	GetName() string
	GetSort() int
}

type NodeID interface {
	// TODO need methods
}
