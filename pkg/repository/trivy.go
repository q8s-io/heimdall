package repository

import (
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"log"
)

func NewJobTrivy(jobScanner entity.JobScanner) {
	jobTrivy := entity.JobTrivy{}
	jobTrivy.JobScanner = jobScanner
	mysql.Client.Create(&jobTrivy)
}

func GetJobTrivy(taskID string) *[]entity.JobScanner {
	jobTrivyDataList := new([]entity.JobScanner)
	rows, err := mysql.Client.Model(&entity.JobTrivy{}).Scopes(mysql.QuerytByTaskID(taskID)).Rows()
	if err != nil {
		log.Print(err)
		return jobTrivyDataList
	}
	defer rows.Close()

	for rows.Next() {
		var jobTrivy entity.JobTrivy
		mysql.Client.ScanRows(rows, &jobTrivy)
		*jobTrivyDataList = append(*jobTrivyDataList, jobTrivy.JobScanner)
	}

	return jobTrivyDataList
}

func UpdateJobTrivy(jobScanner entity.JobScanner) {
	jobTrivy := entity.JobTrivy{}
	jobTrivy.JobScanner = jobScanner
	rows, err := mysql.Client.Model(&entity.JobTrivy{}).Updates(jobTrivy).Scopes(mysql.QuerytByTaskID(jobScanner.JobID)).Rows()
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()
}

func SetJobTrivyStatus(taskID, status string) {
	redis.SetMap(taskID, "trivy", status)
}
