package debugger

import (

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	// Continues execution until specific location is reached.
	ContinueToLocation = "Debugger.continueToLocation"
	
	// Disables debugger for given page.
	Disable = "Debugger.disable"
	
	// Enables debugger for the given page. Clients should not assume that the debugging has been
	// enabled until the result for this command is received.
	Enable = "Debugger.enable"
	
	// Evaluates expression on a given call frame.
	EvaluateOnCallFrame = "Debugger.evaluateOnCallFrame"
	
	// Returns possible locations for breakpoint. scriptId in start and end range locations should be
	// the same.
	GetPossibleBreakpoints = "Debugger.getPossibleBreakpoints"
	
	// Returns source for the script with given id.
	GetScriptSource = "Debugger.getScriptSource"
	
	// Returns stack trace with given `stackTraceId`.
	GetStackTrace = "Debugger.getStackTrace"
	
	// Stops on the next JavaScript statement.
	Pause = "Debugger.pause"
	
	
	PauseOnAsyncCall = "Debugger.pauseOnAsyncCall"
	
	// Removes JavaScript breakpoint.
	RemoveBreakpoint = "Debugger.removeBreakpoint"
	
	// Restarts particular call frame from the beginning.
	RestartFrame = "Debugger.restartFrame"
	
	// Resumes JavaScript execution.
	Resume = "Debugger.resume"
	
	// Searches for given string in script content.
	SearchInContent = "Debugger.searchInContent"
	
	// Enables or disables async call stacks tracking.
	SetAsyncCallStackDepth = "Debugger.setAsyncCallStackDepth"
	
	// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in
	// scripts with url matching one of the patterns. VM will try to leave blackboxed script by
	// performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
	SetBlackboxPatterns = "Debugger.setBlackboxPatterns"
	
	// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted
	// scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
	// Positions array contains positions where blackbox state is changed. First interval isn't
	// blackboxed. Array should be sorted.
	SetBlackboxedRanges = "Debugger.setBlackboxedRanges"
	
	// Sets JavaScript breakpoint at a given location.
	SetBreakpoint = "Debugger.setBreakpoint"
	
	// Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this
	// command is issued, all existing parsed scripts will have breakpoints resolved and returned in
	// `locations` property. Further matching script parsing will result in subsequent
	// `breakpointResolved` events issued. This logical breakpoint will survive page reloads.
	SetBreakpointByUrl = "Debugger.setBreakpointByUrl"
	
	// Sets JavaScript breakpoint before each call to the given function.
	// If another function was created from the same source as a given one,
	// calling it will also trigger the breakpoint.
	SetBreakpointOnFunctionCall = "Debugger.setBreakpointOnFunctionCall"
	
	// Activates / deactivates all breakpoints on the page.
	SetBreakpointsActive = "Debugger.setBreakpointsActive"
	
	// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or
	// no exceptions. Initial pause on exceptions state is `none`.
	SetPauseOnExceptions = "Debugger.setPauseOnExceptions"
	
	// Changes return value in top frame. Available only at return break position.
	SetReturnValue = "Debugger.setReturnValue"
	
	// Edits JavaScript source live.
	SetScriptSource = "Debugger.setScriptSource"
	
	// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
	SetSkipAllPauses = "Debugger.setSkipAllPauses"
	
	// Changes value of variable in a callframe. Object-based scopes are not supported and must be
	// mutated manually.
	SetVariableValue = "Debugger.setVariableValue"
	
	// Steps into the function call.
	StepInto = "Debugger.stepInto"
	
	// Steps out of the function call.
	StepOut = "Debugger.stepOut"
	
	// Steps over the statement.
	StepOver = "Debugger.stepOver"
	
)

// ContinueToLocation parameters
type ContinueToLocationParams struct {
	
	// Location to continue to.
	Location	Location	`json:"location"`
	
	
	TargetCallFrames	string	`json:"targetCallFrames"`
	
}

// ContinueToLocation returns
type ContinueToLocationReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
	// The maximum size in bytes of collected scripts (not referenced by other heap objects)
	// the debugger can hold. Puts no limit if paramter is omitted.
	MaxScriptsCacheSize	float64	`json:"maxScriptsCacheSize,omitempty"`
	
}

// Enable returns
type EnableReturns struct {
	
	// Unique identifier of the debugger.
	DebuggerId	runtime.UniqueDebuggerId	`json:"debuggerId"`
	
}

// EvaluateOnCallFrame parameters
type EvaluateOnCallFrameParams struct {
	
	// Call frame identifier to evaluate on.
	CallFrameId	CallFrameId	`json:"callFrameId"`
	
	// Expression to evaluate.
	Expression	string	`json:"expression"`
	
	// String object group name to put result into (allows rapid releasing resulting object handles
	// using `releaseObjectGroup`).
	ObjectGroup	string	`json:"objectGroup"`
	
	// Specifies whether command line API should be available to the evaluated expression, defaults
	// to false.
	IncludeCommandLineAPI	bool	`json:"includeCommandLineAPI"`
	
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent	bool	`json:"silent"`
	
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue	bool	`json:"returnByValue"`
	
	// Whether preview should be generated for the result.
	GeneratePreview	bool	`json:"generatePreview"`
	
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect	bool	`json:"throwOnSideEffect"`
	
	// Terminate execution after timing out (number of milliseconds).
	Timeout	runtime.TimeDelta	`json:"timeout"`
	
}

