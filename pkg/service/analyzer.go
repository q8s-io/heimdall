package service

import (
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewImageAnalyzer(jobAnalyzerInfo models.JobAnalyzerInfo) error {
	execSQL := fmt.Sprintf("INSERT INTO `job_analyzer` (`task_id`, `job_id`, `job_status`, `image_name`, `image_digest`, `create_time`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
		jobAnalyzerInfo.TaskID, jobAnalyzerInfo.JobID, jobAnalyzerInfo.JobStatus, jobAnalyzerInfo.ImageName, jobAnalyzerInfo.ImageDigest, jobAnalyzerInfo.CreateTime)
	err := mysql.InserData(execSQL)
	return err
}

func SetImageAnalyzerStatus(taskID, status string) {
	redis.SetMap(taskID, "analyzer", status)
}

func SendImageAnalyzerMsg(jobAnalyzerMsg *models.JobAnalyzerMsg) {
	kafka.SyncProducerSendMsg("analyzer", jobAnalyzerMsg)
}

func ConsumerImageAnalyzerMsg(queue chan *sarama.ConsumerMessage) {
	go kafka.ConsumerMsg(queue)
}
