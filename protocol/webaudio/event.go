package webaudio

// Notifies that a new BaseAudioContext has been created.
const ContextCreatedEvent = "WebAudio.contextCreated"
type ContextCreatedParams struct {

	// 
	Context 	BaseAudioContext}



// Notifies that an existing BaseAudioContext will be destroyed.
const ContextWillBeDestroyedEvent = "WebAudio.contextWillBeDestroyed"
type ContextWillBeDestroyedParams struct {

	// 
	ContextId 	GraphObjectId}



// Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
const ContextChangedEvent = "WebAudio.contextChanged"
type ContextChangedParams struct {

	// 
	Context 	BaseAudioContext}



// Notifies that the construction of an AudioListener has finished.
const AudioListenerCreatedEvent = "WebAudio.audioListenerCreated"
type AudioListenerCreatedParams struct {

	// 
	Listener 	AudioListener}



// Notifies that a new AudioListener has been created.
const AudioListenerWillBeDestroyedEvent = "WebAudio.audioListenerWillBeDestroyed"
type AudioListenerWillBeDestroyedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	ListenerId 	GraphObjectId}



// Notifies that a new AudioNode has been created.
const AudioNodeCreatedEvent = "WebAudio.audioNodeCreated"
type AudioNodeCreatedParams struct {

	// 
	Node 	AudioNode}



// Notifies that an existing AudioNode has been destroyed.
const AudioNodeWillBeDestroyedEvent = "WebAudio.audioNodeWillBeDestroyed"
type AudioNodeWillBeDestroyedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	NodeId 	GraphObjectId}



// Notifies that a new AudioParam has been created.
const AudioParamCreatedEvent = "WebAudio.audioParamCreated"
type AudioParamCreatedParams struct {

	// 
	Param 	AudioParam}



// Notifies that an existing AudioParam has been destroyed.
const AudioParamWillBeDestroyedEvent = "WebAudio.audioParamWillBeDestroyed"
type AudioParamWillBeDestroyedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	NodeId 	GraphObjectId
	// 
	ParamId 	GraphObjectId}



// Notifies that two AudioNodes are connected.
const NodesConnectedEvent = "WebAudio.nodesConnected"
type NodesConnectedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	SourceId 	GraphObjectId
	// 
	DestinationId 	GraphObjectId
	// 
	SourceOutputIndex 	float64
	// 
	DestinationInputIndex 	float64}



// Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
const NodesDisconnectedEvent = "WebAudio.nodesDisconnected"
type NodesDisconnectedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	SourceId 	GraphObjectId
	// 
	DestinationId 	GraphObjectId
	// 
	SourceOutputIndex 	float64
	// 
	DestinationInputIndex 	float64}



// Notifies that an AudioNode is connected to an AudioParam.
const NodeParamConnectedEvent = "WebAudio.nodeParamConnected"
type NodeParamConnectedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	SourceId 	GraphObjectId
	// 
	DestinationId 	GraphObjectId
	// 
	SourceOutputIndex 	float64}



// Notifies that an AudioNode is disconnected to an AudioParam.
const NodeParamDisconnectedEvent = "WebAudio.nodeParamDisconnected"
type NodeParamDisconnectedParams struct {

	// 
	ContextId 	GraphObjectId
	// 
	SourceId 	GraphObjectId
	// 
	DestinationId 	GraphObjectId
	// 
	SourceOutputIndex 	float64}

