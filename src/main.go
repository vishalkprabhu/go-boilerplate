package main

import (
	"fmt"
	"log"
	"os"

	"go-boilerplate/src/settings"
	"go-boilerplate/src/storage"
)

func main() {

	// Echo instance

	env := os.Getenv("ENV_NAME")
	yamlPath := fmt.Sprintf("config/%s.yaml", env)

	config, errCfg := settings.GetSettings(yamlPath)

	if errCfg != nil {
		log.Fatal(errCfg)
	}

	var dbCfg storage.DbConfig
	dbCfg.User = config.Database.DbUser
	dbCfg.Password = config.Database.DbPass
	dbCfg.Host = config.Database.DbHost
	dbCfg.Password = config.Database.DbPass
	dbCfg.Port = config.Database.DbPort

	http_app := InitilizeAppRequirements(dbCfg)
	http_app.Run()

}
