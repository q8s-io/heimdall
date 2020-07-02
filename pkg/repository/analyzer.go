package repository

import (
	"fmt"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
)

func NewJobImageAnalyzer(jobImageAnalyze entity.JobImageAnalyzer) {
	execSQL := fmt.Sprintf("INSERT INTO job_analyzer (task_id, job_id, job_status, image_name, image_digest, image_layers, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
		jobImageAnalyze.TaskID, jobImageAnalyze.JobID, jobImageAnalyze.JobStatus, jobImageAnalyze.ImageName, jobImageAnalyze.ImageDigest, jobImageAnalyze.ImageLayers, jobImageAnalyze.CreateTime, jobImageAnalyze.Active)
	_ = mysql.InserData(execSQL)
}

func UpdateJobImageAnalyzer(jobImageAnalyze entity.JobImageAnalyzer) {
	execSQL := fmt.Sprintf("UPDATE job_analyzer SET job_status='%s', image_digest='%s', image_layers='%s' WHERE job_id='%s'",
		jobImageAnalyze.JobStatus, jobImageAnalyze.ImageDigest, jobImageAnalyze.ImageLayers, jobImageAnalyze.JobID)
	_ = mysql.InserData(execSQL)
}

func SetJobImageAnalyzerStatus(taskID, status string) {
	redis.SetMap(taskID, "analyzer", status)
}
