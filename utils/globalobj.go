package utils

import (
	"encoding/json"
	"io/ioutil"
)

type GlobalObject struct {
	Host       string
	TCPPort    string
	ServerName string
}

var GlobalObj *GlobalObject

func (g *GlobalObject) Reload() {
	content, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(content, g); err != nil {
		panic(err)
	}
}

func init() {
	GlobalObj = &GlobalObject{
		Host:       "127.0.0.1",
		TCPPort:    "9627",
		ServerName: "Test",
	}

	GlobalObj.Reload()
}
