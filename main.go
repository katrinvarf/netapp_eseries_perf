package main

import (
	"flag"
	"./config"
	"./login"
	"./getData"
	//"fmt"
)

func main(){
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to the config file")
	flag.Parse()
	conf := config.GetConfig(configPath)
	DeviceID := 1
	for i:=0; i<len(conf.Groups); i++{
		for j:=0; j<len(conf.Groups[i].Arrays); j++{
			DeviceAddress := checkAccessAd(conf.Default.Username, conf.Default.Password, conf.Groups[i].Arrays[j].Address, conf.Default.Port)
			if DeviceAddress!=""{
				getData.GetAllData(conf.Default.Username, conf.Default.Password, conf.Default.Port, DeviceAddress, conf.Groups[i].Arrays[j].Name, DeviceID, conf.Groups[i].Groupname)
			}
		}
	}
}

func checkAccessAd(Username string, Password string, Addresses []string, Port int)(DeviceAddress string){
	for _, address := range Addresses{
		if login.Login(Username, Password, address, Port){
			DeviceAddress = address
			return
		}
	}
	return
}
