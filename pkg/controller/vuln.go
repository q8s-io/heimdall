package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/domain/judge"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

// GetImageVulnData
func GetImageVulnData(c *gin.Context) {
	body := make(map[string]string)
	ginext.ResolveJSON(c, &body)
	imageName := body["image_name"]
	imageDigest := body["image_digest"]

	//judge
	judgeData := judge.Judge(imageName, imageDigest)
	log.Println(judgeData)

	//return
	data := make(map[string]string)
	ginext.Sender(c, 0, "This is status.", data)
}
