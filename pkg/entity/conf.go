package entity

var Config Runtime

type Runtime struct {
	MySQL      mysql
	Redis      redis
	Kafka      kafka
	Docker     docker
	ScanCenter scancenter
	Anchore    anchore
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

type docker struct {
	Host    string `toml:"host"`
	Version string `toml:"version"`
}

type scancenter struct {
	AnalyzerURL string `toml:"analyzer_url"`
	AnchoreURL  string `toml:"anchore_url"`
}

type anchore struct {
	AnchoreURL string `toml:"anchore_url"`
	UserName   string `toml:"username"`
	PassWord   string `toml:"password"`
}
