package repository

import (
	"fmt"
	"log"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
)

func NewJobAnchore(jobScanner entity.JobScanner) {
	execSQL := fmt.Sprintf("INSERT INTO job_anchore (task_id, job_id, job_status, job_data, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
		jobScanner.TaskID, jobScanner.JobID, jobScanner.JobStatus, jobScanner.JobData, jobScanner.ImageName, jobScanner.ImageDigest, jobScanner.CreateTime, jobScanner.Active)
	_ = mysql.InserData(execSQL)
}

func GetJobAnchore(taskID string) *[]entity.JobScanner {
	execSQL := fmt.Sprintf("SELECT task_id, job_id, job_status, job_data FROM job_anchore WHERE task_id='%s'",
		taskID)
	jobAnchoreDataList := new([]entity.JobScanner)
	err := mysql.Client.Select(jobAnchoreDataList, execSQL)
	if err != nil {
		log.Println(err)
	}
	return jobAnchoreDataList
}

func UpdateJobAnchore(jobScanner entity.JobScanner) {
	execSQL := fmt.Sprintf("UPDATE job_anchore SET job_status='%s', job_data='%s' WHERE job_id='%s'",
		jobScanner.JobStatus, jobScanner.JobData, jobScanner.JobID)
	_ = mysql.InserData(execSQL)
}

func SetJobAnchoreStatus(taskID, status string) {
	redis.SetMap(taskID, "anchore", status)
}
