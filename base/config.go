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
	var file *os.File
	var data []byte
	config = &Config{}

	file, err = os.Open(path)
	if err != nil {
		goto Default
	}
	data, err = ioutil.ReadAll(file)
	if err != nil {
		goto Default
	}

	if err = yaml.Unmarshal(data, config); err != nil {
		// nothing to do
	}

Default:
	if config.Priority == 0 {
		config.Priority = DefaultPri
	}

	if config.DESKEY == "" {
		config.DESKEY = DefaultKey
	}

	return config, err
}
