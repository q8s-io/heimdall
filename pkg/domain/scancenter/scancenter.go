package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
	"github.com/q8s-io/heimdall/pkg/service"
)

func PreperScanenter(imageInfoRequest *models.ImageInfoRequest) (interface{}, error) {
	//preper task data
	imageVulnInfo := ConvertPreperTask(imageInfoRequest)
	err := service.NewImageScan(imageVulnInfo)
	if err != nil {
		return nil, err
	}
	//preper job analyzer
	analyzerJobInfo := ConvertPreperJobAnalyzer(&imageVulnInfo)
	service.NewImageAnalyzer(analyzerJobInfo)
	return imageVulnInfo, nil
}

func GetScanTask(imageName string) string {
	//create task id、job id

	//get data by image name, if status is running, return data

	//write to mysql, image name、task id、task status(running)

	//write to redis, task id、job id、status(running)

	//write kafka, analyzer topic, task id、job id、image name

	return ""
}

func GetScanTaskData(taskID string) string {
	return ""
}
