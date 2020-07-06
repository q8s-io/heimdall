package repository

import (
	"fmt"
	"log"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
)

func NewTaskImageScan(taskImageScan entity.TaskImageScan) error {
	execSQL := fmt.Sprintf("INSERT INTO image_vuln (task_id, task_status, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', %d)",
		taskImageScan.TaskID, taskImageScan.TaskStatus, taskImageScan.ImageName, taskImageScan.ImageDigest, taskImageScan.CreateTime, taskImageScan.Active)
	err := mysql.InserData(execSQL)
	return err
}

func GetTaskImageScan(imageRequestInfo model.ImageRequestInfo) *[]entity.TaskImageScan {
	var execSQL string
	if imageRequestInfo.ImageDigest == "" {
		execSQL = fmt.Sprintf("SELECT task_id, task_status, image_name, image_digest, create_time FROM image_vuln WHERE active=1 AND image_name='%s'",
			imageRequestInfo.ImageName)
	} else {
		execSQL = fmt.Sprintf("SELECT task_id, task_status, image_name, image_digest, create_time FROM image_vuln WHERE active=1 AND image_name='%s' AND image_digest='%s'",
			imageRequestInfo.ImageName, imageRequestInfo.ImageDigest)
	}
	taskImageScanList := new([]entity.TaskImageScan)
	err := mysql.Client.Select(taskImageScanList, execSQL)
	if err != nil {
		log.Println(err)
	}
	return taskImageScanList
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
