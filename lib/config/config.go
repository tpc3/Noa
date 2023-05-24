package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Misskey struct {
		Token string
		Host  string
	}
	TextBlacklist []string `yaml:"text_blacklist"`
}

const configFile = "./config.yml"

var Loadconfig Config

func init() {
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Config load faild: ", err)
	}
	err = yaml.Unmarshal(file, &Loadconfig)
	if err != nil {
		log.Fatal("Consig parse faild: ", err)
	}

	if Loadconfig.Misskey.Token == "" {
		log.Fatal("Token is empty")
	}
}
