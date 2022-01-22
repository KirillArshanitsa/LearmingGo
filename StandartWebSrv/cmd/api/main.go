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
	logrus.Println("Parse command line parameters")
	_, err := toml.DecodeFile(pathToConfFile, config)
	if err != nil{
		logrus.Println("Error parse config file ", pathToConfFile," error - ", err, "\nUse default config")
	}
	logrus.Println("Configure application")
	Api := api.NewApi(config)
	logrus.Println("Starting web server")
	logrus.Fatal(Api.Start())
}
