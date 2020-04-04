package browser
// 
type WindowID int

// The state of the browser window.
type WindowState string

// Browser window bounds information
type Bounds  struct {

	// The offset from the left edge of the screen to the window in pixels.
	Left	int	`json:"left"`

	// The offset from the top edge of the screen to the window in pixels.
	Top	int	`json:"top"`

	// The window width in pixels.
	Width	int	`json:"width"`

	// The window height in pixels.
	Height	int	`json:"height"`

	// The window state. Default to normal.
	WindowState	WindowState	`json:"windowState"`
}

// 
type PermissionType string

// 
type PermissionSetting string

// Definition of PermissionDescriptor defined in the Permissions API:
	// https://w3c.github.io/permissions/#dictdef-permissiondescriptor.
type PermissionDescriptor  struct {

	// Name of permission.
	// See https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Name	string	`json:"name"`

	// For "midi" permission, may also specify sysex control.
	Sysex	bool	`json:"sysex"`

	// For "push" permission, may specify userVisibleOnly.
	// Note that userVisibleOnly = true is the only currently supported type.
	UserVisibleOnly	bool	`json:"userVisibleOnly"`

	// For "wake-lock" permission, must specify type as either "screen" or "system".
	Type	string	`json:"type"`
}

// Chrome histogram bucket.
type Bucket  struct {

	// Minimum value (inclusive).
	Low	int	`json:"low"`

	// Maximum value (exclusive).
	High	int	`json:"high"`

	// Number of samples.
	Count	int	`json:"count"`
}

// Chrome histogram.
type Histogram  struct {

	// Name.
	Name	string	`json:"name"`

	// Sum of sample values.
	Sum	int	`json:"sum"`

	// Total number of samples.
	Count	int	`json:"count"`

	// Buckets.
	Buckets	[]*Bucket	`json:"buckets"`
}
