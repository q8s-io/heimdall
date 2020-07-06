package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
)

func TrivyScan(imageName string) model.TrivyScanResult {
	scanResult := model.TrivyScanResult{}
	trivyConfig := model.Config.Trivy

	containerConfig := &container.Config{
		Image: "aquasec/trivy",
		Cmd:   []string{"-f", "json", "-o", "/root/.cache/result.json", imageName},
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeVolume,
				Source:   "trivy",
				Target:   "/root/.cache/",
				ReadOnly: false,
			},
		},
	}
	vilumeType := volumetypes.VolumesCreateBody{}
	vilumeType.Name = "trivy_vol"

	// Create trivy container
	cli, ctx, containerID := docker.CreateContainerWithVolume(trivyConfig.HostURL, trivyConfig.HostURL, containerConfig, hostConfig, vilumeType)
	if containerID == "" {
		return scanResult
	}

	// Run container in id
	runErr := docker.RunContainerWithVolume(cli, ctx, containerID, vilumeType)
	if runErr != nil {
		log.Print("start container failed !!!", runErr)
		return scanResult
	}

	// Get result
	scanResult = getTrivyResults(cli, ctx, containerID)

	// Delete
	docker.DeleceContainerWithVolume(cli, ctx, containerID, vilumeType)

	return scanResult
}

func getTrivyResults(cli *client.Client, ctx context.Context, containerID string) model.TrivyScanResult {
	var data []*model.TrivyScanResult
	result := model.TrivyScanResult{}

	out, cps, err := cli.CopyFromContainer(ctx, containerID, "/root/.cache/result.json")
	fmt.Println(cps)
	if err != nil {
		return result
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, out)

	// 处理前后乱码问题
	str := deletePreAndSufSpace(buf.String())
	if len(str) == 0 {
		return result
	}

	rdr := strings.NewReader(str)
	if err := json.NewDecoder(rdr).Decode(&data); err != nil {
		log.Printf("error deserializing JSON: %v", err)
		return result
	}

	if data != nil && len(data[0].Vulnerabilities) != 0 {
		fmt.Println(data[0].Vulnerabilities[0].Severity)
	}
	result = *data[0]
	return result
}

func deletePreAndSufSpace(str string) string {
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
	return string(strList[lc : rc+1])
}
