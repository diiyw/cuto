package runtime

// Add handler to promise with given promise object id.
const AwaitPromise = "Runtime.awaitPromise"

type AwaitPromiseParams struct {

	// Identifier of the promise.
	PromiseObjectId 	RemoteObjectId	`json:"promiseObjectId"`

	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue 	bool	`json:"returnByValue,omitempty"`

	// Whether preview should be generated for the result.
	GeneratePreview 	bool	`json:"generatePreview,omitempty"`
}

type AwaitPromiseResult struct {

	// Promise result. Will contain rejected value if promise was rejected.
	Result 	RemoteObject	`json:"result"`
	// Exception details if stack strace is available.
	ExceptionDetails 	ExceptionDetails	`json:"exceptionDetails"`
}

// Calls function with given declaration on the given object. Object group of the result is
// inherited from the target object.
const CallFunctionOn = "Runtime.callFunctionOn"

type CallFunctionOnParams struct {

	// Declaration of the function to call.
	FunctionDeclaration 	string	`json:"functionDeclaration"`

	// Identifier of the object to call function on. Either objectId or executionContextId should
	// be specified.
	ObjectId 	RemoteObjectId	`json:"objectId,omitempty"`

	// Call arguments. All call arguments must belong to the same JavaScript world as the target
	// object.
	Arguments 	[]*CallArgument	`json:"arguments,omitempty"`

	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent 	bool	`json:"silent,omitempty"`

	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue 	bool	`json:"returnByValue,omitempty"`

	// Whether preview should be generated for the result.
	GeneratePreview 	bool	`json:"generatePreview,omitempty"`

	// Whether execution should be treated as initiated by user in the UI.
	UserGesture 	bool	`json:"userGesture,omitempty"`

	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise 	bool	`json:"awaitPromise,omitempty"`

	// Specifies execution context which global object will be used to call function on. Either
	// executionContextId or objectId should be specified.
	ExecutionContextId 	ExecutionContextId	`json:"executionContextId,omitempty"`

	// Symbolic group name that can be used to release multiple objects. If objectGroup is not
	// specified and objectId is, objectGroup will be inherited from object.
	ObjectGroup 	string	`json:"objectGroup,omitempty"`
}

type CallFunctionOnResult struct {

	// Call result.
	Result 	RemoteObject	`json:"result"`
	// Exception details.
	ExceptionDetails 	ExceptionDetails	`json:"exceptionDetails"`
}

// Compiles expression.
const CompileScript = "Runtime.compileScript"

type CompileScriptParams struct {

	// Expression to compile.
	Expression 	string	`json:"expression"`

	// Source url to be set for the script.
	SourceURL 	string	`json:"sourceURL"`

	// Specifies whether the compiled script should be persisted.
	PersistScript 	bool	`json:"persistScript"`

	// Specifies in which execution context to perform script run. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ExecutionContextId 	ExecutionContextId	`json:"executionContextId,omitempty"`
}

type CompileScriptResult struct {

	// Id of the script.
	ScriptId 	ScriptId	`json:"scriptId"`
	// Exception details.
	ExceptionDetails 	ExceptionDetails	`json:"exceptionDetails"`
}

// Disables reporting of execution contexts creation.
const Disable = "Runtime.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Discards collected exceptions and console API calls.
const DiscardConsoleEntries = "Runtime.discardConsoleEntries"

type DiscardConsoleEntriesParams struct {
}

type DiscardConsoleEntriesResult struct {

}

// Enables reporting of execution contexts creation by means of `executionContextCreated` event.
// When the reporting gets enabled the event will be sent immediately for each existing execution
// context.
const Enable = "Runtime.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Evaluates expression on global object.
const Evaluate = "Runtime.evaluate"

