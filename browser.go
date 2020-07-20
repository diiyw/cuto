package cuto

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Browser struct {
	// chromium
	binary string
	// remote debug address
	remoteAddr string
	// cache dir，remove all data when shutdown
	dataDir string
	// browser process
	process *os.Process
	// timeout for communicating
	timeout time.Duration
	// debug flag
	debug bool
	// headless
	commands []string
}

// Create chrome client
func Create(options ...Option) (*Browser, error) {
	c := new(Browser)

	for _, filename := range defaultBrowser {
		_, err := os.Stat(filename)
		if err == nil {
			c.binary = filename
			break
		}
	}
	c.remoteAddr = "127.0.0.1:9222"
	c.dataDir = defaultUserDataTmpDir
	c.timeout = 5 * time.Second
	c.commands = append(c.commands, []string{
		"--remote-debugging-port=" + strings.Split(c.remoteAddr, ":")[1],
		"--user-data-dir=" + c.dataDir,
	}...)

	for _, op := range options {
		op(c)
	}

	_, err := os.Stat(c.binary)
	if err != nil {
		return nil, errors.New("Browser not found ")
	}

	cmd := &exec.Cmd{}
	cmd = exec.Command(c.binary, c.commands...)
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("Start chrome with error: %s ", err)
	}

	go func() {
		c.process = cmd.Process
		_ = cmd.Wait()
	}()

	http.DefaultClient.Timeout = c.timeout
	resp, err := http.Get("http://127.0.0.1:9222")
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()

	var s = make(chan os.Signal)
	go func() {
		signal.Notify(s, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
		<-s
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}()
	return c, nil
}

// Open new tab
func (b *Browser) Open(url string) (*Tab, error) {
	r, err := http.Get("http://" + b.remoteAddr + "/json/new?" + url)
	if err != nil {
		return nil, errors.New("Http request error:" + err.Error())
	}
	defer r.Body.Close()
	tab := new(Tab)
	if err := tab.init(r.Body, b.debug); err != nil {
		return nil, err
	}
	return tab, nil
}

// Close chrome
func (b *Browser) Close() error {
	_ = os.RemoveAll(b.dataDir)
	return b.process.Kill()
}

// Catch tab
func (b *Browser) Find(kw string) (*Tab, error) {
	r, err := http.Get("http://" + b.remoteAddr + "/json")
	if err != nil {
		return nil, errors.New("Http request error:" + err.Error())
	}
	defer r.Body.Close()
	var tabs = make([]*Tab, 0, 16)
	if err = json.NewDecoder(r.Body).Decode(&tabs); err != nil {
		return nil, err
	}
	for _, tab := range tabs {
		if strings.Contains(tab.Url, kw) || strings.Contains(tab.Id, kw) || strings.Contains(tab.Title, kw) {
			tab.Channel.events = make(chan []byte, 1024)
			tab.Channel.returns = make(chan []byte, 1024)
			tab.Channel.errors = make(chan []byte, 1024)
			return tab, nil
		}
	}
	return nil, errors.New("tab not found")
}
