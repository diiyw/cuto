package profiler

// 
const ConsoleProfileFinishedEvent = "Profiler.consoleProfileFinished"
type ConsoleProfileFinishedParams struct {

	// 
	Id 	string
	// Location of console.profileEnd().
	Location 	interface{}
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
	Location 	interface{}
	// Profile title passed as an argument to console.profile().
	Title 	string}

