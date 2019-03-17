package config

import (
	"github.com/lotteryjs/configor"
)

// Configuration is stuff that can be configured externally per env variables or config file (config.yml).
type Configuration struct {
	Server struct {
		ListenAddr      string `default:""`
		Port            int    `default:"80"`
		ResponseHeaders map[string]string
	}
	Database struct {
		Dbname     string `default:""`
		Connection string `default:""`
	}
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{EnvironmentPrefix: "TenMinutesApi"}).Load(conf, "config.yml")
	if err != nil {
		panic(err)
	}
	return conf
}
