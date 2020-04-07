package main

import (
	"github.com/diiyw/cuto"
	"github.com/diiyw/cuto/protocol/page"
	"log"
)

func main() {
	browser, err := cuto.Create()
	if err != nil {
		log.Fatal(err)
	}
	defer browser.Close()
	tab, err := browser.Open("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	if err := tab.Wait(); err != nil {
		log.Fatal(err)
	}
	if err := tab.Capture("baidu.png", 80, page.Viewport{X: 20.0, Y: 20.0, Width: 200.0, Height: 200.0}); err != nil {
		log.Fatal(err)
	}
}
