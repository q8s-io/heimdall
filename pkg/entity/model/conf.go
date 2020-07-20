package model

var Config Runtime

type Runtime struct {
	MySQL      mysql
	Redis      redis
	Kafka      kafka
	ScanCenter scancenter
	Docker     docker
	Anchore    anchore
	Trivy      trivy
	Clair      clair
}

type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DB       string `toml:"db"`
	UserName string `toml:"username"`
	PassWord string `toml:"password"`
}

type redis struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	PassWord string `toml:"password"`
}

type kafka struct {
	BrokerList []string `toml:"broker"`
}

type scancenter struct {
	AnalyzerURL string `toml:"analyzer_url"`
	AnchoreURL  string `toml:"anchore_url"`
	TrivyURL    string `toml:"trivy_url"`
	ClairURL    string `toml:"clair_url"`
}

type docker struct {
	Host    string `toml:"docker_url"`
	Version string `toml:"version"`
}

type anchore struct {
	AnchoreURL string `toml:"anchore_url"`
	UserName   string `toml:"username"`
	PassWord   string `toml:"password"`
}

type trivy struct {
	HostURL       string `toml:"trivy_url"`
	Version       string `toml:"version"`
	VolumeName    string `toml:"volume_name"`
	ContainerName string `toml:"containerName"`
	Image         string `toml:"image"`
	// 容器挂载卷的路径
	TargetPath   string   `toml:"target_path"`
	FileName     string   `toml:"file_name"`
	ContainerCMD []string `toml:"container_cmd"`
}

type clair struct {
	HostURL       string `toml:"clair_url"`
	Version       string `toml:"version"`
	ClairADDR     string `toml:"clair_addr"`
	Image         string `toml:"image"`
	ContainerName string `toml:"containerName"`
}
