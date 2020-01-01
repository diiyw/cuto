package fetch

import (

	"github.com/diiyw/chr/protocol/network"

)

// Unique request identifier.
type RequestId string	

// Stages of the request to handle. Request will intercept before the request is
	// sent. Response will intercept after the response is received (but before response
	// body is received.
type RequestStage string	

// 
type RequestPattern struct {
	
	// Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to "*".
	
	UrlPattern	string	`json:"urlPattern"`
	
	// If set, only requests for matching resource types will be intercepted.
	
	ResourceType	network.ResourceType	`json:"resourceType"`
	
	// Stage at wich to begin intercepting requests. Default is Request.
	
	RequestStage	RequestStage	`json:"requestStage"`
	
}	

// Response HTTP header entry
type HeaderEntry struct {
	
	
	
	Name	string	`json:"name"`
	
	
	
	Value	string	`json:"value"`
	
}	

// Authorization challenge for HTTP status code 401 or 407.
type AuthChallenge struct {
	
	// Source of the authentication challenge.
	// Possible value: Server,Proxy,
	Source	string	`json:"source"`
	
	// Origin of the challenger.
	
	Origin	string	`json:"origin"`
	
	// The authentication scheme used, such as basic or digest
	
	Scheme	string	`json:"scheme"`
	
	// The realm of the challenge. May be empty.
	
	Realm	string	`json:"realm"`
	
}	

// Response to an AuthChallenge.
type AuthChallengeResponse struct {
	
	// The decision on what to do in response to the authorization challenge.  Default means
	// deferring to the default behavior of the net stack, which will likely either the Cancel
	// authentication or display a popup dialog box.
	// Possible value: Default,CancelAuth,ProvideCredentials,
	Response	string	`json:"response"`
	
	// The username to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	
	Username	string	`json:"username"`
	
	// The password to provide, possibly empty. Should only be set if response is
	// ProvideCredentials.
	
	Password	string	`json:"password"`
	
}	

