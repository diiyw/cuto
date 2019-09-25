package emulation

import (

	"github.com/diiyw/goc/protocol/dom"

	"github.com/diiyw/goc/protocol/frame"

	"github.com/diiyw/goc/protocol/network"

)
const (
	
	// Tells whether emulation is supported.
	CanEmulate = "Emulation.canEmulate"
	
	// Clears the overriden device metrics.
	ClearDeviceMetricsOverride = "Emulation.clearDeviceMetricsOverride"
	
	// Clears the overriden Geolocation Position and Error.
	ClearGeolocationOverride = "Emulation.clearGeolocationOverride"
	
	// Requests that page scale factor is reset to initial values.
	ResetPageScaleFactor = "Emulation.resetPageScaleFactor"
	
	// Enables or disables simulating a focused and active page.
	SetFocusEmulationEnabled = "Emulation.setFocusEmulationEnabled"
	
	// Enables CPU throttling to emulate slow CPUs.
	SetCPUThrottlingRate = "Emulation.setCPUThrottlingRate"
	
	// Sets or clears an override of the default background color of the frame. This override is used
	// if the content does not specify one.
	SetDefaultBackgroundColorOverride = "Emulation.setDefaultBackgroundColorOverride"
	
	// Overrides the values of device screen dimensions (window.screen.width, window.screen.height,
	// window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
	// query results).
	SetDeviceMetricsOverride = "Emulation.setDeviceMetricsOverride"
	
	
	SetScrollbarsHidden = "Emulation.setScrollbarsHidden"
	
	
	SetDocumentCookieDisabled = "Emulation.setDocumentCookieDisabled"
	
	
	SetEmitTouchEventsForMouse = "Emulation.setEmitTouchEventsForMouse"
	
	// Emulates the given media for CSS media queries.
	SetEmulatedMedia = "Emulation.setEmulatedMedia"
	
	// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position
	// unavailable.
	SetGeolocationOverride = "Emulation.setGeolocationOverride"
	
	// Overrides value returned by the javascript navigator object.
	SetNavigatorOverrides = "Emulation.setNavigatorOverrides"
	
	// Sets a specified page scale factor.
	SetPageScaleFactor = "Emulation.setPageScaleFactor"
	
	// Switches script execution in the page.
	SetScriptExecutionDisabled = "Emulation.setScriptExecutionDisabled"
	
	// Enables touch on platforms which do not support them.
	SetTouchEmulationEnabled = "Emulation.setTouchEmulationEnabled"
	
	// Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets
	// the current virtual time policy.  Note this supersedes any previous time budget.
	SetVirtualTimePolicy = "Emulation.setVirtualTimePolicy"
	
	// Overrides default host system timezone with the specified one.
	SetTimezoneOverride = "Emulation.setTimezoneOverride"
	
	// Resizes the frame/viewport of the page. Note that this does not affect the frame's container
	// (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported
	// on Android.
	SetVisibleSize = "Emulation.setVisibleSize"
	
	// Allows overriding user agent with the given string.
	SetUserAgentOverride = "Emulation.setUserAgentOverride"
	
)

// CanEmulate parameters
type CanEmulateParams struct {
	
}

// CanEmulate returns
type CanEmulateReturns struct {
	
	// True if emulation is supported.
	Result	bool	`json:"result"`
	
}

// ClearDeviceMetricsOverride parameters
type ClearDeviceMetricsOverrideParams struct {
	
}

// ClearDeviceMetricsOverride returns
type ClearDeviceMetricsOverrideReturns struct {
	
}

// ClearGeolocationOverride parameters
type ClearGeolocationOverrideParams struct {
	
}

// ClearGeolocationOverride returns
type ClearGeolocationOverrideReturns struct {
	
}

// ResetPageScaleFactor parameters
type ResetPageScaleFactorParams struct {
	
}

// ResetPageScaleFactor returns
type ResetPageScaleFactorReturns struct {
	
}

// SetFocusEmulationEnabled parameters
type SetFocusEmulationEnabledParams struct {
	
	// Whether to enable to disable focus emulation.
	Enabled	bool	`json:"enabled"`
	
}

// SetFocusEmulationEnabled returns
type SetFocusEmulationEnabledReturns struct {
	
}

// SetCPUThrottlingRate parameters
type SetCPUThrottlingRateParams struct {
	
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate	float64	`json:"rate"`
	
}

// SetCPUThrottlingRate returns
type SetCPUThrottlingRateReturns struct {
	
}

// SetDefaultBackgroundColorOverride parameters
type SetDefaultBackgroundColorOverrideParams struct {
	
	// RGBA of the default background color. If not specified, any existing override will be
	// cleared.
	Color	dom.RGBA	`json:"color"`
	
}

// SetDefaultBackgroundColorOverride returns
type SetDefaultBackgroundColorOverrideReturns struct {
	
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
	ScreenOrientation	ScreenOrientation	`json:"screenOrientation"`
	
	// If set, the visible area of the page will be overridden to this viewport. This viewport
	// change is not observed by the page, e.g. viewport-relative elements do not change positions.
	Viewport	frame.Viewport	`json:"viewport"`
	
}

