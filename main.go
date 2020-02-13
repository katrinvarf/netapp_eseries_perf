package main

import (
	"flag"
	"./config"
	"./login"
	"./getData"
	"time"
	"runtime"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"os"
	"io"
)

func main(){
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to the config file")
	flag.Parse()
	log := logrus.New()

	if err:=config.GetConfig(configPath); err!=nil{
		log.Fatal("Failed to get config file: Error: ", err)
		return
	}
	if err:=initLogger(log); err!=nil{
		log.Warning("Failed to initiate log file: Error: ", err)
	}

	runtime.Gosched()
	DeviceID := 1

	for{
		for i:=0; i<len(config.SanPerfConfig.Groups); i++{
			for j:=0; j<len(config.SanPerfConfig.Groups[i].Arrays); j++{
				DeviceAddress, err := checkAccessAd(log, config.SanPerfConfig.Default.Username, config.SanPerfConfig.Default.Password, config.SanPerfConfig.Groups[i].Arrays[j].Address, config.SanPerfConfig.Default.Port)
				if err!=nil{
					log.Warning("Failed to connect to device: ", config.SanPerfConfig.Groups[i].Arrays[j].Name, " :Error: ", err)
					break
				}
				log.Debug("Successful connect to address: ", DeviceAddress)
					go getData.GetAllData(log, config.SanPerfConfig.Default.Username, config.SanPerfConfig.Default.Password, config.SanPerfConfig.Default.Port, DeviceAddress, config.SanPerfConfig.Groups[i].Arrays[j].Name, DeviceID, config.SanPerfConfig.Groups[i].Groupname)
			}
		}

		time.Sleep(time.Second*time.Duration(config.SanPerfConfig.Default.Interval))
	}
}

func initLogger(log *logrus.Logger) (err error){
	logLevels := map[string]logrus.Level{"trace": logrus.TraceLevel, "debug": logrus.DebugLevel, "info": logrus.InfoLevel, "warn": logrus.WarnLevel, "error": logrus.ErrorLevel, "fatal": logrus.FatalLevel, "panic": logrus.PanicLevel}
	formatters := map[string]logrus.Formatter{"json": &logrus.JSONFormatter{TimestampFormat: "02-01-2006 15:04:05"}, "text": &logrus.TextFormatter{TimestampFormat: "02-01-2006 15:04:05", FullTimestamp: true}}

	if len(config.SanPerfConfig.Loggers)==2{
		setValuesLogrus(log, logLevels[config.SanPerfConfig.Loggers[0].Level], os.Stdout, formatters[config.SanPerfConfig.Loggers[0].Encoding])
		var rotateFileHook logrus.Hook
		rotateFileHook, err = rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
			Filename: config.SanPerfConfig.Loggers[1].File,
			MaxSize: 50, //megabytes
			MaxBackups: 3,
			MaxAge: 28, //days
			Level: logLevels[config.SanPerfConfig.Loggers[1].Level],
			Formatter: formatters[config.SanPerfConfig.Loggers[1].Encoding],
		})
		log.AddHook(rotateFileHook)
	} else if len(config.SanPerfConfig.Loggers)==1{
		if config.SanPerfConfig.Loggers[0].File=="stdout"{
			setValuesLogrus(log, logLevels[config.SanPerfConfig.Loggers[0].Level], os.Stdout, formatters[config.SanPerfConfig.Loggers[0].Encoding])
		} else{
			var file io.Writer
			file, err = os.OpenFile("carbonapi.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err!=nil{
				log.Warning("Failed to initialize log file: Error: ", err)
				return
			}
			setValuesLogrus(log, logLevels[config.SanPerfConfig.Loggers[0].Level], file, formatters[config.SanPerfConfig.Loggers[0].Encoding])
		}
	}

	if err!=nil{
		log.Warning("Failed to initialize file rotate hook: Error: ", err)
		return
	}
	return nil
}

func setValuesLogrus(log *logrus.Logger, level logrus.Level, output io.Writer, formatter logrus.Formatter){
	log.SetLevel(level)
	log.SetOutput(output)
	log.SetFormatter(formatter)
}

func checkAccessAd(log *logrus.Logger, Username string, Password string, Addresses []string, Port int)(DeviceAddress string, err error){
	for _, address := range Addresses{
		if err=login.Login(log, Username, Password, address, Port); err!=nil{
			log.Debug("Failed to connect to address: ", address, " :Error: ", err)
			continue
		}
		DeviceAddress = address
		return DeviceAddress, nil
	}
	return DeviceAddress, err
}

