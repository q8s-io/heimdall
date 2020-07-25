package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"

	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

var imageFullName string

func ImageAnalyzer(imageName string) ([]string, []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	imageFullName = imageName

PULLIMAGE:
	// pull image
	err := PullImage(imageFullName, ctx)
	if err != nil {
		xray.ErrMini(err)
		switch err.Error() {
		case "repository name must be canonical":
			imageFullName = fmt.Sprintf("docker.io/library/%s", imageName)
			goto PULLIMAGE
		default:
			return nil, nil
		}

	}

	// inspect image
	imageID, digest, layers := InspectImage(imageFullName, ctx)

	// remove image
	DeleteImage(imageID, ctx)

	return digest, layers
}

func PullImage(imageName string, ctx context.Context) error {
	out, perr := DClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if perr != nil {
		return perr
	}
	if _, perr = io.Copy(os.Stdout, out); perr != nil {
		return perr
	}
	_ = out.Close()
	return nil
}

func InspectImage(imageName string, ctx context.Context) (string, []string, []string) {
	imageInspect, _, ierr := DClient.ImageInspectWithRaw(ctx, imageName)
	if ierr != nil {
		xray.ErrMini(ierr)
	}
	imageID := imageInspect.ID
	digest := imageInspect.RepoDigests
	layers := imageInspect.RootFS.Layers
	return imageID, digest, layers
}

func InspectImageExist(imageName string, ctx context.Context) error {
	_, _, ierr := DClient.ImageInspectWithRaw(ctx, imageName)
	if ierr != nil {
		return ierr
	}
	return nil
}

func DeleteImage(imageID string, ctx context.Context) {
	var imageRemoveOptions types.ImageRemoveOptions
	imageRemoveOptions.Force = true
	imageRemoveOptions.PruneChildren = true
	_, rerr := DClient.ImageRemove(ctx, imageID, imageRemoveOptions)
	if rerr != nil {
		xray.ErrMini(rerr)
	}
}