// EvaluateOnCallFrame returns
type EvaluateOnCallFrameReturns struct {
	
	// Object wrapper for the evaluation result.
	Result	runtime.RemoteObject	`json:"result"`
	
	// Exception details.
	ExceptionDetails	runtime.ExceptionDetails	`json:"exceptionDetails"`
	
}

// GetPossibleBreakpoints parameters
type GetPossibleBreakpointsParams struct {
	
	// Start of range to search possible breakpoint locations in.
	Start	Location	`json:"start"`
	
	// End of range to search possible breakpoint locations in (excluding). When not specified, end
	// of scripts is used as end of range.
	End	Location	`json:"end"`
	
	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction	bool	`json:"restrictToFunction"`
	
}

// GetPossibleBreakpoints returns
type GetPossibleBreakpointsReturns struct {
	
	// List of the possible breakpoint locations.
	Locations	[]BreakLocation	`json:"locations"`
	
}

// GetScriptSource parameters
type GetScriptSourceParams struct {
	
	// Id of the script to get source for.
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
}

// GetScriptSource returns
type GetScriptSourceReturns struct {
	
	// Script source.
	ScriptSource	string	`json:"scriptSource"`
	
}

// GetStackTrace parameters
type GetStackTraceParams struct {
	
	
	StackTraceId	runtime.StackTraceId	`json:"stackTraceId"`
	
}

// GetStackTrace returns
type GetStackTraceReturns struct {
	
	
	StackTrace	runtime.StackTrace	`json:"stackTrace"`
	
}

// Pause parameters
type PauseParams struct {
	
}

// Pause returns
type PauseReturns struct {
	
}

// PauseOnAsyncCall parameters
type PauseOnAsyncCallParams struct {
	
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceId	runtime.StackTraceId	`json:"parentStackTraceId"`
	
}

// PauseOnAsyncCall returns
type PauseOnAsyncCallReturns struct {
	
}

// RemoveBreakpoint parameters
type RemoveBreakpointParams struct {
	
	
	BreakpointId	BreakpointId	`json:"breakpointId"`
	
}

// RemoveBreakpoint returns
type RemoveBreakpointReturns struct {
	
}

// RestartFrame parameters
type RestartFrameParams struct {
	
	// Call frame identifier to evaluate on.
	CallFrameId	CallFrameId	`json:"callFrameId"`
	
}

// RestartFrame returns
type RestartFrameReturns struct {
	
	// New stack trace.
	CallFrames	[]CallFrame	`json:"callFrames"`
	
	// Async stack trace, if any.
	AsyncStackTrace	runtime.StackTrace	`json:"asyncStackTrace"`
	
	// Async stack trace, if any.
	AsyncStackTraceId	runtime.StackTraceId	`json:"asyncStackTraceId"`
	
}

// Resume parameters
type ResumeParams struct {
	
}

// Resume returns
type ResumeReturns struct {
	
}

// SearchInContent parameters
type SearchInContentParams struct {
	
	// Id of the script to search in.
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	// String to search for.
	Query	string	`json:"query"`
	
	// If true, search is case sensitive.
	CaseSensitive	bool	`json:"caseSensitive"`
	
	// If true, treats string parameter as regex.
	IsRegex	bool	`json:"isRegex"`
	
}

// SearchInContent returns
type SearchInContentReturns struct {
	
	// List of search matches.
	Result	[]SearchMatch	`json:"result"`
	
}

// SetAsyncCallStackDepth parameters
type SetAsyncCallStackDepthParams struct {
	
	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async
	// call stacks (default).
	MaxDepth	int	`json:"maxDepth"`
	
}

// SetAsyncCallStackDepth returns
type SetAsyncCallStackDepthReturns struct {
	
}

// SetBlackboxPatterns parameters
type SetBlackboxPatternsParams struct {
	
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns	[]string	`json:"patterns"`
	
}

// SetBlackboxPatterns returns
type SetBlackboxPatternsReturns struct {
	
}

// SetBlackboxedRanges parameters
type SetBlackboxedRangesParams struct {
	
	// Id of the script.
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	
	Positions	[]ScriptPosition	`json:"positions"`
	
}

// SetBlackboxedRanges returns
type SetBlackboxedRangesReturns struct {
	
}

// SetBreakpoint parameters
type SetBreakpointParams struct {
	
	// Location to set breakpoint in.
	Location	Location	`json:"location"`
	
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition	string	`json:"condition"`
	
}

