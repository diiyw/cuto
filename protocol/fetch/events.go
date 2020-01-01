package fetch

import (

	"github.com/diiyw/chr/protocol/frame"

	"github.com/diiyw/chr/protocol/network"

)
const (
	
	// Issued when the domain is enabled and the request URL matches the
	// specified filter. The request is paused until the client responds
	// with one of continueRequest, failRequest or fulfillRequest.
	// The stage of the request can be determined by presence of responseErrorReason
	// and responseStatusCode -- the request is at the response stage if either
	// of these fields is present and in the request stage otherwise.
	RequestPausedEvent = "Fetch.requestPaused"
	
	// Issued when the domain is enabled with handleAuthRequests set to true.
	// The request is paused until client responds with continueWithAuth.
	AuthRequiredEvent = "Fetch.authRequired"
	
)

// Issued when the domain is enabled and the request URL matches the
	// specified filter. The request is paused until the client responds
	// with one of continueRequest, failRequest or fulfillRequest.
	// The stage of the request can be determined by presence of responseErrorReason
	// and responseStatusCode -- the request is at the response stage if either
	// of these fields is present and in the request stage otherwise.
type RequestPausedParams struct {
	
	// Each request the page makes will have a unique id.
	RequestId	RequestId	`json:"requestId"`
	
	// The details of the request.
	Request	network.Request	`json:"request"`
	
	// The id of the frame that initiated the request.
	FrameId	frame.FrameId	`json:"frameId"`
	
	// How the requested resource will be used.
	ResourceType	network.ResourceType	`json:"resourceType"`
	
	// Response error if intercepted at response stage.
	ResponseErrorReason	network.ErrorReason	`json:"responseErrorReason"`
	
	// Response code if intercepted at response stage.
	ResponseStatusCode	int	`json:"responseStatusCode"`
	
	// Response headers if intercepted at the response stage.
	ResponseHeaders	[]HeaderEntry	`json:"responseHeaders"`
	
	// If the intercepted request had a corresponding Network.requestWillBeSent event fired for it,
	// then this networkId will be the same as the requestId present in the requestWillBeSent event.
	NetworkId	RequestId	`json:"networkId"`
	
}

// Issued when the domain is enabled with handleAuthRequests set to true.
	// The request is paused until client responds with continueWithAuth.
type AuthRequiredParams struct {
	
	// Each request the page makes will have a unique id.
	RequestId	RequestId	`json:"requestId"`
	
	// The details of the request.
	Request	network.Request	`json:"request"`
	
	// The id of the frame that initiated the request.
	FrameId	frame.FrameId	`json:"frameId"`
	
	// How the requested resource will be used.
	ResourceType	network.ResourceType	`json:"resourceType"`
	
	// Details of the Authorization Challenge encountered.
	// If this is set, client should respond with continueRequest that
	// contains AuthChallengeResponse.
	AuthChallenge	AuthChallenge	`json:"authChallenge"`
	
}

