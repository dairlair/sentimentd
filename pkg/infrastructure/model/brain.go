package model

import (
	"github.com/jinzhu/gorm"
)

type Brain struct {
	gorm.Model
	ID int64
	Name string
	Description string
}

func (brain *Brain) GetID() int64 {
	return brain.ID
}

func (brain *Brain) GetName() string {
	return brain.Name
}

func (brain *Brain) GetDescription() string {
	return brain.Description
}