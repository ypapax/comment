package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type Config struct {
	DbType string       `yaml:"db_type"`
	Db     DbConnection `yaml:"db"`
	Api    ApiConf      `yaml:"api"`
}

type DbConnection struct {
	User string `yaml:"user"`
	Addr string `yaml:"addr"`
	Name string `yaml:"name"`
}

type ApiConf struct {
	Bind string `yaml:"bind"`
}

func Parse(reader io.Reader) (*Config, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}