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

func TrivyScan(imageName string) model.TrivyScanResult {
	scanResult := model.TrivyScanResult{}
	trivyConfig := model.Config.Trivy

	// Create a docker client from remote host
	cli, err := client.NewClient(trivyConfig.HostURL, trivyConfig.Version, nil, nil)
	if err != nil {
		log.Println(err)
		return scanResult
	}

	// The runtime of limits is 10 minute
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	// Init config of trivy container
	volumeName := "trivy_vol" // volume name
	containerConfig := &container.Config{
		Image: "aquasec/trivy",
		Cmd:   []string{"-f", "json", "-o", "/root/.cache/result.json", imageName},
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:     mount.TypeVolume,
				Source:   volumeName,
				Target:   "/root/.cache/",
				ReadOnly: false,
			},
		},
	}

	// Create trivy container
	containerID, err := docker.CreateContainerWithVolume(cli, ctx, containerConfig, hostConfig, volumeName)
	if err != nil {
		return scanResult
	}

	// Run trivy container in id
	runErr := docker.RunContainerWithVolume(cli, ctx, containerID, volumeName)
	if runErr != nil {
		return scanResult
	}

	// Get result
	scanResult = getTrivyResults(cli, ctx, containerID)

	// Delete container trivy
	docker.DeleteContainerWithVolume(cli, ctx, containerID, volumeName)

	// Remove volume
	_ = docker.RemoveVolumeByName(cli, ctx, volumeName)

	// Close client
	defer cli.Close()
	// Close context
	defer cancel()

	return scanResult
}

func getTrivyResults(cli *client.Client, ctx context.Context, containerID string) model.TrivyScanResult {
	var data []*model.TrivyScanResult
	result := model.TrivyScanResult{}

	out, cps, err := cli.CopyFromContainer(ctx, containerID, "/root/.cache/result.json")
	log.Println(cps)
	if err != nil {
		log.Println("copy file from container failed !!!", err)
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
