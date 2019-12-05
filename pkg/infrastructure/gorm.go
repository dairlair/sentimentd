package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}