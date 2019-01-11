package browser

import (

	"github.com/diiyw/gator/protocol/target"

)
const (
	
	// Grant specific permissions to the given origin and reject all others.
	GrantPermissions = "Browser.grantPermissions"
	
	// Reset all permission management for all origins.
	ResetPermissions = "Browser.resetPermissions"
	
	// Close browser gracefully.
	Close = "Browser.close"
	
	// Crashes browser on the main thread.
	Crash = "Browser.crash"
	
	// Returns version information.
	GetVersion = "Browser.getVersion"
	
	// Returns the command line switches for the browser process if, and only if
	// --enable-automation is on the commandline.
	GetBrowserCommandLine = "Browser.getBrowserCommandLine"
	
	// Get Chrome histograms.
	GetHistograms = "Browser.getHistograms"
	
	// Get a Chrome histogram by name.
	GetHistogram = "Browser.getHistogram"
	
	// Get position and size of the browser window.
	GetWindowBounds = "Browser.getWindowBounds"
	
	// Get the browser window that contains the devtools target.
	GetWindowForTarget = "Browser.getWindowForTarget"
	
	// Set position and/or size of the browser window.
	SetWindowBounds = "Browser.setWindowBounds"
	
	// Set dock tile details, platform-specific.
	SetDockTile = "Browser.setDockTile"
	
)

// GrantPermissions parameters
type GrantPermissionsParams struct {
	
	
	Origin	string	`json:"origin"`
	
	
	Permissions	[]PermissionType	`json:"permissions"`
	
	// BrowserContext to override permissions. When omitted, default browser context is used.
	BrowserContextId	target.BrowserContextID	`json:"browserContextId,omitempty"`
	
}

// GrantPermissions returns
type GrantPermissionsReturns struct {
	
}

// ResetPermissions parameters
type ResetPermissionsParams struct {
	
	// BrowserContext to reset permissions. When omitted, default browser context is used.
	BrowserContextId	target.BrowserContextID	`json:"browserContextId,omitempty"`
	
}

// ResetPermissions returns
type ResetPermissionsReturns struct {
	
}

// Close parameters
type CloseParams struct {
	
}

// Close returns
type CloseReturns struct {
	
}

// Crash parameters
type CrashParams struct {
	
}

// Crash returns
type CrashReturns struct {
	
}

// GetVersion parameters
type GetVersionParams struct {
	
}

// GetVersion returns
type GetVersionReturns struct {
	
	// Protocol version.
	ProtocolVersion	string	`json:"protocolVersion"`
	
	// Product name.
	Product	string	`json:"product"`
	
	// Product revision.
	Revision	string	`json:"revision"`
	
	// User-Agent.
	UserAgent	string	`json:"userAgent"`
	
	// V8 version.
	JsVersion	string	`json:"jsVersion"`
	
}

// GetBrowserCommandLine parameters
type GetBrowserCommandLineParams struct {
	
}

// GetBrowserCommandLine returns
type GetBrowserCommandLineReturns struct {
	
	// Commandline parameters
	Arguments	[]string	`json:"arguments"`
	
}

// GetHistograms parameters
type GetHistogramsParams struct {
	
	// Requested substring in name. Only histograms which have query as a
	// substring in their name are extracted. An empty or absent query returns
	// all histograms.
	Query	string	`json:"query"`
	
	// If true, retrieve delta since last call.
	Delta	bool	`json:"delta"`
	
}

// GetHistograms returns
type GetHistogramsReturns struct {
	
	// Histograms.
	Histograms	[]Histogram	`json:"histograms"`
	
}

// GetHistogram parameters
type GetHistogramParams struct {
	
	// Requested histogram name.
	Name	string	`json:"name"`
	
	// If true, retrieve delta since last call.
	Delta	bool	`json:"delta"`
	
}

// GetHistogram returns
type GetHistogramReturns struct {
	
	// Histogram.
	Histogram	Histogram	`json:"histogram"`
	
}

// GetWindowBounds parameters
type GetWindowBoundsParams struct {
	
	// Browser window id.
	WindowId	WindowID	`json:"windowId"`
	
}

// GetWindowBounds returns
type GetWindowBoundsReturns struct {
	
	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds	Bounds	`json:"bounds"`
	
}

// GetWindowForTarget parameters
type GetWindowForTargetParams struct {
	
	// Devtools agent host id. If called as a part of the session, associated targetId is used.
	TargetId	target.TargetID	`json:"targetId"`
	
}

// GetWindowForTarget returns
type GetWindowForTargetReturns struct {
	
	// Browser window id.
	WindowId	WindowID	`json:"windowId"`
	
	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds	Bounds	`json:"bounds"`
	
}

// SetWindowBounds parameters
type SetWindowBoundsParams struct {
	
	// Browser window id.
	WindowId	WindowID	`json:"windowId"`
	
	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds	Bounds	`json:"bounds"`
	
}

// SetWindowBounds returns
type SetWindowBoundsReturns struct {
	
}

// SetDockTile parameters
type SetDockTileParams struct {
	
	
	BadgeLabel	string	`json:"badgeLabel"`
	
	// Png encoded image.
	Image	string	`json:"image"`
	
}

// SetDockTile returns
type SetDockTileReturns struct {
	
}

