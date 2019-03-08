package chrome

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/diiyw/goc/protocol/dom"
	"github.com/diiyw/goc/protocol/page"
	"github.com/gorilla/websocket"
	"log"
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
		go func() {
			if err := cmd.Run(); err != nil {
				log.Println(err)
			}
		}()
		break
	}
	_, err := net.DialTimeout("tcp", "127.0.0.1:9222", time.Second*10)
	if err == nil {
		log.Println(err)
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
				var err = ch.Close()
				if err != nil {
					log.Println(err)
				}
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
	if err := os.RemoveAll(chrome.DataDir); err != nil {
		return err
	}
	return chrome.Process.Kill()
}

func newTab(tab *Tab) (*Tab, error) {
	conn, _, err := websocket.DefaultDialer.Dial(tab.WebSocketDebuggerUrl, nil)
	if err != nil {
		return nil, err
	}
	tab.Ipc.Conn = conn
	tab.Ipc.events = make(chan []byte, 1024)
	tab.Ipc.returns = make(chan []byte, 1024)
	tab.Ipc.errors = make(chan []byte, 1024)
	_ = tab.Send(dom.Enable, dom.EnableParams{})
	_ = tab.Send(page.Enable, page.EnableParams{})
	go func() {
		if err := tab.handle(); err != nil {
			log.Println(err)
			return
		}
	}()
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
			tab.Ipc.events = make(chan []byte, 1024)
			tab.Ipc.returns = make(chan []byte, 1024)
			tab.Ipc.errors = make(chan []byte, 1024)
			return tab, nil
		}
	}
	return nil, errors.New("tab not found")
}
