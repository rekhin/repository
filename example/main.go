package main

import (
	"context"
	"log"

	"github.com/rekhin/repository/example/object"
	"github.com/rekhin/repository/example/object/mongodb"
)

func main() {
	ctx := context.Background()

	objectRepository := mongodb.NewObjectRepository()
	if err := objectRepository.CreateObject(ctx,
		mongodb.Object{
			ID:       1,
			ParentID: mongodb.RootObjectID,
			Name:     "object 1",
		},
	); err != nil {
		log.Fatalf("create object failed: %s", err)
	}

	sourceRepository := mongodb.NewSourceRepository()
	if err := sourceRepository.CreateSource(ctx,
		mongodb.Source{
			ID:         mongodb.SourceID{DeviceID: 1, ProtocolID: 1},
			ProtocolID: 1,
			State:      object.StateActive,
		},
		mongodb.Source{
			ID:         mongodb.SourceID{DeviceID: 1, ProtocolID: 2},
			ProtocolID: 2,
			State:      object.StateActive,
		},
	); err != nil {
		log.Fatalf("create object failed: %s", err)
	}

	var sources []object.Source
	if err := sourceRepository.ReadSource(ctx, &sources); err != nil {
		log.Fatalf("read source failed: %s", err)
	}
	log.Printf("sources: %+v", sources)

	deviceRepository := mongodb.NewDeviceRepository()
	if err := deviceRepository.CreateDevice(ctx,
		mongodb.Device{
			ID:       1,
			ObjectID: 1,
			Name:     "device 1",
			SourceIDs: []mongodb.SourceID{
				{DeviceID: 1, ProtocolID: 1},
				{DeviceID: 1, ProtocolID: 2},
			},
			State: object.StateActive,
		},
	); err != nil {
		log.Fatalf("create device failed: %s", err)
	}
}
