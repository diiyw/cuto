package page

import (
	"github.com/diiyw/cuto/protocol/network"
	"github.com/diiyw/cuto/protocol/cdp"
)

// Unique frame identifier.
type FrameId string

// Information about the Frame on the page.
type Frame  struct {

	// Frame unique identifier.
	Id	FrameId	`json:"id"`

	// Parent frame identifier.
	ParentId	string	`json:"parentId,omitempty"`

	// Identifier of the loader associated with this frame.
	LoaderId	network.LoaderId	`json:"loaderId"`

	// Frame's name as specified in the tag.
	Name	string	`json:"name,omitempty"`

	// Frame document's URL without fragment.
	Url	string	`json:"url"`

	// Frame document's URL fragment including the '#'.
	UrlFragment	string	`json:"urlFragment,omitempty"`

	// Frame document's security origin.
	SecurityOrigin	string	`json:"securityOrigin"`

	// Frame document's mimeType as determined by the browser.
	MimeType	string	`json:"mimeType"`

	// If the frame failed to load, this contains the URL that could not be loaded. Note that unlike url above, this URL may contain a fragment.
	UnreachableUrl	string	`json:"unreachableUrl,omitempty"`
}

// Information about the Resource on the page.
type FrameResource  struct {

	// Resource URL.
	Url	string	`json:"url"`

	// Type of this resource.
	Type	network.ResourceType	`json:"type"`

	// Resource mimeType as determined by the browser.
	MimeType	string	`json:"mimeType"`

	// last-modified timestamp as reported by server.
	LastModified	cdp.TimeSinceEpoch	`json:"lastModified,omitempty"`

	// Resource content size.
	ContentSize	float64	`json:"contentSize,omitempty"`

	// True if the resource failed to load.
	Failed	bool	`json:"failed,omitempty"`

	// True if the resource was canceled during loading.
	Canceled	bool	`json:"canceled,omitempty"`
}

// Information about the Frame hierarchy along with their cached resources.
type FrameResourceTree  struct {

	// Frame information for this tree item.
	Frame	Frame	`json:"frame"`

	// Child frames.
	ChildFrames	[]*FrameResourceTree	`json:"childFrames,omitempty"`

	// Information about frame resources.
	Resources	[]*FrameResource	`json:"resources"`
}

// Information about the Frame hierarchy.
type FrameTree  struct {

	// Frame information for this tree item.
	Frame	Frame	`json:"frame"`

	// Child frames.
	ChildFrames	[]*FrameTree	`json:"childFrames,omitempty"`
}

// Unique script identifier.
type ScriptIdentifier string

// Transition type.
type TransitionType string

// Navigation history entry.
type NavigationEntry  struct {

	// Unique id of the navigation history entry.
	Id	int	`json:"id"`

	// URL of the navigation history entry.
	Url	string	`json:"url"`

	// URL that the user typed in the url bar.
	UserTypedURL	string	`json:"userTypedURL"`

	// Title of the navigation history entry.
	Title	string	`json:"title"`

	// Transition type.
	TransitionType	TransitionType	`json:"transitionType"`
}

// Screencast frame metadata.
type ScreencastFrameMetadata  struct {

	// Top offset in DIP.
	OffsetTop	float64	`json:"offsetTop"`

	// Page scale factor.
	PageScaleFactor	float64	`json:"pageScaleFactor"`

	// Device screen width in DIP.
	DeviceWidth	float64	`json:"deviceWidth"`

	// Device screen height in DIP.
	DeviceHeight	float64	`json:"deviceHeight"`

	// Position of horizontal scroll in CSS pixels.
	ScrollOffsetX	float64	`json:"scrollOffsetX"`

	// Position of vertical scroll in CSS pixels.
	ScrollOffsetY	float64	`json:"scrollOffsetY"`

	// Frame swap timestamp.
	Timestamp	cdp.TimeSinceEpoch	`json:"timestamp,omitempty"`
}

// Javascript dialog type.
type DialogType string

// Error while paring app manifest.
type AppManifestError  struct {

	// Error message.
	Message	string	`json:"message"`

	// If criticial, this is a non-recoverable parse error.
	Critical	int	`json:"critical"`

	// Error line.
	Line	int	`json:"line"`

	// Error column.
	Column	int	`json:"column"`
}

// Layout viewport position and dimensions.
type LayoutViewport  struct {

	// Horizontal offset relative to the document (CSS pixels).
	PageX	int	`json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY	int	`json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth	int	`json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight	int	`json:"clientHeight"`
}

// Visual viewport position, dimensions, and scale.
type VisualViewport  struct {

	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX	float64	`json:"offsetX"`

	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY	float64	`json:"offsetY"`

	// Horizontal offset relative to the document (CSS pixels).
	PageX	float64	`json:"pageX"`

	// Vertical offset relative to the document (CSS pixels).
	PageY	float64	`json:"pageY"`

	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth	float64	`json:"clientWidth"`

	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight	float64	`json:"clientHeight"`

	// Scale relative to the ideal viewport (size at width=device-width).
	Scale	float64	`json:"scale"`

	// Page zoom factor (CSS to device independent pixels ratio).
	Zoom	float64	`json:"zoom,omitempty"`
}

// Viewport for capturing screenshot.
type Viewport  struct {

	// X offset in device independent pixels (dip).
	X	float64	`json:"x"`

	// Y offset in device independent pixels (dip).
	Y	float64	`json:"y"`

	// Rectangle width in device independent pixels (dip).
	Width	float64	`json:"width"`

	// Rectangle height in device independent pixels (dip).
	Height	float64	`json:"height"`

	// Page scale factor.
	Scale	float64	`json:"scale"`
}

// Generic font families collection.
type FontFamilies  struct {

	// The standard font-family.
	Standard	string	`json:"standard,omitempty"`

	// The fixed font-family.
	Fixed	string	`json:"fixed,omitempty"`

	// The serif font-family.
	Serif	string	`json:"serif,omitempty"`

	// The sansSerif font-family.
	SansSerif	string	`json:"sansSerif,omitempty"`

	// The cursive font-family.
	Cursive	string	`json:"cursive,omitempty"`

	// The fantasy font-family.
	Fantasy	string	`json:"fantasy,omitempty"`

	// The pictograph font-family.
	Pictograph	string	`json:"pictograph,omitempty"`
}

// Default font sizes.
type FontSizes  struct {

	// Default standard font size.
	Standard	int	`json:"standard,omitempty"`

	// Default fixed font size.
	Fixed	int	`json:"fixed,omitempty"`
}

// 
type ClientNavigationReason string
