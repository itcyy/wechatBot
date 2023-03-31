package config

import (
	"encoding/json"
	"fmt"
	io "io/ioutil"
	"strings"
	"sync"
)

//定义配置文件解析后的结构
type UserInfo struct {
	Key  string `json:key`
	Text string `json:text`
}

var UserIn UserInfo
var file_locker sync.Mutex //config file locker

func Load(filename string) ([]UserInfo, bool) {
	var conf []UserInfo
	file_locker.Lock()
	data, err := io.ReadFile(filename) //read config file
	file_locker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return conf, false
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, false
	}
	return conf, true
}

func FData(str string) (string, bool) {
	conf, ok := Load("./data.json")
	if !ok {
		fmt.Println("load config failed")
		return "", false
	}
	text := ""
	for i := 0; i < len(conf); i++ {
		if strings.Contains(str, conf[i].Key) {
			text = conf[i].Text
			return text, true
		}
	}

	return "", false
}
