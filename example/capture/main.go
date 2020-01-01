package main

import (
	"github.com/diiyw/chr/chrome"
	"github.com/diiyw/chr/protocol/page"
	"log"
)

func main() {
	bs, err := chrome.Create()
	if err != nil {
		log.Fatal(err)
	}
	defer bs.Close()
	tab, err := bs.Open("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	if err := tab.Wait(); err != nil {
		log.Fatal(err)
	}
	if err := tab.Capture("baidu.png", 80, page.Viewport{X: 20, Y: 20, Width: 200, Height: 200}); err != nil {
		log.Fatal(err)
	}
}
