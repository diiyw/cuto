package log

import (
	"github.com/diiyw/cuto/protocol/runtime"
	"github.com/diiyw/cuto/protocol/network"
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
	Timestamp	runtime.Timestamp	`json:"timestamp"`

	// URL of the resource if known.
	Url	string	`json:"url,omitempty"`

	// Line number in the resource.
	LineNumber	int	`json:"lineNumber,omitempty"`

	// JavaScript stack trace.
	StackTrace	runtime.StackTrace	`json:"stackTrace,omitempty"`

	// Identifier of the network request associated with this entry.
	NetworkRequestId	network.RequestId	`json:"networkRequestId,omitempty"`

	// Identifier of the worker associated with this entry.
	WorkerId	string	`json:"workerId,omitempty"`

	// Call arguments.
	Args	[]*runtime.RemoteObject	`json:"args,omitempty"`
}

// Violation configuration setting.
type ViolationSetting  struct {

	// Violation type.
	Name	string	`json:"name"`

	// Time threshold to trigger upon.
	Threshold	float64	`json:"threshold"`
}
