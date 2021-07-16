package mongodb

import (
	"context"
	"fmt"

	"github.com/rekhin/repository/example/object"
)

type ObjectRepository struct{}

func NewObjectRepository() *ObjectRepository {
	return &ObjectRepository{}
}

func (r *ObjectRepository) ReadObject(context.Context, *[]object.Object) error {
	return fmt.Errorf("not implemented")
}
func (r *ObjectRepository) CreateObject(context.Context, ...object.Object) error {
	return fmt.Errorf("not implemented")
}
func (r *ObjectRepository) UpdateObject(context.Context, ...object.Object) error {
	return fmt.Errorf("not implemented")
}
func (r *ObjectRepository) DeleteObject(context.Context, ...object.ObjectID) error {
	return fmt.Errorf("not implemented")
}

type Object struct {
	ID       ObjectID
	ParentID ObjectID
	Name     string
	Sort     int
}

func (o Object) GetNodeID() object.NodeID           { return o.ID }
func (o Object) GetParentID() object.NodeID         { return o.ParentID }
func (o Object) GetName() string                    { return o.Name }
func (o Object) GetSort() int                       { return o.Sort }
func (o Object) GetObjectID() object.ObjectID       { return o.ID }
func (o Object) GetParentObjectID() object.ObjectID { return o.ParentID }

var _ object.Object = Object{}

type ObjectID int

const RootObjectID ObjectID = 0
