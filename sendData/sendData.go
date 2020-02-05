package sendData

import(
	"fmt"
	"gopkg.in/fgrosse/graphigo.v2"
	"../config"
)

func SendObjectPerfs(PerfMap map[string]float64){
	Connection := graphigo.NewClient(config.SanPerfConfig.Default.Graphite.Address)
	Connection.Prefix = config.SanPerfConfig.Default.Graphite.Prefix
	Connection.Connect()
	for name, value := range PerfMap{
		err := Connection.Send(graphigo.Metric{Name: name, Value: value})
		if err!=nil{
			fmt.Println(err)
		}
	}
}
