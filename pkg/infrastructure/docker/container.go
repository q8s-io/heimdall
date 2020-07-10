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

func CreateContainerWithVolume(cli *client.Client, ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName, volumeName string) (string, error) {

	// Create volume
	volErr := createVolume(cli, ctx, volumeName)
	if volErr != nil {
		return "", volErr
	}

	body, createErr := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if createErr != nil {
		_ = RemoveVolumeByName(cli, ctx, volumeName)
		return "", createErr
	}
	log.Printf("create container %s successed !!!", body.ID)
	return body.ID, nil
}

func CreateContainer(cli *client.Client, ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (string, error) {
	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		log.Printf("create container %s failed !!! %s", containerName, err)
		return "", err
	}
	log.Printf("create container %s successed !!!", body.ID)
	return body.ID, nil
}

func DeleteContainerWithVolume(cli *client.Client, ctx context.Context, containerID string, volumeName string) {
	// 删除容器
	_, _ = RemoveContainer(cli, ctx, containerID)
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
	log.Printf("%s status \t %v\n", containerID, info.State.Status)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
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
	log.Printf("%s status \t %v\n", containerID, info.State.Status)
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
	log.Printf("container %s start succeed !!!", containerID)
	return nil
}

func RemoveContainer(cli *client.Client, ctx context.Context, containerID string) (string, error) {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Printf("remove container %s failed !!! %s", containerID, err)
		return "", err
	}
	log.Printf("remove container %s succeed !!!", containerID)
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
	log.Printf("create volume %s successed !!!", volumeName)
	return nil
}

// Remove volume
func RemoveVolumeByName(cli *client.Client, ctx context.Context, volumeName string) error {
	err := cli.VolumeRemove(ctx, volumeName, true)
	if err != nil {
		log.Printf("remove volume %s failed !!! %s", volumeName, err)
		return err
	}
	log.Printf("remove volume %s succeed !!!", volumeName)
	return nil
}

// Copy file from container by giving path
func CopyFileFromContainer(cli *client.Client, ctx context.Context, containerID, path string) (io.ReadCloser, error) {
	out, cps, err := cli.CopyFromContainer(ctx, containerID, path)
	log.Print(cps)
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
	log.Printf("get logs from %s successed !!!", containerID)
	return out, nil
}
