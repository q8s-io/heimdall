package repository

import (
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func NewJobTrivy(jobScanner entity.JobScanner) {
	jobTrivy := entity.JobTrivy{}
	jobTrivy.JobScanner = jobScanner
	mysql.Client.Create(&jobTrivy)
}

func GetJobTrivy(taskID string) *[]entity.JobScanner {
	jobTrivyDataList := new([]entity.JobScanner)

	rows, err := mysql.Client.Model(&entity.JobTrivy{}).
		Scopes(mysql.QueryByTaskID(taskID)).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return jobTrivyDataList
	}
	defer rows.Close()

	for rows.Next() {
		var jobTrivy entity.JobTrivy
		_ = mysql.Client.ScanRows(rows, &jobTrivy)
		*jobTrivyDataList = append(*jobTrivyDataList, jobTrivy.JobScanner)
	}

	return jobTrivyDataList
}

func UpdateJobTrivy(jobScanner entity.JobScanner) {
	jobTrivy := entity.JobTrivy{}
	jobTrivy.JobScanner = jobScanner

	rows, err := mysql.Client.Model(&entity.JobTrivy{}).
		Scopes(mysql.QueryByJobID(jobScanner.JobID)).
		Updates(jobTrivy).
		Rows()
	if err != nil {
		xray.ErrMini(err)
		return
	}
	defer rows.Close()
}

func SetJobTrivyStatus(taskID, status string) {
	redis.SetMap(taskID, "trivy", status)
}
