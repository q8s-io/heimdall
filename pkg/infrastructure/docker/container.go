package docker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func CreateContainerWithVolume(host, version string, config *container.Config, hostConfig *container.HostConfig, vilumeType volumetypes.VolumesCreateBody) (*client.Client, context.Context, string) {
	// Create a client from host
	cli, cerr := client.NewClient(host, version, nil, nil)
	if cerr != nil {
		log.Println(cerr)
	}
	// Close client
	defer cli.Close()

	// The runtime of limits 10 minute
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	// Close context
	defer cancel()

	// Create volume
	volume, volumeErr := cli.VolumeCreate(ctx, vilumeType)
	if volumeErr != nil {
		log.Println(volumeErr)
	}

	body, createErr := cli.ContainerCreate(ctx, config, hostConfig, nil, "trivy")
	if createErr != nil {
		log.Println(volumeErr)
		// 删除卷
		removeVolume(volume.Name, cli, ctx)
		return cli, ctx, ""
	}

	return cli, ctx, body.ID
}

func DeleceContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, vilumeType volumetypes.VolumesCreateBody) {
	// 删除卷
	removeVolume(vilumeType.Name, cli, ctx)
	// 删除容器
	_, _ = removeContainer(containerID, cli, ctx)
}

// 保证容器运行结束, 得到结果
func RunContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, vilumeType volumetypes.VolumesCreateBody) error {
	// start container.
	err := startContainer(containerID, cli, ctx)
	if err != nil {
		log.Println("start trivy container failed")
		// 删除卷
		removeVolume(vilumeType.Name, cli, ctx)
		// 删除容器
		_, _ = removeContainer(containerID, cli, ctx)
		return err
	}

RETYR:
	info, _ := cli.ContainerInspect(ctx, containerID)
	fmt.Printf("%s status \t %v\n", containerID, info.State.Status)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

// 启动
func startContainer(containerID string, cli *client.Client, ctx context.Context) error {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	log.Printf("container %s start succeed", containerID)
	return nil
}

func removeContainer(containerID string, cli *client.Client, ctx context.Context) (string, error) {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Printf("remove container %s failed", containerID)
		return "", err
	}
	log.Printf("remove container %s succeed", containerID)
	return containerID, nil
}

// Remove volume
func removeVolume(volName string, cli *client.Client, ctx context.Context) {
	err := cli.VolumeRemove(ctx, volName, true)
	if err != nil {
		log.Println(err)
	}
	log.Printf("remove volume %s succeed", volName)
}
