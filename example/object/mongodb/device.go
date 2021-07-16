package mongodb

import (
	"context"
	"fmt"

	"github.com/rekhin/repository/example/object"
)

type DeviceRepository struct{}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{}
}

func (r *DeviceRepository) ReadDevice(context.Context, *[]object.Device) error {
	return fmt.Errorf("not implemented")
}
func (r *DeviceRepository) CreateDevice(context.Context, ...object.Device) error {
	return fmt.Errorf("not implemented")
}
func (r *DeviceRepository) UpdateDevice(context.Context, ...object.Device) error {
	return fmt.Errorf("not implemented")
}
func (r *DeviceRepository) DeleteDevice(context.Context, ...object.DeviceID) error {
	return fmt.Errorf("not implemented")
}

type Device struct {
	ID        DeviceID
	ObjectID  object.ObjectID
	Name      string
	SourceIDs []SourceID
	State     object.State
	Sort      int
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
