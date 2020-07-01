package repository

import (
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewJobImageAnalyzer(jobImageAnalyzerData entity.JobImageAnalyzerData) {
	execSQL := fmt.Sprintf("INSERT INTO job_analyzer (task_id, job_id, job_status, image_name, image_digest, image_layers, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
		jobImageAnalyzerData.TaskID, jobImageAnalyzerData.JobID, jobImageAnalyzerData.JobStatus, jobImageAnalyzerData.ImageName, jobImageAnalyzerData.ImageDigest, jobImageAnalyzerData.ImageLayers, jobImageAnalyzerData.CreateTime, jobImageAnalyzerData.Active)
	_ = mysql.InserData(execSQL)
}

func UpdateJobImageAnalyzer(jobImageAnalyzerData entity.JobImageAnalyzerData) {
	execSQL := fmt.Sprintf("UPDATE job_analyzer SET job_status='%s', image_digest='%s', image_layers='%s' WHERE job_id='%s'",
		jobImageAnalyzerData.JobStatus, jobImageAnalyzerData.ImageDigest, jobImageAnalyzerData.ImageLayers, jobImageAnalyzerData.JobID)
	_ = mysql.InserData(execSQL)
}

func SetJobImageAnalyzerStatus(taskID, status string) {
	redis.SetMap(taskID, "analyzer", status)
}

func SendJobImageAnalyzerMsg(jobImageAnalyzerMsg *entity.JobImageAnalyzerMsg) {
	kafka.SyncProducerSendMsg("analyzer", jobImageAnalyzerMsg)
}

func ConsumerJobImageAnalyzerMsg(queue chan *sarama.ConsumerMessage) {
	kafka.ConsumerMsg("analyzer", queue)
}
