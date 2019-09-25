package page

import (

	"github.com/diiyw/goc/protocol/network"

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	
	DomContentEventFiredEvent = "Page.domContentEventFired"
	
	// Emitted only when `page.interceptFileChooser` is enabled.
	FileChooserOpenedEvent = "Page.fileChooserOpened"
	
	// Fired when frame has been attached to its parent.
	FrameAttachedEvent = "Page.frameAttached"
	
	// Fired when frame no longer has a scheduled navigation.
	FrameClearedScheduledNavigationEvent = "Page.frameClearedScheduledNavigation"
	
	// Fired when frame has been detached from its parent.
	FrameDetachedEvent = "Page.frameDetached"
	
	// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
	FrameNavigatedEvent = "Page.frameNavigated"
	
	
	FrameResizedEvent = "Page.frameResized"
	
	// Fired when a renderer-initiated navigation is requested.
	// Navigation may still be cancelled after the event is issued.
	FrameRequestedNavigationEvent = "Page.frameRequestedNavigation"
	
	// Fired when frame schedules a potential navigation.
	FrameScheduledNavigationEvent = "Page.frameScheduledNavigation"
	
	// Fired when frame has started loading.
	FrameStartedLoadingEvent = "Page.frameStartedLoading"
	
	// Fired when frame has stopped loading.
	FrameStoppedLoadingEvent = "Page.frameStoppedLoading"
	
	// Fired when page is about to start a download.
	DownloadWillBeginEvent = "Page.downloadWillBegin"
	
	// Fired when interstitial page was hidden
	InterstitialHiddenEvent = "Page.interstitialHidden"
	
	// Fired when interstitial page was shown
	InterstitialShownEvent = "Page.interstitialShown"
	
	// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been
	// closed.
	JavascriptDialogClosedEvent = "Page.javascriptDialogClosed"
	
	// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to
	// open.
	JavascriptDialogOpeningEvent = "Page.javascriptDialogOpening"
	
	// Fired for top level page lifecycle events such as navigation, load, paint, etc.
	LifecycleEventEvent = "Page.lifecycleEvent"
	
	
	LoadEventFiredEvent = "Page.loadEventFired"
	
	// Fired when same-document navigation happens, e.g. due to history API usage or anchor navigation.
	NavigatedWithinDocumentEvent = "Page.navigatedWithinDocument"
	
	// Compressed image data requested by the `startScreencast`.
	ScreencastFrameEvent = "Page.screencastFrame"
	
	// Fired when the page with currently enabled screencast was shown or hidden `.
	ScreencastVisibilityChangedEvent = "Page.screencastVisibilityChanged"
	
	// Fired when a new window is going to be opened, via window.open(), link click, form submission,
	// etc.
	WindowOpenEvent = "Page.windowOpen"
	
	// Issued for every compilation cache generated. Is only available
	// if Page.setGenerateCompilationCache is enabled.
	CompilationCacheProducedEvent = "Page.compilationCacheProduced"
	
)


type DomContentEventFiredParams struct {
	
	
	Timestamp	network.MonotonicTime	`json:"timestamp"`
	
}

// Emitted only when `page.interceptFileChooser` is enabled.
type FileChooserOpenedParams struct {
	
	
	Mode	string	`json:"mode"`
	
}

// Fired when frame has been attached to its parent.
type FrameAttachedParams struct {
	
	// Id of the frame that has been attached.
	FrameId	FrameId	`json:"frameId"`
	
	// Parent frame identifier.
	ParentFrameId	FrameId	`json:"parentFrameId"`
	
	// JavaScript stack trace of when frame was attached, only set if frame initiated from script.
	Stack	runtime.StackTrace	`json:"stack"`
	
}

// Fired when frame no longer has a scheduled navigation.
type FrameClearedScheduledNavigationParams struct {
	
	// Id of the frame that has cleared its scheduled navigation.
	FrameId	FrameId	`json:"frameId"`
	
}

// Fired when frame has been detached from its parent.
type FrameDetachedParams struct {
	
	// Id of the frame that has been detached.
	FrameId	FrameId	`json:"frameId"`
	
}

// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
type FrameNavigatedParams struct {
	
	// Frame object.
	Frame	Frame	`json:"frame"`
	
}


type FrameResizedParams struct {
	
}

// Fired when a renderer-initiated navigation is requested.
	// Navigation may still be cancelled after the event is issued.
