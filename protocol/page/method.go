package page

import (
	"github.com/diiyw/cuto/protocol/network"
	"github.com/diiyw/cuto/protocol/debugger"
)


// Deprecated, please use addScriptToEvaluateOnNewDocument instead.
const AddScriptToEvaluateOnLoad = "Page.addScriptToEvaluateOnLoad"

type AddScriptToEvaluateOnLoadParams struct {

	// 
	ScriptSource 	string	`json:"scriptSource"`
}

type AddScriptToEvaluateOnLoadResult struct {

	// Identifier of the added script.
	Identifier 	ScriptIdentifier	`json:"identifier"`
}

// Evaluates given script in every frame upon creation (before loading frame's scripts).
const AddScriptToEvaluateOnNewDocument = "Page.addScriptToEvaluateOnNewDocument"

type AddScriptToEvaluateOnNewDocumentParams struct {

	// 
	Source 	string	`json:"source"`

	// If specified, creates an isolated world with the given name and evaluates given script in it.
	// This world name will be used as the ExecutionContextDescription::name when the corresponding
	// event is emitted.
	WorldName 	string	`json:"worldName"`
}

type AddScriptToEvaluateOnNewDocumentResult struct {

	// Identifier of the added script.
	Identifier 	ScriptIdentifier	`json:"identifier"`
}

// Brings page to front (activates tab).
const BringToFront = "Page.bringToFront"

type BringToFrontParams struct {
}

type BringToFrontResult struct {

}

// Capture page screenshot.
const CaptureScreenshot = "Page.captureScreenshot"

type CaptureScreenshotParams struct {

	// Image compression format (defaults to png).
	Format 	string	`json:"format"`

	// Compression quality from range [0..100] (jpeg only).
	Quality 	int	`json:"quality"`

	// Capture the screenshot of a given region only.
	Clip 	Viewport	`json:"clip"`

	// Capture the screenshot from the surface, rather than the view. Defaults to true.
	FromSurface 	bool	`json:"fromSurface"`
}

type CaptureScreenshotResult struct {

	// Base64-encoded image data.
	Data 	[]byte	`json:"data"`
}

// Returns a snapshot of the page as a string. For MHTML format, the serialization includes
// iframes, shadow DOM, external resources, and element-inline styles.
const CaptureSnapshot = "Page.captureSnapshot"

type CaptureSnapshotParams struct {

	// Format (defaults to mhtml).
	Format 	string	`json:"format"`
}

type CaptureSnapshotResult struct {

	// Serialized page data.
	Data 	string	`json:"data"`
}

// Clears the overriden device metrics.
const ClearDeviceMetricsOverride = "Page.clearDeviceMetricsOverride"

type ClearDeviceMetricsOverrideParams struct {
}

type ClearDeviceMetricsOverrideResult struct {

}

// Clears the overridden Device Orientation.
const ClearDeviceOrientationOverride = "Page.clearDeviceOrientationOverride"

type ClearDeviceOrientationOverrideParams struct {
}

type ClearDeviceOrientationOverrideResult struct {

}

// Clears the overriden Geolocation Position and Error.
const ClearGeolocationOverride = "Page.clearGeolocationOverride"

type ClearGeolocationOverrideParams struct {
}

type ClearGeolocationOverrideResult struct {

}

// Creates an isolated world for the given frame.
const CreateIsolatedWorld = "Page.createIsolatedWorld"

type CreateIsolatedWorldParams struct {

	// Id of the frame in which the isolated world should be created.
	FrameId 	FrameId	`json:"frameId"`

	// An optional name which is reported in the Execution Context.
	WorldName 	string	`json:"worldName"`

	// Whether or not universal access should be granted to the isolated world. This is a powerful
	// option, use with caution.
	GrantUniveralAccess 	bool	`json:"grantUniveralAccess"`
}

type CreateIsolatedWorldResult struct {

	// Execution context of the isolated world.
	ExecutionContextId 	interface{}	`json:"executionContextId"`
}

// Deletes browser cookie with given name, domain and path.
const DeleteCookie = "Page.deleteCookie"

type DeleteCookieParams struct {

	// Name of the cookie to remove.
	CookieName 	string	`json:"cookieName"`

	// URL to match cooke domain and path.
	Url 	string	`json:"url"`
}

type DeleteCookieResult struct {

}

// Disables page domain notifications.
const Disable = "Page.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables page domain notifications.
const Enable = "Page.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// 
const GetAppManifest = "Page.getAppManifest"

