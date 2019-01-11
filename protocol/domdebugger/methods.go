package domdebugger

import (

	"github.com/diiyw/goc/protocol/dom"

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	// Returns event listeners of the given object.
	GetEventListeners = "DOMDebugger.getEventListeners"
	
	// Removes DOM breakpoint that was set using `setDOMBreakpoint`.
	RemoveDOMBreakpoint = "DOMDebugger.removeDOMBreakpoint"
	
	// Removes breakpoint on particular DOM event.
	RemoveEventListenerBreakpoint = "DOMDebugger.removeEventListenerBreakpoint"
	
	// Removes breakpoint on particular native event.
	RemoveInstrumentationBreakpoint = "DOMDebugger.removeInstrumentationBreakpoint"
	
	// Removes breakpoint from XMLHttpRequest.
	RemoveXHRBreakpoint = "DOMDebugger.removeXHRBreakpoint"
	
	// Sets breakpoint on particular operation with DOM.
	SetDOMBreakpoint = "DOMDebugger.setDOMBreakpoint"
	
	// Sets breakpoint on particular DOM event.
	SetEventListenerBreakpoint = "DOMDebugger.setEventListenerBreakpoint"
	
	// Sets breakpoint on particular native event.
	SetInstrumentationBreakpoint = "DOMDebugger.setInstrumentationBreakpoint"
	
	// Sets breakpoint on XMLHttpRequest.
	SetXHRBreakpoint = "DOMDebugger.setXHRBreakpoint"
	
)

// GetEventListeners parameters
type GetEventListenersParams struct {
	
	// Identifier of the object to return listeners for.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
	// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth	int	`json:"depth"`
	
	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false). Reports listeners for all contexts if pierce is enabled.
	Pierce	bool	`json:"pierce"`
	
}

// GetEventListeners returns
type GetEventListenersReturns struct {
	
	// Array of relevant listeners.
	Listeners	[]EventListener	`json:"listeners"`
	
}

// RemoveDOMBreakpoint parameters
type RemoveDOMBreakpointParams struct {
	
	// Identifier of the node to remove breakpoint from.
	NodeId	dom.NodeId	`json:"nodeId"`
	
	// Type of the breakpoint to remove.
	Type	DOMBreakpointType	`json:"type"`
	
}

// RemoveDOMBreakpoint returns
type RemoveDOMBreakpointReturns struct {
	
}

// RemoveEventListenerBreakpoint parameters
type RemoveEventListenerBreakpointParams struct {
	
	// Event name.
	EventName	string	`json:"eventName"`
	
	// EventTarget interface name.
	TargetName	string	`json:"targetName"`
	
}

// RemoveEventListenerBreakpoint returns
type RemoveEventListenerBreakpointReturns struct {
	
}

// RemoveInstrumentationBreakpoint parameters
type RemoveInstrumentationBreakpointParams struct {
	
	// Instrumentation name to stop on.
	EventName	string	`json:"eventName"`
	
}

// RemoveInstrumentationBreakpoint returns
type RemoveInstrumentationBreakpointReturns struct {
	
}

// RemoveXHRBreakpoint parameters
type RemoveXHRBreakpointParams struct {
	
	// Resource URL substring.
	Url	string	`json:"url"`
	
}

// RemoveXHRBreakpoint returns
type RemoveXHRBreakpointReturns struct {
	
}

// SetDOMBreakpoint parameters
type SetDOMBreakpointParams struct {
	
	// Identifier of the node to set breakpoint on.
	NodeId	dom.NodeId	`json:"nodeId"`
	
	// Type of the operation to stop upon.
	Type	DOMBreakpointType	`json:"type"`
	
}

// SetDOMBreakpoint returns
type SetDOMBreakpointReturns struct {
	
}

// SetEventListenerBreakpoint parameters
type SetEventListenerBreakpointParams struct {
	
	// DOM Event name to stop on (any DOM event will do).
	EventName	string	`json:"eventName"`
	
	// EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any
	// EventTarget.
	TargetName	string	`json:"targetName"`
	
}

// SetEventListenerBreakpoint returns
type SetEventListenerBreakpointReturns struct {
	
}

// SetInstrumentationBreakpoint parameters
type SetInstrumentationBreakpointParams struct {
	
	// Instrumentation name to stop on.
	EventName	string	`json:"eventName"`
	
}

// SetInstrumentationBreakpoint returns
type SetInstrumentationBreakpointReturns struct {
	
}

// SetXHRBreakpoint parameters
type SetXHRBreakpointParams struct {
	
	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	Url	string	`json:"url"`
	
}

// SetXHRBreakpoint returns
type SetXHRBreakpointReturns struct {
	
}

