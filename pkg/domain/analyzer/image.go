package analyzer

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/q8s-io/heimdall/pkg/models"
)

func ImageAnalyzer(imageName string) ([]string, []string) {
	dockerConfig := models.Config.Docker

	// docker client
	cli, cerr := client.NewClient(dockerConfig.Host, dockerConfig.Version, nil, nil)
	if cerr != nil {
		log.Println(cerr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	// pull image
	PullImage(imageName, cli, ctx)

	// inspect image
	imageID, digest, layers := InspectImage(imageName, cli, ctx)

	// remove image
	DeleteImage(imageID, cli, ctx)

	return digest, layers
}

func PullImage(imageName string, cli *client.Client, ctx context.Context) {
	out, perr := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if perr != nil {
		log.Println(perr)
	}
	if _, perr = io.Copy(os.Stdout, out); perr != nil {
		log.Println(perr)
	}
	out.Close()
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
