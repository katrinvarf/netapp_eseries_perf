package getData

import(
	"fmt"
)

var sections = []string{
	"Controller",
	"Drive",
	"Interface",
	"System",
	"Volume"}
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
	"queueDepthMax"}

func get_drive_stats(){}
func get_controller_stats(){}
func get_interface_stats(){}
func get_system_stats(){}
func get_volume_stats(){}

func GetData(){
	
}
