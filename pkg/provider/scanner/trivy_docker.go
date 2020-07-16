package scanner

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"

	"github.com/q8s-io/heimdall/pkg/entity/model"
)

func TrivyScan(imageName string) (model.TrivyScanResult, error) {
	scanResult := model.TrivyScanResult{}
	trivyConfig := model.Config.Trivy

	// 扫描结果的json文件存储位置，路径 + 文件名
	path := trivyConfig.TargetPath + trivyConfig.FileName
	containerName := trivyConfig.ContainerName
	volumeName := trivyConfig.VolumeName

	// Create a docker client from remote host
	cli, err := client.NewClient(trivyConfig.HostURL, trivyConfig.Version, nil, nil)
	if err != nil {
		log.Println(err)
		return scanResult, err
	}

	// The runtime of limits is 10 minute
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

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
	containerID, crtErr := docker.CreateContainerWithVolume(cli, ctx, containerConfig, hostConfig, containerName, volumeName)
	if crtErr != nil {
		return scanResult, crtErr
	}

	// Run trivy container in id
	runErr := docker.RunContainerWithVolume(cli, ctx, containerID, volumeName)
	if runErr != nil {
		return scanResult, runErr
	}

	// Get result
	scanResult, getErr := getTrivyResults(cli, ctx, containerID, path)
	if getErr != nil {
		return scanResult, getErr
	}
	// Delete container trivy
	docker.DeleteContainerWithVolume(cli, ctx, containerID, volumeName)

	// Remove volume
	_ = docker.RemoveVolumeByName(cli, ctx, volumeName)

	// Close client
	defer cli.Close()
	// Close context
	defer cancel()

	return scanResult, nil
}

func getTrivyResults(cli *client.Client, ctx context.Context, containerID string, path string) (model.TrivyScanResult, error) {
	var data []*model.TrivyScanResult
	result := model.TrivyScanResult{}

	out, cpErr := docker.CopyFileFromContainer(cli, ctx, containerID, path)
	if cpErr != nil {
		return result, cpErr
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, out)

	// 去除前后无用字符
	bytes := deletePreAndSufSpace(buf.String())
	if len(bytes) == 0 {
	}

	if unmarshalErr := json.Unmarshal(bytes, &data); unmarshalErr != nil {
		log.Printf("error deserializing JSON: %v", unmarshalErr)
		return result, unmarshalErr
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
