package fetch

import (
	"github.com/diiyw/cuto/protocol/network"
	"github.com/diiyw/cuto/protocol/cdp"
)


// Issued when the domain is enabled and the request URL matches the
// specified filter. The request is paused until the client responds
// with one of continueRequest, failRequest or fulfillRequest.
// The stage of the request can be determined by presence of responseErrorReason
// and responseStatusCode -- the request is at the response stage if either
// of these fields is present and in the request stage otherwise.
const RequestPausedEvent = "Fetch.requestPaused"
type RequestPausedParams struct {

	// Each request the page makes will have a unique id.
	RequestId 	RequestId
	// The details of the request.
	Request 	network.Request
	// The id of the frame that initiated the request.
	FrameId 	cdp.FrameId
	// How the requested resource will be used.
	ResourceType 	network.ResourceType
	// Response error if intercepted at response stage.
	ResponseErrorReason 	network.ErrorReason
	// Response code if intercepted at response stage.
	ResponseStatusCode 	int
	// Response headers if intercepted at the response stage.
	ResponseHeaders 	[]*HeaderEntry
	// If the intercepted request had a corresponding Network.requestWillBeSent event fired for it,
	// then this networkId will be the same as the requestId present in the requestWillBeSent event.
	NetworkId 	RequestId}



// Issued when the domain is enabled with handleAuthRequests set to true.
// The request is paused until client responds with continueWithAuth.
const AuthRequiredEvent = "Fetch.authRequired"
type AuthRequiredParams struct {

	// Each request the page makes will have a unique id.
	RequestId 	RequestId
	// The details of the request.
	Request 	network.Request
	// The id of the frame that initiated the request.
	FrameId 	cdp.FrameId
	// How the requested resource will be used.
	ResourceType 	network.ResourceType
	// Details of the Authorization Challenge encountered.
	// If this is set, client should respond with continueRequest that
	// contains AuthChallengeResponse.
	AuthChallenge 	AuthChallenge}

