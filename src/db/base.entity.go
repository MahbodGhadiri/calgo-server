package db

import (
	"time"
)

type Base struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
