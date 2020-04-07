package domdebugger
// DOM breakpoint type.
type DOMBreakpointType string

// Object event listener.
type EventListener  struct {

	// `EventListener`'s type.
	Type	string	`json:"type"`

	// `EventListener`'s useCapture.
	UseCapture	bool	`json:"useCapture"`

	// `EventListener`'s passive flag.
	Passive	bool	`json:"passive"`

	// `EventListener`'s once flag.
	Once	bool	`json:"once"`

	// Script id of the handler code.
	ScriptId	interface{}	`json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber	int	`json:"lineNumber"`

	// Column number in the script (0-based).
	ColumnNumber	int	`json:"columnNumber"`

	// Event handler function value.
	Handler	interface{}	`json:"handler"`

	// Event original handler function value.
	OriginalHandler	interface{}	`json:"originalHandler"`

	// Node the listener is added to (if any).
	BackendNodeId	interface{}	`json:"backendNodeId"`
}
