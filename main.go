package main

import (
	"flag"
	"github.com/katrinvarf/netapp_eseries_perf/config"
	"github.com/katrinvarf/netapp_eseries_perf/login"
	"github.com/katrinvarf/netapp_eseries_perf/getData"
	"time"
	"runtime"
	"github.com/sirupsen/logrus"
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
	logLevels := map[string]logrus.Level{"trace": logrus.TraceLevel, "debug": logrus.DebugLevel, "info": logrus.InfoLevel, "warn": logrus.WarnLevel, "error": logrus.ErrorLevel, "fatal": logrus.FatalLevel, "panic": logrus.PanicLevel}
	formatters := map[string]logrus.Formatter{"json": &logrus.JSONFormatter{TimestampFormat: "02-01-2006 15:04:05"}, "text": &logrus.TextFormatter{TimestampFormat: "02-01-2006 15:04:05", FullTimestamp: true}}
	var writers []io.Writer
	var level logrus.Level
	var format logrus.Formatter
	for i, _ := range(config.SanPerfConfig.Loggers){
		if config.SanPerfConfig.Loggers[i].Loggername=="FILE"{
			file, err  := os.OpenFile(config.SanPerfConfig.Loggers[i].File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err!=nil{
				log.Warning("Failed to initialize log file: Error: ", err)
			}
			defer file.Close()
			writers = append(writers, file)
			level = logLevels[config.SanPerfConfig.Loggers[i].Level]
			format = formatters[config.SanPerfConfig.Loggers[i].Encoding]
		} else {
			writers = append(writers, os.Stdout)
			level = logLevels[config.SanPerfConfig.Loggers[i].Level]
			format = formatters[config.SanPerfConfig.Loggers[i].Encoding]
		}
	}

	if len(writers)!=0{
		mw := io.MultiWriter(writers...)
		setValuesLogrus(log, level, mw, format)
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

