package overlay

import (
	"github.com/diiyw/cuto/protocol/dom"
	"github.com/diiyw/cuto/protocol/cdp"
)


// Fired when the node should be inspected. This happens after call to `setInspectMode` or when
// user manually inspects an element.
const InspectNodeRequestedEvent = "Overlay.inspectNodeRequested"
type InspectNodeRequestedParams struct {

	// Id of the node to inspect.
	BackendNodeId 	dom.BackendNodeId}



// Fired when the node should be highlighted. This happens after call to `setInspectMode`.
const NodeHighlightRequestedEvent = "Overlay.nodeHighlightRequested"
type NodeHighlightRequestedParams struct {

	// 
	NodeId 	dom.NodeId}



// Fired when user asks to capture screenshot of some area on the page.
const ScreenshotRequestedEvent = "Overlay.screenshotRequested"
type ScreenshotRequestedParams struct {

	// Viewport to capture, in device independent pixels (dip).
	Viewport 	cdp.Viewport}



// Fired when user cancels the inspect mode.
const InspectModeCanceledEvent = "Overlay.inspectModeCanceled"
type InspectModeCanceledParams struct {
}

