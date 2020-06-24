package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/domain/judge"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
	"github.com/q8s-io/heimdall/pkg/models"
)

func GetImageVulnData(c *gin.Context) {
	imageInfoRequest := new(models.ImageInfoRequest)
	if err := c.ShouldBind(&imageInfoRequest); err != nil {
		return
	}

	//judge
	judgeData, err := judge.Judge(imageInfoRequest)
	if err != nil {
		//return
		ginext.Sender(c, 1, err.Error(), "")
		return
	}

	//return
	ginext.Sender(c, 0, "This is status.", judgeData)
}
