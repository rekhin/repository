package mongodb

import "github.com/rekhin/repository/example/object"

type Protocol struct {
	ID       object.ProtocolID
	DeviceID object.DeviceID
	Name     string
	Sort     int
	State    object.State
}

func (o Protocol) GetNodeID() object.NodeID         { return o.ID }
func (o Protocol) GetParentID() object.NodeID       { return o.DeviceID }
func (o Protocol) GetName() string                  { return o.Name }
func (o Protocol) GetSort() int                     { return o.Sort }
func (o Protocol) GetProtocolID() object.ProtocolID { return o.ID }
func (o Protocol) GetDeviceID() object.DeviceID     { return o.DeviceID }
func (o Protocol) GetState() object.State           { return o.State }

var _ object.Protocol = Protocol{}
