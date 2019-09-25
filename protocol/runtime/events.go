package runtime

const (
	
	// Notification is issued every time when binding is called.
	BindingCalledEvent = "Runtime.bindingCalled"
	
	// Issued when console API was called.
	ConsoleAPICalledEvent = "Runtime.consoleAPICalled"
	
	// Issued when unhandled exception was revoked.
	ExceptionRevokedEvent = "Runtime.exceptionRevoked"
	
	// Issued when exception was thrown and unhandled.
	ExceptionThrownEvent = "Runtime.exceptionThrown"
	
	// Issued when new execution context is created.
	ExecutionContextCreatedEvent = "Runtime.executionContextCreated"
	
	// Issued when execution context is destroyed.
	ExecutionContextDestroyedEvent = "Runtime.executionContextDestroyed"
	
	// Issued when all executionContexts were cleared in browser
	ExecutionContextsClearedEvent = "Runtime.executionContextsCleared"
	
	// Issued when object should be inspected (for example, as a result of inspect() command line API
	// call).
	InspectRequestedEvent = "Runtime.inspectRequested"
	
)

// Notification is issued every time when binding is called.
type BindingCalledParams struct {
	
	
	Name	string	`json:"name"`
	
	
	Payload	string	`json:"payload"`
	
	// Identifier of the context where the call was made.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
	
}

// Issued when console API was called.
type ConsoleAPICalledParams struct {
	
	// Type of the call.
	Type	string	`json:"type"`
	
	// Call arguments.
	Args	[]RemoteObject	`json:"args"`
	
	// Identifier of the context where the call was made.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
	
	// Call timestamp.
	Timestamp	Timestamp	`json:"timestamp"`
	
	// Stack trace captured when the call was made. The async stack chain is automatically reported for
	// the following call types: `assert`, `error`, `trace`, `warning`. For other types the async call
	// chain can be retrieved using `Debugger.getStackTrace` and `stackTrace.parentId` field.
	StackTrace	StackTrace	`json:"stackTrace"`
	
	// Console context descriptor for calls on non-default console context (not console.*):
	// 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call
	// on named context.
	Context	string	`json:"context"`
	
}

// Issued when unhandled exception was revoked.
type ExceptionRevokedParams struct {
	
	// Reason describing why exception was revoked.
	Reason	string	`json:"reason"`
	
	// The id of revoked exception, as reported in `exceptionThrown`.
	ExceptionId	int	`json:"exceptionId"`
	
}

// Issued when exception was thrown and unhandled.
type ExceptionThrownParams struct {
	
	// Timestamp of the exception.
	Timestamp	Timestamp	`json:"timestamp"`
	
	
	ExceptionDetails	ExceptionDetails	`json:"exceptionDetails"`
	
}

// Issued when new execution context is created.
type ExecutionContextCreatedParams struct {
	
	// A newly created execution context.
	Context	ExecutionContextDescription	`json:"context"`
	
}

// Issued when execution context is destroyed.
type ExecutionContextDestroyedParams struct {
	
	// Id of the destroyed context
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
	
}

// Issued when all executionContexts were cleared in browser
type ExecutionContextsClearedParams struct {
	
}

// Issued when object should be inspected (for example, as a result of inspect() command line API
	// call).
type InspectRequestedParams struct {
	
	
	Object	RemoteObject	`json:"object"`
	
	
	Hints	interface{}	`json:"hints"`
	
}

