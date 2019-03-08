package chrome

import (
	"log"
	"testing"
	"time"
)

func TestChrome(t *testing.T) {
	chrome, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1e9 * 5)
	if err := chrome.Close(); err != nil {
		log.Println(err)
	}
}
