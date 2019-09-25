package fetch

import (

	"github.com/diiyw/goc/protocol/io"

	"github.com/diiyw/goc/protocol/network"

)
const (
	
	// Disables the fetch domain.
	Disable = "Fetch.disable"
	
	// Enables issuing of requestPaused events. A request will be paused until client
	// calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
	Enable = "Fetch.enable"
	
	// Causes the request to fail with specified reason.
	FailRequest = "Fetch.failRequest"
	
	// Provides response to the request.
	FulfillRequest = "Fetch.fulfillRequest"
	
	// Continues the request, optionally modifying some of its parameters.
	ContinueRequest = "Fetch.continueRequest"
	
	// Continues a request supplying authChallengeResponse following authRequired event.
	ContinueWithAuth = "Fetch.continueWithAuth"
	
	// Causes the body of the response to be received from the server and
	// returned as a single string. May only be issued for a request that
	// is paused in the Response stage and is mutually exclusive with
	// takeResponseBodyForInterceptionAsStream. Calling other methods that
	// affect the request or disabling fetch domain before body is received
	// results in an undefined behavior.
	GetResponseBody = "Fetch.getResponseBody"
	
	// Returns a handle to the stream representing the response body.
	// The request must be paused in the HeadersReceived stage.
	// Note that after this command the request can't be continued
	// as is -- client either needs to cancel it or to provide the
	// response body.
	// The stream only supports sequential read, IO.read will fail if the position
	// is specified.
	// This method is mutually exclusive with getResponseBody.
	// Calling other methods that affect the request or disabling fetch
	// domain before body is received results in an undefined behavior.
	TakeResponseBodyAsStream = "Fetch.takeResponseBodyAsStream"
	
)

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
	// If specified, only requests matching any of these patterns will produce
	// fetchRequested event and will be paused until clients response. If not set,
	// all requests will be affected.
	Patterns	[]RequestPattern	`json:"patterns"`
	
	// If true, authRequired events will be issued and requests will be paused
	// expecting a call to continueWithAuth.
	HandleAuthRequests	bool	`json:"handleAuthRequests"`
	
}

// Enable returns
type EnableReturns struct {
	
}

// FailRequest parameters
type FailRequestParams struct {
	
	// An id the client received in requestPaused event.
	RequestId	RequestId	`json:"requestId"`
	
	// Causes the request to fail with the given reason.
	ErrorReason	network.ErrorReason	`json:"errorReason"`
	
}

// FailRequest returns
type FailRequestReturns struct {
	
}

// FulfillRequest parameters
type FulfillRequestParams struct {
	
	// An id the client received in requestPaused event.
	RequestId	RequestId	`json:"requestId"`
	
	// An HTTP response code.
	ResponseCode	int	`json:"responseCode"`
	
	// Response headers.
	ResponseHeaders	[]HeaderEntry	`json:"responseHeaders"`
	
	// Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text.
	BinaryResponseHeaders	string	`json:"binaryResponseHeaders"`
	
	// A response body.
	Body	string	`json:"body"`
	
	// A textual representation of responseCode.
	// If absent, a standard phrase matching responseCode is used.
	ResponsePhrase	string	`json:"responsePhrase"`
	
}

// FulfillRequest returns
type FulfillRequestReturns struct {
	
}

// ContinueRequest parameters
type ContinueRequestParams struct {
	
	// An id the client received in requestPaused event.
	RequestId	RequestId	`json:"requestId"`
	
	// If set, the request url will be modified in a way that's not observable by page.
	Url	string	`json:"url"`
	
	// If set, the request method is overridden.
	Method	string	`json:"method"`
	
	// If set, overrides the post data in the request.
	PostData	string	`json:"postData"`
	
	// If set, overrides the request headrts.
	Headers	[]HeaderEntry	`json:"headers"`
	
}

// ContinueRequest returns
type ContinueRequestReturns struct {
	
}

// ContinueWithAuth parameters
type ContinueWithAuthParams struct {
	
	// An id the client received in authRequired event.
	RequestId	RequestId	`json:"requestId"`
	
	// Response to  with an authChallenge.
	AuthChallengeResponse	AuthChallengeResponse	`json:"authChallengeResponse"`
	
}

// ContinueWithAuth returns
type ContinueWithAuthReturns struct {
	
}

// GetResponseBody parameters
type GetResponseBodyParams struct {
	
	// Identifier for the intercepted request to get body for.
	RequestId	RequestId	`json:"requestId"`
	
}

// GetResponseBody returns
type GetResponseBodyReturns struct {
	
	// Response body.
	Body	string	`json:"body"`
	
	// True, if content was sent as base64.
	Base64Encoded	bool	`json:"base64Encoded"`
	
}

// TakeResponseBodyAsStream parameters
type TakeResponseBodyAsStreamParams struct {
	
	
	RequestId	RequestId	`json:"requestId"`
	
}

// TakeResponseBodyAsStream returns
type TakeResponseBodyAsStreamReturns struct {
	
	
	Stream	io.StreamHandle	`json:"stream"`
	
}

