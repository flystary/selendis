package g

import (
	"log"
	"sync"

	"github.com/toolkits/file"
	"gopkg.in/yaml.v3"
)


type HttpConfig struct {
	Enabled bool 	`yaml:"enabled"`
	Listen  string 	`yaml:"listen"`
}


type GlobalConfig struct {
	Debug     bool        `yaml:"debug"`
	Hosts     string      `yaml:"hosts"`
	Database  string      `yaml:"database"`
	MaxConns  int         `yaml:"max-conns"`
	MaxIdle   int         `yaml:"max-idle"`
	Listen    string      `yaml:"listen"`
	Trustable []string    `yaml:"trustable"`
	Http      *HttpConfig `yaml:"http"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = yaml.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}