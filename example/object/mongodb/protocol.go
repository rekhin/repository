package mongodb

import "github.com/rekhin/repository/example/object"

type Protocol struct {
	ID   ProtocolID
	Name string
}

func (o Protocol) GetProtocolID() object.ProtocolID { return o.ID }
func (o Protocol) GetName() string                  { return o.Name }

var _ object.Protocol = Protocol{}

type ProtocolID int
