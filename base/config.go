package base

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	// 优先级用户应该可以自定义，插件本身也应该提供默认值
	Priority uint16					`yaml:"priority"`
	// 应该让用户自定义DES对称加密秘钥
	DESKEY string					`yaml:"des_key"`
}

func ParseConfig(path string) (config *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	config = &Config{}

	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	if config.Priority == 0 {
		config.Priority = DefaultPri
	}

	if config.DESKEY == "" {
		config.DESKEY = DefaultKey
	}

	return config, nil
}
