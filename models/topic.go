package models

import (
	"time"
)

type Topic struct {
	ID        int
	Title     string
	CreatedAt time.Time
}
