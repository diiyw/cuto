package runtime

const (
	
	// Add handler to promise with given promise object id.
	AwaitPromise = "Runtime.awaitPromise"
	
	// Calls function with given declaration on the given object. Object group of the result is
	// inherited from the target object.
	CallFunctionOn = "Runtime.callFunctionOn"
	
	// Compiles expression.
	CompileScript = "Runtime.compileScript"
	
	// Disables reporting of execution contexts creation.
	Disable = "Runtime.disable"
	
	// Discards collected exceptions and console API calls.
	DiscardConsoleEntries = "Runtime.discardConsoleEntries"
	
	// Enables reporting of execution contexts creation by means of `executionContextCreated` event.
	// When the reporting gets enabled the event will be sent immediately for each existing execution
	// context.
	Enable = "Runtime.enable"
	
	// Evaluates expression on global object.
	Evaluate = "Runtime.evaluate"
	
	// Returns the isolate id.
	GetIsolateId = "Runtime.getIsolateId"
	
	// Returns the JavaScript heap usage.
	// It is the total usage of the corresponding isolate not scoped to a particular Runtime.
	GetHeapUsage = "Runtime.getHeapUsage"
	
	// Returns properties of a given object. Object group of the result is inherited from the target
	// object.
	GetProperties = "Runtime.getProperties"
	
	// Returns all let, const and class variables from global scope.
	GlobalLexicalScopeNames = "Runtime.globalLexicalScopeNames"
	
	
	QueryObjects = "Runtime.queryObjects"
	
	// Releases remote object with given id.
	ReleaseObject = "Runtime.releaseObject"
	
	// Releases all remote objects that belong to a given group.
	ReleaseObjectGroup = "Runtime.releaseObjectGroup"
	
	// Tells inspected instance to run if it was waiting for debugger to attach.
	RunIfWaitingForDebugger = "Runtime.runIfWaitingForDebugger"
	
	// Runs script with given id in a given context.
	RunScript = "Runtime.runScript"
	
	// Enables or disables async call stacks tracking.
	SetAsyncCallStackDepth = "Runtime.setAsyncCallStackDepth"
	
	
	SetCustomObjectFormatterEnabled = "Runtime.setCustomObjectFormatterEnabled"
	
	
	SetMaxCallStackSizeToCapture = "Runtime.setMaxCallStackSizeToCapture"
	
	// Terminate current or next JavaScript execution.
	// Will cancel the termination when the outer-most script execution ends.
	TerminateExecution = "Runtime.terminateExecution"
	
	// If executionContextId is empty, adds binding with the given name on the
	// global objects of all inspected contexts, including those created later,
	// bindings survive reloads.
	// If executionContextId is specified, adds binding only on global object of
	// given execution context.
	// Binding function takes exactly one argument, this argument should be string,
	// in case of any other input, function throws an exception.
	// Each binding function call produces Runtime.bindingCalled notification.
	AddBinding = "Runtime.addBinding"
	
	// This method does not remove binding function from global object but
	// unsubscribes current runtime agent from Runtime.bindingCalled notifications.
	RemoveBinding = "Runtime.removeBinding"
	
)

// AwaitPromise parameters
type AwaitPromiseParams struct {
	
	// Identifier of the promise.
	PromiseObjectId	RemoteObjectId	`json:"promiseObjectId"`
	
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue	bool	`json:"returnByValue"`
	
	// Whether preview should be generated for the result.
	GeneratePreview	bool	`json:"generatePreview"`
	
}

// AwaitPromise returns
type AwaitPromiseReturns struct {
	
	// Promise result. Will contain rejected value if promise was rejected.
	Result	RemoteObject	`json:"result"`
	
	// Exception details if stack strace is available.
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
}

