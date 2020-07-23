package scanner

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/q8s-io/heimdall/pkg/infrastructure/xray"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"github.com/q8s-io/heimdall/pkg/infrastructure/docker"
)

func ClairScan(imageName string) (model.ClairScanResult, error) {
	scanResult := model.ClairScanResult{}
	clairConfig := model.Config.Clair
	containerName := clairConfig.ContainerName

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
	containerID, createErr := docker.CreateContainer(ctx, containerConfig, hostConfig, clairConfig.Image, containerName)
	if createErr != nil {
		return scanResult, createErr
	}

	// Start klar container
	runErr := docker.RunContainer(ctx, containerID)
	if runErr != nil {
		return scanResult, runErr
	}

	scanResult, scanErr := getClairResults(ctx, containerID)
	if scanErr != nil {
		return scanResult, scanErr
	}
	// Remove container klar
	_, _ = docker.RemoveContainer(ctx, containerID)

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

func getClairResults(ctx context.Context, containerID string) (model.ClairScanResult, error) {
	result := model.ClairScanResult{}
	out, logErr := docker.GetContainerLogs(ctx, containerID)
	if logErr != nil {
		return result, logErr
	}

	buf := new(strings.Builder)
	_, _ = io.Copy(buf, out)

	bytes := deletePreChar(buf.String())
	if len(bytes) == 0 {
		xray.ErrMini(errors.New("the len of clair result is 0"))
	}
	if unmarshalErr := json.Unmarshal(bytes, &result); unmarshalErr != nil {
		return result, xray.ErrMiniInfo(unmarshalErr)
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
