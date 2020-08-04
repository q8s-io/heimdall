package docker

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	volumeTypes "github.com/docker/docker/api/types/volume"
)

func CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, imageName, containerName string) (string, error) {
	// 先查看是否有镜像。
	imgErr := InspectImageExist(imageName, ctx)
	// 在配置文件中配置镜像全名
	imageFullName = imageName
	if imgErr != nil {
		pullErr := PullImage(imageFullName, ctx)
		if pullErr != nil {
			return "", pullErr
		}
	}

	body, err := DClient.ContainerCreate(ctx, config, hostConfig, nil, containerName)
	if err != nil {
		// 先删除之前的容器
		_, removeErr := RemoveContainer(ctx, containerName)
		if removeErr != nil {
			return "", removeErr
		}
		body, err = DClient.ContainerCreate(ctx, config, hostConfig, nil, containerName)
		if err != nil {
			return "", err
		}
	}
	return body.ID, nil
}

func CreateContainerWithVolume(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, imageName, containerName, volumeName string) (string, error) {
	// Create volume
	volErr := createVolume(ctx, volumeName)
	if volErr != nil {
		return "", volErr
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
		return err
	}
RETYR:
	info, inspectErr := DClient.ContainerInspect(ctx, containerID)
	// 防止获取失败，获取不到就返回。
	if inspectErr != nil {
		return inspectErr
	}
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
	info, inspectErr := DClient.ContainerInspect(ctx, containerID)
	// 防止获取失败，获取不到就返回。
	if inspectErr != nil {
		return inspectErr
	}
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
		return err
	}
	return nil
}

func RemoveContainer(ctx context.Context, containerID string) (string, error) {
	// 必须强制删除。
	err := DClient.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		return "", err
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
		return err
	}
	return nil
}

// Copy file from container by giving path
func CopyFileFromContainer(ctx context.Context, containerID, path string) (io.ReadCloser, error) {
	out, _, err := DClient.CopyFromContainer(ctx, containerID, path)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Get logs from container
func GetContainerLogs(ctx context.Context, containerID string) (io.ReadCloser, error) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := DClient.ContainerLogs(ctx, containerID, options)
	if err != nil {
		return nil, err
	}
	return out, nil
}