type GetAppManifestParams struct {
}

type GetAppManifestResult struct {

	// Manifest location.
	Url 	string	`json:"url"`
	// 
	Errors 	[]*AppManifestError	`json:"errors"`
	// Manifest content.
	Data 	string	`json:"data"`
}

// 
const GetInstallabilityErrors = "Page.getInstallabilityErrors"

type GetInstallabilityErrorsParams struct {
}

type GetInstallabilityErrorsResult struct {

	// 
	Errors 	[]string	`json:"errors"`
}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie
// information in the `cookies` field.
const GetCookies = "Page.getCookies"

type GetCookiesParams struct {
}

type GetCookiesResult struct {

	// Array of cookie objects.
	Cookies 	[]*network.Cookie	`json:"cookies"`
}

// Returns present frame tree structure.
const GetFrameTree = "Page.getFrameTree"

type GetFrameTreeParams struct {
}

type GetFrameTreeResult struct {

	// Present frame tree structure.
	FrameTree 	FrameTree	`json:"frameTree"`
}

// Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
const GetLayoutMetrics = "Page.getLayoutMetrics"

type GetLayoutMetricsParams struct {
}

type GetLayoutMetricsResult struct {

	// Metrics relating to the layout viewport.
	LayoutViewport 	LayoutViewport	`json:"layoutViewport"`
	// Metrics relating to the visual viewport.
	VisualViewport 	VisualViewport	`json:"visualViewport"`
	// Size of scrollable area.
	ContentSize 	interface{}	`json:"contentSize"`
}

// Returns navigation history for the current page.
const GetNavigationHistory = "Page.getNavigationHistory"

type GetNavigationHistoryParams struct {
}

type GetNavigationHistoryResult struct {

	// Index of the current navigation history entry.
	CurrentIndex 	int	`json:"currentIndex"`
	// Array of navigation history entries.
	Entries 	[]*NavigationEntry	`json:"entries"`
}

// Resets navigation history for the current page.
const ResetNavigationHistory = "Page.resetNavigationHistory"

type ResetNavigationHistoryParams struct {
}

type ResetNavigationHistoryResult struct {

}

// Returns content of the given resource.
const GetResourceContent = "Page.getResourceContent"

type GetResourceContentParams struct {

	// Frame id to get resource for.
	FrameId 	FrameId	`json:"frameId"`

	// URL of the resource to get content for.
	Url 	string	`json:"url"`
}

type GetResourceContentResult struct {

	// Resource content.
	Content 	string	`json:"content"`
	// True, if content was served as base64.
	Base64Encoded 	bool	`json:"base64Encoded"`
}

// Returns present frame / resource tree structure.
const GetResourceTree = "Page.getResourceTree"

type GetResourceTreeParams struct {
}

type GetResourceTreeResult struct {

	// Present frame / resource tree structure.
	FrameTree 	FrameResourceTree	`json:"frameTree"`
}

// Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
const HandleJavaScriptDialog = "Page.handleJavaScriptDialog"

type HandleJavaScriptDialogParams struct {

	// Whether to accept or dismiss the dialog.
	Accept 	bool	`json:"accept"`

	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt
	// dialog.
	PromptText 	string	`json:"promptText"`
}

type HandleJavaScriptDialogResult struct {

}

// Navigates current page to the given URL.
const Navigate = "Page.navigate"

type NavigateParams struct {

	// URL to navigate the page to.
	Url 	string	`json:"url"`

	// Referrer URL.
	Referrer 	string	`json:"referrer"`

	// Intended transition type.
	TransitionType 	TransitionType	`json:"transitionType"`

	// Frame id to navigate, if not specified navigates the top frame.
	FrameId 	FrameId	`json:"frameId"`
}

type NavigateResult struct {

	// Frame id that has navigated (or failed to navigate)
	FrameId 	FrameId	`json:"frameId"`
	// Loader identifier.
	LoaderId 	interface{}	`json:"loaderId"`
	// User friendly error message, present if and only if navigation has failed.
	ErrorText 	string	`json:"errorText"`
}

// Navigates current page to the given history entry.
const NavigateToHistoryEntry = "Page.navigateToHistoryEntry"

type NavigateToHistoryEntryParams struct {

	// Unique id of the entry to navigate to.
	EntryId 	int	`json:"entryId"`
}

