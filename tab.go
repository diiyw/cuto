package cuto

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/diiyw/cuto/protocol/dom"
	"github.com/diiyw/cuto/protocol/page"
	"github.com/diiyw/cuto/protocol/runtime"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	math "math"
	"path"
	"strings"
	"time"
)

type (
	Return struct {
		Id     int
		Result interface{}
	}

	Error struct {
		Id    int
		Error struct {
			Code    int
			Data    string
			Message string
		}
	}

	Event struct {
		Method string
		Params interface{}
	}
)

type Channel struct {
	*websocket.Conn

	id      int
	events  chan []byte
	returns chan []byte
	errors  chan []byte
}

type Tab struct {
	Id                   string  `json:"id"`
	Url                  string  `json:"url"`
	Type                 string  `json:"type"`
	Title                string  `json:"title"`
	DevtoolsFrontendUrl  string  `json:"devtoolsFrontendUrl"`
	WebSocketDebuggerUrl string  `json:"webSocketDebuggerUrl"`

	Channel              Channel `json:"-"`

	debug bool `json:"-"`
}

func (tab *Tab) init(body io.Reader, debug bool) error {
	tab.debug = debug
	if err := json.NewDecoder(body).Decode(&tab); err != nil {
		return err
	}
	conn, _, err := websocket.DefaultDialer.Dial(tab.WebSocketDebuggerUrl, nil)
	if err != nil {
		return err
	}
	tab.Channel.Conn = conn
	tab.Channel.events = make(chan []byte, 1024)
	tab.Channel.returns = make(chan []byte, 1024)
	tab.Channel.errors = make(chan []byte, 1024)
	_ = tab.Send(dom.Enable, dom.EnableParams{})
	_ = tab.Send(page.Enable, page.EnableParams{})
	_ = tab.Send(runtime.Enable, nil)
	go func() {
		if err := tab.handle(); err != nil {
			log.Println("Error:", err)
			return
		}
	}()
	return nil
}

// 等待页面加载完成
func (tab *Tab) Wait() error {
	if err := tab.PollEvent(page.LoadEventFiredEvent, nil); err != nil {
		return err
	}
	return nil
}

// 跳转地址
func (tab *Tab) Jump(url string) error {
	if err := tab.Send(page.Navigate, page.NavigateParams{Url: url}); err != nil {
		return err
	}
	var result = page.NavigateResult{}
	if err := tab.GetResult(&result); err != nil {
		return err
	}
	var frameNavigatedParams = page.FrameNavigatedParams{}
	if err := tab.PollEvent(page.FrameNavigatedEvent, &frameNavigatedParams); err != nil {
		return err
	}
	var frameStoppedLoadingParams = page.FrameStoppedLoadingParams{}
	if err := tab.PollEvent(page.FrameStoppedLoadingEvent, &frameStoppedLoadingParams); err != nil {
		return err
	}
	if frameNavigatedParams.Frame.Id == frameStoppedLoadingParams.FrameId {
		return nil
	}
	return errors.New("jump error")
}

// 查询节点
func (tab *Tab) Query(selector string) ([]*dom.NodeId, error) {
	// 初始化整个节点
	if err := tab.Send(dom.GetDocument, dom.GetDocumentParams{}); err != nil {
		return nil, err
	}
	var getDocument = dom.GetDocumentResult{}
	if err := tab.GetResult(&getDocument); err != nil {
		return nil, err
	}
	// 开始搜素节点
	var searchParams = dom.PerformSearchParams{
		Query:                     selector,
		IncludeUserAgentShadowDOM: true,
	}
	if err := tab.Send(dom.PerformSearch, searchParams); err != nil {
		return nil, err
	}
	var searchResult = dom.PerformSearchResult{}
	if err := tab.GetResult(&searchResult); err != nil {
		return nil, err
	}
	// 获取节点结果
	var result = dom.GetSearchResultsParams{
		SearchId:  searchResult.SearchId,
		FromIndex: 0,
		ToIndex:   searchResult.ResultCount,
	}
	if err := tab.Send(dom.GetSearchResults, result); err != nil {
		return nil, err
	}
	var getResult = dom.GetSearchResultsResult{}
	if err := tab.GetResult(&getResult); err != nil {
		return nil, err
	}
	return getResult.NodeIds, nil
}

// 输入值
func (tab *Tab) Input(selector, v string) error {
	if _, err := tab.Js("document.querySelector('"+selector+"').value=\""+v+"\"", 1000); err != nil {
		return nil
	}
	return nil
}

// 获取文本信息
func (tab *Tab) Text(selector string) string {
	if obj, err := tab.Js("document.querySelector('"+selector+"').textContent", 1000); err != nil {
		return obj.Value.(string)
	}
	return ""
}

// 元素值
func (tab *Tab) Value(selector string) string {
	if obj, err := tab.Js("document.querySelector('"+selector+"').value", 1000); err != nil {
		return obj.Value.(string)
	}
	return ""
}

// 选择
func (tab *Tab) Check(selector string, checked bool) error {
	str := "false"
	if checked {
		str = "true"
	}
	obj, err := tab.Js("document.querySelector('"+selector+"').checked = "+str, 1000)
	if err != nil {
		return err
	}
	if obj.Type == "boolean" && obj.Value.(bool) == checked {
		return nil
	}
	return nil
}

// 下拉框
func (tab *Tab) Select(selector, v string) error {
	var expression = "var selector = document.querySelector('" + selector + "');" +
		";for(var i = 0; i <selector.length; i++){" +
		"if(selector[i].value == \"" + v + "\"){selector[i] = true;return;}}"
	_, err := tab.Js(expression, 1000)
	if err != nil {
		return err
	}
	return nil
}

