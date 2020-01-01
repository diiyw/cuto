package webaudio

const (
	
	// Notifies that a new BaseAudioContext has been created.
	ContextCreatedEvent = "WebAudio.contextCreated"
	
	// Notifies that existing BaseAudioContext has been destroyed.
	ContextDestroyedEvent = "WebAudio.contextDestroyed"
	
	// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
	ContextChangedEvent = "WebAudio.contextChanged"
	
)

// Notifies that a new BaseAudioContext has been created.
type ContextCreatedParams struct {
	
	
	Context	BaseAudioContext	`json:"context"`
	
}

// Notifies that existing BaseAudioContext has been destroyed.
type ContextDestroyedParams struct {
	
	
	ContextId	ContextId	`json:"contextId"`
	
}

// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
type ContextChangedParams struct {
	
	
	Context	BaseAudioContext	`json:"context"`
	
}

