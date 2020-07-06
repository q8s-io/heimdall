package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	volumetypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/q8s-io/heimdall/pkg/entity/model"
	"io"
	"log"
	"strings"
	"time"
)

func TrivyScan(imageName string) model.TrivyScanResult {
	scanResult := model.TrivyScanResult{}
	trivyConfig := model.Config.Trivy

	// Create a client from host
	cli, err := client.NewClient(trivyConfig.HostURL, trivyConfig.Version, nil, nil)
	if err != nil {
		log.Println(err)
		return scanResult
	}

	// The runtime of limits 10 minute
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	// Create volume
	volume, volumeErr := cli.VolumeCreate(ctx, volumetypes.VolumesCreateBody{Name: "trivy_vol"})
	if volumeErr != nil {
		log.Print("create volume failed !!!", volumeErr)
		return scanResult
	}

	// Create trivy container
	id, containerErr := createTrivyContainer(imageName, volume.Name, cli, ctx)
	if containerErr != nil {
		log.Print("create trivy container failed !!!", containerErr)
		// 删除卷
		removeVolume(volume.Name, cli, ctx)
		return scanResult
	}

	// Run container in id
	runERR := triggerTrivy(id, cli, ctx)
	if runERR != nil {
		log.Print("start container failed !!!", runERR)
		// 删除卷
		removeVolume(volume.Name, cli, ctx)
		// 删除容器
		removeContainer(id, cli, ctx)
		return scanResult
	}

	// Get result
	scanResult = getTrivyResults(id, cli, ctx)

	// 删除卷
	removeVolume(volume.Name, cli, ctx)
	// 删除容器
	removeContainer(id, cli, ctx)

	// Close client
	defer cli.Close()
	// Close context
	defer cancel()

	return scanResult
}

// 创建trivy容器执行命令
func createTrivyContainer(imageName string, volumeName string, cli *client.Client, ctx context.Context) (string, error) {
	config := &container.Config{
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

	body, err := cli.ContainerCreate(ctx, config, hostConfig, nil, "trivy")
	if err != nil {
		return body.ID, nil
	}

	fmt.Printf("ID: %s\n", body.ID)
	return body.ID, err
}

// 启动
func startContainer(containerID string, cli *client.Client, ctx context.Context) error {
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err == nil {
		log.Print("容器", containerID, "启动成功")
	}
	return err
}

func removeContainer(containerID string, cli *client.Client, ctx context.Context) (string, error) {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err == nil {
		log.Printf("remove container %s succeed", containerID)
	} else {
		log.Printf("remove container %s failed", containerID)
	}
	return containerID, err
}

// Remove volume
func removeVolume(volName string, cli *client.Client, ctx context.Context) {
	err := cli.VolumeRemove(ctx, volName, true)
	if err == nil {
		log.Printf("remove volume %s 成succeed", volName)
	}
}

// 保证容器运行结束, 得到结果
func triggerTrivy(containerID string, cli *client.Client, ctx context.Context) error {
	// start trivy container and scan image in it.
	err := startContainer(containerID, cli, ctx)
	if err != nil {
		log.Print("start trivy container failed !!!")
		return err
	}
RETYR:
	info, _ := cli.ContainerInspect(ctx, containerID)
	fmt.Printf("############# %s ################## status \t %v\n", containerID, info.State.Status)
	if info.State.Status == "running" {
		time.Sleep(3 * time.Second)
		goto RETYR
	}
	return nil
}

func getTrivyResults(containerID string, cli *client.Client, ctx context.Context) model.TrivyScanResult {
	var data []*model.TrivyScanResult
	result := model.TrivyScanResult{}

	out, cps, err := cli.CopyFromContainer(ctx, containerID, "/root/.cache/result.json")
	fmt.Println(cps)
	if err != nil {
		return result
	}

	buf := new(strings.Builder)
	io.Copy(buf, out)

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