// 元素点击
func (tab *Tab) Click(selector string) error {
	if _, err := tab.Js("document.querySelector('"+selector+"').click()", 1000); err != nil {
		return err
	}
	return nil
}

// 运行Javascript
func (tab *Tab) Js(js string, timeout runtime.TimeDelta) (object runtime.RemoteObject, err error) {
	var rParams = runtime.EvaluateParams{
		Expression:            js,
		IncludeCommandLineAPI: true,
		Timeout:               timeout,
	}
	if err := tab.Send(runtime.Evaluate, rParams); err != nil {
		return object, err
	}
	var evalResult = runtime.EvaluateResult{}
	if err := tab.GetResult(&evalResult); err != nil {
		return object, err
	}
	return evalResult.Result, nil
}

// 页面刷新
func (tab *Tab) Refresh() error {
	var reload = page.ReloadParams{}
	if err := tab.Send(page.Reload, reload); err != nil {
		return err
	}
	return tab.Wait()
}

// 关闭标签
func (tab *Tab) Close() error {
	if err := tab.Send(page.Close, page.CloseParams{}); err != nil {
		return err
	}
	err := tab.Channel.Close()
	if err != nil {
		return err
	}
	return nil
}

// 页面截图
func (tab *Tab) Capture(filename string, quality int, viewport page.Viewport) error {
	var capture = page.CaptureScreenshotParams{
		Format:  strings.Trim(path.Ext(filename), "."),
		Clip:    viewport,
		Quality: quality,
	}
	if err := tab.Send(page.CaptureScreenshot, capture); err != nil {
		return err
	}
	var captureScreenshotResult = page.CaptureScreenshotResult{}
	if err := tab.GetResult(&captureScreenshotResult); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, captureScreenshotResult.Data, 0777)
}

// 元素截图
func (tab *Tab) DOMCapture(filename string, quality int, selector string) error {
	nodes, err := tab.Query(selector)
	if err != nil || len(nodes) == 0 {
		return err
	}
	if err := tab.Send(dom.GetBoxModel, dom.GetBoxModelParams{
		NodeId: *nodes[len(nodes)-1],
	}); err != nil {
		return err
	}
	var box dom.GetBoxModelResult
	if err := tab.GetResult(&box); err != nil {
		return err
	}
	return tab.Capture(filename, quality, page.Viewport{
		// Round the dimensions, as otherwise we might
		// lose one pixel in either dimension.
		X:      math.Round(box.Model.Margin[0]),
		Y:      math.Round(box.Model.Margin[1]),
		Width:  math.Round(box.Model.Margin[4] - box.Model.Margin[0]),
		Height: math.Round(box.Model.Margin[5] - box.Model.Margin[1]),
		// This seems to be necessary? Seems to do the
		// right thing regardless of DPI.
		Scale: 1.0,
	})
}

// 获取整个页面的截图
func (tab *Tab) FullCapture(filename string, quality int) error {
	if err := tab.Send(page.GetLayoutMetrics, page.GetLayoutMetricsParams{}); err != nil {
		return err
	}
	var metric page.GetLayoutMetricsResult
	if err := tab.GetResult(&metric); err != nil {
		return err
	}
	width, height := math.Ceil(metric.ContentSize.Width), math.Ceil(metric.ContentSize.Height)
	return tab.Capture(filename, quality, page.Viewport{
		X:      0,
		Y:      0,
		Width:  width,
		Height: height,
		Scale:  1.0,
	})
}

// 发起命令
func (tab *Tab) Send(method string, params interface{}) error {
	tab.Channel.id++
	var request = map[string]interface{}{
		"id":     tab.Channel.id,
		"method": method,
		"params": params,
	}
	data, _ := json.Marshal(request)
	if tab.debug {
		log.Println("Send:", string(data))
	}
	if err := tab.Channel.WriteJSON(request); err != nil {
		return err
	}
	return nil
}

func (tab *Tab) handle() error {
	for {
		_, b, err := tab.Channel.ReadMessage()
		if err != nil {
			return err
		}
		// Error
		if bytes.Contains(b, []byte(`"error"`)) {
			if tab.debug {
				log.Println("Error:", string(b))
			}
			tab.Channel.errors <- b
			continue
		}
		// Event
		if bytes.Contains(b, []byte(`"params"`)) && bytes.Contains(b, []byte(`"method"`)) {
			tab.Channel.events <- b
			continue
		}
		// Result
		tab.Channel.returns <- b
	}
}

func (tab *Tab) GetResult(returns interface{}) error {
	timeout := time.After(time.Second * 15)
	for {
		select {
		case b := <-tab.Channel.returns:
			if tab.debug {
				log.Println("Result:", string(b))
			}
			var ret = Return{Result: returns}
			err := json.Unmarshal(b, &ret)
			if err == nil {
				if ret.Id == tab.Channel.id {
					return nil
				}
			}
		case <-timeout:
			return errors.New("Handle result timeout ")
		}
	}
}

func (tab *Tab) PollEvent(method string, params interface{}) error {
	timeout := time.After(time.Second * 15)
	for {
		select {
		case b := <-tab.Channel.events:
			if tab.debug {
				log.Println("Event:", string(b))
			}
			var event = Event{Params: params}
			err := json.Unmarshal(b, &event)
			if err == nil {
				if event.Method == method {
					return nil
				}
			}
		case <-timeout:
			return errors.New("Handle " + method + " timeout.")
		}
	}
}
