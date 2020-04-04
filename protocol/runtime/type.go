package runtime
// Unique script identifier.
type ScriptId string

// Unique object identifier.
type RemoteObjectId string

// Primitive value which cannot be JSON-stringified. Includes values `-0`, `NaN`, `Infinity`,
	// `-Infinity`, and bigint literals.
type UnserializableValue string

// Mirror object referencing original JavaScript object.
type RemoteObject  struct {

	// Object type.
	Type	string	`json:"type"`

	// Object subtype hint. Specified for `object` type values only.
	Subtype	string	`json:"subtype"`

	// Object class (constructor) name. Specified for `object` type values only.
	ClassName	string	`json:"className"`

	// Remote object value in case of primitive values or JSON values (if it was requested).
	Value	interface{}	`json:"value"`

	// Primitive value which can not be JSON-stringified does not have `value`, but gets this
	// property.
	UnserializableValue	UnserializableValue	`json:"unserializableValue"`

	// String representation of the object.
	Description	string	`json:"description"`

	// Unique object identifier (for non-primitive values).
	ObjectId	RemoteObjectId	`json:"objectId"`

	// Preview containing abbreviated property values. Specified for `object` type values only.
	Preview	ObjectPreview	`json:"preview"`

	// 
	CustomPreview	CustomPreview	`json:"customPreview"`
}

// 
type CustomPreview  struct {

	// The JSON-stringified result of formatter.header(object, config) call.
	// It contains json ML array that represents RemoteObject.
	Header	string	`json:"header"`

	// If formatter returns true as a result of formatter.hasBody call then bodyGetterId will
	// contain RemoteObjectId for the function that returns result of formatter.body(object, config) call.
	// The result value is json ML array.
	BodyGetterId	RemoteObjectId	`json:"bodyGetterId"`
}

// Object containing abbreviated remote object value.
type ObjectPreview  struct {

	// Object type.
	Type	string	`json:"type"`

	// Object subtype hint. Specified for `object` type values only.
	Subtype	string	`json:"subtype"`

	// String representation of the object.
	Description	string	`json:"description"`

	// True iff some of the properties or entries of the original object did not fit.
	Overflow	bool	`json:"overflow"`

	// List of the properties.
	Properties	[]*PropertyPreview	`json:"properties"`

	// List of the entries. Specified for `map` and `set` subtype values only.
	Entries	[]*EntryPreview	`json:"entries"`
}

// 
type PropertyPreview  struct {

	// Property name.
	Name	string	`json:"name"`

	// Object type. Accessor means that the property itself is an accessor property.
	Type	string	`json:"type"`

	// User-friendly property value string.
	Value	string	`json:"value"`

	// Nested value preview.
	ValuePreview	ObjectPreview	`json:"valuePreview"`

	// Object subtype hint. Specified for `object` type values only.
	Subtype	string	`json:"subtype"`
}

// 
type EntryPreview  struct {

	// Preview of the key. Specified for map-like collection entries.
	Key	ObjectPreview	`json:"key"`

	// Preview of the value.
	Value	ObjectPreview	`json:"value"`
}

// Object property descriptor.
type PropertyDescriptor  struct {

	// Property name or symbol description.
	Name	string	`json:"name"`

	// The value associated with the property.
	Value	RemoteObject	`json:"value"`

	// True if the value associated with the property may be changed (data descriptors only).
	Writable	bool	`json:"writable"`

	// A function which serves as a getter for the property, or `undefined` if there is no getter
	// (accessor descriptors only).
	Get	RemoteObject	`json:"get"`

	// A function which serves as a setter for the property, or `undefined` if there is no setter
	// (accessor descriptors only).
	Set	RemoteObject	`json:"set"`

	// True if the type of this property descriptor may be changed and if the property may be
	// deleted from the corresponding object.
	Configurable	bool	`json:"configurable"`

	// True if this property shows up during enumeration of the properties on the corresponding
	// object.
	Enumerable	bool	`json:"enumerable"`

	// True if the result was thrown during the evaluation.
	WasThrown	bool	`json:"wasThrown"`

	// True if the property is owned for the object.
	IsOwn	bool	`json:"isOwn"`

	// Property symbol object, if the property is of the `symbol` type.
	Symbol	RemoteObject	`json:"symbol"`
}

