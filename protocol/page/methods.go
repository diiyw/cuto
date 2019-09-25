package page

import (

	"github.com/diiyw/goc/protocol/debugger"

	"github.com/diiyw/goc/protocol/dom"

	"github.com/diiyw/goc/protocol/emulation"

	"github.com/diiyw/goc/protocol/io"

	"github.com/diiyw/goc/protocol/network"

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	// Deprecated, please use addScriptToEvaluateOnNewDocument instead.
	AddScriptToEvaluateOnLoad = "Page.addScriptToEvaluateOnLoad"
	
	// Evaluates given script in every frame upon creation (before loading frame's scripts).
	AddScriptToEvaluateOnNewDocument = "Page.addScriptToEvaluateOnNewDocument"
	
	// Brings page to front (activates tab).
	BringToFront = "Page.bringToFront"
	
	// Capture page screenshot.
	CaptureScreenshot = "Page.captureScreenshot"
	
	// Returns a snapshot of the page as a string. For MHTML format, the serialization includes
	// iframes, shadow DOM, external resources, and element-inline styles.
	CaptureSnapshot = "Page.captureSnapshot"
	
	// Clears the overriden device metrics.
	ClearDeviceMetricsOverride = "Page.clearDeviceMetricsOverride"
	
	// Clears the overridden Device Orientation.
	ClearDeviceOrientationOverride = "Page.clearDeviceOrientationOverride"
	
	// Clears the overriden Geolocation Position and Error.
	ClearGeolocationOverride = "Page.clearGeolocationOverride"
	
	// Creates an isolated world for the given frame.
	CreateIsolatedWorld = "Page.createIsolatedWorld"
	
	// Deletes browser cookie with given name, domain and path.
	DeleteCookie = "Page.deleteCookie"
	
	// Disables page domain notifications.
	Disable = "Page.disable"
	
	// Enables page domain notifications.
	Enable = "Page.enable"
	
	
	GetAppManifest = "Page.getAppManifest"
	
	
	GetInstallabilityErrors = "Page.getInstallabilityErrors"
	
	// Returns all browser cookies. Depending on the backend support, will return detailed cookie
	// information in the `cookies` field.
	GetCookies = "Page.getCookies"
	
	// Returns present frame tree structure.
	GetFrameTree = "Page.getFrameTree"
	
	// Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
	GetLayoutMetrics = "Page.getLayoutMetrics"
	
	// Returns navigation history for the current page.
	GetNavigationHistory = "Page.getNavigationHistory"
	
	// Resets navigation history for the current page.
	ResetNavigationHistory = "Page.resetNavigationHistory"
	
	// Returns content of the given resource.
	GetResourceContent = "Page.getResourceContent"
	
	// Returns present frame / resource tree structure.
	GetResourceTree = "Page.getResourceTree"
	
	// Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
	HandleJavaScriptDialog = "Page.handleJavaScriptDialog"
	
	// Navigates current page to the given URL.
	Navigate = "Page.navigate"
	
	// Navigates current page to the given history entry.
	NavigateToHistoryEntry = "Page.navigateToHistoryEntry"
	
	// Print page as PDF.
	PrintToPDF = "Page.printToPDF"
	
	// Reloads given page optionally ignoring the cache.
	Reload = "Page.reload"
	
	// Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
	RemoveScriptToEvaluateOnLoad = "Page.removeScriptToEvaluateOnLoad"
	
	// Removes given script from the list.
	RemoveScriptToEvaluateOnNewDocument = "Page.removeScriptToEvaluateOnNewDocument"
	
	// Acknowledges that a screencast frame has been received by the frontend.
	ScreencastFrameAck = "Page.screencastFrameAck"
	
	// Searches for given string in resource content.
	SearchInResource = "Page.searchInResource"
	
	// Enable Chrome's experimental ad filter on all sites.
	SetAdBlockingEnabled = "Page.setAdBlockingEnabled"
	
	// Enable page Content Security Policy by-passing.
	SetBypassCSP = "Page.setBypassCSP"
	
	// Overrides the values of device screen dimensions (window.screen.width, window.screen.height,
	// window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
	// query results).
	SetDeviceMetricsOverride = "Page.setDeviceMetricsOverride"
	
	// Overrides the Device Orientation.
	SetDeviceOrientationOverride = "Page.setDeviceOrientationOverride"
	
	// Set generic font families.
	SetFontFamilies = "Page.setFontFamilies"
	
	// Set default font sizes.
	SetFontSizes = "Page.setFontSizes"
	
	// Sets given markup as the document's HTML.
	SetDocumentContent = "Page.setDocumentContent"
	
	// Set the behavior when downloading a file.
	SetDownloadBehavior = "Page.setDownloadBehavior"
	
	// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position
	// unavailable.
	SetGeolocationOverride = "Page.setGeolocationOverride"
	
	// Controls whether page will emit lifecycle events.
	SetLifecycleEventsEnabled = "Page.setLifecycleEventsEnabled"
	
	// Toggles mouse event-based touch event emulation.
	SetTouchEmulationEnabled = "Page.setTouchEmulationEnabled"
	
	// Starts sending each frame using the `screencastFrame` event.
	StartScreencast = "Page.startScreencast"
	
	// Force the page stop all navigations and pending resource fetches.
	StopLoading = "Page.stopLoading"
	
	// Crashes renderer on the IO thread, generates minidumps.
	Crash = "Page.crash"
	
	// Tries to close page, running its beforeunload hooks, if any.
	Close = "Page.close"
	
	// Tries to update the web lifecycle state of the page.
	// It will transition the page to the given state according to:
	// https://github.com/WICG/web-lifecycle/
	SetWebLifecycleState = "Page.setWebLifecycleState"
	
	// Stops sending each frame in the `screencastFrame`.
	StopScreencast = "Page.stopScreencast"
	
	// Forces compilation cache to be generated for every subresource script.
	SetProduceCompilationCache = "Page.setProduceCompilationCache"
	
	// Seeds compilation cache for given url. Compilation cache does not survive
	// cross-process navigation.
	AddCompilationCache = "Page.addCompilationCache"
	
	// Clears seeded compilation cache.
	ClearCompilationCache = "Page.clearCompilationCache"
	
	// Generates a report for testing.
	GenerateTestReport = "Page.generateTestReport"
	
	// Pauses page execution. Can be resumed using generic Runtime.runIfWaitingForDebugger.
	WaitForDebugger = "Page.waitForDebugger"
	
	// Intercept file chooser requests and transfer control to protocol clients.
	// When file chooser interception is enabled, native file chooser dialog is not shown.
	// Instead, a protocol event `Page.fileChooserOpened` is emitted.
	// File chooser can be handled with `page.handleFileChooser` command.
	SetInterceptFileChooserDialog = "Page.setInterceptFileChooserDialog"
	
	// Accepts or cancels an intercepted file chooser dialog.
	HandleFileChooser = "Page.handleFileChooser"
	
)

