package mongodb

import "github.com/rekhin/repository/example/object"

type Device struct {
	ID        DeviceID
	ObjectID  object.ObjectID
	Name      string
	SourceIDs []SourceID
	Sort      int
	State     object.State
}

func (o Device) GetNodeID() object.NodeID        { return o.ID }
func (o Device) GetParentID() object.NodeID      { return o.ObjectID }
func (o Device) GetName() string                 { return o.Name }
func (o Device) GetSort() int                    { return o.Sort }
func (o Device) GetDeviceID() object.DeviceID    { return o.ID }
func (o Device) GetObjectID() object.ObjectID    { return o.ObjectID }
func (o Device) GetSourceIDs() []object.SourceID { return convertSourceIDs(o.SourceIDs) }
func (o Device) GetState() object.State          { return o.State }

func convertSourceIDs(from []SourceID) []object.SourceID {
	var to []object.SourceID = make([]object.SourceID, len(from))
	for i := range from {
		to[i] = from[i]
	}
	return to
}

var _ object.Device = Device{}

type DeviceID int

// TODO move to protocol.go
type ProtocolID int

// TODO move to source.go
type SourceID struct {
	DeviceID
	ProtocolID
}
