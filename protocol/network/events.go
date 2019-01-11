package network

import (

	"github.com/diiyw/goc/protocol/frame"

)
const (
	
	// Fired when data chunk was received over the network.
	DataReceivedEvent = "Network.dataReceived"
	
	// Fired when EventSource message is received.
	EventSourceMessageReceivedEvent = "Network.eventSourceMessageReceived"
	
	// Fired when HTTP request has failed to load.
	LoadingFailedEvent = "Network.loadingFailed"
	
	// Fired when HTTP request has finished loading.
	LoadingFinishedEvent = "Network.loadingFinished"
	
	// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or
	// mocked.
	RequestInterceptedEvent = "Network.requestIntercepted"
	
	// Fired if request ended up loading from cache.
	RequestServedFromCacheEvent = "Network.requestServedFromCache"
	
	// Fired when page is about to send HTTP request.
	RequestWillBeSentEvent = "Network.requestWillBeSent"
	
	// Fired when resource loading priority is changed
	ResourceChangedPriorityEvent = "Network.resourceChangedPriority"
	
	// Fired when a signed exchange was received over the network
	SignedExchangeReceivedEvent = "Network.signedExchangeReceived"
	
	// Fired when HTTP response is available.
	ResponseReceivedEvent = "Network.responseReceived"
	
	// Fired when WebSocket is closed.
	WebSocketClosedEvent = "Network.webSocketClosed"
	
	// Fired upon WebSocket creation.
	WebSocketCreatedEvent = "Network.webSocketCreated"
	
	// Fired when WebSocket message error occurs.
	WebSocketFrameErrorEvent = "Network.webSocketFrameError"
	
	// Fired when WebSocket message is received.
	WebSocketFrameReceivedEvent = "Network.webSocketFrameReceived"
	
	// Fired when WebSocket message is sent.
	WebSocketFrameSentEvent = "Network.webSocketFrameSent"
	
	// Fired when WebSocket handshake response becomes available.
	WebSocketHandshakeResponseReceivedEvent = "Network.webSocketHandshakeResponseReceived"
	
	// Fired when WebSocket is about to initiate handshake.
	WebSocketWillSendHandshakeRequestEvent = "Network.webSocketWillSendHandshakeRequest"
	
)

// Fired when data chunk was received over the network.
type DataReceivedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// Data chunk length.
	DataLength	int	`json:"dataLength"`
	
	// Actual bytes received (might be less than dataLength for compressed encodings).
	EncodedDataLength	int	`json:"encodedDataLength"`
	
}

// Fired when EventSource message is received.
type EventSourceMessageReceivedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// Message type.
	EventName	string	`json:"eventName"`
	
	// Message identifier.
	EventId	string	`json:"eventId"`
	
	// Message content.
	Data	string	`json:"data"`
	
}

// Fired when HTTP request has failed to load.
type LoadingFailedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// Resource type.
	Type	ResourceType	`json:"type"`
	
	// User friendly error message.
	ErrorText	string	`json:"errorText"`
	
	// True if loading was canceled.
	Canceled	bool	`json:"canceled"`
	
	// The reason why loading was blocked, if any.
	BlockedReason	BlockedReason	`json:"blockedReason"`
	
}

// Fired when HTTP request has finished loading.
type LoadingFinishedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// Total number of bytes received for this request.
	EncodedDataLength	float64	`json:"encodedDataLength"`
	
	// Set when 1) response was blocked by Cross-Origin Read Blocking and also
	// 2) this needs to be reported to the DevTools console.
	ShouldReportCorbBlocking	bool	`json:"shouldReportCorbBlocking"`
	
}

// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or
	// mocked.
