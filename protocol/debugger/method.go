package debugger

import (
	"github.com/diiyw/cuto/protocol/runtime"
)


// Continues execution until specific location is reached.
const ContinueToLocation = "Debugger.continueToLocation"

type ContinueToLocationParams struct {

	// Location to continue to.
	Location 	Location	`json:"location"`

	// 
	TargetCallFrames 	string	`json:"targetCallFrames,omitempty"`
}

type ContinueToLocationResult struct {

}

// Disables debugger for given page.
const Disable = "Debugger.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables debugger for the given page. Clients should not assume that the debugging has been
// enabled until the result for this command is received.
const Enable = "Debugger.enable"

type EnableParams struct {

	// The maximum size in bytes of collected scripts (not referenced by other heap objects)
	// the debugger can hold. Puts no limit if paramter is omitted.
	MaxScriptsCacheSize 	float64	`json:"maxScriptsCacheSize,omitempty"`
}

type EnableResult struct {

	// Unique identifier of the debugger.
	DebuggerId 	runtime.UniqueDebuggerId	`json:"debuggerId"`
}

// Evaluates expression on a given call frame.
const EvaluateOnCallFrame = "Debugger.evaluateOnCallFrame"

type EvaluateOnCallFrameParams struct {

	// Call frame identifier to evaluate on.
	CallFrameId 	CallFrameId	`json:"callFrameId"`

	// Expression to evaluate.
	Expression 	string	`json:"expression"`

	// String object group name to put result into (allows rapid releasing resulting object handles
	// using `releaseObjectGroup`).
	ObjectGroup 	string	`json:"objectGroup,omitempty"`

	// Specifies whether command line API should be available to the evaluated expression, defaults
	// to false.
	IncludeCommandLineAPI 	bool	`json:"includeCommandLineAPI,omitempty"`

	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent 	bool	`json:"silent,omitempty"`

	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue 	bool	`json:"returnByValue,omitempty"`

	// Whether preview should be generated for the result.
	GeneratePreview 	bool	`json:"generatePreview,omitempty"`

	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect 	bool	`json:"throwOnSideEffect,omitempty"`

	// Terminate execution after timing out (number of milliseconds).
	Timeout 	runtime.TimeDelta	`json:"timeout,omitempty"`
}

type EvaluateOnCallFrameResult struct {

	// Object wrapper for the evaluation result.
	Result 	runtime.RemoteObject	`json:"result"`
	// Exception details.
	ExceptionDetails 	runtime.ExceptionDetails	`json:"exceptionDetails"`
}

// Returns possible locations for breakpoint. scriptId in start and end range locations should be
// the same.
const GetPossibleBreakpoints = "Debugger.getPossibleBreakpoints"

type GetPossibleBreakpointsParams struct {

	// Start of range to search possible breakpoint locations in.
	Start 	Location	`json:"start"`

	// End of range to search possible breakpoint locations in (excluding). When not specified, end
	// of scripts is used as end of range.
	End 	Location	`json:"end,omitempty"`

	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction 	bool	`json:"restrictToFunction,omitempty"`
}

type GetPossibleBreakpointsResult struct {

	// List of the possible breakpoint locations.
	Locations 	[]*BreakLocation	`json:"locations"`
}

// Returns source for the script with given id.
const GetScriptSource = "Debugger.getScriptSource"

type GetScriptSourceParams struct {

	// Id of the script to get source for.
	ScriptId 	runtime.ScriptId	`json:"scriptId"`
}

type GetScriptSourceResult struct {

	// Script source.
	ScriptSource 	string	`json:"scriptSource"`
}

// Returns bytecode for the WebAssembly script with given id.
const GetWasmBytecode = "Debugger.getWasmBytecode"

type GetWasmBytecodeParams struct {

	// Id of the Wasm script to get source for.
	ScriptId 	runtime.ScriptId	`json:"scriptId"`
}

type GetWasmBytecodeResult struct {

	// Script source.
	Bytecode 	[]byte	`json:"bytecode"`
}

// Returns stack trace with given `stackTraceId`.
const GetStackTrace = "Debugger.getStackTrace"

type GetStackTraceParams struct {

	// 
	StackTraceId 	runtime.StackTraceId	`json:"stackTraceId"`
}

type GetStackTraceResult struct {

	// 
	StackTrace 	runtime.StackTrace	`json:"stackTrace"`
}

// Stops on the next JavaScript statement.
const Pause = "Debugger.pause"

type PauseParams struct {
}

type PauseResult struct {

}

// 
const PauseOnAsyncCall = "Debugger.pauseOnAsyncCall"

type PauseOnAsyncCallParams struct {

	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceId 	runtime.StackTraceId	`json:"parentStackTraceId"`
}

