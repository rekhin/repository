package mongodb

import (
	"context"
	"fmt"

	"github.com/rekhin/repository/example/object"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) ReadObject(context.Context, *[]object.Object) error {
	return fmt.Errorf("Not implemented")
}
func (r *Repository) CreateObject(context.Context, ...object.Object) error {
	return fmt.Errorf("Not implemented")
}
func (r *Repository) UpdateObject(context.Context, ...object.Object) error {
	return fmt.Errorf("Not implemented")
}
func (r *Repository) DeleteObject(context.Context, ...object.ObjectID) error {
	return fmt.Errorf("Not implemented")
}

func (r *Repository) ReadDevice(context.Context, *[]object.Device) error {
	return fmt.Errorf("Not implemented")
}
func (r *Repository) CreateDevice(context.Context, ...object.Device) error {
	return fmt.Errorf("Not implemented")
}
func (r *Repository) UpdateDevice(context.Context, ...object.Device) error {
	return fmt.Errorf("Not implemented")
}
func (r *Repository) DeleteDevice(context.Context, ...object.DeviceID) error {
	return fmt.Errorf("Not implemented")
}
