package scanner

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
)

func TrivyScan(imageName string, scanTime int) (model.TrivyScanResult, error) {
	scanResult := model.TrivyScanResult{}
	trivyConfig := model.Config.Trivy

	// 扫描结果的json文件存储位置，路径 + 文件名
	path := trivyConfig.TargetPath + trivyConfig.FileName
	containerName := trivyConfig.ContainerName
	volumeName := trivyConfig.VolumeName

	// The limits of container runtime is ScanTime minute
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*time.Duration(scanTime))
	// Initialize config of trivy container
	containerConfig := &container.Config{
		Image: trivyConfig.Image,
		Cmd:   append(trivyConfig.ContainerCMD, path, imageName),
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeVolume,
				Source:   volumeName,
				Target:   trivyConfig.TargetPath,
				ReadOnly: false,
			},
		},
	}

	// Create trivy container
	containerID, crtErr := docker.CreateContainerWithVolume(ctx, containerConfig, hostConfig, trivyConfig.Image, containerName, volumeName)
	if crtErr != nil {
		return scanResult, crtErr
	}

	// Run trivy container in id
	runErr := docker.RunContainerWithVolume(ctx, containerID, volumeName)
	if runErr != nil {
		return scanResult, runErr
	}

	// Get result
	scanResult, getErr := getTrivyResults(ctx, containerID, path)
	if getErr != nil {
		return scanResult, getErr
	}
	// Delete container trivy
	docker.DeleteContainerWithVolume(ctx, containerID, volumeName)

	// Remove volume
	_ = docker.RemoveVolumeByName(ctx, volumeName)

	// Close client
	defer func() {
		closeErr := docker.DClient.Close()
		if closeErr != nil {
			xray.ErrMini(closeErr)
		}
	}()
	// Close context
	defer cancel()

	return scanResult, nil
}

func getTrivyResults(ctx context.Context, containerID string, path string) (model.TrivyScanResult, error) {
	var data []*model.TrivyScanResult
	result := model.TrivyScanResult{}

	out, cpErr := docker.CopyFileFromContainer(ctx, containerID, path)
	if cpErr != nil {
		return result, cpErr
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, out)

	// 去除前后无用字符
	bytes := deletePreAndSufSpace(buf.String())
	if len(bytes) == 0 {
		xray.ErrMini(errors.New("the len of trivy result is 0"))
	}

	if unmarshalErr := json.Unmarshal(bytes, &data); unmarshalErr != nil {
		return result, xray.ErrMiniInfo(unmarshalErr)
	}

	result = *data[0]
	return result, nil
}

func deletePreAndSufSpace(str string) []byte {
	strList := []byte(str)
	lc, rc := 0, len(strList)-1
	for ; lc <= rc; lc++ {
		if strList[lc] == '[' {
			break
		}
	}
	for ; rc >= lc; rc-- {
		if strList[rc] == ']' {
			break
		}
	}
	return strList[lc : rc+1]
}