type FrameRequestedNavigationParams struct {
	
	// Id of the frame that is being navigated.
	FrameId	FrameId	`json:"frameId"`
	
	// The reason for the navigation.
	Reason	ClientNavigationReason	`json:"reason"`
	
	// The destination URL for the requested navigation.
	Url	string	`json:"url"`
	
}

// Fired when frame schedules a potential navigation.
type FrameScheduledNavigationParams struct {
	
	// Id of the frame that has scheduled a navigation.
	FrameId	FrameId	`json:"frameId"`
	
	// Delay (in seconds) until the navigation is scheduled to begin. The navigation is not
	// guaranteed to start.
	Delay	float64	`json:"delay"`
	
	// The reason for the navigation.
	Reason	string	`json:"reason"`
	
	// The destination URL for the scheduled navigation.
	Url	string	`json:"url"`
	
}

// Fired when frame has started loading.
type FrameStartedLoadingParams struct {
	
	// Id of the frame that has started loading.
	FrameId	FrameId	`json:"frameId"`
	
}

// Fired when frame has stopped loading.
type FrameStoppedLoadingParams struct {
	
	// Id of the frame that has stopped loading.
	FrameId	FrameId	`json:"frameId"`
	
}

// Fired when page is about to start a download.
type DownloadWillBeginParams struct {
	
	// Id of the frame that caused download to begin.
	FrameId	FrameId	`json:"frameId"`
	
	// URL of the resource being downloaded.
	Url	string	`json:"url"`
	
}

// Fired when interstitial page was hidden
type InterstitialHiddenParams struct {
	
}

// Fired when interstitial page was shown
type InterstitialShownParams struct {
	
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been
	// closed.
type JavascriptDialogClosedParams struct {
	
	// Whether dialog was confirmed.
	Result	bool	`json:"result"`
	
	// User input in case of prompt.
	UserInput	string	`json:"userInput"`
	
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to
	// open.
type JavascriptDialogOpeningParams struct {
	
	// Frame url.
	Url	string	`json:"url"`
	
	// Message that will be displayed by the dialog.
	Message	string	`json:"message"`
	
	// Dialog type.
	Type	DialogType	`json:"type"`
	
	// True iff browser is capable showing or acting on the given dialog. When browser has no
	// dialog handler for given target, calling alert while Page domain is engaged will stall
	// the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
	HasBrowserHandler	bool	`json:"hasBrowserHandler"`
	
	// Default dialog prompt.
	DefaultPrompt	string	`json:"defaultPrompt"`
	
}

// Fired for top level page lifecycle events such as navigation, load, paint, etc.
type LifecycleEventParams struct {
	
	// Id of the frame.
	FrameId	FrameId	`json:"frameId"`
	
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId	network.LoaderId	`json:"loaderId"`
	
	
	Name	string	`json:"name"`
	
	
	Timestamp	network.MonotonicTime	`json:"timestamp"`
	
}


type LoadEventFiredParams struct {
	
	
	Timestamp	network.MonotonicTime	`json:"timestamp"`
	
}

// Fired when same-document navigation happens, e.g. due to history API usage or anchor navigation.
type NavigatedWithinDocumentParams struct {
	
	// Id of the frame.
	FrameId	FrameId	`json:"frameId"`
	
	// Frame's new url.
	Url	string	`json:"url"`
	
}

// Compressed image data requested by the `startScreencast`.
type ScreencastFrameParams struct {
	
	// Base64-encoded compressed image.
	Data	string	`json:"data"`
	
	// Screencast frame metadata.
	Metadata	ScreencastFrameMetadata	`json:"metadata"`
	
	// Frame number.
	SessionId	int	`json:"sessionId"`
	
}

// Fired when the page with currently enabled screencast was shown or hidden `.
type ScreencastVisibilityChangedParams struct {
	
	// True if the page is visible.
	Visible	bool	`json:"visible"`
	
}

// Fired when a new window is going to be opened, via window.open(), link click, form submission,
	// etc.
type WindowOpenParams struct {
	
	// The URL for the new window.
	Url	string	`json:"url"`
	
	// Window name.
	WindowName	string	`json:"windowName"`
	
	// An array of enabled window features.
	WindowFeatures	[]string	`json:"windowFeatures"`
	
	// Whether or not it was triggered by user gesture.
	UserGesture	bool	`json:"userGesture"`
	
}

// Issued for every compilation cache generated. Is only available
	// if Page.setGenerateCompilationCache is enabled.
type CompilationCacheProducedParams struct {
	
	
	Url	string	`json:"url"`
	
	// Base64-encoded data
	Data	string	`json:"data"`
	
}