type PauseOnAsyncCallResult struct {

}

// Removes JavaScript breakpoint.
const RemoveBreakpoint = "Debugger.removeBreakpoint"

type RemoveBreakpointParams struct {

	// 
	BreakpointId 	BreakpointId	`json:"breakpointId"`
}

type RemoveBreakpointResult struct {

}

// Restarts particular call frame from the beginning.
const RestartFrame = "Debugger.restartFrame"

type RestartFrameParams struct {

	// Call frame identifier to evaluate on.
	CallFrameId 	CallFrameId	`json:"callFrameId"`
}

type RestartFrameResult struct {

	// New stack trace.
	CallFrames 	[]*CallFrame	`json:"callFrames"`
	// Async stack trace, if any.
	AsyncStackTrace 	runtime.StackTrace	`json:"asyncStackTrace"`
	// Async stack trace, if any.
	AsyncStackTraceId 	runtime.StackTraceId	`json:"asyncStackTraceId"`
}

// Resumes JavaScript execution.
const Resume = "Debugger.resume"

type ResumeParams struct {
}

type ResumeResult struct {

}

// Searches for given string in script content.
const SearchInContent = "Debugger.searchInContent"

type SearchInContentParams struct {

	// Id of the script to search in.
	ScriptId 	runtime.ScriptId	`json:"scriptId"`

	// String to search for.
	Query 	string	`json:"query"`

	// If true, search is case sensitive.
	CaseSensitive 	bool	`json:"caseSensitive,omitempty"`

	// If true, treats string parameter as regex.
	IsRegex 	bool	`json:"isRegex,omitempty"`
}

type SearchInContentResult struct {

	// List of search matches.
	Result 	[]*SearchMatch	`json:"result"`
}

// Enables or disables async call stacks tracking.
const SetAsyncCallStackDepth = "Debugger.setAsyncCallStackDepth"

type SetAsyncCallStackDepthParams struct {

	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async
	// call stacks (default).
	MaxDepth 	int	`json:"maxDepth"`
}

type SetAsyncCallStackDepthResult struct {

}

// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in
// scripts with url matching one of the patterns. VM will try to leave blackboxed script by
// performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
const SetBlackboxPatterns = "Debugger.setBlackboxPatterns"

type SetBlackboxPatternsParams struct {

	// Array of regexps that will be used to check script url for blackbox state.
	Patterns 	[]string	`json:"patterns"`
}

type SetBlackboxPatternsResult struct {

}

// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted
// scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// Positions array contains positions where blackbox state is changed. First interval isn't
// blackboxed. Array should be sorted.
const SetBlackboxedRanges = "Debugger.setBlackboxedRanges"

type SetBlackboxedRangesParams struct {

	// Id of the script.
	ScriptId 	runtime.ScriptId	`json:"scriptId"`

	// 
	Positions 	[]*ScriptPosition	`json:"positions"`
}

type SetBlackboxedRangesResult struct {

}

// Sets JavaScript breakpoint at a given location.
const SetBreakpoint = "Debugger.setBreakpoint"

type SetBreakpointParams struct {

	// Location to set breakpoint in.
	Location 	Location	`json:"location"`

	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition 	string	`json:"condition,omitempty"`
}

type SetBreakpointResult struct {

	// Id of the created breakpoint for further reference.
	BreakpointId 	BreakpointId	`json:"breakpointId"`
	// Location this breakpoint resolved into.
	ActualLocation 	Location	`json:"actualLocation"`
}

// Sets instrumentation breakpoint.
const SetInstrumentationBreakpoint = "Debugger.setInstrumentationBreakpoint"

type SetInstrumentationBreakpointParams struct {

	// Instrumentation name.
	Instrumentation 	string	`json:"instrumentation"`
}

type SetInstrumentationBreakpointResult struct {

	// Id of the created breakpoint for further reference.
	BreakpointId 	BreakpointId	`json:"breakpointId"`
}

// Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this
// command is issued, all existing parsed scripts will have breakpoints resolved and returned in
// `locations` property. Further matching script parsing will result in subsequent
// `breakpointResolved` events issued. This logical breakpoint will survive page reloads.
const SetBreakpointByUrl = "Debugger.setBreakpointByUrl"

type SetBreakpointByUrlParams struct {

	// Line number to set breakpoint at.
	LineNumber 	int	`json:"lineNumber"`

	// URL of the resources to set breakpoint on.
	Url 	string	`json:"url,omitempty"`

	// Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or
	// `urlRegex` must be specified.
	UrlRegex 	string	`json:"urlRegex,omitempty"`

	// Script hash of the resources to set breakpoint on.
	ScriptHash 	string	`json:"scriptHash,omitempty"`

	// Offset in the line to set breakpoint at.
	ColumnNumber 	int	`json:"columnNumber,omitempty"`

	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition 	string	`json:"condition,omitempty"`
}

