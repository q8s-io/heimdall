package repository

import (
	"log"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
)

func NewJobAnchore(jobScanner entity.JobScanner) {
	jobAnchore := entity.JobAnchore{}
	jobAnchore.JobScanner = jobScanner
	mysql.Client.Create(&jobAnchore)
}

func GetJobAnchore(taskID string) *[]entity.JobScanner {
	rows, err := mysql.Client.Model(&entity.JobAnchore{}).Scopes(mysql.QueryByTaskID(taskID)).Rows()
	jobAnchoreDataList := new([]entity.JobScanner)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var jobAnchore entity.JobAnchore
		_ = mysql.Client.ScanRows(rows, &jobAnchore)
		*jobAnchoreDataList = append(*jobAnchoreDataList, jobAnchore.JobScanner)
	}
	return jobAnchoreDataList
}

func UpdateJobAnchore(jobScanner entity.JobScanner) {
	jobAnchore := entity.JobAnchore{}
	jobAnchore.JobScanner = jobScanner
	rows, err := mysql.Client.Model(&entity.JobAnchore{}).Updates(jobAnchore).
		Scopes(mysql.QueryByJobID(jobScanner.JobID)).Rows()
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()
}

func SetJobAnchoreStatus(taskID, status string) {
	redis.SetMap(taskID, "anchore", status)
}
