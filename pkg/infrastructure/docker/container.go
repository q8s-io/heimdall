package docker

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func CreateContainer(cli *client.Client, ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (string, error) {
	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		log.Printf("container %s create error !!! %s", containerName, err)
		// 先删除之前的容器
		_, removeErr := RemoveContainer(cli, ctx, containerName)
		if removeErr != nil {
			return "", removeErr
		}

		body, err = cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
		if err != nil {
			log.Printf("删除之前的容器后再创建%s容器还有问题！！！", containerName)
			return "", err
		}
	}
	return body.ID, nil
}

func CreateContainerWithVolume(cli *client.Client, ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName, volumeName string) (string, error) {
	// Create volume
	volErr := createVolume(cli, ctx, volumeName)
	if volErr != nil {
		return "", volErr
	}
	id, createErr := CreateContainer(cli, ctx, config, hostConfig, containerName)
	if createErr != nil {
		_ = RemoveVolumeByName(cli, ctx, volumeName)
		return "", createErr
	}
	return id, nil
}

func DeleteContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, volumeName string) {
	// 删除容器
	_, _ = RemoveContainer(cli, ctx, containerID)
}

func RunContainer(cli *client.Client, ctx context.Context, containerID string) error {
	// start container.
	err := StartContainer(cli, ctx, containerID)
	if err != nil {
		// 删除容器
		_, _ = RemoveContainer(cli, ctx, containerID)
		return err
	}

RETYR:
	info, _ := cli.ContainerInspect(ctx, containerID)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

// 保证容器运行结束, 得到结果
func RunContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, volumeName string) error {
	// start container.
	err := StartContainer(cli, ctx, containerID)
	if err != nil {
		// 删除容器
		_, _ = RemoveContainer(cli, ctx, containerID)
		// 删除卷
		_ = RemoveVolumeByName(cli, ctx, volumeName)
		return err
	}

RETYR:
	info, _ := cli.ContainerInspect(ctx, containerID)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

// 启动
func StartContainer(cli *client.Client, ctx context.Context, containerID string) error {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("container %s start failed !!! %s", containerID, err)
		return err
	}
	return nil
}

func RemoveContainer(cli *client.Client, ctx context.Context, containerID string) (string, error) {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Printf("remove container %s failed !!! %s", containerID, err)
		return "", err
	}
	return containerID, nil
}

// Create volume
func createVolume(cli *client.Client, ctx context.Context, volumeName string) error {
	volumeType := volumetypes.VolumesCreateBody{Name: volumeName}

	_, err := cli.VolumeCreate(ctx, volumeType)
	if err != nil {
		log.Printf("create volume %s failed !!! %s", volumeName, err)
		return err
	}
	return nil
}

// Remove volume
func RemoveVolumeByName(cli *client.Client, ctx context.Context, volumeName string) error {
	err := cli.VolumeRemove(ctx, volumeName, true)
	if err != nil {
		log.Printf("remove volume %s failed !!! %s", volumeName, err)
		return err
	}
	return nil
}

// Copy file from container by giving path
func CopyFileFromContainer(cli *client.Client, ctx context.Context, containerID, path string) (io.ReadCloser, error) {
	out, _, err := cli.CopyFromContainer(ctx, containerID, path)
	// log.Print(cps)
	if err != nil {
		log.Printf("copy file from container %s failed !!! %s", containerID, err)
		return nil, err
	}
	return out, nil
}

// Get logs from container
func GetContainerLogs(cli *client.Client, ctx context.Context, containerID string) (io.ReadCloser, error) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		log.Printf("get logs from %s failed !!! %s", containerID, err)
		return nil, err
	}
	return out, nil
}
