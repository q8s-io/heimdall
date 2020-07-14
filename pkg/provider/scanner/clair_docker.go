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

func ClairScan(imageName string) model.ClairScanResult {
	scanResult := model.ClairScanResult{}
	trivyConfig := model.Config.Trivy

	// Create a docker client from remote host
	cli, err := client.NewClient(trivyConfig.HostURL, trivyConfig.Version, nil, nil)
	if err != nil {
		log.Println(err)
		return scanResult
	}

	// The limits of container runtime is 10 minute
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	containerConfig := &container.Config{
		Image: "klar",
		Cmd:   []string{imageName},
		Env:   []string{"CLAIR_ADDR=127.0.0.1", "JSON_OUTPUT=true"},
	}
	hostConfig := &container.HostConfig{
		NetworkMode: "host",
	}

	// Create klar container
	containerID, createErr := docker.CreateContainer(cli, ctx, containerConfig, hostConfig, "")
	if createErr != nil {
		log.Printf("create container %s failed", containerID)
		return scanResult
	}

	// Start klar container
	runErr := docker.RunContainer(cli, ctx, containerID)
	if runErr != nil {
		return scanResult
	}

	scanResult = getClairResults(cli, ctx, containerID)

	// Remove container klar
	_, _ = docker.RemoveContainer(cli, ctx, containerID)

	// Close client
	defer cli.Close()
	// Close context
	defer cancel()

	return scanResult
}

func getClairResults(cli *client.Client, ctx context.Context, containerID string) model.ClairScanResult {
	result := model.ClairScanResult{}

	out, logErr := docker.GetContainerLogs(cli, ctx, containerID)
	if logErr != nil {
		return result
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, out)

	bytes := deletePreChar(buf.String())
	if len(bytes) == 0 {
		log.Print("日志长度为0 !!!")
		return result
	}
	if unmarshalErr := json.Unmarshal(bytes, &result); unmarshalErr != nil {
		log.Printf("error deserializing JSON: %s", unmarshalErr)
		return result
	}
	return result
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
		// case ch == '\r':
		// case ch == '\n':
		// case ch == '\t':
		case ch < ' ':
			chList[i] = ' '
		}
	}
	return chList[lc : rc+1]
}
