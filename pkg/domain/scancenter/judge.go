package scancenter

import (
	"github.com/q8s-io/heimdall/pkg/models"
)

func Judge(imageRequestInfo *models.ImageRequestInfo) (interface{}, error) {
	// get data by imageName & imageDigest

	// if data is empty, run scan center
	return CreateTaskImageScan(imageRequestInfo)

	// if status is running, return data

	// if imageDigest is empty, return data

	// if imageDigest is db.imageDigest, return data

	// if imageDigest not is db.imageDigest, run scan center
}
