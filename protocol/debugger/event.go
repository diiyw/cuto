package debugger

import (
	"github.com/diiyw/cuto/protocol/runtime"
)


// Fired when breakpoint is resolved to an actual script and location.
const BreakpointResolvedEvent = "Debugger.breakpointResolved"
type BreakpointResolvedParams struct {

	// Breakpoint unique identifier.
	BreakpointId 	BreakpointId
	// Actual breakpoint location.
	Location 	Location}



// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
const PausedEvent = "Debugger.paused"
type PausedParams struct {

	// Call stack the virtual machine stopped on.
	CallFrames 	[]*CallFrame
	// Pause reason.
	Reason 	string
	// Object containing break-specific auxiliary properties.
	Data 	interface{}
	// Hit breakpoints IDs
	HitBreakpoints 	[]string
	// Async stack trace, if any.
	AsyncStackTrace 	runtime.StackTrace
	// Async stack trace, if any.
	AsyncStackTraceId 	runtime.StackTraceId
	// Never present, will be removed.
	AsyncCallStackTraceId 	runtime.StackTraceId}



// Fired when the virtual machine resumed execution.
const ResumedEvent = "Debugger.resumed"
type ResumedParams struct {
}



// Fired when virtual machine fails to parse the script.
const ScriptFailedToParseEvent = "Debugger.scriptFailedToParse"
type ScriptFailedToParseParams struct {

	// Identifier of the script parsed.
	ScriptId 	runtime.ScriptId
	// URL or name of the script parsed (if any).
	Url 	string
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine 	int
	// Column offset of the script within the resource with given URL.
	StartColumn 	int
	// Last line of the script.
	EndLine 	int
	// Length of the last line of the script.
	EndColumn 	int
	// Specifies script creation context.
	ExecutionContextId 	runtime.ExecutionContextId
	// Content hash of the script.
	Hash 	string
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData 	interface{}
	// URL of source map associated with script (if any).
	SourceMapURL 	string
	// True, if this script has sourceURL.
	HasSourceURL 	bool
	// True, if this script is ES6 module.
	IsModule 	bool
	// This script length.
	Length 	int
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	StackTrace 	runtime.StackTrace}



// Fired when virtual machine parses script. This event is also fired for all known and uncollected
// scripts upon enabling debugger.
const ScriptParsedEvent = "Debugger.scriptParsed"
type ScriptParsedParams struct {

	// Identifier of the script parsed.
	ScriptId 	runtime.ScriptId
	// URL or name of the script parsed (if any).
	Url 	string
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine 	int
	// Column offset of the script within the resource with given URL.
	StartColumn 	int
	// Last line of the script.
	EndLine 	int
	// Length of the last line of the script.
	EndColumn 	int
	// Specifies script creation context.
	ExecutionContextId 	runtime.ExecutionContextId
	// Content hash of the script.
	Hash 	string
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData 	interface{}
	// True, if this script is generated as a result of the live edit operation.
	IsLiveEdit 	bool
	// URL of source map associated with script (if any).
	SourceMapURL 	string
	// True, if this script has sourceURL.
	HasSourceURL 	bool
	// True, if this script is ES6 module.
	IsModule 	bool
	// This script length.
	Length 	int
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	StackTrace 	runtime.StackTrace}

