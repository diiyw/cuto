package browser

import (
	"github.com/diiyw/cuto/protocol/target"
)


// Set permission settings for given origin.
const SetPermission = "Browser.setPermission"

type SetPermissionParams struct {

	// Origin the permission applies to.
	Origin 	string	`json:"origin"`

	// Descriptor of permission to override.
	Permission 	PermissionDescriptor	`json:"permission"`

	// Setting of the permission.
	Setting 	PermissionSetting	`json:"setting"`

	// Context to override. When omitted, default browser context is used.
	BrowserContextId 	target.TargetID	`json:"browserContextId,omitempty"`
}

type SetPermissionResult struct {

}

// Grant specific permissions to the given origin and reject all others.
const GrantPermissions = "Browser.grantPermissions"

type GrantPermissionsParams struct {

	// 
	Origin 	string	`json:"origin"`

	// 
	Permissions 	[]*PermissionType	`json:"permissions"`

	// BrowserContext to override permissions. When omitted, default browser context is used.
	BrowserContextId 	target.BrowserContextID	`json:"browserContextId,omitempty"`
}

type GrantPermissionsResult struct {

}

// Reset all permission management for all origins.
const ResetPermissions = "Browser.resetPermissions"

type ResetPermissionsParams struct {

	// BrowserContext to reset permissions. When omitted, default browser context is used.
	BrowserContextId 	target.BrowserContextID	`json:"browserContextId,omitempty"`
}

type ResetPermissionsResult struct {

}

// Close browser gracefully.
const Close = "Browser.close"

type CloseParams struct {
}

type CloseResult struct {

}

// Crashes browser on the main thread.
const Crash = "Browser.crash"

type CrashParams struct {
}

type CrashResult struct {

}

// Crashes GPU process.
const CrashGpuProcess = "Browser.crashGpuProcess"

type CrashGpuProcessParams struct {
}

type CrashGpuProcessResult struct {

}

// Returns version information.
const GetVersion = "Browser.getVersion"

type GetVersionParams struct {
}

type GetVersionResult struct {

	// Protocol version.
	ProtocolVersion 	string	`json:"protocolVersion"`
	// Product name.
	Product 	string	`json:"product"`
	// Product revision.
	Revision 	string	`json:"revision"`
	// User-Agent.
	UserAgent 	string	`json:"userAgent"`
	// V8 version.
	JsVersion 	string	`json:"jsVersion"`
}

// Returns the command line switches for the browser process if, and only if
// --enable-automation is on the commandline.
const GetBrowserCommandLine = "Browser.getBrowserCommandLine"

type GetBrowserCommandLineParams struct {
}

type GetBrowserCommandLineResult struct {

	// Commandline parameters
	Arguments 	[]string	`json:"arguments"`
}

// Get Chrome histograms.
const GetHistograms = "Browser.getHistograms"

type GetHistogramsParams struct {

	// Requested substring in name. Only histograms which have query as a
	// substring in their name are extracted. An empty or absent query returns
	// all histograms.
	Query 	string	`json:"query,omitempty"`

	// If true, retrieve delta since last call.
	Delta 	bool	`json:"delta,omitempty"`
}

type GetHistogramsResult struct {

	// Histograms.
	Histograms 	[]*Histogram	`json:"histograms"`
}

// Get a Chrome histogram by name.
const GetHistogram = "Browser.getHistogram"

type GetHistogramParams struct {

	// Requested histogram name.
	Name 	string	`json:"name"`

	// If true, retrieve delta since last call.
	Delta 	bool	`json:"delta,omitempty"`
}

type GetHistogramResult struct {

	// Histogram.
	Histogram 	Histogram	`json:"histogram"`
}

// Get position and size of the browser window.
const GetWindowBounds = "Browser.getWindowBounds"

type GetWindowBoundsParams struct {

	// Browser window id.
	WindowId 	WindowID	`json:"windowId"`
}

type GetWindowBoundsResult struct {

	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds 	Bounds	`json:"bounds"`
}

// Get the browser window that contains the devtools target.
const GetWindowForTarget = "Browser.getWindowForTarget"

type GetWindowForTargetParams struct {

	// Devtools agent host id. If called as a part of the session, associated targetId is used.
	TargetId 	target.TargetID	`json:"targetId,omitempty"`
}

type GetWindowForTargetResult struct {

	// Browser window id.
	WindowId 	WindowID	`json:"windowId"`
	// Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds 	Bounds	`json:"bounds"`
}

// Set position and/or size of the browser window.
const SetWindowBounds = "Browser.setWindowBounds"

type SetWindowBoundsParams struct {

	// Browser window id.
	WindowId 	WindowID	`json:"windowId"`

	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds 	Bounds	`json:"bounds"`
}

type SetWindowBoundsResult struct {

}

// Set dock tile details, platform-specific.
const SetDockTile = "Browser.setDockTile"

type SetDockTileParams struct {

	// 
	BadgeLabel 	string	`json:"badgeLabel,omitempty"`

	// Png encoded image.
	Image 	[]byte	`json:"image,omitempty"`
}

type SetDockTileResult struct {

}