// AddScriptToEvaluateOnLoad parameters
type AddScriptToEvaluateOnLoadParams struct {
	
	
	ScriptSource	string	`json:"scriptSource"`
	
}

// AddScriptToEvaluateOnLoad returns
type AddScriptToEvaluateOnLoadReturns struct {
	
	// Identifier of the added script.
	Identifier	ScriptIdentifier	`json:"identifier"`
	
}

// AddScriptToEvaluateOnNewDocument parameters
type AddScriptToEvaluateOnNewDocumentParams struct {
	
	
	Source	string	`json:"source"`
	
	// If specified, creates an isolated world with the given name and evaluates given script in it.
	// This world name will be used as the ExecutionContextDescription::name when the corresponding
	// event is emitted.
	WorldName	string	`json:"worldName"`
	
}

// AddScriptToEvaluateOnNewDocument returns
type AddScriptToEvaluateOnNewDocumentReturns struct {
	
	// Identifier of the added script.
	Identifier	ScriptIdentifier	`json:"identifier"`
	
}

// BringToFront parameters
type BringToFrontParams struct {
	
}

// BringToFront returns
type BringToFrontReturns struct {
	
}

// CaptureScreenshot parameters
type CaptureScreenshotParams struct {
	
	// Image compression format (defaults to png).
	Format	string	`json:"format"`
	
	// Compression quality from range [0..100] (jpeg only).
	Quality	int	`json:"quality"`
	
	// Capture the screenshot of a given region only.
	Clip	Viewport	`json:"clip"`
	
	// Capture the screenshot from the surface, rather than the view. Defaults to true.
	FromSurface	bool	`json:"fromSurface"`
	
}

