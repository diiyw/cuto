package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("protocol.json")
	var proto Protocol
	if err := json.Unmarshal(b, &proto); err != nil {
		panic(err)
	}
	var protocol = "../protocol/"
	_ = os.RemoveAll(protocol)
	for _, domain := range proto.Domains {
		dirname := protocol + strings.ToLower(domain.Domain)
		if err := os.MkdirAll(dirname, 0755); err != nil {
			panic(err)
		}
		_ = ioutil.WriteFile(dirname+"/type.go", domain.AllTypes(), 0755)
		_ = ioutil.WriteFile(dirname+"/method.go", domain.AllMethods(), 0755)
		_ = ioutil.WriteFile(dirname+"/event.go", domain.AllEvents(), 0755)
	}
}