type NavigateToHistoryEntryResult struct {

}

// Print page as PDF.
const PrintToPDF = "Page.printToPDF"

type PrintToPDFParams struct {

	// Paper orientation. Defaults to false.
	Landscape 	bool	`json:"landscape"`

	// Display header and footer. Defaults to false.
	DisplayHeaderFooter 	bool	`json:"displayHeaderFooter"`

	// Print background graphics. Defaults to false.
	PrintBackground 	bool	`json:"printBackground"`

	// Scale of the webpage rendering. Defaults to 1.
	Scale 	float64	`json:"scale"`

	// Paper width in inches. Defaults to 8.5 inches.
	PaperWidth 	float64	`json:"paperWidth"`

	// Paper height in inches. Defaults to 11 inches.
	PaperHeight 	float64	`json:"paperHeight"`

	// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop 	float64	`json:"marginTop"`

	// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom 	float64	`json:"marginBottom"`

	// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft 	float64	`json:"marginLeft"`

	// Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight 	float64	`json:"marginRight"`

	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means
	// print all pages.
	PageRanges 	string	`json:"pageRanges"`

	// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'.
	// Defaults to false.
	IgnoreInvalidPageRanges 	bool	`json:"ignoreInvalidPageRanges"`

	// HTML template for the print header. Should be valid HTML markup with following
	// classes used to inject printing values into them:
	// - `date`: formatted print date
	// - `title`: document title
	// - `url`: document location
	// - `pageNumber`: current page number
	// - `totalPages`: total pages in the document
	// 
	// For example, `<span class=title></span>` would generate span containing the title.
	HeaderTemplate 	string	`json:"headerTemplate"`

	// HTML template for the print footer. Should use the same format as the `headerTemplate`.
	FooterTemplate 	string	`json:"footerTemplate"`

	// Whether or not to prefer page size as defined by css. Defaults to false,
	// in which case the content will be scaled to fit the paper size.
	PreferCSSPageSize 	bool	`json:"preferCSSPageSize"`

	// return as stream
	TransferMode 	string	`json:"transferMode"`
}

type PrintToPDFResult struct {

	// Base64-encoded pdf data. Empty if |returnAsStream| is specified.
	Data 	[]byte	`json:"data"`
	// A handle of the stream that holds resulting PDF data.
	Stream 	interface{}	`json:"stream"`
}

// Reloads given page optionally ignoring the cache.
const Reload = "Page.reload"

type ReloadParams struct {

	// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache 	bool	`json:"ignoreCache"`

	// If set, the script will be injected into all frames of the inspected page after reload.
	// Argument will be ignored if reloading dataURL origin.
	ScriptToEvaluateOnLoad 	string	`json:"scriptToEvaluateOnLoad"`
}

type ReloadResult struct {

}

// Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
const RemoveScriptToEvaluateOnLoad = "Page.removeScriptToEvaluateOnLoad"

type RemoveScriptToEvaluateOnLoadParams struct {

	// 
	Identifier 	ScriptIdentifier	`json:"identifier"`
}

type RemoveScriptToEvaluateOnLoadResult struct {

}

// Removes given script from the list.
const RemoveScriptToEvaluateOnNewDocument = "Page.removeScriptToEvaluateOnNewDocument"

type RemoveScriptToEvaluateOnNewDocumentParams struct {

	// 
	Identifier 	ScriptIdentifier	`json:"identifier"`
}

type RemoveScriptToEvaluateOnNewDocumentResult struct {

}

// Acknowledges that a screencast frame has been received by the frontend.
const ScreencastFrameAck = "Page.screencastFrameAck"

type ScreencastFrameAckParams struct {

	// Frame number.
	SessionId 	int	`json:"sessionId"`
}

type ScreencastFrameAckResult struct {

}

// Searches for given string in resource content.
const SearchInResource = "Page.searchInResource"

type SearchInResourceParams struct {

	// Frame id for resource to search in.
	FrameId 	FrameId	`json:"frameId"`

	// URL of the resource to search in.
	Url 	string	`json:"url"`

	// String to search for.
	Query 	string	`json:"query"`

	// If true, search is case sensitive.
	CaseSensitive 	bool	`json:"caseSensitive"`

	// If true, treats string parameter as regex.
	IsRegex 	bool	`json:"isRegex"`
}

type SearchInResourceResult struct {

	// List of search matches.
	Result 	[]*debugger.SearchMatch	`json:"result"`
}

