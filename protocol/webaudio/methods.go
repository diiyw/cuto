package webaudio

const (
	
	// Enables the WebAudio domain and starts sending context lifetime events.
	Enable = "WebAudio.enable"
	
	// Disables the WebAudio domain.
	Disable = "WebAudio.disable"
	
	// Fetch the realtime data from the registered contexts.
	GetRealtimeData = "WebAudio.getRealtimeData"
	
)

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// GetRealtimeData parameters
type GetRealtimeDataParams struct {
	
	
	ContextId	ContextId	`json:"contextId"`
	
}

// GetRealtimeData returns
type GetRealtimeDataReturns struct {
	
	
	RealtimeData	ContextRealtimeData	`json:"realtimeData"`
	
}

