package scancenter

import (
	"fmt"
	"log"

	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewTaskImageScan(taskImageScanData entity.TaskImageScanData) error {
	execSQL := fmt.Sprintf("INSERT INTO image_vuln (task_id, task_status, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', %d)",
		taskImageScanData.TaskID, taskImageScanData.TaskStatus, taskImageScanData.ImageName, taskImageScanData.ImageDigest, taskImageScanData.CreateTime, taskImageScanData.Active)
	err := mysql.InserData(execSQL)
	return err
}

func GetTaskImageScan(imageRequestInfo entity.ImageRequestInfo) *[]entity.TaskImageScanData {
	var execSQL string
	if imageRequestInfo.ImageDigest == "" {
		execSQL = fmt.Sprintf("SELECT task_id, task_status, image_name, image_digest, create_time FROM image_vuln WHERE active=1 AND image_name='%s'",
			imageRequestInfo.ImageName)
	} else {
		execSQL = fmt.Sprintf("SELECT task_id, task_status, image_name, image_digest, create_time FROM image_vuln WHERE active=1 AND image_name='%s' AND image_digest='%s'",
			imageRequestInfo.ImageName, imageRequestInfo.ImageDigest)
	}
	taskImageScanDataList := new([]entity.TaskImageScanData)
	err := mysql.Client.Select(taskImageScanDataList, execSQL)
	if err != nil {
		log.Println(err)
	}
	return taskImageScanDataList
}

func UpdateTaskImageScanDigest(taskID, igest string) {
	execSQL := fmt.Sprintf("UPDATE image_vuln SET image_digest='%s' WHERE task_id='%s'",
		igest, taskID)
	_ = mysql.InserData(execSQL)
}

func UpdateTaskImageScanStatus(taskID, status string) {
	execSQL := fmt.Sprintf("UPDATE image_vuln SET task_status='%s' WHERE task_id='%s'",
		status, taskID)
	_ = mysql.InserData(execSQL)
}

func UpdateTaskImageScanActive(imageName string, active int) {
	execSQL := fmt.Sprintf("UPDATE image_vuln SET active=%d WHERE image_name='%s'",
		active, imageName)
	_ = mysql.InserData(execSQL)
}

func GetTaskStatus(taskID string) map[string]string {
	return redis.GetMapAll(taskID)
}

func DeleteTask(taskID string) {
	redis.DelKV(taskID)
}
