package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	CONFIG_ENV = "AUTH_SERVICE_CONFIG"
)

var (
	config         *ServiceConfig
	configFilepath string
)

type MySqlConnConfig struct {
	Host            string `json:"host"`
	Port            uint16 `json:"port"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Schema          string `json:"schema"`
	MaxOpenConn     int    `json:"max_open_conn"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	MaxConnLifeTime int    `json:"max_conn_life_time"`
}

type Database struct {
	MaxReconnectRetry  int              `json:"max_reconnect_retry"`
	Mysql              *MySqlConnConfig `json:"mysql"`
	ContextTimeoutInMs int64            `json:"context_timeout_in_ms"`
}

type LogConfig struct {
	Level string `json:"level"`
}

type HttpServerConfig struct {
	Host              string        `json:"host"`
	Port              string        `json:"port"`
	IdleTimeoutInSec  int           `json:"idle_timeout_in_seconds"`
	IdleTimeout       time.Duration `json:"-"`
	ReadTimeoutInSec  int           `json:"read_timeout_in_seconds"`
	ReadTimeout       time.Duration `json:"-"`
	WriteTimeoutInSec int           `json:"write_timeout_in_seconds"`
	WriteTimeout      time.Duration `json:"-"`
}

type ServiceConfig struct {
	Log        LogConfig        `json:"log"`
	HttpServer HttpServerConfig `json:"http_server"`
	MySql      Database         `json:"database"`
	Secret     string           `json:"secret"`
}

func LoadServiceConfig(ctx context.Context) (*ServiceConfig, error) {
	configFilepath = os.Getenv(CONFIG_ENV)
	if len(configFilepath) < 1 {
		return nil, fmt.Errorf("can't load config file '%s'", configFilepath)
	}

	cfg, err := loadConfigFile(configFilepath)
	if err != nil {
		return nil, err
	}
	config = cfg

	return cfg, nil
}

func loadConfigFile(f string) (c *ServiceConfig, err error) {
	var content []byte
	content, err = ioutil.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("can't load config file '%s': %v", f, err)
	}
	err = json.Unmarshal(content, &c)
	if err != nil {
		return nil, fmt.Errorf("can't parse config file '%s': %v", f, err)
	}

	c.HttpServer.IdleTimeout = time.Duration(c.HttpServer.IdleTimeoutInSec) * time.Second
	c.HttpServer.ReadTimeout = time.Duration(c.HttpServer.ReadTimeoutInSec) * time.Second
	c.HttpServer.WriteTimeout = time.Duration(c.HttpServer.WriteTimeoutInSec) * time.Second

	return c, nil
}
