package main

import (
	"github.com/diiyw/goc/chrome"
	"github.com/diiyw/goc/protocol/page"
	"log"
)

func main() {
	bs, err := chrome.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer bs.Close()
	tab, err := bs.Open("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	if err := tab.Capture("baidu.png", 100, page.Viewport{}); err != nil {
		log.Fatal(err)
	}
}
