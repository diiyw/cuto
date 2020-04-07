package log

import (
	"github.com/diiyw/cuto/protocol/runtime"
)

// Log entry.
type LogEntry  struct {

	// Log entry source.
	Source	string	`json:"source"`

	// Log entry severity.
	Level	string	`json:"level"`

	// Logged text.
	Text	string	`json:"text"`

	// Timestamp when this entry was added.
	Timestamp	interface{}	`json:"timestamp"`

	// URL of the resource if known.
	Url	string	`json:"url"`

	// Line number in the resource.
	LineNumber	int	`json:"lineNumber"`

	// JavaScript stack trace.
	StackTrace	interface{}	`json:"stackTrace"`

	// Identifier of the network request associated with this entry.
	NetworkRequestId	interface{}	`json:"networkRequestId"`

	// Identifier of the worker associated with this entry.
	WorkerId	string	`json:"workerId"`

	// Call arguments.
	Args	[]*runtime.RemoteObject	`json:"args"`
}

// Violation configuration setting.
type ViolationSetting  struct {

	// Violation type.
	Name	string	`json:"name"`

	// Time threshold to trigger upon.
	Threshold	float64	`json:"threshold"`
}
