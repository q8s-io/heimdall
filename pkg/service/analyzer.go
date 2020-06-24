package service

import (
	"fmt"

	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewImageAnalyzer(analyzerJobInfo models.AnalyzerJobInfo) error {
	execSQL := fmt.Sprintf("INSERT INTO `job_analyzer` (`task_id`, `job_id`, `job_status`, `image_name`, `image_digest`, `create_time`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
		analyzerJobInfo.TaskID, analyzerJobInfo.JobID, analyzerJobInfo.JobStatus, analyzerJobInfo.ImageName, analyzerJobInfo.ImageDigest, analyzerJobInfo.CreateTime)
	err := mysql.InserData(execSQL)
    return err
}
