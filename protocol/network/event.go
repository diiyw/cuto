package network

import (
	"github.com/diiyw/cuto/protocol/cdp"
)


// Fired when data chunk was received over the network.
const DataReceivedEvent = "Network.dataReceived"
type DataReceivedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// Data chunk length.
	DataLength 	int
	// Actual bytes received (might be less than dataLength for compressed encodings).
	EncodedDataLength 	int}



// Fired when EventSource message is received.
const EventSourceMessageReceivedEvent = "Network.eventSourceMessageReceived"
type EventSourceMessageReceivedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// Message type.
	EventName 	string
	// Message identifier.
	EventId 	string
	// Message content.
	Data 	string}



// Fired when HTTP request has failed to load.
const LoadingFailedEvent = "Network.loadingFailed"
type LoadingFailedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// Resource type.
	Type 	ResourceType
	// User friendly error message.
	ErrorText 	string
	// True if loading was canceled.
	Canceled 	bool
	// The reason why loading was blocked, if any.
	BlockedReason 	BlockedReason}



// Fired when HTTP request has finished loading.
const LoadingFinishedEvent = "Network.loadingFinished"
type LoadingFinishedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// Total number of bytes received for this request.
	EncodedDataLength 	float64
	// Set when 1) response was blocked by Cross-Origin Read Blocking and also
	// 2) this needs to be reported to the DevTools console.
	ShouldReportCorbBlocking 	bool}



// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or
// mocked.
// Deprecated, use Fetch.requestPaused instead.
const RequestInterceptedEvent = "Network.requestIntercepted"
type RequestInterceptedParams struct {

	// Each request the page makes will have a unique id, however if any redirects are encountered
	// while processing that fetch, they will be reported with the same id as the original fetch.
	// Likewise if HTTP authentication is needed then the same fetch id will be used.
	InterceptionId 	InterceptionId
	// 
	Request 	Request
	// The id of the frame that initiated the request.
	FrameId 	cdp.FrameId
	// How the requested resource will be used.
	ResourceType 	ResourceType
	// Whether this is a navigation request, which can abort the navigation completely.
	IsNavigationRequest 	bool
	// Set if the request is a navigation that will result in a download.
	// Only present after response is received from the server (i.e. HeadersReceived stage).
	IsDownload 	bool
	// Redirect location, only sent if a redirect was intercepted.
	RedirectUrl 	string
	// Details of the Authorization Challenge encountered. If this is set then
	// continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge 	AuthChallenge
	// Response error if intercepted at response stage or if redirect occurred while intercepting
	// request.
	ResponseErrorReason 	ErrorReason
	// Response code if intercepted at response stage or if redirect occurred while intercepting
	// request or auth retry occurred.
	ResponseStatusCode 	int
	// Response headers if intercepted at the response stage or if redirect occurred while
	// intercepting request or auth retry occurred.
	ResponseHeaders 	Headers
	// If the intercepted request had a corresponding requestWillBeSent event fired for it, then
	// this requestId will be the same as the requestId present in the requestWillBeSent event.
	RequestId 	RequestId}



// Fired if request ended up loading from cache.
const RequestServedFromCacheEvent = "Network.requestServedFromCache"
type RequestServedFromCacheParams struct {

	// Request identifier.
	RequestId 	RequestId}



// Fired when page is about to send HTTP request.
const RequestWillBeSentEvent = "Network.requestWillBeSent"
type RequestWillBeSentParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId 	LoaderId
	// URL of the document this request is loaded for.
	DocumentURL 	string
	// Request data.
	Request 	Request
	// Timestamp.
	Timestamp 	MonotonicTime
	// Timestamp.
	WallTime 	TimeSinceEpoch
	// Request initiator.
	Initiator 	Initiator
	// Redirect response data.
	RedirectResponse 	Response
	// Type of this resource.
	Type 	ResourceType
	// Frame identifier.
	FrameId 	cdp.FrameId
	// Whether the request is initiated by a user gesture. Defaults to false.
	HasUserGesture 	bool}



// Fired when resource loading priority is changed
const ResourceChangedPriorityEvent = "Network.resourceChangedPriority"
type ResourceChangedPriorityParams struct {

	// Request identifier.
	RequestId 	RequestId
	// New priority
	NewPriority 	ResourcePriority
	// Timestamp.
	Timestamp 	MonotonicTime}



