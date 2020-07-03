package docker

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/q8s-io/heimdall/pkg/entity/model"
)

var imageFullName string

func ImageAnalyzer(imageName string) ([]string, []string) {
	dockerConfig := model.Config.Docker

	// docker client
	cli, cerr := client.NewClient(dockerConfig.Host, dockerConfig.Version, nil, nil)
	if cerr != nil {
		log.Println(cerr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	imageFullName = imageName

PULLIMAGE:
	// pull image
	err := PullImage(imageFullName, cli, ctx)
	if err != nil {
		switch err.Error() {
		case "repository name must be canonical":
			imageFullName = fmt.Sprintf("docker.io/library/%s", imageName)
			goto PULLIMAGE
		}
	}

	// inspect image
	imageID, digest, layers := InspectImage(imageFullName, cli, ctx)

	// remove image
	DeleteImage(imageID, cli, ctx)

	return digest, layers
}

func PullImage(imageName string, cli *client.Client, ctx context.Context) error {
	out, perr := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if perr != nil {
		log.Println(perr)
		return perr
	}
	if _, perr = io.Copy(os.Stdout, out); perr != nil {
		log.Println(perr)
		return perr
	}
	out.Close()
	return nil
}

func InspectImage(imageName string, cli *client.Client, ctx context.Context) (string, []string, []string) {
	imageInspect, _, ierr := cli.ImageInspectWithRaw(ctx, imageName)
	if ierr != nil {
		log.Println(ierr)
	}
	imageID := imageInspect.ID
	digest := imageInspect.RepoDigests
	layers := imageInspect.RootFS.Layers
	return imageID, digest, layers
}

func DeleteImage(imageID string, cli *client.Client, ctx context.Context) {
	var imageRemoveOptions types.ImageRemoveOptions
	imageRemoveOptions.Force = true
	imageRemoveOptions.PruneChildren = true
	_, rerr := cli.ImageRemove(ctx, imageID, imageRemoveOptions)
	if rerr != nil {
		log.Println(rerr)
	}
}
