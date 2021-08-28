package config

import "fmt"

const (
	DBUser     = "user"
	DBPassword = "user"
	DBName     = "mydb"
	DBHost     = "mysql"
	DBPort     = "3306"
	DBtype     = "mysql"
)

func GetDBType() string {
	return DBtype
}

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return dataBase
}
