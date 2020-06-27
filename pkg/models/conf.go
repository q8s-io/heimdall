package models

var Config Runtime

type Runtime struct {
	MySQL   mysql
	Redis   redis
	Kafka   kafka
	Anchore anchore
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

type anchore struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	UserName string `toml:"username"`
	PassWord string `toml:"password"`
}
