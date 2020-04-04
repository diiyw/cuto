package backgroundservice

// Enables event updates for the service.
const StartObserving = "BackgroundService.startObserving"

type StartObservingParams struct {

	// 
	Service 	ServiceName	`json:"service"`
}

type StartObservingResult struct {

}

// Disables event updates for the service.
const StopObserving = "BackgroundService.stopObserving"

type StopObservingParams struct {

	// 
	Service 	ServiceName	`json:"service"`
}

type StopObservingResult struct {

}

// Set the recording state for the service.
const SetRecording = "BackgroundService.setRecording"

type SetRecordingParams struct {

	// 
	ShouldRecord 	bool	`json:"shouldRecord"`

	// 
	Service 	ServiceName	`json:"service"`
}

type SetRecordingResult struct {

}

// Clears all stored data for the service.
const ClearEvents = "BackgroundService.clearEvents"

type ClearEventsParams struct {

	// 
	Service 	ServiceName	`json:"service"`
}

type ClearEventsResult struct {

}