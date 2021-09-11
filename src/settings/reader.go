package settings

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Database struct {
		DbName string `yaml:"name"`
		DbUser string `yaml:"username"`
		DbPass string `yaml:"password"`
		DbHost string `yaml:"host"`
		DbPort string `yaml:"port"`
	}
}

func GetSettings(yamlPath string) (*YamlConfig, error) {

	file, err := ioutil.ReadFile(yamlPath)

	if err != nil {
		log.Println("Reading settings from config file failed", yamlPath, err)
		return nil, err

	}

	var config YamlConfig

	err = yaml.Unmarshal(file, &config)

	if err != nil {
		log.Println("Parsing error for yaml", &yamlPath)
		return nil, err
	}
	return &config, nil
}
