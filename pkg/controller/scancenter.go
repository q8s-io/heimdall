package controller

import (
	"github.com/gin-gonic/gin"
	
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	"github.com/q8s-io/heimdall/pkg/provider/scancenter"
)

func GetImageVulnData(c *gin.Context) {
	imageRequestInfo := new(model.ImageRequestInfo)
	if err := c.ShouldBind(&imageRequestInfo); err != nil {
		return
	}

	// judge
	judgeData, err := scancenter.JudgeTask(imageRequestInfo)

	// return
	if err != nil {
		ginext.Sender(c, 1, err.Error(), "")
		return
	}
	ginext.Sender(c, 0, "", judgeData)
}

func UpdateImageAnalyzerData(c *gin.Context) {
	jobImageAnalyzerInfo := new(model.JobImageAnalyzerInfo)
	if err := c.ShouldBind(&jobImageAnalyzerInfo); err != nil {
		return
	}

	// update job
	scancenter.TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo)

	ginext.Sender(c, 0, "", "")
}

func UpdateAnchoreData(c *gin.Context) {
	jobScannerInfo := new(model.JobScannerInfo)
	if err := c.ShouldBind(&jobScannerInfo); err != nil {
		return
	}

	// update job
	scancenter.TaskImageScanRotaryAnchore(jobScannerInfo)

	ginext.Sender(c, 0, "", "")
}
