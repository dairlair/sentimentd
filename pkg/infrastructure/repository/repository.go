package repository

import (
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}