// SetDeviceMetricsOverride returns
type SetDeviceMetricsOverrideReturns struct {
	
}

// SetScrollbarsHidden parameters
type SetScrollbarsHiddenParams struct {
	
	// Whether scrollbars should be always hidden.
	Hidden	bool	`json:"hidden"`
	
}

// SetScrollbarsHidden returns
type SetScrollbarsHiddenReturns struct {
	
}

// SetDocumentCookieDisabled parameters
type SetDocumentCookieDisabledParams struct {
	
	// Whether document.coookie API should be disabled.
	Disabled	bool	`json:"disabled"`
	
}

// SetDocumentCookieDisabled returns
type SetDocumentCookieDisabledReturns struct {
	
}

// SetEmitTouchEventsForMouse parameters
type SetEmitTouchEventsForMouseParams struct {
	
	// Whether touch emulation based on mouse input should be enabled.
	Enabled	bool	`json:"enabled"`
	
	// Touch/gesture events configuration. Default: current platform.
	Configuration	string	`json:"configuration"`
	
}

// SetEmitTouchEventsForMouse returns
type SetEmitTouchEventsForMouseReturns struct {
	
}

// SetEmulatedMedia parameters
type SetEmulatedMediaParams struct {
	
	// Media type to emulate. Empty string disables the override.
	Media	string	`json:"media"`
	
}

// SetEmulatedMedia returns
type SetEmulatedMediaReturns struct {
	
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

// SetNavigatorOverrides parameters
type SetNavigatorOverridesParams struct {
	
	// The platform navigator.platform should return.
	Platform	string	`json:"platform"`
	
}

// SetNavigatorOverrides returns
type SetNavigatorOverridesReturns struct {
	
}

// SetPageScaleFactor parameters
type SetPageScaleFactorParams struct {
	
	// Page scale factor.
	PageScaleFactor	float64	`json:"pageScaleFactor"`
	
}

// SetPageScaleFactor returns
type SetPageScaleFactorReturns struct {
	
}

// SetScriptExecutionDisabled parameters
type SetScriptExecutionDisabledParams struct {
	
	// Whether script execution should be disabled in the page.
	Value	bool	`json:"value"`
	
}

// SetScriptExecutionDisabled returns
type SetScriptExecutionDisabledReturns struct {
	
}

// SetTouchEmulationEnabled parameters
type SetTouchEmulationEnabledParams struct {
	
	// Whether the touch event emulation should be enabled.
	Enabled	bool	`json:"enabled"`
	
	// Maximum touch points supported. Defaults to one.
	MaxTouchPoints	int	`json:"maxTouchPoints"`
	
}

// SetTouchEmulationEnabled returns
type SetTouchEmulationEnabledReturns struct {
	
}

// SetVirtualTimePolicy parameters
type SetVirtualTimePolicyParams struct {
	
	
	Policy	VirtualTimePolicy	`json:"policy"`
	
	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a
	// virtualTimeBudgetExpired event is sent.
	Budget	float64	`json:"budget"`
	
	// If set this specifies the maximum number of tasks that can be run before virtual is forced
	// forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount	int	`json:"maxVirtualTimeTaskStarvationCount"`
	
	// If set the virtual time policy change should be deferred until any frame starts navigating.
	// Note any previous deferred policy change is superseded.
	WaitForNavigation	bool	`json:"waitForNavigation"`
	
	// If set, base::Time::Now will be overriden to initially return this value.
	InitialVirtualTime	network.TimeSinceEpoch	`json:"initialVirtualTime"`
	
}

// SetVirtualTimePolicy returns
type SetVirtualTimePolicyReturns struct {
	
	// Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
	VirtualTimeTicksBase	float64	`json:"virtualTimeTicksBase"`
	
}

// SetTimezoneOverride parameters
type SetTimezoneOverrideParams struct {
	
	// The timezone identifier. If empty, disables the override and
	// restores default host system timezone.
	TimezoneId	string	`json:"timezoneId"`
	
}

// SetTimezoneOverride returns
type SetTimezoneOverrideReturns struct {
	
}

// SetVisibleSize parameters
type SetVisibleSizeParams struct {
	
	// Frame width (DIP).
	Width	int	`json:"width"`
	
	// Frame height (DIP).
	Height	int	`json:"height"`
	
}

// SetVisibleSize returns
type SetVisibleSizeReturns struct {
	
}

// SetUserAgentOverride parameters
type SetUserAgentOverrideParams struct {
	
	// User agent to use.
	UserAgent	string	`json:"userAgent"`
	
	// Browser langugage to emulate.
	AcceptLanguage	string	`json:"acceptLanguage"`
	
	// The platform navigator.platform should return.
	Platform	string	`json:"platform"`
	
}

// SetUserAgentOverride returns
type SetUserAgentOverrideReturns struct {
	
}

