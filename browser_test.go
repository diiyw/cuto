package cuto

import (
	"log"
	"testing"
	"time"
)

func Test_browser(t *testing.T) {
	browser, err := NewBrowser(nil)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1e9 * 5)
	if err := browser.Close(); err != nil {
		log.Println(err)
	}
}
