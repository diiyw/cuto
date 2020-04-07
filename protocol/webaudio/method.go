package webaudio

// Enables the WebAudio domain and starts sending context lifetime events.
const Enable = "WebAudio.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Disables the WebAudio domain.
const Disable = "WebAudio.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Fetch the realtime data from the registered contexts.
const GetRealtimeData = "WebAudio.getRealtimeData"

type GetRealtimeDataParams struct {

	// 
	ContextId 	GraphObjectId	`json:"contextId"`
}

type GetRealtimeDataResult struct {

	// 
	RealtimeData 	ContextRealtimeData	`json:"realtimeData"`
}