// CallFunctionOn parameters
type CallFunctionOnParams struct {
	
	// Declaration of the function to call.
	FunctionDeclaration	string	`json:"functionDeclaration"`
	
	// Identifier of the object to call function on. Either objectId or executionContextId should
	// be specified.
	ObjectId	RemoteObjectId	`json:"objectId"`
	
	// Call arguments. All call arguments must belong to the same JavaScript world as the target
	// object.
	Arguments	[]CallArgument	`json:"arguments"`
	
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent	bool	`json:"silent"`
	
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue	bool	`json:"returnByValue"`
	
	// Whether preview should be generated for the result.
	GeneratePreview	bool	`json:"generatePreview"`
	
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture	bool	`json:"userGesture"`
	
	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise	bool	`json:"awaitPromise"`
	
	// Specifies execution context which global object will be used to call function on. Either
	// executionContextId or objectId should be specified.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
	
	// Symbolic group name that can be used to release multiple objects. If objectGroup is not
	// specified and objectId is, objectGroup will be inherited from object.
	ObjectGroup	string	`json:"objectGroup"`
	
}

// CallFunctionOn returns
type CallFunctionOnReturns struct {
	
	// Call result.
	Result	RemoteObject	`json:"result"`
	
	// Exception details.
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
}

// CompileScript parameters
type CompileScriptParams struct {
	
	// Expression to compile.
	Expression	string	`json:"expression"`
	
	// Source url to be set for the script.
	SourceURL	string	`json:"sourceURL"`
	
	// Specifies whether the compiled script should be persisted.
	PersistScript	bool	`json:"persistScript"`
	
	// Specifies in which execution context to perform script run. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId,omitempty"`
	
}

