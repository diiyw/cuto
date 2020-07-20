package emulation

import (
	"github.com/diiyw/cuto/protocol/cdp"
)


// Tells whether emulation is supported.
const CanEmulate = "Emulation.canEmulate"

type CanEmulateParams struct {
}

type CanEmulateResult struct {

	// True if emulation is supported.
	Result 	bool	`json:"result"`
}

// Clears the overriden device metrics.
const ClearDeviceMetricsOverride = "Emulation.clearDeviceMetricsOverride"

type ClearDeviceMetricsOverrideParams struct {
}

type ClearDeviceMetricsOverrideResult struct {

}

// Clears the overriden Geolocation Position and Error.
const ClearGeolocationOverride = "Emulation.clearGeolocationOverride"

type ClearGeolocationOverrideParams struct {
}

type ClearGeolocationOverrideResult struct {

}

// Requests that page scale factor is reset to initial values.
const ResetPageScaleFactor = "Emulation.resetPageScaleFactor"

type ResetPageScaleFactorParams struct {
}

type ResetPageScaleFactorResult struct {

}

// Enables or disables simulating a focused and active page.
const SetFocusEmulationEnabled = "Emulation.setFocusEmulationEnabled"

type SetFocusEmulationEnabledParams struct {

	// Whether to enable to disable focus emulation.
	Enabled 	bool	`json:"enabled"`
}

type SetFocusEmulationEnabledResult struct {

}

// Enables CPU throttling to emulate slow CPUs.
const SetCPUThrottlingRate = "Emulation.setCPUThrottlingRate"

type SetCPUThrottlingRateParams struct {

	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate 	float64	`json:"rate"`
}

type SetCPUThrottlingRateResult struct {

}

// Sets or clears an override of the default background color of the frame. This override is used
// if the content does not specify one.
const SetDefaultBackgroundColorOverride = "Emulation.setDefaultBackgroundColorOverride"

type SetDefaultBackgroundColorOverrideParams struct {

	// RGBA of the default background color. If not specified, any existing override will be
	// cleared.
	Color 	cdp.RGBA	`json:"color,omitempty"`
}

type SetDefaultBackgroundColorOverrideResult struct {

}

// Overrides the values of device screen dimensions (window.screen.width, window.screen.height,
// window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
const SetDeviceMetricsOverride = "Emulation.setDeviceMetricsOverride"

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
	Scale 	float64	`json:"scale,omitempty"`

	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenWidth 	int	`json:"screenWidth,omitempty"`

	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	ScreenHeight 	int	`json:"screenHeight,omitempty"`

	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionX 	int	`json:"positionX,omitempty"`

	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	PositionY 	int	`json:"positionY,omitempty"`

	// Do not set visible view size, rely upon explicit setVisibleSize call.
	DontSetVisibleSize 	bool	`json:"dontSetVisibleSize,omitempty"`

	// Screen orientation override.
	ScreenOrientation 	ScreenOrientation	`json:"screenOrientation,omitempty"`

	// If set, the visible area of the page will be overridden to this viewport. This viewport
	// change is not observed by the page, e.g. viewport-relative elements do not change positions.
	Viewport 	cdp.Viewport	`json:"viewport,omitempty"`
}

type SetDeviceMetricsOverrideResult struct {

}

// 
const SetScrollbarsHidden = "Emulation.setScrollbarsHidden"

type SetScrollbarsHiddenParams struct {

	// Whether scrollbars should be always hidden.
	Hidden 	bool	`json:"hidden"`
}

type SetScrollbarsHiddenResult struct {

}

// 
const SetDocumentCookieDisabled = "Emulation.setDocumentCookieDisabled"

type SetDocumentCookieDisabledParams struct {

	// Whether document.coookie API should be disabled.
	Disabled 	bool	`json:"disabled"`
}

type SetDocumentCookieDisabledResult struct {

}

// 
const SetEmitTouchEventsForMouse = "Emulation.setEmitTouchEventsForMouse"

type SetEmitTouchEventsForMouseParams struct {

	// Whether touch emulation based on mouse input should be enabled.
	Enabled 	bool	`json:"enabled"`

	// Touch/gesture events configuration. Default: current platform.
	Configuration 	string	`json:"configuration,omitempty"`
}

type SetEmitTouchEventsForMouseResult struct {

}

// Emulates the given media type or media feature for CSS media queries.
const SetEmulatedMedia = "Emulation.setEmulatedMedia"

