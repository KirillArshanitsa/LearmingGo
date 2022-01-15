package main

import (
	"StandartWebServer/internal/api"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"net/http"

	//	"net/http"
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
	_, err := toml.Decode(pathToConfFile, &config)
	if err != nil{
		logrus.Printf("Error parse config file %s - %s\nUse default config", pathToConfFile, err)
	}
	logrus.Println("Configure application")
	Api := api.NewApi(config)
	logrus.Fatal(http.ListenAndServe("localhost:" + Api.Config.Port, Api.Start()))
}
