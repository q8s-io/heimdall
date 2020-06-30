package service

import (
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/q8s-io/heimdall/pkg/infrastructure/kafka"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewJobAnchore(jobAnchoreData models.JobAnchoreData) {
	execSQL := fmt.Sprintf("INSERT INTO job_anchore (task_id, job_id, job_status, job_data, image_name, image_digest, create_time, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
		jobAnchoreData.TaskID, jobAnchoreData.JobID, jobAnchoreData.JobStatus, jobAnchoreData.JobData, jobAnchoreData.ImageName, jobAnchoreData.ImageDigest, jobAnchoreData.CreateTime, jobAnchoreData.Active)
	_ = mysql.InserData(execSQL)
}

func UpdateJobAnchore(jobAnchoreData models.JobAnchoreData) {
	execSQL := fmt.Sprintf("UPDATE job_anchore SET job_status='%s', job_data='%s' WHERE job_id='%s'",
		jobAnchoreData.JobStatus, jobAnchoreData.JobData, jobAnchoreData.JobID)
	_ = mysql.InserData(execSQL)
}

func SetJobAnchoreStatus(taskID, status string) {
	redis.SetMap(taskID, "anchore", status)
}

func SendJobAnchoreMsg(jobAnchoreMsg *models.JobAnchoreMsg) {
	kafka.SyncProducerSendMsg("anchore", jobAnchoreMsg)
}

func ConsumerJobAnchoreMsg(queue chan *sarama.ConsumerMessage) {
	kafka.ConsumerMsg("anchore", queue)
}
