package profiler

import (

	"github.com/diiyw/chr/protocol/debugger"

)
const (
	
	
	ConsoleProfileFinishedEvent = "Profiler.consoleProfileFinished"
	
	// Sent when new profile recording is started using console.profile() call.
	ConsoleProfileStartedEvent = "Profiler.consoleProfileStarted"
	
)


type ConsoleProfileFinishedParams struct {
	
	
	Id	string	`json:"id"`
	
	// Location of console.profileEnd().
	Location	debugger.Location	`json:"location"`
	
	
	Profile	Profile	`json:"profile"`
	
	// Profile title passed as an argument to console.profile().
	Title	string	`json:"title"`
	
}

// Sent when new profile recording is started using console.profile() call.
type ConsoleProfileStartedParams struct {
	
	
	Id	string	`json:"id"`
	
	// Location of console.profile().
	Location	debugger.Location	`json:"location"`
	
	// Profile title passed as an argument to console.profile().
	Title	string	`json:"title"`
	
}

