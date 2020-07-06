package repository

import (
	"fmt"
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"log"
)

func NewJobTrivy(jobScanner entity.JobScanner) {
	execSQL := fmt.Sprintf("INSERT INTO job_trivy (task_id, job_id, job_status, job_data, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
		jobScanner.TaskID, jobScanner.JobID, jobScanner.JobStatus, jobScanner.JobData, jobScanner.ImageName, jobScanner.ImageDigest, jobScanner.CreateTime, jobScanner.Active)
	_ = mysql.InserData(execSQL)
}

func GetJobTrivy(taskID string) *[]entity.JobScanner {
	execSQL := fmt.Sprintf("SELECT task_id, job_id, job_status, job_data FROM job_trivy WHERE task_id='%s'",
		taskID)
	jobTrivyDataList := new([]entity.JobScanner)
	err := mysql.Client.Select(jobTrivyDataList, execSQL)
	if err != nil {
		log.Println(err)
	}
	return jobTrivyDataList
}

func UpdateJobTrivy(jobScanner entity.JobScanner) {
	execSQL := fmt.Sprintf("UPDATE job_trivy SET job_status='%s', job_data='%s' WHERE job_id='%s'",
		jobScanner.JobStatus, jobScanner.JobData, jobScanner.JobID)
	_ = mysql.InserData(execSQL)
}

func SetJobTrivyStatus(taskID, status string) {
	redis.SetMap(taskID, "anchore", status)
}