// CaptureScreenshot returns
type CaptureScreenshotReturns struct {
	
	// Base64-encoded image data.
	Data	string	`json:"data"`
	
}

// CaptureSnapshot parameters
type CaptureSnapshotParams struct {
	
	// Format (defaults to mhtml).
	Format	string	`json:"format"`
	
}

// CaptureSnapshot returns
type CaptureSnapshotReturns struct {
	
	// Serialized page data.
	Data	string	`json:"data"`
	
}

// ClearDeviceMetricsOverride parameters
type ClearDeviceMetricsOverrideParams struct {
	
}

// ClearDeviceMetricsOverride returns
type ClearDeviceMetricsOverrideReturns struct {
	
}

// ClearDeviceOrientationOverride parameters
type ClearDeviceOrientationOverrideParams struct {
	
}

// ClearDeviceOrientationOverride returns
type ClearDeviceOrientationOverrideReturns struct {
	
}

// ClearGeolocationOverride parameters
type ClearGeolocationOverrideParams struct {
	
}

// ClearGeolocationOverride returns
type ClearGeolocationOverrideReturns struct {
	
}

// CreateIsolatedWorld parameters
type CreateIsolatedWorldParams struct {
	
	// Id of the frame in which the isolated world should be created.
	FrameId	FrameId	`json:"frameId"`
	
	// An optional name which is reported in the Execution Context.
	WorldName	string	`json:"worldName"`
	
	// Whether or not universal access should be granted to the isolated world. This is a powerful
	// option, use with caution.
	GrantUniveralAccess	bool	`json:"grantUniveralAccess"`
	
}

// CreateIsolatedWorld returns
type CreateIsolatedWorldReturns struct {
	
	// Execution context of the isolated world.
	ExecutionContextId	runtime.ExecutionContextId	`json:"executionContextId"`
	
}

// DeleteCookie parameters
type DeleteCookieParams struct {
	
	// Name of the cookie to remove.
	CookieName	string	`json:"cookieName"`
	
	// URL to match cooke domain and path.
	Url	string	`json:"url"`
	
}

// DeleteCookie returns
type DeleteCookieReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// GetAppManifest parameters
type GetAppManifestParams struct {
	
}

// GetAppManifest returns
type GetAppManifestReturns struct {
	
	// Manifest location.
	Url	string	`json:"url"`
	
	
	Errors	[]AppManifestError	`json:"errors"`
	
	// Manifest content.
	Data	string	`json:"data"`
	
}

// GetInstallabilityErrors parameters
type GetInstallabilityErrorsParams struct {
	
}

// GetInstallabilityErrors returns
type GetInstallabilityErrorsReturns struct {
	
	
	Errors	[]string	`json:"errors"`
	
}

// GetCookies parameters
type GetCookiesParams struct {
	
}

// GetCookies returns
type GetCookiesReturns struct {
	
	// Array of cookie objects.
	Cookies	[]network.Cookie	`json:"cookies"`
	
}

// GetFrameTree parameters
type GetFrameTreeParams struct {
	
}

