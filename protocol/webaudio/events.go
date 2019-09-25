package webaudio

const (
	
	// Notifies that a new BaseAudioContext has been created.
	ContextCreatedEvent = "WebAudio.contextCreated"
	
	// Notifies that an existing BaseAudioContext will be destroyed.
	ContextWillBeDestroyedEvent = "WebAudio.contextWillBeDestroyed"
	
	// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
	ContextChangedEvent = "WebAudio.contextChanged"
	
	// Notifies that the construction of an AudioListener has finished.
	AudioListenerCreatedEvent = "WebAudio.audioListenerCreated"
	
	// Notifies that a new AudioListener has been created.
	AudioListenerWillBeDestroyedEvent = "WebAudio.audioListenerWillBeDestroyed"
	
	// Notifies that a new AudioNode has been created.
	AudioNodeCreatedEvent = "WebAudio.audioNodeCreated"
	
	// Notifies that an existing AudioNode has been destroyed.
	AudioNodeWillBeDestroyedEvent = "WebAudio.audioNodeWillBeDestroyed"
	
	// Notifies that a new AudioParam has been created.
	AudioParamCreatedEvent = "WebAudio.audioParamCreated"
	
	// Notifies that an existing AudioParam has been destroyed.
	AudioParamWillBeDestroyedEvent = "WebAudio.audioParamWillBeDestroyed"
	
	// Notifies that two AudioNodes are connected.
	NodesConnectedEvent = "WebAudio.nodesConnected"
	
	// Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
	NodesDisconnectedEvent = "WebAudio.nodesDisconnected"
	
	// Notifies that an AudioNode is connected to an AudioParam.
	NodeParamConnectedEvent = "WebAudio.nodeParamConnected"
	
	// Notifies that an AudioNode is disconnected to an AudioParam.
	NodeParamDisconnectedEvent = "WebAudio.nodeParamDisconnected"
	
)

// Notifies that a new BaseAudioContext has been created.
type ContextCreatedParams struct {
	
	
	Context	BaseAudioContext	`json:"context"`
	
}

// Notifies that an existing BaseAudioContext will be destroyed.
type ContextWillBeDestroyedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
}

// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
type ContextChangedParams struct {
	
	
	Context	BaseAudioContext	`json:"context"`
	
}

// Notifies that the construction of an AudioListener has finished.
type AudioListenerCreatedParams struct {
	
	
	Listener	AudioListener	`json:"listener"`
	
}

// Notifies that a new AudioListener has been created.
type AudioListenerWillBeDestroyedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	ListenerId	GraphObjectId	`json:"listenerId"`
	
}

// Notifies that a new AudioNode has been created.
type AudioNodeCreatedParams struct {
	
	
	Node	AudioNode	`json:"node"`
	
}

// Notifies that an existing AudioNode has been destroyed.
type AudioNodeWillBeDestroyedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	NodeId	GraphObjectId	`json:"nodeId"`
	
}

// Notifies that a new AudioParam has been created.
type AudioParamCreatedParams struct {
	
	
	Param	AudioParam	`json:"param"`
	
}

// Notifies that an existing AudioParam has been destroyed.
type AudioParamWillBeDestroyedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	NodeId	GraphObjectId	`json:"nodeId"`
	
	
	ParamId	GraphObjectId	`json:"paramId"`
	
}

// Notifies that two AudioNodes are connected.
type NodesConnectedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	SourceId	GraphObjectId	`json:"sourceId"`
	
	
	DestinationId	GraphObjectId	`json:"destinationId"`
	
	
	SourceOutputIndex	float64	`json:"sourceOutputIndex"`
	
	
	DestinationInputIndex	float64	`json:"destinationInputIndex"`
	
}

// Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
type NodesDisconnectedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	SourceId	GraphObjectId	`json:"sourceId"`
	
	
	DestinationId	GraphObjectId	`json:"destinationId"`
	
	
	SourceOutputIndex	float64	`json:"sourceOutputIndex"`
	
	
	DestinationInputIndex	float64	`json:"destinationInputIndex"`
	
}

// Notifies that an AudioNode is connected to an AudioParam.
type NodeParamConnectedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	SourceId	GraphObjectId	`json:"sourceId"`
	
	
	DestinationId	GraphObjectId	`json:"destinationId"`
	
	
	SourceOutputIndex	float64	`json:"sourceOutputIndex"`
	
}

// Notifies that an AudioNode is disconnected to an AudioParam.
type NodeParamDisconnectedParams struct {
	
	
	ContextId	GraphObjectId	`json:"contextId"`
	
	
	SourceId	GraphObjectId	`json:"sourceId"`
	
	
	DestinationId	GraphObjectId	`json:"destinationId"`
	
	
	SourceOutputIndex	float64	`json:"sourceOutputIndex"`
	
}

