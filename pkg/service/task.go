package service

import (
	"fmt"

	"github.com/q8s-io/heimdall/pkg/infrastructure/mysql"
	"github.com/q8s-io/heimdall/pkg/models"
)

func NewImageScan(imageVulnInfo models.ImageVulnInfo) error {
	execSQL := fmt.Sprintf("INSERT INTO `image_vuln` (`task_id`, `task_status`, `image_name`, `image_digest`, `create_time`) VALUES ('%s', '%s', '%s', '%s', '%s')",
		imageVulnInfo.TaskID, imageVulnInfo.TaskStatus, imageVulnInfo.ImageName, imageVulnInfo.ImageDigest, imageVulnInfo.CreateTime)
	err := mysql.InserData(execSQL)
	return err
}
