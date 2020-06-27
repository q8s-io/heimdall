package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/domain/scancenter"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	"github.com/q8s-io/heimdall/pkg/models"
)

func GetImageVulnData(c *gin.Context) {
	imageRequestInfo := new(models.ImageRequestInfo)
	if err := c.ShouldBind(&imageRequestInfo); err != nil {
		return
	}

	// judge
	judgeData, err := scancenter.Judge(imageRequestInfo)

	// return
	if err != nil {
		ginext.Sender(c, 1, err.Error(), "")
		return
	}
	ginext.Sender(c, 0, "", judgeData)
}