type SetEmulatedMediaParams struct {

	// Media type to emulate. Empty string disables the override.
	Media 	string	`json:"media,omitempty"`

	// Media features to emulate.
	Features 	[]*MediaFeature	`json:"features,omitempty"`
}

type SetEmulatedMediaResult struct {

}

// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position
// unavailable.
const SetGeolocationOverride = "Emulation.setGeolocationOverride"

type SetGeolocationOverrideParams struct {

	// Mock latitude
	Latitude 	float64	`json:"latitude,omitempty"`

	// Mock longitude
	Longitude 	float64	`json:"longitude,omitempty"`

	// Mock accuracy
	Accuracy 	float64	`json:"accuracy,omitempty"`
}

type SetGeolocationOverrideResult struct {

}

// Overrides value returned by the javascript navigator object.
const SetNavigatorOverrides = "Emulation.setNavigatorOverrides"

type SetNavigatorOverridesParams struct {

	// The platform navigator.platform should return.
	Platform 	string	`json:"platform"`
}

type SetNavigatorOverridesResult struct {

}

// Sets a specified page scale factor.
const SetPageScaleFactor = "Emulation.setPageScaleFactor"

type SetPageScaleFactorParams struct {

	// Page scale factor.
	PageScaleFactor 	float64	`json:"pageScaleFactor"`
}

type SetPageScaleFactorResult struct {

}

// Switches script execution in the page.
const SetScriptExecutionDisabled = "Emulation.setScriptExecutionDisabled"

type SetScriptExecutionDisabledParams struct {

	// Whether script execution should be disabled in the page.
	Value 	bool	`json:"value"`
}

type SetScriptExecutionDisabledResult struct {

}

// Enables touch on platforms which do not support them.
const SetTouchEmulationEnabled = "Emulation.setTouchEmulationEnabled"

type SetTouchEmulationEnabledParams struct {

	// Whether the touch event emulation should be enabled.
	Enabled 	bool	`json:"enabled"`

	// Maximum touch points supported. Defaults to one.
	MaxTouchPoints 	int	`json:"maxTouchPoints,omitempty"`
}

type SetTouchEmulationEnabledResult struct {

}

// Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets
// the current virtual time policy.  Note this supersedes any previous time budget.
const SetVirtualTimePolicy = "Emulation.setVirtualTimePolicy"

type SetVirtualTimePolicyParams struct {

	// 
	Policy 	VirtualTimePolicy	`json:"policy"`

	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a
	// virtualTimeBudgetExpired event is sent.
	Budget 	float64	`json:"budget,omitempty"`

	// If set this specifies the maximum number of tasks that can be run before virtual is forced
	// forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount 	int	`json:"maxVirtualTimeTaskStarvationCount,omitempty"`

	// If set the virtual time policy change should be deferred until any frame starts navigating.
	// Note any previous deferred policy change is superseded.
	WaitForNavigation 	bool	`json:"waitForNavigation,omitempty"`

	// If set, base::Time::Now will be overriden to initially return this value.
	InitialVirtualTime 	cdp.TimeSinceEpoch	`json:"initialVirtualTime,omitempty"`
}

type SetVirtualTimePolicyResult struct {

	// Absolute timestamp at which virtual time was first enabled (up time in milliseconds).
	VirtualTimeTicksBase 	float64	`json:"virtualTimeTicksBase"`
}

// Overrides default host system timezone with the specified one.
const SetTimezoneOverride = "Emulation.setTimezoneOverride"

type SetTimezoneOverrideParams struct {

	// The timezone identifier. If empty, disables the override and
	// restores default host system timezone.
	TimezoneId 	string	`json:"timezoneId"`
}

type SetTimezoneOverrideResult struct {

}

// Resizes the frame/viewport of the page. Note that this does not affect the frame's container
// (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported
// on Android.
const SetVisibleSize = "Emulation.setVisibleSize"

type SetVisibleSizeParams struct {

	// Frame width (DIP).
	Width 	int	`json:"width"`

	// Frame height (DIP).
	Height 	int	`json:"height"`
}

type SetVisibleSizeResult struct {

}

// Allows overriding user agent with the given string.
const SetUserAgentOverride = "Emulation.setUserAgentOverride"

type SetUserAgentOverrideParams struct {

	// User agent to use.
	UserAgent 	string	`json:"userAgent"`

	// Browser langugage to emulate.
	AcceptLanguage 	string	`json:"acceptLanguage,omitempty"`

	// The platform navigator.platform should return.
	Platform 	string	`json:"platform,omitempty"`
}

type SetUserAgentOverrideResult struct {

}