type SetBreakpointByUrlResult struct {

	// Id of the created breakpoint for further reference.
	BreakpointId 	BreakpointId	`json:"breakpointId"`
	// List of the locations this breakpoint resolved into upon addition.
	Locations 	[]*Location	`json:"locations"`
}

// Sets JavaScript breakpoint before each call to the given function.
// If another function was created from the same source as a given one,
// calling it will also trigger the breakpoint.
const SetBreakpointOnFunctionCall = "Debugger.setBreakpointOnFunctionCall"

type SetBreakpointOnFunctionCallParams struct {

	// Function object id.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId"`

	// Expression to use as a breakpoint condition. When specified, debugger will
	// stop on the breakpoint if this expression evaluates to true.
	Condition 	string	`json:"condition,omitempty"`
}

type SetBreakpointOnFunctionCallResult struct {

	// Id of the created breakpoint for further reference.
	BreakpointId 	BreakpointId	`json:"breakpointId"`
}

// Activates / deactivates all breakpoints on the page.
const SetBreakpointsActive = "Debugger.setBreakpointsActive"

type SetBreakpointsActiveParams struct {

	// New value for breakpoints active state.
	Active 	bool	`json:"active"`
}

type SetBreakpointsActiveResult struct {

}

// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or
// no exceptions. Initial pause on exceptions state is `none`.
const SetPauseOnExceptions = "Debugger.setPauseOnExceptions"

type SetPauseOnExceptionsParams struct {

	// Pause on exceptions mode.
	State 	string	`json:"state"`
}

type SetPauseOnExceptionsResult struct {

}

// Changes return value in top frame. Available only at return break position.
const SetReturnValue = "Debugger.setReturnValue"

type SetReturnValueParams struct {

	// New return value.
	NewValue 	runtime.CallArgument	`json:"newValue"`
}

type SetReturnValueResult struct {

}

// Edits JavaScript source live.
const SetScriptSource = "Debugger.setScriptSource"

type SetScriptSourceParams struct {

	// Id of the script to edit.
	ScriptId 	runtime.ScriptId	`json:"scriptId"`

	// New content of the script.
	ScriptSource 	string	`json:"scriptSource"`

	// If true the change will not actually be applied. Dry run may be used to get result
	// description without actually modifying the code.
	DryRun 	bool	`json:"dryRun,omitempty"`
}

type SetScriptSourceResult struct {

	// New stack trace in case editing has happened while VM was stopped.
	CallFrames 	[]*CallFrame	`json:"callFrames"`
	// Whether current call stack  was modified after applying the changes.
	StackChanged 	bool	`json:"stackChanged"`
	// Async stack trace, if any.
	AsyncStackTrace 	runtime.StackTrace	`json:"asyncStackTrace"`
	// Async stack trace, if any.
	AsyncStackTraceId 	runtime.StackTraceId	`json:"asyncStackTraceId"`
	// Exception details if any.
	ExceptionDetails 	runtime.ExceptionDetails	`json:"exceptionDetails"`
}

// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
const SetSkipAllPauses = "Debugger.setSkipAllPauses"

type SetSkipAllPausesParams struct {

	// New value for skip pauses state.
	Skip 	bool	`json:"skip"`
}

type SetSkipAllPausesResult struct {

}

// Changes value of variable in a callframe. Object-based scopes are not supported and must be
// mutated manually.
const SetVariableValue = "Debugger.setVariableValue"

type SetVariableValueParams struct {

	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch'
	// scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber 	int	`json:"scopeNumber"`

	// Variable name.
	VariableName 	string	`json:"variableName"`

	// New variable value.
	NewValue 	runtime.CallArgument	`json:"newValue"`

	// Id of callframe that holds variable.
	CallFrameId 	CallFrameId	`json:"callFrameId"`
}

type SetVariableValueResult struct {

}

// Steps into the function call.
const StepInto = "Debugger.stepInto"

type StepIntoParams struct {

	// Debugger will pause on the execution of the first async task which was scheduled
	// before next pause.
	BreakOnAsyncCall 	bool	`json:"breakOnAsyncCall,omitempty"`
}

type StepIntoResult struct {

}

// Steps out of the function call.
const StepOut = "Debugger.stepOut"

type StepOutParams struct {
}

type StepOutResult struct {

}

// Steps over the statement.
const StepOver = "Debugger.stepOver"

type StepOverParams struct {
}

type StepOverResult struct {

}