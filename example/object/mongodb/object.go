package mongodb

import "github.com/rekhin/repository/example/object"

type Object struct {
	ID       object.ObjectID
	ParentID int
	Name     string
	Sort     int
}

func (o Object) GetNodeID() object.NodeID   { return o.ID }
func (o Object) GetParentID() object.NodeID { return o.ParentID }
func (o Object) GetName() string            { return o.Name }
func (o Object) GetSort() int               { return o.Sort }
func (o Object) ObjectID() object.ObjectID  { return o.ID }
