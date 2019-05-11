package backgroundservice

const (
	
	// Called when the recording state for the service has been updated.
	RecordingStateChangedEvent = "BackgroundService.recordingStateChanged"
	
	// Called with all existing backgroundServiceEvents when enabled, and all new
	// events afterwards if enabled and recording.
	BackgroundServiceEventReceivedEvent = "BackgroundService.backgroundServiceEventReceived"
	
)

// Called when the recording state for the service has been updated.
type RecordingStateChangedParams struct {
	
	
	IsRecording	bool	`json:"isRecording"`
	
	
	Service	ServiceName	`json:"service"`
	
}

// Called with all existing backgroundServiceEvents when enabled, and all new
	// events afterwards if enabled and recording.
type BackgroundServiceEventReceivedParams struct {
	
	
	BackgroundServiceEvent	BackgroundServiceEvent	`json:"backgroundServiceEvent"`
	
}