// Fired when a signed exchange was received over the network
const SignedExchangeReceivedEvent = "Network.signedExchangeReceived"
type SignedExchangeReceivedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Information about the signed exchange response.
	Info 	SignedExchangeInfo}



// Fired when HTTP response is available.
const ResponseReceivedEvent = "Network.responseReceived"
type ResponseReceivedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId 	LoaderId
	// Timestamp.
	Timestamp 	MonotonicTime
	// Resource type.
	Type 	ResourceType
	// Response data.
	Response 	Response
	// Frame identifier.
	FrameId 	cdp.FrameId}



// Fired when WebSocket is closed.
const WebSocketClosedEvent = "Network.webSocketClosed"
type WebSocketClosedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime}



// Fired upon WebSocket creation.
const WebSocketCreatedEvent = "Network.webSocketCreated"
type WebSocketCreatedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// WebSocket request URL.
	Url 	string
	// Request initiator.
	Initiator 	Initiator}



// Fired when WebSocket message error occurs.
const WebSocketFrameErrorEvent = "Network.webSocketFrameError"
type WebSocketFrameErrorParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// WebSocket error message.
	ErrorMessage 	string}



// Fired when WebSocket message is received.
const WebSocketFrameReceivedEvent = "Network.webSocketFrameReceived"
type WebSocketFrameReceivedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// WebSocket response data.
	Response 	WebSocketFrame}



// Fired when WebSocket message is sent.
const WebSocketFrameSentEvent = "Network.webSocketFrameSent"
type WebSocketFrameSentParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// WebSocket response data.
	Response 	WebSocketFrame}



// Fired when WebSocket handshake response becomes available.
const WebSocketHandshakeResponseReceivedEvent = "Network.webSocketHandshakeResponseReceived"
type WebSocketHandshakeResponseReceivedParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// WebSocket response data.
	Response 	WebSocketResponse}



// Fired when WebSocket is about to initiate handshake.
const WebSocketWillSendHandshakeRequestEvent = "Network.webSocketWillSendHandshakeRequest"
type WebSocketWillSendHandshakeRequestParams struct {

	// Request identifier.
	RequestId 	RequestId
	// Timestamp.
	Timestamp 	MonotonicTime
	// UTC Timestamp.
	WallTime 	TimeSinceEpoch
	// WebSocket request data.
	Request 	WebSocketRequest}



// Fired when additional information about a requestWillBeSent event is available from the
// network stack. Not every requestWillBeSent event will have an additional
// requestWillBeSentExtraInfo fired for it, and there is no guarantee whether requestWillBeSent
// or requestWillBeSentExtraInfo will be fired first for the same request.
const RequestWillBeSentExtraInfoEvent = "Network.requestWillBeSentExtraInfo"
type RequestWillBeSentExtraInfoParams struct {

	// Request identifier. Used to match this information to an existing requestWillBeSent event.
	RequestId 	RequestId
	// A list of cookies which will not be sent with this request along with corresponding reasons
	// for blocking.
	BlockedCookies 	[]*BlockedCookieWithReason
	// Raw request headers as they will be sent over the wire.
	Headers 	Headers}



// Fired when additional information about a responseReceived event is available from the network
// stack. Not every responseReceived event will have an additional responseReceivedExtraInfo for
// it, and responseReceivedExtraInfo may be fired before or after responseReceived.
const ResponseReceivedExtraInfoEvent = "Network.responseReceivedExtraInfo"
type ResponseReceivedExtraInfoParams struct {

	// Request identifier. Used to match this information to another responseReceived event.
	RequestId 	RequestId
	// A list of cookies which were not stored from the response along with the corresponding
	// reasons for blocking. The cookies here may not be valid due to syntax errors, which
	// are represented by the invalid cookie line string instead of a proper cookie.
	BlockedCookies 	[]*BlockedSetCookieWithReason
	// Raw response headers as they were received over the wire.
	Headers 	Headers
	// Raw response header text as it was received over the wire. The raw text may not always be
	// available, such as in the case of HTTP/2 or QUIC.
	HeadersText 	string}

