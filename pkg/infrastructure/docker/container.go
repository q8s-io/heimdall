package docker

import (
	"context"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	volumeTypes "github.com/docker/docker/api/types/volume"
)

func CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, imageName, containerName string) (string, error) {
	// 先查看是否有镜像。
	imgErr := InspectImageExist(imageName, ctx)
	// 在配置文件中最好配置全名
	imageFullName = imageName
	if imgErr != nil {
		pullErr := PullImage(imageFullName, ctx)
		if pullErr != nil {
			return "", xray.ErrMiniInfo(pullErr)
		}
	}

	body, err := DClient.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		// 先删除之前的容器
		_, removeErr := RemoveContainer(ctx, containerName)
		if removeErr != nil {
			return "", xray.ErrMiniInfo(err)
		}
		body, err = DClient.ContainerCreate(ctx, config, hostConfig, nil, containerName)
		if err != nil {
			return "", xray.ErrMiniInfo(err)
		}
	}
	return body.ID, nil
}

func CreateContainerWithVolume(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, imageName, containerName, volumeName string) (string, error) {
	// Create volume
	volErr := createVolume(ctx, volumeName)
	if volErr != nil {
		return "", xray.ErrMiniInfo(volErr)
	}
	id, createErr := CreateContainer(ctx, config, hostConfig, imageName, containerName)
	if createErr != nil {
		_ = RemoveVolumeByName(ctx, volumeName)
		return "", createErr
	}
	return id, nil
}

func DeleteContainerWithVolume(ctx context.Context, containerID string, volumeName string) {
	// 删除容器
	_, _ = RemoveContainer(ctx, containerID)
}

func RunContainer(ctx context.Context, containerID string) error {
	// start container.
	err := StartContainer(ctx, containerID)
	if err != nil {
		// 删除容器
		_, _ = RemoveContainer(ctx, containerID)
		return xray.ErrMiniInfo(err)
	}

RETYR:
	info, _ := DClient.ContainerInspect(ctx, containerID)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

// 保证容器运行结束, 得到结果
func RunContainerWithVolume(ctx context.Context, containerID string, volumeName string) error {
	// start container.
	err := StartContainer(ctx, containerID)
	if err != nil {
		// 删除容器
		_, _ = RemoveContainer(ctx, containerID)
		// 删除卷
		_ = RemoveVolumeByName(ctx, volumeName)
		return err
	}

RETYR:
	info, _ := DClient.ContainerInspect(ctx, containerID)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

// 启动
func StartContainer(ctx context.Context, containerID string) error {
	err := DClient.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return xray.ErrMiniInfo(err)
	}
	return nil
}

func RemoveContainer(ctx context.Context, containerID string) (string, error) {
	err := DClient.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		return "", xray.ErrMiniInfo(err)
	}
	return containerID, nil
}

// Create volume
func createVolume(ctx context.Context, volumeName string) error {
	volumeType := volumeTypes.VolumesCreateBody{Name: volumeName}
	_, err := DClient.VolumeCreate(ctx, volumeType)
	if err != nil {
		return err
	}
	return nil
}

// Remove volume
func RemoveVolumeByName(ctx context.Context, volumeName string) error {
	err := DClient.VolumeRemove(ctx, volumeName, true)
	if err != nil {
		return xray.ErrMiniInfo(err)
	}
	return nil
}

// Copy file from container by giving path
func CopyFileFromContainer(ctx context.Context, containerID, path string) (io.ReadCloser, error) {
	out, _, err := DClient.CopyFromContainer(ctx, containerID, path)
	if err != nil {
		return nil, xray.ErrMiniInfo(err)
	}
	return out, nil
}

// Get logs from container
func GetContainerLogs(ctx context.Context, containerID string) (io.ReadCloser, error) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := DClient.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return nil, xray.ErrMiniInfo(err)
	}
	return out, nil
}
