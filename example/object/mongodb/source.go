package mongodb

import (
	"context"
	"fmt"

	"github.com/rekhin/repository/example/object"
)

type SourceRepository struct{}

func NewSourceRepository() *SourceRepository {
	return &SourceRepository{}
}

func (r *SourceRepository) ReadSource(context.Context, *[]object.Source) error {
	return fmt.Errorf("not implemented")
}
func (r *SourceRepository) CreateSource(context.Context, ...object.Source) error {
	return fmt.Errorf("not implemented")
}
func (r *SourceRepository) UpdateSource(context.Context, ...object.Source) error {
	return fmt.Errorf("not implemented")
}
func (r *SourceRepository) DeleteSource(context.Context, ...object.SourceID) error {
	return fmt.Errorf("not implemented")
}

type Source struct {
	ID         SourceID
	ProtocolID ProtocolID
	State      object.State
}

func (o Source) GetSourceID() object.SourceID     { return o.ID }
func (o Source) GetProtocolID() object.ProtocolID { return o.ProtocolID }
func (o Source) GetDeviceIDs() []object.DeviceID  { return []object.DeviceID{o.ID.DeviceID} }
func (o Source) GetState() object.State           { return o.State }

// func convertDeviceIDs(from []DeviceID) []object.DeviceID {
// 	var to []object.DeviceID = make([]object.DeviceID, len(from))
// 	for i := range from {
// 		to[i] = from[i]
// 	}
// 	return to
// }

var _ object.Source = Source{}

type SourceID struct {
	DeviceID
	ProtocolID
}
