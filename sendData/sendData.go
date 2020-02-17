package sendData

import(
	"gopkg.in/fgrosse/graphigo.v2"
	"github.com/sirupsen/logrus"
	"github.com/katrinvarf/netapp_eseries_perf/config"
)

func SendObjectPerfs(log *logrus.Logger, PerfMap map[string]float64){
	Connection := graphigo.NewClient(config.SanPerfConfig.Default.Graphite.Address)
	Connection.Prefix = config.SanPerfConfig.Default.Graphite.Prefix
	Connection.Connect()
	for name, value := range PerfMap{
		err := Connection.Send(graphigo.Metric{Name: name, Value: value})
		if err!=nil{
			log.Warning("Failed to send metric: ", name, " = ", value, " :Error: ", err)
			continue
		}
		log.Debug("Metric sent successfully: ", name, " = ", value)
	}
}
