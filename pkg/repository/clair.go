package repository

import (
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func NewJobClair(jobScanner entity.JobScanner) {
	jobClair := entity.JobClair{}
	jobClair.JobScanner = jobScanner
	mysql.Client.Create(&jobClair)
}

func GetJobClair(taskID string) *[]entity.JobScanner {
	jobClairDataList := new([]entity.JobScanner)

	rows, err := mysql.Client.Model(&entity.JobClair{}).
		Scopes(mysql.QueryByTaskID(taskID)).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return jobClairDataList
	}
	defer rows.Close()

	for rows.Next() {
		var jobClair entity.JobClair
		_ = mysql.Client.ScanRows(rows, &jobClair)
		*jobClairDataList = append(*jobClairDataList, jobClair.JobScanner)
	}
	return jobClairDataList
}

func UpdateJobClair(jobScanner entity.JobScanner) {
	jobClair := entity.JobClair{}
	jobClair.JobScanner = jobScanner

	rows, err := mysql.Client.Model(&entity.JobClair{}).
		Scopes(mysql.QueryByJobID(jobScanner.JobID)).
		Updates(jobClair).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return
	}
	defer rows.Close()
}

func SetJobClairStatus(taskID, status string) {
	redis.SetMap(taskID, "clair", status)
}