// Object internal property descriptor. This property isn't normally visible in JavaScript code.
type InternalPropertyDescriptor  struct {

	// Conventional property name.
	Name	string	`json:"name"`

	// The value associated with the property.
	Value	RemoteObject	`json:"value"`
}

// Object private field descriptor.
type PrivatePropertyDescriptor  struct {

	// Private property name.
	Name	string	`json:"name"`

	// The value associated with the private property.
	Value	RemoteObject	`json:"value"`
}

// Represents function call argument. Either remote object id `objectId`, primitive `value`,
	// unserializable primitive value or neither of (for undefined) them should be specified.
type CallArgument  struct {

	// Primitive value or serializable javascript object.
	Value	interface{}	`json:"value"`

	// Primitive value which can not be JSON-stringified.
	UnserializableValue	UnserializableValue	`json:"unserializableValue"`

	// Remote object handle.
	ObjectId	RemoteObjectId	`json:"objectId"`
}

// Id of an execution context.
type ExecutionContextId int

// Description of an isolated world.
type ExecutionContextDescription  struct {

	// Unique id of the execution context. It can be used to specify in which execution context
	// script evaluation should be performed.
	Id	ExecutionContextId	`json:"id"`

	// Execution context origin.
	Origin	string	`json:"origin"`

	// Human readable name describing given context.
	Name	string	`json:"name"`

	// Embedder-specific auxiliary data.
	AuxData	interface{}	`json:"auxData"`
}

// Detailed information about exception (or error) that was thrown during script compilation or
	// execution.
type ExceptionDetails  struct {

	// Exception id.
	ExceptionId	int	`json:"exceptionId"`

	// Exception text, which should be used together with exception object when available.
	Text	string	`json:"text"`

	// Line number of the exception location (0-based).
	LineNumber	int	`json:"lineNumber"`

	// Column number of the exception location (0-based).
	ColumnNumber	int	`json:"columnNumber"`

	// Script ID of the exception location.
	ScriptId	ScriptId	`json:"scriptId"`

	// URL of the exception location, to be used when the script was not reported.
	Url	string	`json:"url"`

	// JavaScript stack trace if available.
	StackTrace	StackTrace	`json:"stackTrace"`

	// Exception object if available.
	Exception	RemoteObject	`json:"exception"`

	// Identifier of the context where exception happened.
	ExecutionContextId	ExecutionContextId	`json:"executionContextId"`
}

// Number of milliseconds since epoch.
type Timestamp float64

// Number of milliseconds.
type TimeDelta float64

// Stack entry for runtime errors and assertions.
type CallFrame  struct {

	// JavaScript function name.
	FunctionName	string	`json:"functionName"`

	// JavaScript script id.
	ScriptId	ScriptId	`json:"scriptId"`

	// JavaScript script name or url.
	Url	string	`json:"url"`

	// JavaScript script line number (0-based).
	LineNumber	int	`json:"lineNumber"`

	// JavaScript script column number (0-based).
	ColumnNumber	int	`json:"columnNumber"`
}

// Call frames for assertions or error messages.
type StackTrace  struct {

	// String label of this stack trace. For async traces this may be a name of the function that
	// initiated the async call.
	Description	string	`json:"description"`

	// JavaScript function name.
	CallFrames	[]*CallFrame	`json:"callFrames"`

	// Asynchronous JavaScript stack trace that preceded this stack, if available.
	Parent	*StackTrace	`json:"parent"`

	// Asynchronous JavaScript stack trace that preceded this stack, if available.
	ParentId	StackTraceId	`json:"parentId"`
}

// Unique identifier of current debugger.
type UniqueDebuggerId string

// If `debuggerId` is set stack trace comes from another debugger and can be resolved there. This
	// allows to track cross-debugger calls. See `Runtime.StackTrace` and `Debugger.paused` for usages.
type StackTraceId  struct {

	// 
	Id	string	`json:"id"`

	// 
	DebuggerId	UniqueDebuggerId	`json:"debuggerId"`
}