// GetFrameTree returns
type GetFrameTreeReturns struct {
	
	// Present frame tree structure.
	FrameTree	FrameTree	`json:"frameTree"`
	
}

// GetLayoutMetrics parameters
type GetLayoutMetricsParams struct {
	
}

// GetLayoutMetrics returns
type GetLayoutMetricsReturns struct {
	
	// Metrics relating to the layout viewport.
	LayoutViewport	LayoutViewport	`json:"layoutViewport"`
	
	// Metrics relating to the visual viewport.
	VisualViewport	VisualViewport	`json:"visualViewport"`
	
	// Size of scrollable area.
	ContentSize	dom.Rect	`json:"contentSize"`
	
}

// GetNavigationHistory parameters
type GetNavigationHistoryParams struct {
	
}

// GetNavigationHistory returns
type GetNavigationHistoryReturns struct {
	
	// Index of the current navigation history entry.
	CurrentIndex	int	`json:"currentIndex"`
	
	// Array of navigation history entries.
	Entries	[]NavigationEntry	`json:"entries"`
	
}

// ResetNavigationHistory parameters
type ResetNavigationHistoryParams struct {
	
}

// ResetNavigationHistory returns
type ResetNavigationHistoryReturns struct {
	
}

// GetResourceContent parameters
type GetResourceContentParams struct {
	
	// Frame id to get resource for.
	FrameId	FrameId	`json:"frameId"`
	
	// URL of the resource to get content for.
	Url	string	`json:"url"`
	
}

// GetResourceContent returns
type GetResourceContentReturns struct {
	
	// Resource content.
	Content	string	`json:"content"`
	
	// True, if content was served as base64.
	Base64Encoded	bool	`json:"base64Encoded"`
	
}

// GetResourceTree parameters
type GetResourceTreeParams struct {
	
}

// GetResourceTree returns
type GetResourceTreeReturns struct {
	
	// Present frame / resource tree structure.
	FrameTree	FrameResourceTree	`json:"frameTree"`
	
}

// HandleJavaScriptDialog parameters
type HandleJavaScriptDialogParams struct {
	
	// Whether to accept or dismiss the dialog.
	Accept	bool	`json:"accept"`
	
	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt
	// dialog.
	PromptText	string	`json:"promptText"`
	
}

// HandleJavaScriptDialog returns
type HandleJavaScriptDialogReturns struct {
	
}

// Navigate parameters
type NavigateParams struct {
	
	// URL to navigate the page to.
	Url	string	`json:"url"`
	
	// Referrer URL.
	Referrer	string	`json:"referrer"`
	
	// Intended transition type.
	TransitionType	TransitionType	`json:"transitionType"`
	
	// Frame id to navigate, if not specified navigates the top frame.
	FrameId	FrameId	`json:"frameId"`
	
}

// Navigate returns
type NavigateReturns struct {
	
	// Frame id that has navigated (or failed to navigate)
	FrameId	FrameId	`json:"frameId"`
	
	// Loader identifier.
	LoaderId	network.LoaderId	`json:"loaderId"`
	
	// User friendly error message, present if and only if navigation has failed.
	ErrorText	string	`json:"errorText"`
	
}

// NavigateToHistoryEntry parameters
type NavigateToHistoryEntryParams struct {
	
	// Unique id of the entry to navigate to.
	EntryId	int	`json:"entryId"`
	
}

// NavigateToHistoryEntry returns
type NavigateToHistoryEntryReturns struct {
	
}

