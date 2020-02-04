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
}

var SanPerfConfig = TSanPerfConfig{
	Default: TDefaultSanPerfConfig{
		Username: "",
		Password: "",
		Port: 8443,
		Interval: 60,
		Graphite: TGraphiteConfig{
			Address: "0.0.0.0:2003",
			Prefix: "storage.eseries",
		},
	},
}

func GetConfig(configPath string){//(configFile *TSanPerfConfig){
	if configPath!="" {
		buff, err := ioutil.ReadFile(configPath)
		if err!=nil {
			fmt.Println("Error while read config", err)
		}
		yaml.Unmarshal(buff, &SanPerfConfig)
	}
	//return
}
