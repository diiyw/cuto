package log

import (

	"github.com/diiyw/gator/protocol/network"

	"github.com/diiyw/gator/protocol/runtime"

)

// Log entry.
type LogEntry struct {
	
	// Log entry source.
	// Possible value: xml,javascript,network,storage,appcache,rendering,security,deprecation,worker,violation,intervention,recommendation,other,
	Source	string	`json:"source"`
	
	// Log entry severity.
	// Possible value: verbose,info,warning,error,
	Level	string	`json:"level"`
	
	// Logged text.
	
	Text	string	`json:"text"`
	
	// Timestamp when this entry was added.
	
	Timestamp	runtime.Timestamp	`json:"timestamp"`
	
	// URL of the resource if known.
	
	Url	string	`json:"url"`
	
	// Line number in the resource.
	
	LineNumber	int	`json:"lineNumber"`
	
	// JavaScript stack trace.
	
	StackTrace	runtime.StackTrace	`json:"stackTrace"`
	
	// Identifier of the network request associated with this entry.
	
	NetworkRequestId	network.RequestId	`json:"networkRequestId"`
	
	// Identifier of the worker associated with this entry.
	
	WorkerId	string	`json:"workerId"`
	
	// Call arguments.
	
	Args	[]runtime.RemoteObject	`json:"args"`
	
}	

// Violation configuration setting.
type ViolationSetting struct {
	
	// Violation type.
	// Possible value: longTask,longLayout,blockedEvent,blockedParser,discouragedAPIUse,handler,recurringHandler,
	Name	string	`json:"name"`
	
	// Time threshold to trigger upon.
	
	Threshold	float64	`json:"threshold"`
	
}	

