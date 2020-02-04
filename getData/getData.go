package getData

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"strconv"
	"encoding/json"
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

func GetAllData(Username string, Password string, DevicePort int, DeviceAddress string, DeviceName string, DeviceID int, GroupName string){
	for _, section := range sections{
		if section=="Controller"{
			fmt.Println(getSection(Username, Password, section, DevicePort, DeviceAddress, DeviceName, DeviceID, GroupName))
			break
		}
	}
}

func getSectionPerfData(Username string, Password string, SectionAPI string, DevicePort int, DeviceAddress string, DeviceName string, DeviceID int) interface{} {
	urlString := "https://" + DeviceAddress + ":" + strconv.Itoa(DevicePort) + "/devmgr/v2/storage-systems/" + strconv.Itoa(DeviceID) + "/" + SectionAPI
	tr := &http.Transport{
		TLSClientConfig : &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	request, _ := http.NewRequest("GET", urlString, nil)
	request.SetBasicAuth(Username, Password)
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	var buf []byte
	buf, _ = ioutil.ReadAll(resp.Body)
	var raw interface{}
	json.Unmarshal(buf, &raw)
	return raw
}

func getSection(Username string, Password string, Section string, DevicePort int, DeviceAddress string, DeviceName string, DeviceID int, GroupName string) (result map[string]float64) {
	result = make(map[string]float64)
	var perf_data interface{}
	switch Section{
		case "Controller":
			perf_data = getSectionPerfData(Username, Password, "analysed-controller-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["controllerId"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return
		case "Drive":
			perf_data = getSectionPerfData(Username, Password, "analysed-drive-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["diskId"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 11, 12}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return
		case "Interface":
			perf_data = getSectionPerfData(Username, Password, "analysed-interface-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["interfaceId"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 11, 12}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return
		case "System":
			perf_data = getSectionPerfData(Username, Password, "analysed-system-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["storageSystemName"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return
		case "Volume":
			perf_data = getSectionPerfData(Username, Password, "analysed-volume-statistics", DevicePort, DeviceAddress, DeviceName, DeviceID)
			for _, item := range perf_data.([]interface{}){
				Name := item.(map[string]interface{})["volumeName"].(string)
				for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 11, 12}{
					metricName := statisticName[num]
					metricValue := item.(map[string]interface{})[statisticName[num]].(float64)
					result[GroupName + "." + DeviceName + "." + Section + "." + Name + "." + metricName] = metricValue
				}
			}
			return
		case "Pool":
			perf_data = getSectionPerfData(Username, Password, "", DevicePort, DeviceAddress, DeviceName, DeviceID)
			for _, num := range []int{13, 14, 15}{
				metricName := statisticName[num]
				metricValue := perf_data.(map[string]interface{})[statisticName[num]].(float64)
				result[GroupName + "." + DeviceName + "." + Section + "." + metricName] = metricValue
			}
			return
	}
	return
}