// PrintToPDF parameters
type PrintToPDFParams struct {
	
	// Paper orientation. Defaults to false.
	Landscape	bool	`json:"landscape"`
	
	// Display header and footer. Defaults to false.
	DisplayHeaderFooter	bool	`json:"displayHeaderFooter"`
	
	// Print background graphics. Defaults to false.
	PrintBackground	bool	`json:"printBackground"`
	
	// Scale of the webpage rendering. Defaults to 1.
	Scale	float64	`json:"scale"`
	
	// Paper width in inches. Defaults to 8.5 inches.
	PaperWidth	float64	`json:"paperWidth"`
	
	// Paper height in inches. Defaults to 11 inches.
	PaperHeight	float64	`json:"paperHeight"`
	
	// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop	float64	`json:"marginTop"`
	
	// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom	float64	`json:"marginBottom"`
	
	// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft	float64	`json:"marginLeft"`
	
	// Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight	float64	`json:"marginRight"`
	
	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means
	// print all pages.
	PageRanges	string	`json:"pageRanges"`
	
	// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'.
	// Defaults to false.
	IgnoreInvalidPageRanges	bool	`json:"ignoreInvalidPageRanges"`
	
	// HTML template for the print header. Should be valid HTML markup with following
	// classes used to inject printing values into them:
	// - `date`: formatted print date
	// - `title`: document title
	// - `url`: document location
	// - `pageNumber`: current page number
	// - `totalPages`: total pages in the document
	// 
	// For example, `<span class=title></span>` would generate span containing the title.
	HeaderTemplate	string	`json:"headerTemplate"`
	
	// HTML template for the print footer. Should use the same format as the `headerTemplate`.
	FooterTemplate	string	`json:"footerTemplate"`
	
	// Whether or not to prefer page size as defined by css. Defaults to false,
	// in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize	bool	`json:"preferCSSPageSize"`
	
	// return as stream
	TransferMode	string	`json:"transferMode"`
	
}

// PrintToPDF returns
type PrintToPDFReturns struct {
	
	// Base64-encoded pdf data. Empty if |returnAsStream| is specified.
	Data	string	`json:"data"`
	
	// A handle of the stream that holds resulting PDF data.
	Stream	io.StreamHandle	`json:"stream"`
	
}

// Reload parameters
type ReloadParams struct {
	
	// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache	bool	`json:"ignoreCache"`
	
	// If set, the script will be injected into all frames of the inspected page after reload.
	// Argument will be ignored if reloading dataURL origin.
	ScriptToEvaluateOnLoad	string	`json:"scriptToEvaluateOnLoad"`
	
}

// Reload returns
type ReloadReturns struct {
	
}

// RemoveScriptToEvaluateOnLoad parameters
type RemoveScriptToEvaluateOnLoadParams struct {
	
	
	Identifier	ScriptIdentifier	`json:"identifier"`
	
}

// RemoveScriptToEvaluateOnLoad returns
type RemoveScriptToEvaluateOnLoadReturns struct {
	
}

// RemoveScriptToEvaluateOnNewDocument parameters
type RemoveScriptToEvaluateOnNewDocumentParams struct {
	
	
	Identifier	ScriptIdentifier	`json:"identifier"`
	
}

// RemoveScriptToEvaluateOnNewDocument returns
type RemoveScriptToEvaluateOnNewDocumentReturns struct {
	
}

// ScreencastFrameAck parameters
type ScreencastFrameAckParams struct {
	
	// Frame number.
	SessionId	int	`json:"sessionId"`
	
}

// ScreencastFrameAck returns
type ScreencastFrameAckReturns struct {
	
}

// SearchInResource parameters
type SearchInResourceParams struct {
	
	// Frame id for resource to search in.
	FrameId	FrameId	`json:"frameId"`
	
	// URL of the resource to search in.
	Url	string	`json:"url"`
	
	// String to search for.
	Query	string	`json:"query"`
	
	// If true, search is case sensitive.
	CaseSensitive	bool	`json:"caseSensitive"`
	
	// If true, treats string parameter as regex.
	IsRegex	bool	`json:"isRegex"`
	
}

// SearchInResource returns
type SearchInResourceReturns struct {
	
	// List of search matches.
	Result	[]debugger.SearchMatch	`json:"result"`
	
}