// Enable Chrome's experimental ad filter on all sites.
const SetAdBlockingEnabled = "Page.setAdBlockingEnabled"

type SetAdBlockingEnabledParams struct {

	// Whether to block ads.
	Enabled 	bool	`json:"enabled"`
}

type SetAdBlockingEnabledResult struct {

}

// Enable page Content Security Policy by-passing.
const SetBypassCSP = "Page.setBypassCSP"

type SetBypassCSPParams struct {

	// Whether to bypass page CSP.
	Enabled 	bool	`json:"enabled"`
}

type SetBypassCSPResult struct {

}

// Overrides the values of device screen dimensions (window.screen.width, window.screen.height,
// window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
const SetDeviceMetricsOverride = "Page.setDeviceMetricsOverride"

type SetDeviceMetricsOverrideParams struct {

	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width 	int	`json:"width"`

	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height 	int	`json:"height"`

	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor 	float64	`json:"deviceScaleFactor"`

	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text
	// autosizing and more.
	Mobile 	bool	`json:"mobile"`

	// Scale to apply to resulting view image.
	Scale 	float64	`json:"scale"`

	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenWidth 	int	`json:"screenWidth"`

	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	ScreenHeight 	int	`json:"screenHeight"`

	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionX 	int	`json:"positionX"`

	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	PositionY 	int	`json:"positionY"`

	// Do not set visible view size, rely upon explicit setVisibleSize call.
	DontSetVisibleSize 	bool	`json:"dontSetVisibleSize"`

	// Screen orientation override.
	ScreenOrientation 	interface{}	`json:"screenOrientation"`

	// The viewport dimensions and scale. If not set, the override is cleared.
	Viewport 	Viewport	`json:"viewport"`
}

type SetDeviceMetricsOverrideResult struct {

}

// Overrides the Device Orientation.
const SetDeviceOrientationOverride = "Page.setDeviceOrientationOverride"

type SetDeviceOrientationOverrideParams struct {

	// Mock alpha
	Alpha 	float64	`json:"alpha"`

	// Mock beta
	Beta 	float64	`json:"beta"`

	// Mock gamma
	Gamma 	float64	`json:"gamma"`
}

type SetDeviceOrientationOverrideResult struct {

}

// Set generic font families.
const SetFontFamilies = "Page.setFontFamilies"

type SetFontFamiliesParams struct {

	// Specifies font families to set. If a font family is not specified, it won't be changed.
	FontFamilies 	FontFamilies	`json:"fontFamilies"`
}

type SetFontFamiliesResult struct {

}

// Set default font sizes.
const SetFontSizes = "Page.setFontSizes"

type SetFontSizesParams struct {

	// Specifies font sizes to set. If a font size is not specified, it won't be changed.
	FontSizes 	FontSizes	`json:"fontSizes"`
}

type SetFontSizesResult struct {

}

// Sets given markup as the document's HTML.
const SetDocumentContent = "Page.setDocumentContent"

type SetDocumentContentParams struct {

	// Frame id to set HTML for.
	FrameId 	FrameId	`json:"frameId"`

	// HTML content to set.
	Html 	string	`json:"html"`
}

type SetDocumentContentResult struct {

}

// Set the behavior when downloading a file.
const SetDownloadBehavior = "Page.setDownloadBehavior"

type SetDownloadBehaviorParams struct {

	// Whether to allow all or deny all download requests, or use default Chrome behavior if
	// available (otherwise deny).
	Behavior 	string	`json:"behavior"`

	// The default path to save downloaded files to. This is requred if behavior is set to 'allow'
	DownloadPath 	string	`json:"downloadPath"`
}

type SetDownloadBehaviorResult struct {

}

// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position
// unavailable.
const SetGeolocationOverride = "Page.setGeolocationOverride"

type SetGeolocationOverrideParams struct {

	// Mock latitude
	Latitude 	float64	`json:"latitude"`

	// Mock longitude
	Longitude 	float64	`json:"longitude"`

	// Mock accuracy
	Accuracy 	float64	`json:"accuracy"`
}

type SetGeolocationOverrideResult struct {

}

// Controls whether page will emit lifecycle events.
const SetLifecycleEventsEnabled = "Page.setLifecycleEventsEnabled"

type SetLifecycleEventsEnabledParams struct {

	// If true, starts emitting lifecycle events.
	Enabled 	bool	`json:"enabled"`
}

