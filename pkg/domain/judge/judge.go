package judge

import (
	"log"

	"github.com/q8s-io/heimdall/pkg/domain/scancenter"
	"github.com/q8s-io/heimdall/pkg/models"
)

// Judge
func Judge(imageInfoRequest *models.ImageInfoRequest) interface{} {
	log.Println(imageInfoRequest)
	//get data by imageName & imageDigest

	//if data is empty, run scan center
	return scancenter.PreperScanenter(imageInfoRequest)

	//if status is running, return data

	//if imageDigest is empty, return data

	//if imageDigest is db.imageDigest, return data

	//if imageDigest not is db.imageDigest, run scan center
}
