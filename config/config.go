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
	Address string `yaml:"address"`
}

type TSanPerfConfig struct {
	Default TDefaultSanPerfConfig
	Groups []TGroupConfig
}

func GetConfig(configPath string)(configFile *Config){
	if configPath!="" {
		buff, err := ioutil.ReadFile(configPath)
		if err!=nil {
			fmt.Println("Error while read config", err)
		}
		yaml.Unmarshal(buff, &configFile)
	}
}
