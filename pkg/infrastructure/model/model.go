package model

import "time"

type Model struct {
	ID        int64
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