type EvaluateParams struct {

	// Expression to evaluate.
	Expression 	string	`json:"expression"`

	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup 	string	`json:"objectGroup,omitempty"`

	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI 	bool	`json:"includeCommandLineAPI,omitempty"`

	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent 	bool	`json:"silent,omitempty"`

	// Specifies in which execution context to perform evaluation. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ContextId 	ExecutionContextId	`json:"contextId,omitempty"`

	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue 	bool	`json:"returnByValue,omitempty"`

	// Whether preview should be generated for the result.
	GeneratePreview 	bool	`json:"generatePreview,omitempty"`

	// Whether execution should be treated as initiated by user in the UI.
	UserGesture 	bool	`json:"userGesture,omitempty"`

	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise 	bool	`json:"awaitPromise,omitempty"`

	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	// This implies `disableBreaks` below.
	ThrowOnSideEffect 	bool	`json:"throwOnSideEffect,omitempty"`

	// Terminate execution after timing out (number of milliseconds).
	Timeout 	TimeDelta	`json:"timeout,omitempty"`

	// Disable breakpoints during execution.
	DisableBreaks 	bool	`json:"disableBreaks,omitempty"`
}

type EvaluateResult struct {

	// Evaluation result.
	Result 	RemoteObject	`json:"result"`
	// Exception details.
	ExceptionDetails 	ExceptionDetails	`json:"exceptionDetails"`
}

// Returns the isolate id.
const GetIsolateId = "Runtime.getIsolateId"

type GetIsolateIdParams struct {
}

type GetIsolateIdResult struct {

	// The isolate id.
	Id 	string	`json:"id"`
}

// Returns the JavaScript heap usage.
// It is the total usage of the corresponding isolate not scoped to a particular Runtime.
const GetHeapUsage = "Runtime.getHeapUsage"

type GetHeapUsageParams struct {
}

type GetHeapUsageResult struct {

	// Used heap size in bytes.
	UsedSize 	float64	`json:"usedSize"`
	// Allocated heap size in bytes.
	TotalSize 	float64	`json:"totalSize"`
}

// Returns properties of a given object. Object group of the result is inherited from the target
// object.
const GetProperties = "Runtime.getProperties"

type GetPropertiesParams struct {

	// Identifier of the object to return properties for.
	ObjectId 	RemoteObjectId	`json:"objectId"`

	// If true, returns properties belonging only to the element itself, not to its prototype
	// chain.
	OwnProperties 	bool	`json:"ownProperties,omitempty"`

	// If true, returns accessor properties (with getter/setter) only; internal properties are not
	// returned either.
	AccessorPropertiesOnly 	bool	`json:"accessorPropertiesOnly,omitempty"`

	// Whether preview should be generated for the results.
	GeneratePreview 	bool	`json:"generatePreview,omitempty"`
}

type GetPropertiesResult struct {

	// Object properties.
	Result 	[]*PropertyDescriptor	`json:"result"`
	// Internal object properties (only of the element itself).
	InternalProperties 	[]*InternalPropertyDescriptor	`json:"internalProperties"`
	// Object private properties.
	PrivateProperties 	[]*PrivatePropertyDescriptor	`json:"privateProperties"`
	// Exception details.
	ExceptionDetails 	ExceptionDetails	`json:"exceptionDetails"`
}

// Returns all let, const and class variables from global scope.
const GlobalLexicalScopeNames = "Runtime.globalLexicalScopeNames"

type GlobalLexicalScopeNamesParams struct {

	// Specifies in which execution context to lookup global scope variables.
	ExecutionContextId 	ExecutionContextId	`json:"executionContextId,omitempty"`
}

type GlobalLexicalScopeNamesResult struct {

	// 
	Names 	[]string	`json:"names"`
}

// 
const QueryObjects = "Runtime.queryObjects"

type QueryObjectsParams struct {

	// Identifier of the prototype to return objects for.
	PrototypeObjectId 	RemoteObjectId	`json:"prototypeObjectId"`

	// Symbolic group name that can be used to release the results.
	ObjectGroup 	string	`json:"objectGroup,omitempty"`
}

type QueryObjectsResult struct {

	// Array with objects.
	Objects 	RemoteObject	`json:"objects"`
}

// Releases remote object with given id.
const ReleaseObject = "Runtime.releaseObject"

type ReleaseObjectParams struct {

	// Identifier of the object to release.
	ObjectId 	RemoteObjectId	`json:"objectId"`
}

