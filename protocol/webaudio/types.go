package webaudio


// Context's UUID in string
type ContextId string	

// Enum of BaseAudioContext types
type ContextType string	

// Enum of AudioContextState from the spec
type ContextState string	

// Fields in AudioContext that change in real-time. These are not updated
	// on OfflineAudioContext.
type ContextRealtimeData struct {
	
	// The current context time in second in BaseAudioContext.
	
	CurrentTime	float64	`json:"currentTime"`
	
	// The time spent on rendering graph divided by render qunatum duration,
	// and multiplied by 100. 100 means the audio renderer reached the full
	// capacity and glitch may occur.
	
	RenderCapacity	float64	`json:"renderCapacity"`
	
}	

// Protocol object for BaseAudioContext
type BaseAudioContext struct {
	
	
	
	ContextId	ContextId	`json:"contextId"`
	
	
	
	ContextType	ContextType	`json:"contextType"`
	
	
	
	ContextState	ContextState	`json:"contextState"`
	
	
	
	RealtimeData	ContextRealtimeData	`json:"realtimeData"`
	
	// Platform-dependent callback buffer size.
	
	CallbackBufferSize	float64	`json:"callbackBufferSize"`
	
	// Number of output channels supported by audio hardware in use.
	
	MaxOutputChannelCount	float64	`json:"maxOutputChannelCount"`
	
	// Context sample rate.
	
	SampleRate	float64	`json:"sampleRate"`
	
}	

