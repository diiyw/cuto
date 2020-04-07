package domdebugger

// Returns event listeners of the given object.
const GetEventListeners = "DOMDebugger.getEventListeners"

type GetEventListenersParams struct {

	// Identifier of the object to return listeners for.
	ObjectId 	interface{}	`json:"objectId"`

	// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the
	// entire subtree or provide an integer larger than 0.
	Depth 	int	`json:"depth"`

	// Whether or not iframes and shadow roots should be traversed when returning the subtree
	// (default is false). Reports listeners for all contexts if pierce is enabled.
	Pierce 	bool	`json:"pierce"`
}

type GetEventListenersResult struct {

	// Array of relevant listeners.
	Listeners 	[]*EventListener	`json:"listeners"`
}

// Removes DOM breakpoint that was set using `setDOMBreakpoint`.
const RemoveDOMBreakpoint = "DOMDebugger.removeDOMBreakpoint"

type RemoveDOMBreakpointParams struct {

	// Identifier of the node to remove breakpoint from.
	NodeId 	interface{}	`json:"nodeId"`

	// Type of the breakpoint to remove.
	Type 	DOMBreakpointType	`json:"type"`
}

type RemoveDOMBreakpointResult struct {

}

// Removes breakpoint on particular DOM event.
const RemoveEventListenerBreakpoint = "DOMDebugger.removeEventListenerBreakpoint"

type RemoveEventListenerBreakpointParams struct {

	// Event name.
	EventName 	string	`json:"eventName"`

	// EventTarget interface name.
	TargetName 	string	`json:"targetName"`
}

type RemoveEventListenerBreakpointResult struct {

}

// Removes breakpoint on particular native event.
const RemoveInstrumentationBreakpoint = "DOMDebugger.removeInstrumentationBreakpoint"

type RemoveInstrumentationBreakpointParams struct {

	// Instrumentation name to stop on.
	EventName 	string	`json:"eventName"`
}

type RemoveInstrumentationBreakpointResult struct {

}

// Removes breakpoint from XMLHttpRequest.
const RemoveXHRBreakpoint = "DOMDebugger.removeXHRBreakpoint"

type RemoveXHRBreakpointParams struct {

	// Resource URL substring.
	Url 	string	`json:"url"`
}

type RemoveXHRBreakpointResult struct {

}

// Sets breakpoint on particular operation with DOM.
const SetDOMBreakpoint = "DOMDebugger.setDOMBreakpoint"

type SetDOMBreakpointParams struct {

	// Identifier of the node to set breakpoint on.
	NodeId 	interface{}	`json:"nodeId"`

	// Type of the operation to stop upon.
	Type 	DOMBreakpointType	`json:"type"`
}

type SetDOMBreakpointResult struct {

}

// Sets breakpoint on particular DOM event.
const SetEventListenerBreakpoint = "DOMDebugger.setEventListenerBreakpoint"

type SetEventListenerBreakpointParams struct {

	// DOM Event name to stop on (any DOM event will do).
	EventName 	string	`json:"eventName"`

	// EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any
	// EventTarget.
	TargetName 	string	`json:"targetName"`
}

type SetEventListenerBreakpointResult struct {

}

// Sets breakpoint on particular native event.
const SetInstrumentationBreakpoint = "DOMDebugger.setInstrumentationBreakpoint"

type SetInstrumentationBreakpointParams struct {

	// Instrumentation name to stop on.
	EventName 	string	`json:"eventName"`
}

type SetInstrumentationBreakpointResult struct {

}

// Sets breakpoint on XMLHttpRequest.
const SetXHRBreakpoint = "DOMDebugger.setXHRBreakpoint"

type SetXHRBreakpointParams struct {

	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	Url 	string	`json:"url"`
}

type SetXHRBreakpointResult struct {

}