type ReleaseObjectResult struct {

}

// Releases all remote objects that belong to a given group.
const ReleaseObjectGroup = "Runtime.releaseObjectGroup"

type ReleaseObjectGroupParams struct {

	// Symbolic object group name.
	ObjectGroup 	string	`json:"objectGroup"`
}

type ReleaseObjectGroupResult struct {

}

// Tells inspected instance to run if it was waiting for debugger to attach.
const RunIfWaitingForDebugger = "Runtime.runIfWaitingForDebugger"

type RunIfWaitingForDebuggerParams struct {
}

type RunIfWaitingForDebuggerResult struct {

}

// Runs script with given id in a given context.
const RunScript = "Runtime.runScript"

type RunScriptParams struct {

	// Id of the script to run.
	ScriptId 	ScriptId	`json:"scriptId"`

	// Specifies in which execution context to perform script run. If the parameter is omitted the
	// evaluation will be performed in the context of the inspected page.
	ExecutionContextId 	ExecutionContextId	`json:"executionContextId,omitempty"`

	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup 	string	`json:"objectGroup,omitempty"`

	// In silent mode exceptions thrown during evaluation are not reported and do not pause
	// execution. Overrides `setPauseOnException` state.
	Silent 	bool	`json:"silent,omitempty"`

	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI 	bool	`json:"includeCommandLineAPI,omitempty"`

	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue 	bool	`json:"returnByValue,omitempty"`

	// Whether preview should be generated for the result.
	GeneratePreview 	bool	`json:"generatePreview,omitempty"`

	// Whether execution should `await` for resulting value and return once awaited promise is
	// resolved.
	AwaitPromise 	bool	`json:"awaitPromise,omitempty"`
}

type RunScriptResult struct {

	// Run result.
	Result 	RemoteObject	`json:"result"`
	// Exception details.
	ExceptionDetails 	ExceptionDetails	`json:"exceptionDetails"`
}

// Enables or disables async call stacks tracking.
const SetAsyncCallStackDepth = "Runtime.setAsyncCallStackDepth"

type SetAsyncCallStackDepthParams struct {

	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async
	// call stacks (default).
	MaxDepth 	int	`json:"maxDepth"`
}

type SetAsyncCallStackDepthResult struct {

}

// 
const SetCustomObjectFormatterEnabled = "Runtime.setCustomObjectFormatterEnabled"

type SetCustomObjectFormatterEnabledParams struct {

	// 
	Enabled 	bool	`json:"enabled"`
}

type SetCustomObjectFormatterEnabledResult struct {

}

// 
const SetMaxCallStackSizeToCapture = "Runtime.setMaxCallStackSizeToCapture"

type SetMaxCallStackSizeToCaptureParams struct {

	// 
	Size 	int	`json:"size"`
}

type SetMaxCallStackSizeToCaptureResult struct {

}

// Terminate current or next JavaScript execution.
// Will cancel the termination when the outer-most script execution ends.
const TerminateExecution = "Runtime.terminateExecution"

type TerminateExecutionParams struct {
}

type TerminateExecutionResult struct {

}

// If executionContextId is empty, adds binding with the given name on the
// global objects of all inspected contexts, including those created later,
// bindings survive reloads.
// If executionContextId is specified, adds binding only on global object of
// given execution context.
// Binding function takes exactly one argument, this argument should be string,
// in case of any other input, function throws an exception.
// Each binding function call produces Runtime.bindingCalled notification.
const AddBinding = "Runtime.addBinding"

type AddBindingParams struct {

	// 
	Name 	string	`json:"name"`

	// 
	ExecutionContextId 	ExecutionContextId	`json:"executionContextId,omitempty"`
}

type AddBindingResult struct {

}

// This method does not remove binding function from global object but
// unsubscribes current runtime agent from Runtime.bindingCalled notifications.
const RemoveBinding = "Runtime.removeBinding"

type RemoveBindingParams struct {

	// 
	Name 	string	`json:"name"`
}

type RemoveBindingResult struct {

}