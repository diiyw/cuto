package chr

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTab(t *testing.T) {
	browser, err := Create()
	if err != nil {
		log.Println(err)
	}
	defer browser.Close()
	tab, err := browser.Open("https://www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	if err := tab.Wait(); err != nil {
		log.Println(err)
	}
	nodes, err := tab.Query("//*[@id=\"kw\"]")
	if err != nil {
		log.Fatal(err)
	}
	if nodes == nil {
		log.Fatal("nodes not found!")
	}
	fmt.Println(nodes)
	if err := tab.Input("#kw", "百度一下", ); err != nil {
		log.Println(err)
	}
	if err := tab.Click("#su"); err != nil {
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
}