// SetBreakpoint returns
type SetBreakpointReturns struct {
	
	// Id of the created breakpoint for further reference.
	BreakpointId	BreakpointId	`json:"breakpointId"`
	
	// Location this breakpoint resolved into.
	ActualLocation	Location	`json:"actualLocation"`
	
}

// SetBreakpointByUrl parameters
type SetBreakpointByUrlParams struct {
	
	// Line number to set breakpoint at.
	LineNumber	int	`json:"lineNumber"`
	
	// URL of the resources to set breakpoint on.
	Url	string	`json:"url"`
	
	// Regex pattern for the URLs of the resources to set breakpoints on. Either `url` or
	// `urlRegex` must be specified.
	UrlRegex	string	`json:"urlRegex"`
	
	// Script hash of the resources to set breakpoint on.
	ScriptHash	string	`json:"scriptHash"`
	
	// Offset in the line to set breakpoint at.
	ColumnNumber	int	`json:"columnNumber"`
	
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the
	// breakpoint if this expression evaluates to true.
	Condition	string	`json:"condition"`
	
}

// SetBreakpointByUrl returns
type SetBreakpointByUrlReturns struct {
	
	// Id of the created breakpoint for further reference.
	BreakpointId	BreakpointId	`json:"breakpointId"`
	
	// List of the locations this breakpoint resolved into upon addition.
	Locations	[]Location	`json:"locations"`
	
}

// SetBreakpointOnFunctionCall parameters
type SetBreakpointOnFunctionCallParams struct {
	
	// Function object id.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
	// Expression to use as a breakpoint condition. When specified, debugger will
	// stop on the breakpoint if this expression evaluates to true.
	Condition	string	`json:"condition"`
	
}

// SetBreakpointOnFunctionCall returns
type SetBreakpointOnFunctionCallReturns struct {
	
	// Id of the created breakpoint for further reference.
	BreakpointId	BreakpointId	`json:"breakpointId"`
	
}

// SetBreakpointsActive parameters
type SetBreakpointsActiveParams struct {
	
	// New value for breakpoints active state.
	Active	bool	`json:"active"`
	
}

// SetBreakpointsActive returns
type SetBreakpointsActiveReturns struct {
	
}

// SetPauseOnExceptions parameters
type SetPauseOnExceptionsParams struct {
	
	// Pause on exceptions mode.
	State	string	`json:"state"`
	
}

// SetPauseOnExceptions returns
type SetPauseOnExceptionsReturns struct {
	
}

// SetReturnValue parameters
type SetReturnValueParams struct {
	
	// New return value.
	NewValue	runtime.CallArgument	`json:"newValue"`
	
}

// SetReturnValue returns
type SetReturnValueReturns struct {
	
}

// SetScriptSource parameters
type SetScriptSourceParams struct {
	
	// Id of the script to edit.
	ScriptId	runtime.ScriptId	`json:"scriptId"`
	
	// New content of the script.
	ScriptSource	string	`json:"scriptSource"`
	
	// If true the change will not actually be applied. Dry run may be used to get result
	// description without actually modifying the code.
	DryRun	bool	`json:"dryRun"`
	
}

// SetScriptSource returns
type SetScriptSourceReturns struct {
	
	// New stack trace in case editing has happened while VM was stopped.
	CallFrames	[]CallFrame	`json:"callFrames"`
	
	// Whether current call stack  was modified after applying the changes.
	StackChanged	bool	`json:"stackChanged"`
	
	// Async stack trace, if any.
	AsyncStackTrace	runtime.StackTrace	`json:"asyncStackTrace"`
	
	// Async stack trace, if any.
	AsyncStackTraceId	runtime.StackTraceId	`json:"asyncStackTraceId"`
	
	// Exception details if any.
	ExceptionDetails	runtime.ExceptionDetails	`json:"exceptionDetails"`
	
}

// SetSkipAllPauses parameters
type SetSkipAllPausesParams struct {
	
	// New value for skip pauses state.
	Skip	bool	`json:"skip"`
	
}

// SetSkipAllPauses returns
type SetSkipAllPausesReturns struct {
	
}

// SetVariableValue parameters
type SetVariableValueParams struct {
	
	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch'
	// scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber	int	`json:"scopeNumber"`
	
	// Variable name.
	VariableName	string	`json:"variableName"`
	
	// New variable value.
	NewValue	runtime.CallArgument	`json:"newValue"`
	
	// Id of callframe that holds variable.
	CallFrameId	CallFrameId	`json:"callFrameId"`
	
}

// SetVariableValue returns
type SetVariableValueReturns struct {
	
}

// StepInto parameters
type StepIntoParams struct {
	
	// Debugger will issue additional Debugger.paused notification if any async task is scheduled
	// before next pause.
	BreakOnAsyncCall	bool	`json:"breakOnAsyncCall"`
	
}

// StepInto returns
type StepIntoReturns struct {
	
}

// StepOut parameters
type StepOutParams struct {
	
}

// StepOut returns
type StepOutReturns struct {
	
}

// StepOver parameters
type StepOverParams struct {
	
}

// StepOver returns
type StepOverReturns struct {
	
}

