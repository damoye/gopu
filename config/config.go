package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

// Config ...
type Config struct {
	HTTPAddress    string
	DatabaseString string
	RedisAddress   string
	RedisPrefix    string
}

var (
	// Conf ...
	Conf          Config
	configPath    = flag.String("config", "", "config file path")
	defaultConfig = Config{
		HTTPAddress:    ":8080",
		DatabaseString: "root:@/gopu",
		RedisAddress:   "127.0.0.1:6379",
		RedisPrefix:    "gopu:",
	}
)

func init() {
	flag.Parse()
	if *configPath == "" {
		Conf = defaultConfig
		return
	}
	b, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &Conf)
	if err != nil {
		panic(err)
	}
}
