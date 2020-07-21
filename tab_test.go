package cuto

import (
	"log"
	"testing"
	"time"
)

func TestTab(t *testing.T) {
	browser, err := NewBrowser(
		Debug())
	if err != nil {
		log.Println(err)
	}
	defer browser.Close()
	// 打开百度首页
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
	// 搜索
	if err := tab.Input("#kw", "百度一下"); err != nil {
		log.Println(err)
	}
	// 点击搜索
	if err := tab.Click("#su"); err != nil {
		log.Println(err)
	}
	time.Sleep(5 * time.Second)
}

func TestTabJump(t *testing.T) {
	browser, err := NewBrowser(
		Debug())
	if err != nil {
		log.Println(err)
	}
	defer browser.Close()
	// 打开百度首页
	tab, err := browser.Open("https://www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	if err := tab.Wait(); err != nil {
		log.Println(err)
	}
	if err := tab.Jump("https://www.qq.com"); err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
}
