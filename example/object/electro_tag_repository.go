package object

import (
	"context"
)

type ElectroTagRepository interface {
	ReadElectroTag(context.Context, *[]ElectroTag) error
	CreateElectroTag(context.Context, ...ElectroTag) error
	UpdateElectroTag(context.Context, ...ElectroTag) error
	DeleteElectroTag(context.Context, ...ElectroTagID) error
}

type ElectroTag interface {
	Tag
	GetElectroTagID() ElectroTagID
	GetElectroTagType() ElectroTagType
}

type ElectroTagID interface {
	TagID
}

type ElectroTagType int

const (
	ElectroTagTypeImportActivePower ElectroTagType = iota
	ElectroTagTypeExportReactivePower
	ElectroTagTypeExportActivePower
	ElectroTagTypeImportReactivePower
)
