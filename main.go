package main

import (
	"flag"
	"./config"
)

func main(){
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to the config file")
	flag.Parse()
	conf := config.GetConfig(configPath)
}
