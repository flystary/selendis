package g

import (
	"log"
	"os"
	"sync"


	"github.com/toolkits/file"
	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	Debug    		bool    		`yaml:"debug"`
	Hostname 		string			`yaml:"hostname"`
	IP				string			`yaml:"ip"`
	Plugin			*PluginConfig	`yaml:"plugin"`
	Transfer 		*TransferConfig	`yaml:"transfer"`
	Http			*HttpConfig		`yaml:"http"`
}

type TransferConfig struct {
	Enabled  bool     `yaml:"enabled"`
	Addrs    []string `yaml:"addrs"`
	Interval int      `yaml:"interval"`
	Timeout  int      `yaml:"timeout"`
}

type PluginConfig struct {
	Enabled bool   `yaml:"enabled"`
	Dir     string `yaml:"dir"`
	Git     string `yaml:"git"`
	LogDir  string `yaml:"logs"`
}

type HttpConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Listen   string `yaml:"listen"`
	Backdoor bool   `yaml:"backdoor"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}


func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}
	if os.Getenv("FALCON_ENDPOINT") != "" {
		hostname = os.Getenv("FALCON_ENDPOINT")
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
	}
	return hostname, err
}

func IP() string {
	ip := Config().IP
	if ip != "" {
		return ip
	}
	if len(LocalIp) > 0 {
		ip = LocalIp
	}
	return ip
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

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
