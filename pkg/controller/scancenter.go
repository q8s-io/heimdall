package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/domain/scancenter"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	"github.com/q8s-io/heimdall/pkg/models"
)

func GetImageVulnData(c *gin.Context) {
	imageRequestInfo := new(entity.ImageRequestInfo)
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
	jobImageAnalyzerInfo := new(entity.JobImageAnalyzerInfo)
	if err := c.ShouldBind(&jobImageAnalyzerInfo); err != nil {
		return
	}

	// update job
	scancenter.TaskImageScanRotaryAnalyzer(jobImageAnalyzerInfo)

	ginext.Sender(c, 0, "", "")
}

func UpdateAnchoreData(c *gin.Context) {
	jobAnchoreInfo := new(entity.JobAnchoreInfo)
	if err := c.ShouldBind(&jobAnchoreInfo); err != nil {
		return
	}

	// update job
	scancenter.TaskImageScanRotaryAnchore(jobAnchoreInfo)

	ginext.Sender(c, 0, "", "")
}
