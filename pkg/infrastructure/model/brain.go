package model

import (
	"github.com/jinzhu/gorm"
)

type Brain struct {
	gorm.Model
	BrainID int64
	Name string
	Description string
}

func (brain *Brain) GetID() int64 {
	return brain.BrainID
}

func (brain *Brain) GetName() string {
	return brain.Name
}

func (brain *Brain) GetDescription() string {
	return brain.Description
}