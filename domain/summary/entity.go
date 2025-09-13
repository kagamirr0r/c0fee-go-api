package summary

import (
	"github.com/google/uuid"
)

type Country struct {
	ID   uint
	Name string
}

type Area struct {
	ID        uint
	CountryID uint
	Name      string
}

type Farm struct {
	ID     uint
	AreaID uint
	Name   string
}

type Farmer struct {
	ID   uint
	Name string
}

type User struct {
	ID   uuid.UUID
	Name string
}

type Roaster struct {
	ID   uint
	Name string
}

type ProcessMethod struct {
	ID   uint
	Name string
}

type RoastLevel struct {
	ID    uint
	Name  string
	Level int
}

type Variety struct {
	ID   uint
	Name string
}

type Bean struct {
	ID   uint
	Name *string
}

type BeanRating struct {
	ID         uint
	UserID     uuid.UUID
	User       User
	Bitterness int
	Acidity    int
	Body       int
	FlavorNote string
}
