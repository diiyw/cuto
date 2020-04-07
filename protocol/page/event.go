package page

// 
const DomContentEventFiredEvent = "Page.domContentEventFired"
type DomContentEventFiredParams struct {

	// 
	Timestamp 	interface{}}



// Emitted only when `page.interceptFileChooser` is enabled.
const FileChooserOpenedEvent = "Page.fileChooserOpened"
type FileChooserOpenedParams struct {

	// 
	Mode 	string}



// Fired when frame has been attached to its parent.
const FrameAttachedEvent = "Page.frameAttached"
type FrameAttachedParams struct {

	// Id of the frame that has been attached.
	FrameId 	FrameId
	// Parent frame identifier.
	ParentFrameId 	FrameId
	// JavaScript stack trace of when frame was attached, only set if frame initiated from script.
	Stack 	interface{}}



// Fired when frame no longer has a scheduled navigation.
const FrameClearedScheduledNavigationEvent = "Page.frameClearedScheduledNavigation"
type FrameClearedScheduledNavigationParams struct {

	// Id of the frame that has cleared its scheduled navigation.
	FrameId 	FrameId}



// Fired when frame has been detached from its parent.
const FrameDetachedEvent = "Page.frameDetached"
type FrameDetachedParams struct {

	// Id of the frame that has been detached.
	FrameId 	FrameId}



// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
const FrameNavigatedEvent = "Page.frameNavigated"
type FrameNavigatedParams struct {

	// Frame object.
	Frame 	Frame}



// 
const FrameResizedEvent = "Page.frameResized"
type FrameResizedParams struct {
}



// Fired when a renderer-initiated navigation is requested.
// Navigation may still be cancelled after the event is issued.
const FrameRequestedNavigationEvent = "Page.frameRequestedNavigation"
type FrameRequestedNavigationParams struct {

	// Id of the frame that is being navigated.
	FrameId 	FrameId
	// The reason for the navigation.
	Reason 	ClientNavigationReason
	// The destination URL for the requested navigation.
	Url 	string}



// Fired when frame schedules a potential navigation.
const FrameScheduledNavigationEvent = "Page.frameScheduledNavigation"
type FrameScheduledNavigationParams struct {

	// Id of the frame that has scheduled a navigation.
	FrameId 	FrameId
	// Delay (in seconds) until the navigation is scheduled to begin. The navigation is not
	// guaranteed to start.
	Delay 	float64
	// The reason for the navigation.
	Reason 	string
	// The destination URL for the scheduled navigation.
	Url 	string}



// Fired when frame has started loading.
const FrameStartedLoadingEvent = "Page.frameStartedLoading"
type FrameStartedLoadingParams struct {

	// Id of the frame that has started loading.
	FrameId 	FrameId}



// Fired when frame has stopped loading.
const FrameStoppedLoadingEvent = "Page.frameStoppedLoading"
type FrameStoppedLoadingParams struct {

	// Id of the frame that has stopped loading.
	FrameId 	FrameId}



// Fired when page is about to start a download.
const DownloadWillBeginEvent = "Page.downloadWillBegin"
type DownloadWillBeginParams struct {

	// Id of the frame that caused download to begin.
	FrameId 	FrameId
	// URL of the resource being downloaded.
	Url 	string}



// Fired when interstitial page was hidden
const InterstitialHiddenEvent = "Page.interstitialHidden"
type InterstitialHiddenParams struct {
}



// Fired when interstitial page was shown
const InterstitialShownEvent = "Page.interstitialShown"
type InterstitialShownParams struct {
}



// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been
// closed.
const JavascriptDialogClosedEvent = "Page.javascriptDialogClosed"
type JavascriptDialogClosedParams struct {

	// Whether dialog was confirmed.
	Result 	bool
	// User input in case of prompt.
	UserInput 	string}



// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to
// open.
const JavascriptDialogOpeningEvent = "Page.javascriptDialogOpening"
type JavascriptDialogOpeningParams struct {

	// Frame url.
	Url 	string
	// Message that will be displayed by the dialog.
	Message 	string
	// Dialog type.
	Type 	DialogType
	// True iff browser is capable showing or acting on the given dialog. When browser has no
	// dialog handler for given target, calling alert while Page domain is engaged will stall
	// the page execution. Execution can be resumed via calling Page.handleJavaScriptDialog.
	HasBrowserHandler 	bool
	// Default dialog prompt.
	DefaultPrompt 	string}



// Fired for top level page lifecycle events such as navigation, load, paint, etc.
const LifecycleEventEvent = "Page.lifecycleEvent"
type LifecycleEventParams struct {

	// Id of the frame.
	FrameId 	FrameId
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId 	interface{}
	// 
	Name 	string
	// 
	Timestamp 	interface{}}



// 
const LoadEventFiredEvent = "Page.loadEventFired"
type LoadEventFiredParams struct {

	// 
	Timestamp 	interface{}}



// Fired when same-document navigation happens, e.g. due to history API usage or anchor navigation.
const NavigatedWithinDocumentEvent = "Page.navigatedWithinDocument"
type NavigatedWithinDocumentParams struct {

	// Id of the frame.
	FrameId 	FrameId
	// Frame's new url.
	Url 	string}



// Compressed image data requested by the `startScreencast`.
const ScreencastFrameEvent = "Page.screencastFrame"
type ScreencastFrameParams struct {

	// Base64-encoded compressed image.
	Data 	[]byte
	// Screencast frame metadata.
	Metadata 	ScreencastFrameMetadata
	// Frame number.
	SessionId 	int}



// Fired when the page with currently enabled screencast was shown or hidden `.
const ScreencastVisibilityChangedEvent = "Page.screencastVisibilityChanged"
type ScreencastVisibilityChangedParams struct {

	// True if the page is visible.
	Visible 	bool}



// Fired when a new window is going to be opened, via window.open(), link click, form submission,
// etc.
const WindowOpenEvent = "Page.windowOpen"
type WindowOpenParams struct {

	// The URL for the new window.
	Url 	string
	// Window name.
	WindowName 	string
	// An array of enabled window features.
	WindowFeatures 	[]string
	// Whether or not it was triggered by user gesture.
	UserGesture 	bool}



// Issued for every compilation cache generated. Is only available
// if Page.setGenerateCompilationCache is enabled.
const CompilationCacheProducedEvent = "Page.compilationCacheProduced"
type CompilationCacheProducedParams struct {

	// 
	Url 	string
	// Base64-encoded data
	Data 	[]byte}

