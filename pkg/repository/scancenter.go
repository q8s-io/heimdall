package repository

import (
	"database/sql"
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"log"
)

func NewTaskImageScan(taskImageScan entity.TaskImageScan) error {
	imageVuln := entity.ImageVuln{}
	imageVuln.TaskImageScan = taskImageScan
	mysql.Client.Create(&imageVuln)
	return nil
}

func GetTaskImageScan(imageRequestInfo model.ImageRequestInfo) *[]entity.TaskImageScan {
	taskImageScanList := new([]entity.TaskImageScan)
	rows := new(sql.Rows)
	var err error

	if imageRequestInfo.ImageDigest == "" {
		rows, err = mysql.Client.Model(&entity.ImageVuln{}).Where("active = 1 AND image_name = ?", imageRequestInfo.ImageName).Rows()
	} else {
		rows, err = mysql.Client.Model(&entity.ImageVuln{}).Where("active = 1 AND image_name = ?  AND image_digest= ?", imageRequestInfo.ImageName, imageRequestInfo.ImageDigest).Rows()
	}
	if err != nil {
		log.Print(err)
		return taskImageScanList
	}
	defer rows.Close()

	for rows.Next() {
		var imageVuln entity.ImageVuln
		// 将sql.Rows扫描到entity中
		mysql.Client.ScanRows(rows, &imageVuln)
		*taskImageScanList = append(*taskImageScanList, imageVuln.TaskImageScan)
	}
	return taskImageScanList
}

func UpdateTaskImageScanDigest(taskID, igest string) {
	rows, err := mysql.Client.Model(&entity.ImageVuln{}).Update("image_digest", igest).Scopes(mysql.QuerytByTaskID(taskID)).Rows()
	if err != nil {
		log.Print(err)
		return
	}
	rows.Close()
}

func UpdateTaskImageScanStatus(taskID, status string) {
	rows, err := mysql.Client.Model(&entity.ImageVuln{}).Update("task_status", status).Scopes(mysql.QuerytByTaskID(taskID)).Rows()
	if err != nil {
		log.Print(err)
		return
	}
	rows.Close()
}

func UpdateTaskImageScanActive(imageName string, active int) {
	rows, err := mysql.Client.Model(&entity.ImageVuln{}).Update("active", active).Where("image_name = ?", imageName).Rows()
	if err != nil {
		log.Print(err)
		return
	}
	rows.Close()
}

func GetTaskStatus(taskID string) map[string]string {
	return redis.GetMapAll(taskID)
}

func DeleteTask(taskID string) {
	redis.DelKV(taskID)
}
