package repository

import (
	"fmt"
	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"log"
)

func NewJobClair(jobScanner entity.JobScanner) {
	execSQL := fmt.Sprintf("INSERT INTO job_clair (task_id, job_id, job_status, job_data, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
		jobScanner.TaskID, jobScanner.JobID, jobScanner.JobStatus, jobScanner.JobData, jobScanner.ImageName, jobScanner.ImageDigest, jobScanner.CreateTime, jobScanner.Active)
	_ = mysql.InserData(execSQL)
}

func GetJobClair(taskID string) *[]entity.JobScanner {
	execSQL := fmt.Sprintf("SELECT task_id, job_id, job_status, job_data FROM job_clair WHERE task_id='%s'",
		taskID)
	jobClairDataList := new([]entity.JobScanner)
	err := mysql.Client.Select(jobClairDataList, execSQL)
	if err != nil {
		log.Println(err)
	}
	return jobClairDataList
}

func UpdateJobClair(jobScanner entity.JobScanner) {
	execSQL := fmt.Sprintf("UPDATE job_clair SET job_status='%s', job_data='%s' WHERE job_id='%s'",
		jobScanner.JobStatus, jobScanner.JobData, jobScanner.JobID)
	_ = mysql.InserData(execSQL)
}

func SetJobClairStatus(taskID, status string) {
	redis.SetMap(taskID, "clair", status)
}
