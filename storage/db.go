package storage

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/vishalkprabhu/go-boilerplate/config"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
