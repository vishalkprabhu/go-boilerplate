package storage

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConfig struct {
	User     string
	Password string
	Db       string
	Host     string
	Port     string
}

func (t DbConfig) GetDBType() string {
	return "mysql"
}

func NewDB(config DbConfig) *gorm.DB {
	var err error
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)

	log.Print(dataBase)

	DB, err := gorm.Open(config.GetDBType(), dataBase)

	if err != nil {
		log.Panic(err)
	}

	return DB
}
