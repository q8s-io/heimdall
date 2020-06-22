package scancenter

import (
	"log"
)

// CreateScan
func CreateScan(imageName string) string {
	log.Println(imageName)

	//create task id、job id

	//get data by image name, if status is running, return data

	//write to mysql, image name、task id、task status(running)

	//write to redis, task id、job id、status(running)

	//write kafka, analyzer topic, task id、job id、image name

	return ""
}

// GetScanTaskData
func GetScanTaskData(taskID string) string {
	log.Println(taskID)
	return ""
}
