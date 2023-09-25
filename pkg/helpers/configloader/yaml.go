package configloader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"proyecto-golang/pkg/helpers/environment"
)

const PrefixConfigFile = "config"
const ConfigDir = "config"

type Config struct {
	Conn struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"databaseName"`
	} `yaml:"postgresBD"`
}

func ReadConfig() Config {
	var config Config

	// Get File
	filename := getConfigFileNames()
	// Open YAML file
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			log.Println(err.Error())
		}
	}

	return config
}

func getConfigFileNames() string {

	env := environment.GetEnv()

	yamlFileName := fmt.Sprintf("%s/%s.yaml", ConfigDir, PrefixConfigFile)
	yamlFileNameEnv := fmt.Sprintf("%s/%s-%s.yaml", ConfigDir, PrefixConfigFile, env)

	if _, err := os.Stat(yamlFileNameEnv); os.IsNotExist(err) {
		yamlFileNameEnv = yamlFileName
	}

	return yamlFileNameEnv
}
