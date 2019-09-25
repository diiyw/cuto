package debugger

import (

	"github.com/diiyw/goc/protocol/runtime"

)

// Breakpoint identifier.
type BreakpointId string	

// Call frame identifier.
type CallFrameId string	

// Location in the source code.
type Location struct {
	
	// Script identifier as reported in the `Debugger.scriptParsed`.
	
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	// Line number in the script (0-based).
	
	LineNumber	int	`json:"lineNumber"`
	
	// Column number in the script (0-based).
	
	ColumnNumber	int	`json:"columnNumber"`
	
}	

// Location in the source code.
type ScriptPosition struct {
	
	
	
	LineNumber	int	`json:"lineNumber"`
	
	
	
	ColumnNumber	int	`json:"columnNumber"`
	
}	

// JavaScript call frame. Array of call frames form the call stack.
type CallFrame struct {
	
	// Call frame identifier. This identifier is only valid while the virtual machine is paused.
	
	CallFrameId	CallFrameId	`json:"callFrameId"`
	
	// Name of the JavaScript function called on this call frame.
	
	FunctionName	string	`json:"functionName"`
	
	// Location in the source code.
	
	FunctionLocation	Location	`json:"functionLocation"`
	
	// Location in the source code.
	
	Location	Location	`json:"location"`
	
	// JavaScript script name or url.
	
	Url	string	`json:"url"`
	
	// Scope chain for this call frame.
	
	ScopeChain	[]Scope	`json:"scopeChain"`
	
	// `this` object for this call frame.
	
	This	runtime.RemoteObject	`json:"this"`
	
	// The value being returned, if the function is at return point.
	
	ReturnValue	runtime.RemoteObject	`json:"returnValue"`
	
}	

// Scope description.
type Scope struct {
	
	// Scope type.
	// Possible value: global,local,with,closure,catch,block,script,eval,module,
	Type	string	`json:"type"`
	
	// Object representing the scope. For `global` and `with` scopes it represents the actual
	// object; for the rest of the scopes, it is artificial transient object enumerating scope
	// variables as its properties.
	
	Object	runtime.RemoteObject	`json:"object"`
	
	
	
	Name	string	`json:"name"`
	
	// Location in the source code where scope starts
	
	StartLocation	Location	`json:"startLocation"`
	
	// Location in the source code where scope ends
	
	EndLocation	Location	`json:"endLocation"`
	
}	

// Search match for resource.
type SearchMatch struct {
	
	// Line number in resource content.
	
	LineNumber	float64	`json:"lineNumber"`
	
	// Line with match content.
	
	LineContent	string	`json:"lineContent"`
	
}	

// 
type BreakLocation struct {
	
	// Script identifier as reported in the `Debugger.scriptParsed`.
	
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	// Line number in the script (0-based).
	
	LineNumber	int	`json:"lineNumber"`
	
	// Column number in the script (0-based).
	
	ColumnNumber	int	`json:"columnNumber"`
	
	
	// Possible value: debuggerStatement,call,return,
	Type	string	`json:"type"`
	
}	

