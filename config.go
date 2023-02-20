package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigFile struct {
	Accesskey   string   `json:accesskey`
	Secretkey   string   `json:secretkey`
	Region      string   `json:region`
	Instance_id []string `json:instance_id`
}

func (c *ConfigFile) ConfigRead(filename string) {
	//var config ConfigFile
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(content, c)
	if err != nil {
		fmt.Println("json unmarshal error happened")
	}
}
