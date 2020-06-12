package models

var Config Runtime

type Runtime struct {
	MySQL mysql
}

type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DB       string `toml:"db"`
	UserName string `toml:"username"`
	PassWord string `toml:"password"`
}
