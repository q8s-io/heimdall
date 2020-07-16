package scanner

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
)

func ClairScan(imageName string) (model.ClairScanResult, error) {
	scanResult := model.ClairScanResult{}
	clairConfig := model.Config.Clair
	containerName := clairConfig.ContainerName

	// Create a docker client from remote host
	cli, err := client.NewClient(clairConfig.HostURL, clairConfig.Version, nil, nil)
	if err != nil {
		log.Println(err)
		return scanResult, err
	}

	// The limits of container runtime is 10 minute
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	containerConfig := &container.Config{
		Image: clairConfig.Image,
		Cmd:   []string{imageName},
		Env:   []string{clairConfig.ClairADDR, model.ClairJsonType},
	}
	hostConfig := &container.HostConfig{
		NetworkMode: container.NetworkMode(model.ClairNetworkMode),
	}

	// Create klar container
	containerID, createErr := docker.CreateContainer(cli, ctx, containerConfig, hostConfig, containerName)
	if createErr != nil {
		return scanResult, createErr
	}

	// Start klar container
	runErr := docker.RunContainer(cli, ctx, containerID)
	if runErr != nil {
		return scanResult, runErr
	}

	scanResult, scanErr := getClairResults(cli, ctx, containerID)
	if scanErr != nil {
		return scanResult, scanErr
	}
	// Remove container klar
	_, _ = docker.RemoveContainer(cli, ctx, containerID)

	// Close client
	defer cli.Close()
	// Close context
	defer cancel()

	return scanResult, nil
}

func getClairResults(cli *client.Client, ctx context.Context, containerID string) (model.ClairScanResult, error) {
	result := model.ClairScanResult{}

	out, logErr := docker.GetContainerLogs(cli, ctx, containerID)
	if logErr != nil {
		return result, logErr
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, out)

	bytes := deletePreChar(buf.String())
	if len(bytes) == 0 {
		log.Print("日志长度为0 !!!")
	}
	if unmarshalErr := json.Unmarshal(bytes, &result); unmarshalErr != nil {
		log.Printf("error deserializing JSON: %s", unmarshalErr)
		return result, unmarshalErr
	}
	return result, nil
}

func deletePreChar(str string) []byte {
	chList := []byte(str)
	lc, rc := 0, len(chList)-1
	// 处理首尾无效字符，直接舍弃
	for ; lc <= rc; lc++ {
		if chList[lc] == '{' {
			break
		}
	}
	for ; rc >= lc; rc-- {
		if chList[rc] == '}' {
			break
		}
	}
	// 处理字符流中的无效字符，替换为‘ ’
	for i := lc; i <= rc; i++ {
		ch := chList[i]
		switch {
		case ch > '~':
			chList[i] = ' '
		case ch == '\r':
		case ch == '\n':
		case ch == '\t':
		case ch < ' ':
			chList[i] = ' '
		}
	}
	return chList[lc : rc+1]
}
