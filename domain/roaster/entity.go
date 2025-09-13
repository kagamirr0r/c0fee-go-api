package roaster

import (
	"c0fee-api/domain/bean"
	"time"
)

type Entity struct {
	ID        uint
	Name      string
	Address   string
	WebURL    string
	ImageKey  *string
	Beans     []bean.Summary
	CreatedAt time.Time
	UpdatedAt time.Time
}
