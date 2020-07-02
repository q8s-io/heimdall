package convert

import (
	"time"
	
	"github.com/q8s-io/heimdall/pkg/entity/model"
)

func AnchoreRequestInfo(jobScannerMsg *model.JobScannerMsg) *model.AnchoreRequestInfo {
	anchoreRequestInfo := new(model.AnchoreRequestInfo)
	anchoreRequestInfo.ImageName = jobScannerMsg.ImageName
	anchoreRequestInfo.ImageDigest = jobScannerMsg.ImageDigest
	anchoreRequestInfo.CreateTime = time.Now().Format("2006-01-02T15:04:05Z")
	return anchoreRequestInfo
}
