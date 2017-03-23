package main

import (
	"encoding/json"
	"io/ioutil"
)

type gopuConfig struct {
	HTTPAddress    string
	DatabaseString string
	RedisAddress   string
	RedisPrefix    string
	QueueExpire    int
	QueueMaxLength int
}

var config gopuConfig

func init() {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// file := path.Join(path.Dir(ex), "config.json")
	file := "/Users/mo/go/src/github.com/damoye/gopu/config.json"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
}
