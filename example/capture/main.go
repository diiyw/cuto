package main

import (
	"github.com/diiyw/cuto"
	"github.com/diiyw/cuto/protocol/page"
	"log"
)

func main() {
	browser, err := cuto.Create(cuto.Debug(),cuto.Headless())
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
	if err := tab.Capture("baidu.png", 80, page.Viewport{X: 200.0, Y: 200.0, Width: 100.0, Height: 100.0, Scale: 1.0}); err != nil {
		log.Fatal(err)
	}
	if err := tab.DOMCapture("baidu_dom.png", 80, "#lg"); err != nil {
		log.Fatal(err)
	}
	if err := tab.FullCapture("full_baidu.png", 80); err != nil {
		log.Fatal(err)
	}
}