// SetAdBlockingEnabled parameters
type SetAdBlockingEnabledParams struct {
	
	// Whether to block ads.
	Enabled	bool	`json:"enabled"`
	
}

// SetAdBlockingEnabled returns
type SetAdBlockingEnabledReturns struct {
	
}

// SetBypassCSP parameters
type SetBypassCSPParams struct {
	
	// Whether to bypass page CSP.
	Enabled	bool	`json:"enabled"`
	
}

// SetBypassCSP returns
type SetBypassCSPReturns struct {
	
}

// SetDeviceMetricsOverride parameters
type SetDeviceMetricsOverrideParams struct {
	
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width	int	`json:"width"`
	
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height	int	`json:"height"`
	
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor	float64	`json:"deviceScaleFactor"`
	
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text
	// autosizing and more.
	Mobile	bool	`json:"mobile"`
	
	// Scale to apply to resulting view image.
	Scale	float64	`json:"scale"`
	
	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenWidth	int	`json:"screenWidth"`
	
	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	ScreenHeight	int	`json:"screenHeight"`
	
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionX	int	`json:"positionX"`
	
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	PositionY	int	`json:"positionY"`
	
	// Do not set visible view size, rely upon explicit setVisibleSize call.
	DontSetVisibleSize	bool	`json:"dontSetVisibleSize"`
	
	// Screen orientation override.
	ScreenOrientation	emulation.ScreenOrientation	`json:"screenOrientation"`
	
	// The viewport dimensions and scale. If not set, the override is cleared.
	Viewport	Viewport	`json:"viewport"`
	
}

// SetDeviceMetricsOverride returns
type SetDeviceMetricsOverrideReturns struct {
	
}

// SetDeviceOrientationOverride parameters
type SetDeviceOrientationOverrideParams struct {
	
	// Mock alpha
	Alpha	float64	`json:"alpha"`
	
	// Mock beta
	Beta	float64	`json:"beta"`
	
	// Mock gamma
	Gamma	float64	`json:"gamma"`
	
}

// SetDeviceOrientationOverride returns
type SetDeviceOrientationOverrideReturns struct {
	
}

// SetFontFamilies parameters
type SetFontFamiliesParams struct {
	
	// Specifies font families to set. If a font family is not specified, it won't be changed.
	FontFamilies	FontFamilies	`json:"fontFamilies"`
	
}

// SetFontFamilies returns
type SetFontFamiliesReturns struct {
	
}

// SetFontSizes parameters
type SetFontSizesParams struct {
	
	// Specifies font sizes to set. If a font size is not specified, it won't be changed.
	FontSizes	FontSizes	`json:"fontSizes"`
	
}

// SetFontSizes returns
type SetFontSizesReturns struct {
	
}

// SetDocumentContent parameters
type SetDocumentContentParams struct {
	
	// Frame id to set HTML for.
	FrameId	FrameId	`json:"frameId"`
	
	// HTML content to set.
	Html	string	`json:"html"`
	
}

// SetDocumentContent returns
type SetDocumentContentReturns struct {
	
}

// SetDownloadBehavior parameters
type SetDownloadBehaviorParams struct {
	
	// Whether to allow all or deny all download requests, or use default Chrome behavior if
	// available (otherwise deny).
	Behavior	string	`json:"behavior"`
	
	// The default path to save downloaded files to. This is requred if behavior is set to 'allow'
	DownloadPath	string	`json:"downloadPath"`
	
}

// SetDownloadBehavior returns
type SetDownloadBehaviorReturns struct {
	
}

// SetGeolocationOverride parameters
type SetGeolocationOverrideParams struct {
	
	// Mock latitude
	Latitude	float64	`json:"latitude"`
	
	// Mock longitude
	Longitude	float64	`json:"longitude"`
	
	// Mock accuracy
	Accuracy	float64	`json:"accuracy"`
	
}

// SetGeolocationOverride returns
type SetGeolocationOverrideReturns struct {
	
}

