package main

import (
	"StandartWebServer/internal/api"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var(
	pathToConfFile string
)

func init() {
	flag.StringVar(&pathToConfFile, "pathToConfigFile", "config/config.toml", "path to config file")
}

func main(){
	logrus.Println("Start application")
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(pathToConfFile, config)
	if err != nil{
		logrus.Printf("Error parse config file %s - %s\nUse default config", pathToConfFile, err)
	}
	logrus.Println("Port " + config.Port)
	logrus.Println("Log level " + config.LogLevel)
	logrus.Println("Configure application")
	Api := api.NewApi(config)
	logrus.Fatal(Api.Start())
}
