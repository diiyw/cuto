package overlay

import (
	"github.com/diiyw/cuto/protocol/runtime"
	"github.com/diiyw/cuto/protocol/dom"
	"github.com/diiyw/cuto/protocol/cdp"
)


// Disables domain notifications.
const Disable = "Overlay.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables domain notifications.
const Enable = "Overlay.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// For testing.
const GetHighlightObjectForTest = "Overlay.getHighlightObjectForTest"

type GetHighlightObjectForTestParams struct {

	// Id of the node to get highlight object for.
	NodeId 	dom.NodeId	`json:"nodeId"`

	// Whether to include distance info.
	IncludeDistance 	bool	`json:"includeDistance,omitempty"`

	// Whether to include style info.
	IncludeStyle 	bool	`json:"includeStyle,omitempty"`
}

type GetHighlightObjectForTestResult struct {

	// Highlight data for the node.
	Highlight 	interface{}	`json:"highlight"`
}

// Hides any highlight.
const HideHighlight = "Overlay.hideHighlight"

type HideHighlightParams struct {
}

type HideHighlightResult struct {

}

// Highlights owner element of the frame with given id.
const HighlightFrame = "Overlay.highlightFrame"

type HighlightFrameParams struct {

	// Identifier of the frame to highlight.
	FrameId 	cdp.FrameId	`json:"frameId"`

	// The content box highlight fill color (default: transparent).
	ContentColor 	cdp.RGBA	`json:"contentColor,omitempty"`

	// The content box highlight outline color (default: transparent).
	ContentOutlineColor 	cdp.RGBA	`json:"contentOutlineColor,omitempty"`
}

type HighlightFrameResult struct {

}

// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or
// objectId must be specified.
const HighlightNode = "Overlay.highlightNode"

type HighlightNodeParams struct {

	// A descriptor for the highlight appearance.
	HighlightConfig 	HighlightConfig	`json:"highlightConfig"`

	// Identifier of the node to highlight.
	NodeId 	dom.NodeId	`json:"nodeId,omitempty"`

	// Identifier of the backend node to highlight.
	BackendNodeId 	dom.BackendNodeId	`json:"backendNodeId,omitempty"`

	// JavaScript object id of the node to be highlighted.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId,omitempty"`

	// Selectors to highlight relevant nodes.
	Selector 	string	`json:"selector,omitempty"`
}

type HighlightNodeResult struct {

}

// Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
const HighlightQuad = "Overlay.highlightQuad"

type HighlightQuadParams struct {

	// Quad to highlight
	Quad 	dom.Quad	`json:"quad"`

	// The highlight fill color (default: transparent).
	Color 	cdp.RGBA	`json:"color,omitempty"`

	// The highlight outline color (default: transparent).
	OutlineColor 	cdp.RGBA	`json:"outlineColor,omitempty"`
}

type HighlightQuadResult struct {

}

// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
const HighlightRect = "Overlay.highlightRect"

type HighlightRectParams struct {

	// X coordinate
	X 	int	`json:"x"`

	// Y coordinate
	Y 	int	`json:"y"`

	// Rectangle width
	Width 	int	`json:"width"`

	// Rectangle height
	Height 	int	`json:"height"`

	// The highlight fill color (default: transparent).
	Color 	cdp.RGBA	`json:"color,omitempty"`

	// The highlight outline color (default: transparent).
	OutlineColor 	cdp.RGBA	`json:"outlineColor,omitempty"`
}

type HighlightRectResult struct {

}

// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted.
// Backend then generates 'inspectNodeRequested' event upon element selection.
const SetInspectMode = "Overlay.setInspectMode"

type SetInspectModeParams struct {

	// Set an inspection mode.
	Mode 	InspectMode	`json:"mode"`

	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if `enabled
	// == false`.
	HighlightConfig 	HighlightConfig	`json:"highlightConfig,omitempty"`
}

type SetInspectModeResult struct {

}

// Highlights owner element of all frames detected to be ads.
const SetShowAdHighlights = "Overlay.setShowAdHighlights"

type SetShowAdHighlightsParams struct {

	// True for showing ad highlights
	Show 	bool	`json:"show"`
}

type SetShowAdHighlightsResult struct {

}

// 
const SetPausedInDebuggerMessage = "Overlay.setPausedInDebuggerMessage"

type SetPausedInDebuggerMessageParams struct {

	// The message to display, also triggers resume and step over controls.
	Message 	string	`json:"message,omitempty"`
}

type SetPausedInDebuggerMessageResult struct {

}

// Requests that backend shows debug borders on layers
const SetShowDebugBorders = "Overlay.setShowDebugBorders"

type SetShowDebugBordersParams struct {

	// True for showing debug borders
	Show 	bool	`json:"show"`
}

type SetShowDebugBordersResult struct {

}

// Requests that backend shows the FPS counter
const SetShowFPSCounter = "Overlay.setShowFPSCounter"

type SetShowFPSCounterParams struct {

	// True for showing the FPS counter
	Show 	bool	`json:"show"`
}

type SetShowFPSCounterResult struct {

}

// Requests that backend shows paint rectangles
const SetShowPaintRects = "Overlay.setShowPaintRects"

type SetShowPaintRectsParams struct {

	// True for showing paint rectangles
	Result 	bool	`json:"result"`
}

type SetShowPaintRectsResult struct {

}

// Requests that backend shows layout shift regions
const SetShowLayoutShiftRegions = "Overlay.setShowLayoutShiftRegions"

type SetShowLayoutShiftRegionsParams struct {

	// True for showing layout shift regions
	Result 	bool	`json:"result"`
}

type SetShowLayoutShiftRegionsResult struct {

}

// Requests that backend shows scroll bottleneck rects
const SetShowScrollBottleneckRects = "Overlay.setShowScrollBottleneckRects"

type SetShowScrollBottleneckRectsParams struct {

	// True for showing scroll bottleneck rects
	Show 	bool	`json:"show"`
}

type SetShowScrollBottleneckRectsResult struct {

}

// Requests that backend shows hit-test borders on layers
const SetShowHitTestBorders = "Overlay.setShowHitTestBorders"

type SetShowHitTestBordersParams struct {

	// True for showing hit-test borders
	Show 	bool	`json:"show"`
}

type SetShowHitTestBordersResult struct {

}

// Paints viewport size upon main frame resize.
const SetShowViewportSizeOnResize = "Overlay.setShowViewportSizeOnResize"

type SetShowViewportSizeOnResizeParams struct {

	// Whether to paint size or not.
	Show 	bool	`json:"show"`
}

type SetShowViewportSizeOnResizeResult struct {

}