// SetLifecycleEventsEnabled parameters
type SetLifecycleEventsEnabledParams struct {
	
	// If true, starts emitting lifecycle events.
	Enabled	bool	`json:"enabled"`
	
}

// SetLifecycleEventsEnabled returns
type SetLifecycleEventsEnabledReturns struct {
	
}

// SetTouchEmulationEnabled parameters
type SetTouchEmulationEnabledParams struct {
	
	// Whether the touch event emulation should be enabled.
	Enabled	bool	`json:"enabled"`
	
	// Touch/gesture events configuration. Default: current platform.
	Configuration	string	`json:"configuration"`
	
}

// SetTouchEmulationEnabled returns
type SetTouchEmulationEnabledReturns struct {
	
}

// StartScreencast parameters
type StartScreencastParams struct {
	
	// Image compression format.
	Format	string	`json:"format"`
	
	// Compression quality from range [0..100].
	Quality	int	`json:"quality"`
	
	// Maximum screenshot width.
	MaxWidth	int	`json:"maxWidth"`
	
	// Maximum screenshot height.
	MaxHeight	int	`json:"maxHeight"`
	
	// Send every n-th frame.
	EveryNthFrame	int	`json:"everyNthFrame"`
	
}

// StartScreencast returns
type StartScreencastReturns struct {
	
}

// StopLoading parameters
type StopLoadingParams struct {
	
}

// StopLoading returns
type StopLoadingReturns struct {
	
}

// Crash parameters
type CrashParams struct {
	
}

// Crash returns
type CrashReturns struct {
	
}

// Close parameters
type CloseParams struct {
	
}

// Close returns
type CloseReturns struct {
	
}

// SetWebLifecycleState parameters
type SetWebLifecycleStateParams struct {
	
	// Target lifecycle state
	State	string	`json:"state"`
	
}

// SetWebLifecycleState returns
type SetWebLifecycleStateReturns struct {
	
}

// StopScreencast parameters
type StopScreencastParams struct {
	
}

// StopScreencast returns
type StopScreencastReturns struct {
	
}

// SetProduceCompilationCache parameters
type SetProduceCompilationCacheParams struct {
	
	
	Enabled	bool	`json:"enabled"`
	
}

// SetProduceCompilationCache returns
type SetProduceCompilationCacheReturns struct {
	
}

// AddCompilationCache parameters
type AddCompilationCacheParams struct {
	
	
	Url	string	`json:"url"`
	
	// Base64-encoded data
	Data	string	`json:"data"`
	
}

// AddCompilationCache returns
type AddCompilationCacheReturns struct {
	
}

// ClearCompilationCache parameters
type ClearCompilationCacheParams struct {
	
}

// ClearCompilationCache returns
type ClearCompilationCacheReturns struct {
	
}

// GenerateTestReport parameters
type GenerateTestReportParams struct {
	
	// Message to be displayed in the report.
	Message	string	`json:"message"`
	
	// Specifies the endpoint group to deliver the report to.
	Group	string	`json:"group"`
	
}

// GenerateTestReport returns
type GenerateTestReportReturns struct {
	
}

// WaitForDebugger parameters
type WaitForDebuggerParams struct {
	
}

// WaitForDebugger returns
type WaitForDebuggerReturns struct {
	
}

// SetInterceptFileChooserDialog parameters
type SetInterceptFileChooserDialogParams struct {
	
	
	Enabled	bool	`json:"enabled"`
	
}

// SetInterceptFileChooserDialog returns
type SetInterceptFileChooserDialogReturns struct {
	
}

// HandleFileChooser parameters
type HandleFileChooserParams struct {
	
	
	Action	string	`json:"action"`
	
	// Array of absolute file paths to set, only respected with `accept` action.
	Files	[]string	`json:"files"`
	
}

// HandleFileChooser returns
type HandleFileChooserReturns struct {
	
}

