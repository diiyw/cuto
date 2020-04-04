package backgroundservice

// Called when the recording state for the service has been updated.
const RecordingStateChangedEvent = "BackgroundService.recordingStateChanged"
type RecordingStateChangedParams struct {

	// 
	IsRecording 	bool
	// 
	Service 	ServiceName}



// Called with all existing backgroundServiceEvents when enabled, and all new
// events afterwards if enabled and recording.
const BackgroundServiceEventReceivedEvent = "BackgroundService.backgroundServiceEventReceived"
type BackgroundServiceEventReceivedParams struct {

	// 
	BackgroundServiceEvent 	BackgroundServiceEvent}

