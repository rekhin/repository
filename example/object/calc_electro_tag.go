package object

import (
	"context"
)

type CalcElectroTagRepository interface {
	ReadCalcElectroTag(context.Context, *[]CalcElectroTag) error
	CreateCalcElectroTag(context.Context, ...CalcElectroTag) error
	UpdateCalcElectroTag(context.Context, ...CalcElectroTag) error
	DeleteCalcElectroTag(context.Context, ...CalcElectroTagID) error
}

type CalcElectroTag interface {
	ElectroTag
	GetCalcElectroTagID() CalcElectroTagID
}

type CalcElectroTagID int
