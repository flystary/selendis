package g

import (
	"log"
	"sync"


	"github.com/toolkits/file"
	"gopkg.in/yaml.v3"
)


type HttpConfig struct {
	Enabled bool   `yaml:"enabled"`
	Listen  string `yaml:"listen"`
}

type RpcConfig struct {
	Enabled bool   `yaml:"enabled"`
	Listen  string `yaml:"listen"`
}

type TransferConfig struct {
	Enabled     bool              `yaml:"enabled"`
	Batch       int               `yaml:"batch"`
	ConnTimeout int               `yaml:"connTimeout"`
	CallTimeout int               `yaml:"callTimeout"`
	MaxConns    int               `yaml:"maxConns"`
	MaxIdle     int               `yaml:"maxIdle"`
	MaxRetry    int               `yaml:"retry"`
	Cluster     map[string]string `yaml:"cluster"`
}

type InfluxdbConfig struct {
	Enabled   bool   `yaml:"enabled"`
	Batch     int    `yaml:"batch"`
	MaxRetry  int    `yaml:"retry"`
	MaxConns  int    `yaml:"maxConns"`
	Timeout   int    `yaml:"timeout"`
	Address   string `yaml:"address"`
	Database  string `yaml:"db"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Precision string `yaml:"precision"`
}

type GlobalConfig struct {
	Debug    bool            `yaml:"debug"`
	MinStep  int             `yaml:"minStep"` //最小周期,单位sec
	Http     *HttpConfig     `yaml:"http"`
	Rpc      *RpcConfig      `yaml:"rpc"`
	Transfer *TransferConfig `yaml:"transfer"`
	Influxdb *InfluxdbConfig `yaml:"influxdb"`
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
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.yaml cfg.yaml`")
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

	log.Println("g.ParseConfig ok, file ", cfg)
}