// CompileScript returns
type CompileScriptReturns struct {
	
	// Id of the script.
	ScriptId	ScriptId	`json:"scriptId"`
	
	// Exception details.
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// DiscardConsoleEntries parameters
type DiscardConsoleEntriesParams struct {
	
}

// DiscardConsoleEntries returns
type DiscardConsoleEntriesReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// Evaluate parameters
type EvaluateParams struct {
	
	// Expression to evaluate.
	Expression	string	`json:"expression"`
	
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup	string	`json:"objectGroup"`
	
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI	bool	`json:"includeCommandLineAPI"`
	
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent	bool	`json:"silent"`
	
	// Specifies in which execution context to perform evaluation. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ContextId	ExecutionContextId	`json:"contextId,omitempty"`
	
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue	bool	`json:"returnByValue"`
	
	// Whether preview should be generated for the result.
	GeneratePreview	bool	`json:"generatePreview"`
	
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture	bool	`json:"userGesture"`
	
	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise	bool	`json:"awaitPromise"`
	
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect	bool	`json:"throwOnSideEffect"`
	
	// Terminate execution after timing out (number of milliseconds).
	Timeout	TimeDelta	`json:"timeout"`
	
}

// Evaluate returns
type EvaluateReturns struct {
	
	// Evaluation result.
	Result	RemoteObject	`json:"result"`
	
	// Exception details.
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
}

// GetIsolateId parameters
type GetIsolateIdParams struct {
	
}

// GetIsolateId returns
type GetIsolateIdReturns struct {
	
	// The isolate id.
	Id	string	`json:"id"`
	
}

// GetHeapUsage parameters
type GetHeapUsageParams struct {
	
}

// GetHeapUsage returns
type GetHeapUsageReturns struct {
	
	// Used heap size in bytes.
	UsedSize	float64	`json:"usedSize"`
	
	// Allocated heap size in bytes.
	TotalSize	float64	`json:"totalSize"`
	
}

// GetProperties parameters
type GetPropertiesParams struct {
	
	// Identifier of the object to return properties for.
	ObjectId	RemoteObjectId	`json:"objectId"`
	
	// If true, returns properties belonging only to the element itself, not to its prototype
	// chain.
	OwnProperties	bool	`json:"ownProperties"`
	
	// If true, returns accessor properties (with getter/setter) only; internal properties are not
	// returned either.
	AccessorPropertiesOnly	bool	`json:"accessorPropertiesOnly"`
	
	// Whether preview should be generated for the results.
	GeneratePreview	bool	`json:"generatePreview"`
	
}

// GetProperties returns
type GetPropertiesReturns struct {
	
	// Object properties.
	Result	[]PropertyDescriptor	`json:"result"`
	
	// Internal object properties (only of the element itself).
	InternalProperties	[]InternalPropertyDescriptor	`json:"internalProperties"`
	
	// Object private properties.
	PrivateProperties	[]PrivatePropertyDescriptor	`json:"privateProperties"`
	
	// Exception details.
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
}

// GlobalLexicalScopeNames parameters
type GlobalLexicalScopeNamesParams struct {
	
	// Specifies in which execution context to lookup global scope variables.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
	
}

// GlobalLexicalScopeNames returns
type GlobalLexicalScopeNamesReturns struct {
	
	
	Names	[]string	`json:"names"`
	
}

// QueryObjects parameters
type QueryObjectsParams struct {
	
	// Identifier of the prototype to return objects for.
	PrototypeObjectId	RemoteObjectId	`json:"prototypeObjectId"`
	
	// Symbolic group name that can be used to release the results.
	ObjectGroup	string	`json:"objectGroup"`
	
}

// QueryObjects returns
type QueryObjectsReturns struct {
	
	// Array with objects.
	Objects	RemoteObject	`json:"objects"`
	
}

// ReleaseObject parameters
type ReleaseObjectParams struct {
	
	// Identifier of the object to release.
	ObjectId	RemoteObjectId	`json:"objectId"`
	
}

// ReleaseObject returns
type ReleaseObjectReturns struct {
	
}

// ReleaseObjectGroup parameters
type ReleaseObjectGroupParams struct {
	
	// Symbolic object group name.
	ObjectGroup	string	`json:"objectGroup"`
	
}

// ReleaseObjectGroup returns
type ReleaseObjectGroupReturns struct {
	
}

// RunIfWaitingForDebugger parameters
type RunIfWaitingForDebuggerParams struct {
	
}

// RunIfWaitingForDebugger returns
type RunIfWaitingForDebuggerReturns struct {
	
}

// RunScript parameters
type RunScriptParams struct {
	
	// Id of the script to run.
	ScriptId	ScriptId	`json:"scriptId"`
	
	// Specifies in which execution context to perform script run. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId,omitempty"`
	
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup	string	`json:"objectGroup"`
	
	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent	bool	`json:"silent"`
	
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI	bool	`json:"includeCommandLineAPI"`
	
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue	bool	`json:"returnByValue"`
	
	// Whether preview should be generated for the result.
	GeneratePreview	bool	`json:"generatePreview"`
	
	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise	bool	`json:"awaitPromise"`
	
}

// RunScript returns
type RunScriptReturns struct {
	
	// Run result.
	Result	RemoteObject	`json:"result"`
	
	// Exception details.
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
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

// SetCustomObjectFormatterEnabled parameters
type SetCustomObjectFormatterEnabledParams struct {
	
	
	Enabled	bool	`json:"enabled"`
	
}

// SetCustomObjectFormatterEnabled returns
type SetCustomObjectFormatterEnabledReturns struct {
	
}

// SetMaxCallStackSizeToCapture parameters
type SetMaxCallStackSizeToCaptureParams struct {
	
	
	Size	int	`json:"size"`
	
}

// SetMaxCallStackSizeToCapture returns
type SetMaxCallStackSizeToCaptureReturns struct {
	
}

// TerminateExecution parameters
type TerminateExecutionParams struct {
	
}

// TerminateExecution returns
type TerminateExecutionReturns struct {
	
}

// AddBinding parameters
type AddBindingParams struct {
	
	
	Name	string	`json:"name"`
	
	
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
	
}

// AddBinding returns
type AddBindingReturns struct {
	
}

// RemoveBinding parameters
type RemoveBindingParams struct {
	
	
	Name	string	`json:"name"`
	
}

// RemoveBinding returns
type RemoveBindingReturns struct {
	
}

