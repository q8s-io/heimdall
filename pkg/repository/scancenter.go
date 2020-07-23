package repository

import (
	"database/sql"
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func NewTaskImageScan(taskImageScan entity.TaskImageScan) {
	imageVuln := entity.ImageVuln{}
	imageVuln.TaskImageScan = taskImageScan
	mysql.Client.Create(&imageVuln)
}

func GetTaskImageScan(imageRequestInfo model.ImageRequestInfo) *[]entity.TaskImageScan {
	taskImageScanList := new([]entity.TaskImageScan)
	rows := new(sql.Rows)
	var err error
	rows, err = mysql.Client.Model(&entity.ImageVuln{}).
		Where("active = 1 AND image_name = ?", imageRequestInfo.ImageName).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return taskImageScanList
	}
	defer rows.Close()

	for rows.Next() {
		var imageVuln entity.ImageVuln
		// 将 sql.Rows 扫描到 entity 中
		_ = mysql.Client.ScanRows(rows, &imageVuln)
		*taskImageScanList = append(*taskImageScanList, imageVuln.TaskImageScan)
	}
	return taskImageScanList
}

func UpdateTaskImageScanDigest(taskID, igest string) {
	rows, err := mysql.Client.Model(&entity.ImageVuln{}).
		Update("image_digest", igest).
		Scopes(mysql.QueryByTaskID(taskID)).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return
	}
	defer rows.Close()
}

func UpdateTaskImageScanStatus(taskID, status string) {
	rows, err := mysql.Client.Model(&entity.ImageVuln{}).
		Update("task_status", status).
		Scopes(mysql.QueryByTaskID(taskID)).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return
	}
	defer rows.Close()
}

func UpdateTaskImageScanActive(imageName string, active int) {
	rows, err := mysql.Client.Model(&entity.ImageVuln{}).
		Updates(map[string]interface{}{"active": active, "task_status": model.StatusSucceed}).
		Where("image_name = ?", imageName).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return
	}
	defer rows.Close()
}

func GetTaskStatus(taskID string) map[string]string {
	return redis.GetMapAll(taskID)
}

func DeleteTask(taskID string) {
	redis.DelKV(taskID)
}
