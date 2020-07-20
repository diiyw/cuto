package profiler

import (
	"github.com/diiyw/cuto/protocol/debugger"
)


// 
const ConsoleProfileFinishedEvent = "Profiler.consoleProfileFinished"
type ConsoleProfileFinishedParams struct {

	// 
	Id 	string
	// Location of console.profileEnd().
	Location 	debugger.Location
	// 
	Profile 	Profile
	// Profile title passed as an argument to console.profile().
	Title 	string}



// Sent when new profile recording is started using console.profile() call.
const ConsoleProfileStartedEvent = "Profiler.consoleProfileStarted"
type ConsoleProfileStartedParams struct {

	// 
	Id 	string
	// Location of console.profile().
	Location 	debugger.Location
	// Profile title passed as an argument to console.profile().
	Title 	string}

