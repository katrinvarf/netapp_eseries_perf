package config

import(
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type TGraphiteConfig struct{
	Address string `yaml:"address"`
	Prefix string `yaml:"prefix"`
}

type TDefaultSanPerfConfig struct{
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port int `yaml:"port"`
	Interval int `yaml:"interval"`
	Graphite TGraphiteConfig `yaml:"graphite"`
}

type TGroupConfig struct {
	Groupname string `yaml:"groupname"`
	Arrays []TSanArray `yaml:"arrays"`
}

type TSanArray struct {
	Name string `yaml:"name"`
	Address []string `yaml:"address"`
}

type TSanPerfConfig struct {
	Default TDefaultSanPerfConfig `yaml:"default"`
	Groups []TGroupConfig `yaml:"groups"`
	Loggers []TLoggingConfig `yaml:"logging"`
}

type TLoggingConfig struct {
	Loggername string `yaml:"logger"`
	File string `yaml:"file"`
	Level string `yaml:"level"`
	Encoding string `yaml:"encoding"`
}

var SanPerfConfig = TSanPerfConfig{}

func GetConfig(configPath string) (err error){
	var buff []byte
	buff, err = ioutil.ReadFile(configPath)
	if err!=nil{
		fmt.Println("Failed to read config", err)
		return
	}
	err = yaml.Unmarshal(buff, &SanPerfConfig)
	if err!=nil{
		fmt.Println("Failed to decode document", err)
		return
	}
	return nil
}