type SetLifecycleEventsEnabledResult struct {

}

// Toggles mouse event-based touch event emulation.
const SetTouchEmulationEnabled = "Page.setTouchEmulationEnabled"

type SetTouchEmulationEnabledParams struct {

	// Whether the touch event emulation should be enabled.
	Enabled 	bool	`json:"enabled"`

	// Touch/gesture events configuration. Default: current platform.
	Configuration 	string	`json:"configuration"`
}

type SetTouchEmulationEnabledResult struct {

}

// Starts sending each frame using the `screencastFrame` event.
const StartScreencast = "Page.startScreencast"

type StartScreencastParams struct {

	// Image compression format.
	Format 	string	`json:"format"`

	// Compression quality from range [0..100].
	Quality 	int	`json:"quality"`

	// Maximum screenshot width.
	MaxWidth 	int	`json:"maxWidth"`

	// Maximum screenshot height.
	MaxHeight 	int	`json:"maxHeight"`

	// Send every n-th frame.
	EveryNthFrame 	int	`json:"everyNthFrame"`
}

type StartScreencastResult struct {

}

// Force the page stop all navigations and pending resource fetches.
const StopLoading = "Page.stopLoading"

type StopLoadingParams struct {
}

type StopLoadingResult struct {

}

// Crashes renderer on the IO thread, generates minidumps.
const Crash = "Page.crash"

type CrashParams struct {
}

type CrashResult struct {

}

// Tries to close page, running its beforeunload hooks, if any.
const Close = "Page.close"

type CloseParams struct {
}

type CloseResult struct {

}

// Tries to update the web lifecycle state of the page.
// It will transition the page to the given state according to:
// https://github.com/WICG/web-lifecycle/
const SetWebLifecycleState = "Page.setWebLifecycleState"

type SetWebLifecycleStateParams struct {

	// Target lifecycle state
	State 	string	`json:"state"`
}

type SetWebLifecycleStateResult struct {

}

// Stops sending each frame in the `screencastFrame`.
const StopScreencast = "Page.stopScreencast"

type StopScreencastParams struct {
}

type StopScreencastResult struct {

}

// Forces compilation cache to be generated for every subresource script.
const SetProduceCompilationCache = "Page.setProduceCompilationCache"

type SetProduceCompilationCacheParams struct {

	// 
	Enabled 	bool	`json:"enabled"`
}

type SetProduceCompilationCacheResult struct {

}

// Seeds compilation cache for given url. Compilation cache does not survive
// cross-process navigation.
const AddCompilationCache = "Page.addCompilationCache"

type AddCompilationCacheParams struct {

	// 
	Url 	string	`json:"url"`

	// Base64-encoded data
	Data 	[]byte	`json:"data"`
}

type AddCompilationCacheResult struct {

}

// Clears seeded compilation cache.
const ClearCompilationCache = "Page.clearCompilationCache"

type ClearCompilationCacheParams struct {
}

type ClearCompilationCacheResult struct {

}

// Generates a report for testing.
const GenerateTestReport = "Page.generateTestReport"

type GenerateTestReportParams struct {

	// Message to be displayed in the report.
	Message 	string	`json:"message"`

	// Specifies the endpoint group to deliver the report to.
	Group 	string	`json:"group"`
}

type GenerateTestReportResult struct {

}

// Pauses page execution. Can be resumed using generic Runtime.runIfWaitingForDebugger.
const WaitForDebugger = "Page.waitForDebugger"

type WaitForDebuggerParams struct {
}

type WaitForDebuggerResult struct {

}

// Intercept file chooser requests and transfer control to protocol clients.
// When file chooser interception is enabled, native file chooser dialog is not shown.
// Instead, a protocol event `Page.fileChooserOpened` is emitted.
// File chooser can be handled with `page.handleFileChooser` command.
const SetInterceptFileChooserDialog = "Page.setInterceptFileChooserDialog"

type SetInterceptFileChooserDialogParams struct {

	// 
	Enabled 	bool	`json:"enabled"`
}

type SetInterceptFileChooserDialogResult struct {

}

// Accepts or cancels an intercepted file chooser dialog.
const HandleFileChooser = "Page.handleFileChooser"

type HandleFileChooserParams struct {

	// 
	Action 	string	`json:"action"`

	// Array of absolute file paths to set, only respected with `accept` action.
	Files 	[]string	`json:"files"`
}

type HandleFileChooserResult struct {

}