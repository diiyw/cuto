package overlay

import (

	"github.com/diiyw/goc/protocol/dom"

	"github.com/diiyw/goc/protocol/frame"

)
const (
	
	// Fired when the node should be inspected. This happens after call to `setInspectMode` or when
	// user manually inspects an element.
	InspectNodeRequestedEvent = "Overlay.inspectNodeRequested"
	
	// Fired when the node should be highlighted. This happens after call to `setInspectMode`.
	NodeHighlightRequestedEvent = "Overlay.nodeHighlightRequested"
	
	// Fired when user asks to capture screenshot of some area on the page.
	ScreenshotRequestedEvent = "Overlay.screenshotRequested"
	
)

// Fired when the node should be inspected. This happens after call to `setInspectMode` or when
	// user manually inspects an element.
type InspectNodeRequestedParams struct {
	
	// Id of the node to inspect.
	BackendNodeId	dom.BackendNodeId	`json:"backendNodeId"`
	
}

// Fired when the node should be highlighted. This happens after call to `setInspectMode`.
type NodeHighlightRequestedParams struct {
	
	
	NodeId	dom.NodeId	`json:"nodeId"`
	
}

// Fired when user asks to capture screenshot of some area on the page.
type ScreenshotRequestedParams struct {
	
	// Viewport to capture, in CSS.
	Viewport	frame.Viewport	`json:"viewport"`
	
}