type RequestInterceptedParams struct {
	
	// Each request the page makes will have a unique id, however if any redirects are encountered
	// while processing that fetch, they will be reported with the same id as the original fetch.
	// Likewise if HTTP authentication is needed then the same fetch id will be used.
	InterceptionId	InterceptionId	`json:"interceptionId"`
	
	
	Request	Request	`json:"request"`
	
	// The id of the frame that initiated the request.
	FrameId	frame.FrameId	`json:"frameId"`
	
	// How the requested resource will be used.
	ResourceType	ResourceType	`json:"resourceType"`
	
	// Whether this is a navigation request, which can abort the navigation completely.
	IsNavigationRequest	bool	`json:"isNavigationRequest"`
	
	// Set if the request is a navigation that will result in a download.
	// Only present after response is received from the server (i.e. HeadersReceived stage).
	IsDownload	bool	`json:"isDownload"`
	
	// Redirect location, only sent if a redirect was intercepted.
	RedirectUrl	string	`json:"redirectUrl"`
	
	// Details of the Authorization Challenge encountered. If this is set then
	// continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge	AuthChallenge	`json:"authChallenge"`
	
	// Response error if intercepted at response stage or if redirect occurred while intercepting
	// request.
	ResponseErrorReason	ErrorReason	`json:"responseErrorReason"`
	
	// Response code if intercepted at response stage or if redirect occurred while intercepting
	// request or auth retry occurred.
	ResponseStatusCode	int	`json:"responseStatusCode"`
	
	// Response headers if intercepted at the response stage or if redirect occurred while
	// intercepting request or auth retry occurred.
	ResponseHeaders	Headers	`json:"responseHeaders"`
	
}

// Fired if request ended up loading from cache.
type RequestServedFromCacheParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
}

// Fired when page is about to send HTTP request.
type RequestWillBeSentParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId	LoaderId	`json:"loaderId"`
	
	// URL of the document this request is loaded for.
	DocumentURL	string	`json:"documentURL"`
	
	// Request data.
	Request	Request	`json:"request"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// Timestamp.
	WallTime	TimeSinceEpoch	`json:"wallTime"`
	
	// Request initiator.
	Initiator	Initiator	`json:"initiator"`
	
	// Redirect response data.
	RedirectResponse	Response	`json:"redirectResponse"`
	
	// Type of this resource.
	Type	ResourceType	`json:"type"`
	
	// Frame identifier.
	FrameId	frame.FrameId	`json:"frameId"`
	
	// Whether the request is initiated by a user gesture. Defaults to false.
	HasUserGesture	bool	`json:"hasUserGesture"`
	
}

// Fired when resource loading priority is changed
type ResourceChangedPriorityParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// New priority
	NewPriority	ResourcePriority	`json:"newPriority"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
}

// Fired when a signed exchange was received over the network
type SignedExchangeReceivedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Information about the signed exchange response.
	Info	SignedExchangeInfo	`json:"info"`
	
}

// Fired when HTTP response is available.
type ResponseReceivedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId	LoaderId	`json:"loaderId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// Resource type.
	Type	ResourceType	`json:"type"`
	
	// Response data.
	Response	Response	`json:"response"`
	
	// Frame identifier.
	FrameId	frame.FrameId	`json:"frameId"`
	
}

// Fired when WebSocket is closed.
type WebSocketClosedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
}

// Fired upon WebSocket creation.
type WebSocketCreatedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// WebSocket request URL.
	Url	string	`json:"url"`
	
	// Request initiator.
	Initiator	Initiator	`json:"initiator"`
	
}

// Fired when WebSocket message error occurs.
type WebSocketFrameErrorParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// WebSocket error message.
	ErrorMessage	string	`json:"errorMessage"`
	
}

// Fired when WebSocket message is received.
type WebSocketFrameReceivedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// WebSocket response data.
	Response	WebSocketFrame	`json:"response"`
	
}

// Fired when WebSocket message is sent.
type WebSocketFrameSentParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// WebSocket response data.
	Response	WebSocketFrame	`json:"response"`
	
}

// Fired when WebSocket handshake response becomes available.
type WebSocketHandshakeResponseReceivedParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// WebSocket response data.
	Response	WebSocketResponse	`json:"response"`
	
}

// Fired when WebSocket is about to initiate handshake.
type WebSocketWillSendHandshakeRequestParams struct {
	
	// Request identifier.
	RequestId	RequestId	`json:"requestId"`
	
	// Timestamp.
	Timestamp	MonotonicTime	`json:"timestamp"`
	
	// UTC Timestamp.
	WallTime	TimeSinceEpoch	`json:"wallTime"`
	
	// WebSocket request data.
	Request	WebSocketRequest	`json:"request"`
	
}

