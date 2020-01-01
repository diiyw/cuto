package overlay

import (

	"github.com/diiyw/chr/protocol/dom"

	"github.com/diiyw/chr/protocol/frame"

	"github.com/diiyw/chr/protocol/runtime"

)
const (
	
	// Disables domain notifications.
	Disable = "Overlay.disable"
	
	// Enables domain notifications.
	Enable = "Overlay.enable"
	
	// For testing.
	GetHighlightObjectForTest = "Overlay.getHighlightObjectForTest"
	
	// Hides any highlight.
	HideHighlight = "Overlay.hideHighlight"
	
	// Highlights owner element of the frame with given id.
	HighlightFrame = "Overlay.highlightFrame"
	
	// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or
	// objectId must be specified.
	HighlightNode = "Overlay.highlightNode"
	
	// Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
	HighlightQuad = "Overlay.highlightQuad"
	
	// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
	HighlightRect = "Overlay.highlightRect"
	
	// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted.
	// Backend then generates 'inspectNodeRequested' event upon element selection.
	SetInspectMode = "Overlay.setInspectMode"
	
	// Highlights owner element of all frames detected to be ads.
	SetShowAdHighlights = "Overlay.setShowAdHighlights"
	
	
	SetPausedInDebuggerMessage = "Overlay.setPausedInDebuggerMessage"
	
	// Requests that backend shows debug borders on layers
	SetShowDebugBorders = "Overlay.setShowDebugBorders"
	
	// Requests that backend shows the FPS counter
	SetShowFPSCounter = "Overlay.setShowFPSCounter"
	
	// Requests that backend shows paint rectangles
	SetShowPaintRects = "Overlay.setShowPaintRects"
	
	// Requests that backend shows scroll bottleneck rects
	SetShowScrollBottleneckRects = "Overlay.setShowScrollBottleneckRects"
	
	// Requests that backend shows hit-test borders on layers
	SetShowHitTestBorders = "Overlay.setShowHitTestBorders"
	
	// Paints viewport size upon main frame resize.
	SetShowViewportSizeOnResize = "Overlay.setShowViewportSizeOnResize"
	
)

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

// GetHighlightObjectForTest parameters
type GetHighlightObjectForTestParams struct {
	
	// Id of the node to get highlight object for.
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// GetHighlightObjectForTest returns
type GetHighlightObjectForTestReturns struct {
	
	// Highlight data for the node.
	Highlight	interface{}	`json:"highlight"`
	
}

// HideHighlight parameters
type HideHighlightParams struct {
	
}

// HideHighlight returns
type HideHighlightReturns struct {
	
}

// HighlightFrame parameters
type HighlightFrameParams struct {
	
	// Identifier of the frame to highlight.
	FrameId	frame.FrameId	`json:"frameId"`
	
	// The content box highlight fill color (default: transparent).
	ContentColor	dom.RGBA	`json:"contentColor"`
	
	// The content box highlight outline color (default: transparent).
	ContentOutlineColor	dom.RGBA	`json:"contentOutlineColor"`
	
}

// HighlightFrame returns
type HighlightFrameReturns struct {
	
}

// HighlightNode parameters
type HighlightNodeParams struct {
	
	// A descriptor for the highlight appearance.
	HighlightConfig	HighlightConfig	`json:"highlightConfig"`
	
	// Identifier of the node to highlight.
	NodeId	dom.NodeId	`json:"nodeId"`
	
	// Identifier of the backend node to highlight.
	BackendNodeId	dom.BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node to be highlighted.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
	// Selectors to highlight relevant nodes.
	Selector	string	`json:"selector"`
	
}

// HighlightNode returns
type HighlightNodeReturns struct {
	
}

// HighlightQuad parameters
type HighlightQuadParams struct {
	
	// Quad to highlight
	Quad	dom.Quad	`json:"quad"`
	
	// The highlight fill color (default: transparent).
	Color	dom.RGBA	`json:"color"`
	
	// The highlight outline color (default: transparent).
	OutlineColor	dom.RGBA	`json:"outlineColor"`
	
}

// HighlightQuad returns
type HighlightQuadReturns struct {
	
}

// HighlightRect parameters
type HighlightRectParams struct {
	
	// X coordinate
	X	int	`json:"x"`
	
	// Y coordinate
	Y	int	`json:"y"`
	
	// Rectangle width
	Width	int	`json:"width"`
	
	// Rectangle height
	Height	int	`json:"height"`
	
	// The highlight fill color (default: transparent).
	Color	dom.RGBA	`json:"color"`
	
	// The highlight outline color (default: transparent).
	OutlineColor	dom.RGBA	`json:"outlineColor"`
	
}

// HighlightRect returns
type HighlightRectReturns struct {
	
}

// SetInspectMode parameters
type SetInspectModeParams struct {
	
	// Set an inspection mode.
	Mode	InspectMode	`json:"mode"`
	
	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled
	// == false`.
	HighlightConfig	HighlightConfig	`json:"highlightConfig,omitempty"`
	
}

// SetInspectMode returns
type SetInspectModeReturns struct {
	
}

// SetShowAdHighlights parameters
type SetShowAdHighlightsParams struct {
	
	// True for showing ad highlights
	Show	bool	`json:"show"`
	
}

// SetShowAdHighlights returns
type SetShowAdHighlightsReturns struct {
	
}

// SetPausedInDebuggerMessage parameters
type SetPausedInDebuggerMessageParams struct {
	
	// The message to display, also triggers resume and step over controls.
	Message	string	`json:"message"`
	
}

// SetPausedInDebuggerMessage returns
type SetPausedInDebuggerMessageReturns struct {
	
}

// SetShowDebugBorders parameters
type SetShowDebugBordersParams struct {
	
	// True for showing debug borders
	Show	bool	`json:"show"`
	
}

// SetShowDebugBorders returns
type SetShowDebugBordersReturns struct {
	
}

// SetShowFPSCounter parameters
type SetShowFPSCounterParams struct {
	
	// True for showing the FPS counter
	Show	bool	`json:"show"`
	
}

// SetShowFPSCounter returns
type SetShowFPSCounterReturns struct {
	
}

// SetShowPaintRects parameters
type SetShowPaintRectsParams struct {
	
	// True for showing paint rectangles
	Result	bool	`json:"result"`
	
}

// SetShowPaintRects returns
type SetShowPaintRectsReturns struct {
	
}

// SetShowScrollBottleneckRects parameters
type SetShowScrollBottleneckRectsParams struct {
	
	// True for showing scroll bottleneck rects
	Show	bool	`json:"show"`
	
}

// SetShowScrollBottleneckRects returns
type SetShowScrollBottleneckRectsReturns struct {
	
}

// SetShowHitTestBorders parameters
type SetShowHitTestBordersParams struct {
	
	// True for showing hit-test borders
	Show	bool	`json:"show"`
	
}

// SetShowHitTestBorders returns
type SetShowHitTestBordersReturns struct {
	
}

// SetShowViewportSizeOnResize parameters
type SetShowViewportSizeOnResizeParams struct {
	
	// Whether to paint size or not.
	Show	bool	`json:"show"`
	
}

// SetShowViewportSizeOnResize returns
type SetShowViewportSizeOnResizeReturns struct {
	
}

