package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/q8s-io/heimdall/pkg/domain/scancenter"
	"github.com/q8s-io/heimdall/pkg/infrastructure/ginext"
)

// CreateScanTask
func CreateScanTask(c *gin.Context) {
	body := make(map[string]string)
	ginext.ResolveJSON(c, &body)
	imageName := body["image_name"]

	//scan
	scanData := scancenter.CreateScan(imageName)
	log.Println(scanData)

	//return
	data := make(map[string]string)
	ginext.Sender(c, 0, "This is status.", data)
}

// GetScanTaskData
func GetScanTaskData(c *gin.Context) {
	taskID := c.Param("taskid")

	//scan
	scanData := scancenter.GetScanTaskData(taskID)
	log.Println(scanData)

	//return
	data := make(map[string]string)
	ginext.Sender(c, 0, "This is status.", data)
}
