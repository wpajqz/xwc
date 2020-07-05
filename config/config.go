package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config 配置文件数据结构
type Config struct {
	// Enviroment 环境变量配置字段
	Enviroment []map[string]string `yaml:"enviroment"`
	// Command 命令配置字段
	Command map[string]string `yaml:"command"`
}

// LoadConfigFile 加载配置文件
func (c *Config) LoadConfigFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, c)
}

// StoreConfigFile 存储配置文件
func (c *Config) StoreConfigFile(filename string) error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, out, os.ModePerm)
}

// IsExists 判断配置文件是否存在
func (c *Config) IsExists(filename string) bool {
	_, err := os.Lstat(filename)

	return !os.IsNotExist(err)
}
