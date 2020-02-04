package main

import (
	"flag"
	"./config"
	"./login"
	"./getData"
	"time"
	"runtime"
	//"fmt"
)

func main(){
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to the config file")
	flag.Parse()
	config.GetConfig(configPath)
	DeviceID := 1
	runtime.Gosched()
	for i:=0; i<len(config.SanPerfConfig.Groups); i++{
		for j:=0; j<len(config.SanPerfConfig.Groups[i].Arrays); j++{
			DeviceAddress := checkAccessAd(config.SanPerfConfig.Default.Username, config.SanPerfConfig.Default.Password, config.SanPerfConfig.Groups[i].Arrays[j].Address, config.SanPerfConfig.Default.Port)
			if DeviceAddress!=""{
				go worker(config.SanPerfConfig.Default.Username, config.SanPerfConfig.Default.Password, config.SanPerfConfig.Default.Port, DeviceAddress, config.SanPerfConfig.Groups[i].Arrays[j].Name, DeviceID, config.SanPerfConfig.Groups[i].Groupname)
			}
		}
	}
}

func worker(Username string, Password string, Port int, Address string, DeviceName string, DeviceID int, GroupName string){
	for{
		getData.GetAllData(Username, Password, Port, Address, DeviceName, DeviceID, GroupName)
		time.Sleep(time.Second*time.Duration(config.SanPerfConfig.Default.Interval))
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


