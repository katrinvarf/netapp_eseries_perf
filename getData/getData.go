package getData

import(
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"strconv"
	"../sendData"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

var sections = []string{
	"Controller",
	"Drive",
	"Interface",
	"System",
	"Volume",
	"Pool"}
var statisticName =[]string{
	"readIOps",
	"writeIOps",
	"combinedIOps",
	"readResponseTime",
	"writeResponseTime",
	"combinedResponseTime",
	"readThroughput",
	"writeThroughput",
	"combinedThroughput",
	"maxCpuUtilization",
	"cpuAvgUtilization",
	"queueDepthTotal",
	"queueDepthMax",
	"freePoolSpace",
	"unconfiguredSpace",
	"usedPoolSpace"}

func GetAllData(log *logrus.Logger, Username string, Password string, DevicePort int, DeviceAddress string, DeviceName string, DeviceID int, GroupName string){
	for _, section := range sections{
		metrics, err := getSection(log, Username, Password, section, DevicePort, DeviceAddress, DeviceName, DeviceID, GroupName)
		if err!=nil{
			log.Warning("Failed to get ", section, " metrics, device: ", DeviceName, "; Error: ", err)
			continue
		}
		go sendData.SendObjectPerfs(log, metrics)
	}
}

func getSectionPerfData(log *logrus.Logger, Username string, Password string, SectionAPI string, DevicePort int, DeviceAddress string, DeviceName string, DeviceID int) (interface{}, error) {
	urlString := "https://" + DeviceAddress + ":" + strconv.Itoa(DevicePort) + "/devmgr/v2/storage-systems/" + strconv.Itoa(DeviceID) + "/" + SectionAPI
	tr := &http.Transport{
		TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	request, err := http.NewRequest("GET", urlString, nil)
	if err!=nil{
		log.Warning("Failed to create http request: Error: ", err)
		return nil, err
	}
	request.SetBasicAuth(Username, Password)
	resp, err := client.Do(request)
	if err!=nil{
		log.Warning("Failed to do client request: Error: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	var buf []byte
	buf, err = ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Warning("Failed to read response body: Error: ", err)
		return nil, err
	}
	var raw interface{}
	json.Unmarshal(buf, &raw)
	return raw, nil
}

func getSection(log *logrus.Logger, Username string, Password string, Section string, DevicePort int, DeviceAddress string, DeviceName string, DeviceID int, GroupName string) (map[string]float64, error) {
	result := make(map[string]float64)
	switch Section{
		case "Controller":
			perf_data, err := getSectionPerfData(log, Username, Password, "analysed-controller-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			if err!=nil{
				return result, err
			}
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["controllerId"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return result, nil
		case "Drive":
			perf_data, err := getSectionPerfData(log, Username, Password, "analysed-drive-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			if err!=nil{
				return result, err
			}
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["diskId"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 12}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return result, nil
		case "Interface":
			perf_data, err := getSectionPerfData(log, Username, Password, "analysed-interface-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			if err!=nil{
				return result, err
			}
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["interfaceId"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 11, 12}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return result, nil
		case "System":
			perf_data, err := getSectionPerfData(log, Username, Password, "analysed-system-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			if err!=nil{
				return result, err
			}
			Name := perf_data.(map[string]interface{})["storageSystemName"].(string)
			for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}{
				metricName := statisticName[num]
				metricValue := perf_data.(map[string]interface{})[statisticName[num]].(float64)
				result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
			}
			return result, nil
		case "Volume":
			perf_data, err := getSectionPerfData(log, Username, Password, "analysed-volume-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			if err!=nil{
				return result, err
			}
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["volumeName"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 11, 12}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return result, nil
		case "Pool":
			perf_data, err := getSectionPerfData(log, Username, Password, "", DevicePort, DeviceAddress, DeviceName, DeviceID)
			if err!=nil{
				return result, err
			}
			for _, num := range []int{13, 14, 15}{
				metricName := statisticName[num]
				metricValue, _ := strconv.ParseFloat(perf_data.(map[string]interface{})[statisticName[num]].(string), 64)
				result[GroupName + "." + DeviceName + "." + Section + "." + metricName] = metricValue
			}
			return result, nil
	}
	return result, nil
}
