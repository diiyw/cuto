package chrome

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/diiyw/gator/protocol/dom"
	"github.com/diiyw/gator/protocol/page"
	"github.com/diiyw/gator/protocol/runtime"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
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
	id int
	*websocket.Conn
	message chan string
}

type Tab struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	DevUrl string `json:"devtoolsFrontendUrl"`
	WsUrl  string `json:"webSocketDebuggerUrl"`
	Ipc    ipc    `json:"-"`
}

// 跳转地址
func (tab *Tab) Jump(url string) error {
	if err := tab.Send(page.Navigate, page.NavigateParams{Url: url}); err != nil {
		return err
	}
	err := tab.Recv(func(m []byte) bool {
		var result = page.NavigateReturns{}
		var returns = Return{Result: &result}
		if err := json.Unmarshal(m, &returns); err != nil {
			return true
		}
		if returns.Id != 0 && result.FrameId != "" {
			return true
		}
		return false
	})
	if err != nil {
		return err
	}
	return tab.Recv(func(m []byte) bool {
		var navigatedParams = &page.FrameNavigatedParams{}
		var navigated = Event{Params: navigatedParams}
		if err := json.Unmarshal(m, &navigated); err != nil {
			return true
		}
		if navigated.Method == page.FrameNavigatedEvent {
			_ = tab.Recv(func(m []byte) bool {
				var stoppedLoadingParams = &page.FrameStoppedLoadingParams{}
				var stoppedLoading = Event{Params: stoppedLoadingParams}
				if err := json.Unmarshal(m, &stoppedLoading); err != nil {
					return true
				}
				if stoppedLoading.Method == page.FrameStoppedLoadingEvent {
					if navigatedParams.Frame.Id == string(stoppedLoadingParams.FrameId) {
						return true
					}
				}
				return false
			})
			return true
		}
		return false
	})
}

// 等待页面加载完成
func (tab *Tab) Wait() (err error) {
	return tab.Recv(func(m []byte) bool {
		var params = page.LoadEventFiredParams{}
		var loadEvent = Event{Params: params}
		if err = json.Unmarshal(m, &loadEvent); err != nil {
			return true
		}
		if loadEvent.Method == page.LoadEventFiredEvent {
			return true
		}
		return false
	})
}

// 查询节点
func (tab *Tab) Query(selector string) ([]dom.NodeId, error) {
	var err error
	// 初始化整个节点
	if err := tab.Send(dom.GetDocument, dom.GetDocumentParams{}); err != nil {
		return nil, err
	}
	err = tab.Recv(func(m []byte) bool {
		var getDocument = &dom.GetDocumentReturns{}
		var returns = Return{Result: getDocument}
		if err = json.Unmarshal(m, &returns); err != nil {
			return true
		}
		if returns.Id == tab.Ipc.id {
			return true
		}
		return false
	})
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
	var searchResult = &dom.PerformSearchReturns{}
	err = tab.Recv(func(m []byte) bool {
		var returns = Return{Result: searchResult}
		if err := json.Unmarshal(m, &returns); err != nil {
			return true
		}
		if searchResult.SearchId != "" {
			return true
		}
		return false
	})
	// 获取节点结果
	var result = dom.GetSearchResultsParams{
		SearchId:  searchResult.SearchId,
		FromIndex: 0,
		ToIndex:   searchResult.ResultCount,
	}
	if err := tab.Send(dom.GetSearchResults, result); err != nil {
		return nil, err
	}
	var getResult = &dom.GetSearchResultsReturns{}
	err = tab.Recv(func(m []byte) bool {
		var returns = Return{Result: getResult}
		if err := json.Unmarshal(m, &returns); err != nil {
			return true
		}
		if returns.Id == tab.Ipc.id {
			return true
		}
		return false
	})
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
	var evalResult = &runtime.EvaluateReturns{}
	err = tab.Recv(func(m []byte) bool {
		var returns = Return{Result: evalResult}
		if err = json.Unmarshal(m, &returns); err != nil {
			return true
		}
		if returns.Id == tab.Ipc.id {
			return true
		}
		return false
	})
	if err != nil {
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
	return nil
}

// 页面截图
func (tab *Tab) Capture(filename string, quality int, viewport page.Viewport) error {
	var err error
	err = tab.Recv(func(m []byte) bool {
		var params = &page.DomContentEventFiredParams{}
		var contentEvent = Event{Params: params}
		if err = json.Unmarshal(m, &contentEvent); err != nil {
			return true
		}
		if contentEvent.Method == page.DomContentEventFiredEvent && params.Timestamp != 0 {
			return true
		}
		return false
	})
	var capture = page.CaptureScreenshotParams{
		Format:  strings.Trim(path.Ext(filename), "."),
		Clip:    viewport,
		Quality: quality,
	}
	if err := tab.Send(page.CaptureScreenshot, capture); err != nil {
		return err
	}
	err = tab.Recv(func(m []byte) bool {
		var result = &page.CaptureScreenshotReturns{}
		var returns = Return{Result: result}
		if err = json.Unmarshal(m, &returns); err != nil {
			return true
		}
		if result.Data != "" && returns.Id != 0 {
			b, _ := base64.StdEncoding.DecodeString(result.Data)
			err = ioutil.WriteFile(filename, b, 0777)
			return true
		}
		return false
	})
	return err
}

// 发起命令
func (tab *Tab) Send(method string, params interface{}) error {
	tab.Ipc.id ++
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

func (tab *Tab) Recv(f func(m []byte) bool) error {
	for {
		var e Error
		_, b, err := tab.Ipc.ReadMessage()
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &e); err != nil {
			return err
		}
		if e.Id != 0 && e.Error.Code != 0 {
			var code = "[CODE]:" + strconv.Itoa(e.Error.Code) + "\n[DATA]:"
			return errors.New(code + e.Error.Data + "\n[MESSAGE]:" + e.Error.Message)
		}
		if f(b) {
			break
		}
	}
	return nil
}
