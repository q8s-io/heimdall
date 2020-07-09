package docker

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func CreateContainerWithVolume(cli *client.Client, ctx context.Context, config *container.Config, hostConfig *container.HostConfig, volumeName string) (string, error) {

	// Create volume
	volErr := createVolume(cli, ctx, volumeName)
	if volErr != nil {
		return "", volErr
	}

	body, createErr := cli.ContainerCreate(ctx, config, hostConfig, nil, "trivy")
	if createErr != nil {
		log.Print("create container trivy failed !!!")
		RemoveVolumeByName(cli, ctx, volumeName)
		return "", createErr
	}
	log.Printf("create container %s successed !!!", body.ID)
	return body.ID, nil
}

func DeleteContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, volumeName string) {
	// 删除容器
	removeContainer(cli, ctx, containerID)
}

// 保证容器运行结束, 得到结果
func RunContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, volumeName string) error {
	// start container.
	err := startContainer(cli, ctx, containerID)
	if err != nil {
		// 删除容器
		_, _ = removeContainer(cli, ctx, containerID)
		// 删除卷
		RemoveVolumeByName(cli, ctx, volumeName)
		return err
	}

RETYR:
	info, _ := cli.ContainerInspect(ctx, containerID)
	log.Printf("%s status \t %v\n", containerID, info.State.Status)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

// 启动
func startContainer(cli *client.Client, ctx context.Context, containerID string) error {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("container %s start failed !!!", containerID)
		return err
	}
	log.Printf("container %s start succeed !!!", containerID)
	return nil
}

func removeContainer(cli *client.Client, ctx context.Context, containerID string) (string, error) {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Printf("remove container %s failed !!!", containerID)
		return "", err
	}
	log.Printf("remove container %s succeed !!!", containerID)
	return containerID, nil
}

// Create volume
func createVolume(cli *client.Client, ctx context.Context, volumeName string) error {
	volumeType := volumetypes.VolumesCreateBody{Name: volumeName}

	_, volumeErr := cli.VolumeCreate(ctx, volumeType)
	if volumeErr != nil {
		log.Printf("create volume %s failed !!!", volumeName)
		return volumeErr
	}
	log.Printf("create volume %s successed !!!", volumeName)
	return nil
}

// Remove volume
func RemoveVolumeByName(cli *client.Client, ctx context.Context, volumeName string) error {
	err := cli.VolumeRemove(ctx, volumeName, true)
	if err != nil {
		log.Printf("remove volume %s failed !!!", volumeName)
		log.Print(err)
		return err
	}
	log.Printf("remove volume %s succeed !!!", volumeName)
	return nil
}
