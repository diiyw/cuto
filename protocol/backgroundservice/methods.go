package backgroundservice

const (
	
	// Enables event updates for the service.
	StartObserving = "BackgroundService.startObserving"
	
	// Disables event updates for the service.
	StopObserving = "BackgroundService.stopObserving"
	
	// Set the recording state for the service.
	SetRecording = "BackgroundService.setRecording"
	
	// Clears all stored data for the service.
	ClearEvents = "BackgroundService.clearEvents"
	
)

// StartObserving parameters
type StartObservingParams struct {
	
	
	Service	ServiceName	`json:"service"`
	
}

// StartObserving returns
type StartObservingReturns struct {
	
}

// StopObserving parameters
type StopObservingParams struct {
	
	
	Service	ServiceName	`json:"service"`
	
}

// StopObserving returns
type StopObservingReturns struct {
	
}

// SetRecording parameters
type SetRecordingParams struct {
	
	
	ShouldRecord	bool	`json:"shouldRecord"`
	
	
	Service	ServiceName	`json:"service"`
	
}

// SetRecording returns
type SetRecordingReturns struct {
	
}

// ClearEvents parameters
type ClearEventsParams struct {
	
	
	Service	ServiceName	`json:"service"`
	
}

// ClearEvents returns
type ClearEventsReturns struct {
	
}

