package object

import (
	"context"
)

type TagRepository interface {
	ReadTag(context.Context, *[]Tag) error
	CreateTag(context.Context, ...Tag) error
	UpdateTag(context.Context, ...Tag) error
	DeleteTag(context.Context, ...TagID) error
}

type Tag interface {
	Node
	GetTagID() TagID
	GetProfile() Profile
	GetState() State
}

type TagID interface {
	// TODO need methods
}

type Profile int

const (
	ProfileRTD      Profile = 0
	Profile1Minute  Profile = minute
	Profile2Minute  Profile = 2 * minute
	Profile3Minute  Profile = 3 * minute
	Profile4Minute  Profile = 4 * minute
	Profile5Minute  Profile = 5 * minute
	Profile6Minute  Profile = 6 * minute
	Profile10Minute Profile = 10 * minute
	Profile12Minute Profile = 12 * minute
	Profile15Minute Profile = 15 * minute
	Profile20Minute Profile = 20 * minute
	Profile30Minute Profile = 30 * minute
	Profile1Hour    Profile = hour
	Profile1Day     Profile = day
	Profile1Week    Profile = week
	Profile1Month   Profile = month
	Profile1Year    Profile = year
)

const (
	second = 1
	minute = 60 * second
	hour   = 60 * minute
	day    = 24 * hour
	week   = 7 * day
	month  = year / 12
	year   = 365.2425 * day
)
