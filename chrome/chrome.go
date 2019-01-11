package chrome

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/diiyw/goc/protocol/dom"
	"github.com/diiyw/goc/protocol/page"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

type Chrome struct {
	Addr    string
	DataDir string
	Process *os.Process
}

// Create chrome client
func New(opts []string) (*Chrome, error) {
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	s := []rune(hex.EncodeToString(h.Sum(nil)))
	cmd := &exec.Cmd{}
	for _, filename := range defaultChrome {
		_, err := os.Stat(filename)
		if err != nil {
			continue
		}
		cmd = exec.Command(filename, append([]string{
			"--remote-debugging-port=9222",
			"--user-data-dir=" + defaultUserDataTmpDir + "/" + string(s[:8]),
		}, opts...)...)
		go cmd.Run()
		break
	}
	for {
		c, err := net.Dial("tcp", "127.0.0.1:9222")
		if err == nil {
			c.Close()
			break
		}
	}
	ch := &Chrome{
		"127.0.0.1:9222",
		defaultUserDataTmpDir + "/" + string(s[:8]),
		cmd.Process,
	}
	var c = make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, os.Interrupt, os.Kill)
		for {
			select {
			case <-c:
				ch.Close()
			}
		}
	}()
	return ch, nil
}

// Open new tab
func (chrome *Chrome) Open(url string) (*Tab, error) {
	r, err := http.Get("http://" + chrome.Addr + "/json/new?" + url)
	if err != nil {
		return nil, errors.New("Http request error:" + err.Error())
	}
	defer r.Body.Close()
	tab := &Tab{}
	if err := json.NewDecoder(r.Body).Decode(&tab); err != nil {
		return nil, errors.New("Open new window error:" + err.Error())
	}
	return newTab(tab)
}

// Close chrome
func (chrome *Chrome) Close() error {
	_ = os.RemoveAll(chrome.DataDir)
	return chrome.Process.Kill()
}

func newTab(tab *Tab) (*Tab, error) {
	conn, _, err := websocket.DefaultDialer.Dial(tab.WsUrl, nil)
	if err != nil {
		return nil, err
	}
	tab.Ipc.Conn = conn
	_ = tab.Send(dom.Enable, dom.EnableParams{})
	_ = tab.Send(page.Enable, page.EnableParams{})
	return tab, nil
}

// Catch tab
func (chrome *Chrome) Find(kw string) (*Tab, error) {
	r, err := http.Get("http://" + chrome.Addr + "/json")
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
			return tab, nil
		}
	}
	return nil, errors.New("tab not found")
}
