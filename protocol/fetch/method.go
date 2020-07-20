package fetch

import (
	"github.com/diiyw/cuto/protocol/network"
	"github.com/diiyw/cuto/protocol/io"
)


// Disables the fetch domain.
const Disable = "Fetch.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables issuing of requestPaused events. A request will be paused until client
// calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
const Enable = "Fetch.enable"

type EnableParams struct {

	// If specified, only requests matching any of these patterns will produce
	// fetchRequested event and will be paused until clients response. If not set,
	// all requests will be affected.
	Patterns 	[]*RequestPattern	`json:"patterns,omitempty"`

	// If true, authRequired events will be issued and requests will be paused
	// expecting a call to continueWithAuth.
	HandleAuthRequests 	bool	`json:"handleAuthRequests,omitempty"`
}

type EnableResult struct {

}

// Causes the request to fail with specified reason.
const FailRequest = "Fetch.failRequest"

type FailRequestParams struct {

	// An id the client received in requestPaused event.
	RequestId 	RequestId	`json:"requestId"`

	// Causes the request to fail with the given reason.
	ErrorReason 	network.ErrorReason	`json:"errorReason"`
}

type FailRequestResult struct {

}

// Provides response to the request.
const FulfillRequest = "Fetch.fulfillRequest"

type FulfillRequestParams struct {

	// An id the client received in requestPaused event.
	RequestId 	RequestId	`json:"requestId"`

	// An HTTP response code.
	ResponseCode 	int	`json:"responseCode"`

	// Response headers.
	ResponseHeaders 	[]*HeaderEntry	`json:"responseHeaders,omitempty"`

	// Alternative way of specifying response headers as a \0-separated
	// series of name: value pairs. Prefer the above method unless you
	// need to represent some non-UTF8 values that can't be transmitted
	// over the protocol as text.
	BinaryResponseHeaders 	[]byte	`json:"binaryResponseHeaders,omitempty"`

	// A response body.
	Body 	[]byte	`json:"body,omitempty"`

	// A textual representation of responseCode.
	// If absent, a standard phrase matching responseCode is used.
	ResponsePhrase 	string	`json:"responsePhrase,omitempty"`
}

type FulfillRequestResult struct {

}

// Continues the request, optionally modifying some of its parameters.
const ContinueRequest = "Fetch.continueRequest"

type ContinueRequestParams struct {

	// An id the client received in requestPaused event.
	RequestId 	RequestId	`json:"requestId"`

	// If set, the request url will be modified in a way that's not observable by page.
	Url 	string	`json:"url,omitempty"`

	// If set, the request method is overridden.
	Method 	string	`json:"method,omitempty"`

	// If set, overrides the post data in the request.
	PostData 	string	`json:"postData,omitempty"`

	// If set, overrides the request headrts.
	Headers 	[]*HeaderEntry	`json:"headers,omitempty"`
}

type ContinueRequestResult struct {

}

// Continues a request supplying authChallengeResponse following authRequired event.
const ContinueWithAuth = "Fetch.continueWithAuth"

type ContinueWithAuthParams struct {

	// An id the client received in authRequired event.
	RequestId 	RequestId	`json:"requestId"`

	// Response to  with an authChallenge.
	AuthChallengeResponse 	AuthChallengeResponse	`json:"authChallengeResponse"`
}

type ContinueWithAuthResult struct {

}

// Causes the body of the response to be received from the server and
// returned as a single string. May only be issued for a request that
// is paused in the Response stage and is mutually exclusive with
// takeResponseBodyForInterceptionAsStream. Calling other methods that
// affect the request or disabling fetch domain before body is received
// results in an undefined behavior.
const GetResponseBody = "Fetch.getResponseBody"

type GetResponseBodyParams struct {

	// Identifier for the intercepted request to get body for.
	RequestId 	RequestId	`json:"requestId"`
}

type GetResponseBodyResult struct {

	// Response body.
	Body 	string	`json:"body"`
	// True, if content was sent as base64.
	Base64Encoded 	bool	`json:"base64Encoded"`
}

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
const TakeResponseBodyAsStream = "Fetch.takeResponseBodyAsStream"

type TakeResponseBodyAsStreamParams struct {

	// 
	RequestId 	RequestId	`json:"requestId"`
}

type TakeResponseBodyAsStreamResult struct {

	// 
	Stream 	io.StreamHandle	`json:"stream"`
}