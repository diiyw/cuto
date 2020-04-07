package runtime

// Notification is issued every time when binding is called.
const BindingCalledEvent = "Runtime.bindingCalled"
type BindingCalledParams struct {

	// 
	Name 	string
	// 
	Payload 	string
	// Identifier of the context where the call was made.
	ExecutionContextId 	ExecutionContextId}



// Issued when console API was called.
const ConsoleAPICalledEvent = "Runtime.consoleAPICalled"
type ConsoleAPICalledParams struct {

	// Type of the call.
	Type 	string
	// Call arguments.
	Args 	[]*RemoteObject
	// Identifier of the context where the call was made.
	ExecutionContextId 	ExecutionContextId
	// Call timestamp.
	Timestamp 	Timestamp
	// Stack trace captured when the call was made. The async stack chain is automatically reported for
	// the following call types: `assert`, `error`, `trace`, `warning`. For other types the async call
	// chain can be retrieved using `Debugger.getStackTrace` and `stackTrace.parentId` field.
	StackTrace 	StackTrace
	// Console context descriptor for calls on non-default console context (not console.*):
	// 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call
	// on named context.
	Context 	string}



// Issued when unhandled exception was revoked.
const ExceptionRevokedEvent = "Runtime.exceptionRevoked"
type ExceptionRevokedParams struct {

	// Reason describing why exception was revoked.
	Reason 	string
	// The id of revoked exception, as reported in `exceptionThrown`.
	ExceptionId 	int}



// Issued when exception was thrown and unhandled.
const ExceptionThrownEvent = "Runtime.exceptionThrown"
type ExceptionThrownParams struct {

	// Timestamp of the exception.
	Timestamp 	Timestamp
	// 
	ExceptionDetails 	ExceptionDetails}



// Issued when new execution context is created.
const ExecutionContextCreatedEvent = "Runtime.executionContextCreated"
type ExecutionContextCreatedParams struct {

	// A newly created execution context.
	Context 	ExecutionContextDescription}



// Issued when execution context is destroyed.
const ExecutionContextDestroyedEvent = "Runtime.executionContextDestroyed"
type ExecutionContextDestroyedParams struct {

	// Id of the destroyed context
	ExecutionContextId 	ExecutionContextId}



// Issued when all executionContexts were cleared in browser
const ExecutionContextsClearedEvent = "Runtime.executionContextsCleared"
type ExecutionContextsClearedParams struct {
}



// Issued when object should be inspected (for example, as a result of inspect() command line API
// call).
const InspectRequestedEvent = "Runtime.inspectRequested"
type InspectRequestedParams struct {

	// 
	Object 	RemoteObject
	// 
	Hints 	interface{}}

