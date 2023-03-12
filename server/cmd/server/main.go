package main

import (
	"flag"
	//"net/http"
	"github.com/t67y110v/software-engineering/internal/app/logging"
	"github.com/t67y110v/software-engineering/internal/app/server"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/configs.toml", "path to config file")
}

// @title Food Delivery API
// @version 1.0
// @description This API have endpoints for food delivery site

// @host localhost:4000
// @BasePath /

func main() {

	logging.Init()
	l := logging.GetLogger()
	l.Infoln("Parsing flag")
	flag.Parse()
	l.Infoln("Config initialization")
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		l.Fatal(err)
	}
	l.Infof("Starting apiserver addr : %s\n", config.BindAddr)
	if err := server.Start(config); err != nil {
		l.Fatal(err)
	}

}
