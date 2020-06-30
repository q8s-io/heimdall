package service

import (
	"fmt"

	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewTaskImageScan(imageVulnData models.ImageVulnData) error {
	execSQL := fmt.Sprintf("INSERT INTO image_vuln (task_id, task_status, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', %d)",
		imageVulnData.TaskID, imageVulnData.TaskStatus, imageVulnData.ImageName, imageVulnData.ImageDigest, imageVulnData.CreateTime, imageVulnData.Active)
	err := mysql.InserData(execSQL)
	return err
}

func UpdateTaskImageScan(taskID, status string) {
	execSQL := fmt.Sprintf("UPDATE image_vuln SET task_status='%s' WHERE task_id='%s'",
		status, taskID)
	_ = mysql.InserData(execSQL)
}

func GetTaskStatus(taskID string) map[string]string {
	return redis.GetMapAll(taskID)
}

func DeleteTask(taskID string) {
	redis.DelKV(taskID)
}
