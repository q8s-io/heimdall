package repository

import (
	"log"

	"github.com/q8s-io/heimdall/pkg/entity"
	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/infrastructure/redis"
)

func NewJobImageAnalyzer(jobImageAnalyzer entity.JobImageAnalyzer) {
	jobAnalyzer := entity.JobAnalyzer{}
	jobAnalyzer.JobImageAnalyzer = jobImageAnalyzer
	mysql.Client.Create(&jobAnalyzer)
}

func UpdateJobImageAnalyzer(jobImageAnalyzer entity.JobImageAnalyzer) {
	var jobAnalyzer entity.JobAnalyzer
	jobAnalyzer.JobImageAnalyzer = jobImageAnalyzer
	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
	rows, err := mysql.Client.Model(&entity.JobAnalyzer{}).Updates(jobAnalyzer).
		Scopes(mysql.QueryByJobID(jobImageAnalyzer.JobID)).Rows()
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()
}

func SetJobImageAnalyzerStatus(taskID, status string) {
	redis.SetMap(taskID, "analyzer", status)
}
