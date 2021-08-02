package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server struct{
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}

var Conf *Config

func InitConf(){
	f, err := os.Open("./config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read from config file. Error %+v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Conf)
	if err != nil {
		log.Fatal("Failed to decode yml")
	}
}