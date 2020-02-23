package chr

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/diiyw/chr/protocol/dom"
	"github.com/diiyw/chr/protocol/page"
	"github.com/diiyw/chr/protocol/runtime"
	"github.com/gorilla/websocket"
	"io/ioutil"
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

type ipc struct {
	*websocket.Conn

	id      int
	events  chan []byte
	returns chan []byte
	errors  chan []byte
}

type Tab struct {
	Id                   string `json:"id"`
	Url                  string `json:"url"`
	Type                 string `json:"type"`
	Title                string `json:"title"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
	Ipc                  ipc    `json:"-"`
}

// 等待页面加载完成
func (tab *Tab) Wait() error {
	if err := tab.HandleEvent(page.LoadEventFiredEvent, nil); err != nil {
		return err
	}
	return nil
}

// 跳转地址
func (tab *Tab) Jump(url string) error {
	if err := tab.Send(page.Navigate, page.NavigateParams{Url: url}); err != nil {
		return err
	}
	var result = page.NavigateReturns{}
	if err := tab.HandleResult(&result); err != nil {
		return err
	}
	var frameNavigatedParams = page.FrameNavigatedParams{}
	if err := tab.HandleEvent(page.FrameNavigatedEvent, &frameNavigatedParams); err != nil {
		return err
	}
	var frameStoppedLoadingParams = page.FrameStoppedLoadingParams{}
	if err := tab.HandleEvent(page.FrameStoppedLoadingEvent, &frameStoppedLoadingParams); err != nil {
		return err
	}
	if frameNavigatedParams.Frame.Id == string(frameStoppedLoadingParams.FrameId) {
		return nil
	}
	return errors.New("jump error")
}

// 查询节点
func (tab *Tab) Query(selector string) ([]dom.NodeId, error) {
	var err error
	// 初始化整个节点
	if err := tab.Send(dom.GetDocument, dom.GetDocumentParams{}); err != nil {
		return nil, err
	}
	var getDocument = dom.GetDocumentReturns{}
	if err := tab.HandleResult(&getDocument); err != nil {
		return nil, err
	}
	if err != nil {
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
	var searchResult = dom.PerformSearchReturns{}
	if err := tab.HandleResult(&searchResult); err != nil {
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
	var getResult = dom.GetSearchResultsReturns{}
	if err := tab.HandleResult(&getResult); err != nil {
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
	var evalResult = runtime.EvaluateReturns{}
	if err := tab.HandleResult(&evalResult); err != nil {
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
	err := tab.Ipc.Close()
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
	var captureScreenshotReturns = page.CaptureScreenshotReturns{}
	if err := tab.HandleResult(&captureScreenshotReturns); err != nil {
		return err
	}
	b, _ := base64.StdEncoding.DecodeString(captureScreenshotReturns.Data)
	return ioutil.WriteFile(filename, b, 0777)
}

// 发起命令
func (tab *Tab) Send(method string, params interface{}) error {
	tab.Ipc.id++
	var request = map[string]interface{}{
		"id":     tab.Ipc.id,
		"method": method,
		"params": params,
	}
	if err := tab.Ipc.WriteJSON(request); err != nil {
		return err
	}
	return nil
}

func (tab *Tab) handle() error {
	for {
		_, b, err := tab.Ipc.ReadMessage()
		if err != nil {
			return err
		}
		// Error
		if bytes.Contains(b, []byte(`"error"`)) {
			tab.Ipc.errors <- b
			continue
		}
		// Event
		if bytes.Contains(b, []byte(`"params"`)) && bytes.Contains(b, []byte(`"method"`)) {
			tab.Ipc.events <- b
			continue
		}
		// Result
		tab.Ipc.returns <- b
	}
}

func (tab *Tab) HandleResult(returns interface{}) error {
	timeout := time.After(time.Second * 15)
	for {
		select {
		case b := <-tab.Ipc.returns:
			var ret = Return{Result: returns}
			err := json.Unmarshal(b, &ret)
			if err == nil {
				if ret.Id == tab.Ipc.id {
					return nil
				}
			}
		case <-timeout:
			return errors.New("Handle result timeout ")
		}
	}
}

func (tab *Tab) HandleEvent(method string, params interface{}) error {
	timeout := time.After(time.Second * 15)
	for {
		select {
		case b := <-tab.Ipc.events:
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
