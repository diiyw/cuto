package debugger

import (

	"github.com/diiyw/gator/protocol/runtime"

)
const (
	
	// Fired when breakpoint is resolved to an actual script and location.
	BreakpointResolvedEvent = "Debugger.breakpointResolved"
	
	// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
	PausedEvent = "Debugger.paused"
	
	// Fired when the virtual machine resumed execution.
	ResumedEvent = "Debugger.resumed"
	
	// Fired when virtual machine fails to parse the script.
	ScriptFailedToParseEvent = "Debugger.scriptFailedToParse"
	
	// Fired when virtual machine parses script. This event is also fired for all known and uncollected
	// scripts upon enabling debugger.
	ScriptParsedEvent = "Debugger.scriptParsed"
	
)

// Fired when breakpoint is resolved to an actual script and location.
type BreakpointResolvedParams struct {
	
	// Breakpoint unique identifier.
	BreakpointId	BreakpointId	`json:"breakpointId"`
	
	// Actual breakpoint location.
	Location	Location	`json:"location"`
	
}

// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
type PausedParams struct {
	
	// Call stack the virtual machine stopped on.
	CallFrames	[]CallFrame	`json:"callFrames"`
	
	// Pause reason.
	Reason	string	`json:"reason"`
	
	// Object containing break-specific auxiliary properties.
	Data	interface{}	`json:"data"`
	
	// Hit breakpoints IDs
	HitBreakpoints	[]string	`json:"hitBreakpoints"`
	
	// Async stack trace, if any.
	AsyncStackTrace	runtime.StackTrace	`json:"asyncStackTrace"`
	
	// Async stack trace, if any.
	AsyncStackTraceId	runtime.StackTraceId	`json:"asyncStackTraceId"`
	
	// Just scheduled async call will have this stack trace as parent stack during async execution.
	// This field is available only after `Debugger.stepInto` call with `breakOnAsynCall` flag.
	AsyncCallStackTraceId	runtime.StackTraceId	`json:"asyncCallStackTraceId"`
	
}

// Fired when the virtual machine resumed execution.
type ResumedParams struct {
	
}

// Fired when virtual machine fails to parse the script.
type ScriptFailedToParseParams struct {
	
	// Identifier of the script parsed.
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	// URL or name of the script parsed (if any).
	Url	string	`json:"url"`
	
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine	int	`json:"startLine"`
	
	// Column offset of the script within the resource with given URL.
	StartColumn	int	`json:"startColumn"`
	
	// Last line of the script.
	EndLine	int	`json:"endLine"`
	
	// Length of the last line of the script.
	EndColumn	int	`json:"endColumn"`
	
	// Specifies script creation context.
	ExecutionContextId	runtime.ExecutionContextId	`json:"executionContextId"`
	
	// Content hash of the script.
	Hash	string	`json:"hash"`
	
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData	interface{}	`json:"executionContextAuxData"`
	
	// URL of source map associated with script (if any).
	SourceMapURL	string	`json:"sourceMapURL"`
	
	// True, if this script has sourceURL.
	HasSourceURL	bool	`json:"hasSourceURL"`
	
	// True, if this script is ES6 module.
	IsModule	bool	`json:"isModule"`
	
	// This script length.
	Length	int	`json:"length"`
	
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	StackTrace	runtime.StackTrace	`json:"stackTrace"`
	
}

// Fired when virtual machine parses script. This event is also fired for all known and uncollected
	// scripts upon enabling debugger.
type ScriptParsedParams struct {
	
	// Identifier of the script parsed.
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	// URL or name of the script parsed (if any).
	Url	string	`json:"url"`
	
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine	int	`json:"startLine"`
	
	// Column offset of the script within the resource with given URL.
	StartColumn	int	`json:"startColumn"`
	
	// Last line of the script.
	EndLine	int	`json:"endLine"`
	
	// Length of the last line of the script.
	EndColumn	int	`json:"endColumn"`
	
	// Specifies script creation context.
	ExecutionContextId	runtime.ExecutionContextId	`json:"executionContextId"`
	
	// Content hash of the script.
	Hash	string	`json:"hash"`
	
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData	interface{}	`json:"executionContextAuxData"`
	
	// True, if this script is generated as a result of the live edit operation.
	IsLiveEdit	bool	`json:"isLiveEdit"`
	
	// URL of source map associated with script (if any).
	SourceMapURL	string	`json:"sourceMapURL"`
	
	// True, if this script has sourceURL.
	HasSourceURL	bool	`json:"hasSourceURL"`
	
	// True, if this script is ES6 module.
	IsModule	bool	`json:"isModule"`
	
	// This script length.
	Length	int	`json:"length"`
	
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	StackTrace	runtime.StackTrace	`json:"stackTrace"`
	
}

