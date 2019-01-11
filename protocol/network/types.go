package network

import (

	"github.com/diiyw/goc/protocol/runtime"

	"github.com/diiyw/goc/protocol/security"

)

// Resource type as it was perceived by the rendering engine.
type ResourceType string	

// Unique loader identifier.
type LoaderId string	

// Unique request identifier.
type RequestId string	

// Unique intercepted request identifier.
type InterceptionId string	

// Network level fetch failure reason.
type ErrorReason string	

// UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch float64	

// Monotonically increasing time in seconds since an arbitrary point in the past.
type MonotonicTime float64	

// Request / response headers as keys / values of JSON object.
type Headers struct {
	
}	

// The underlying connection technology that the browser is supposedly using.
type ConnectionType string	

// Represents the cookie's 'SameSite' status:
	// https://tools.ietf.org/html/draft-west-first-party-cookies
type CookieSameSite string	

// Timing information for the request.
type ResourceTiming struct {
	
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in
	// milliseconds relatively to this requestTime.
	
	RequestTime	float64	`json:"requestTime"`
	
	// Started resolving proxy.
	
	ProxyStart	float64	`json:"proxyStart"`
	
	// Finished resolving proxy.
	
	ProxyEnd	float64	`json:"proxyEnd"`
	
	// Started DNS address resolve.
	
	DnsStart	float64	`json:"dnsStart"`
	
	// Finished DNS address resolve.
	
	DnsEnd	float64	`json:"dnsEnd"`
	
	// Started connecting to the remote host.
	
	ConnectStart	float64	`json:"connectStart"`
	
	// Connected to the remote host.
	
	ConnectEnd	float64	`json:"connectEnd"`
	
	// Started SSL handshake.
	
	SslStart	float64	`json:"sslStart"`
	
	// Finished SSL handshake.
	
	SslEnd	float64	`json:"sslEnd"`
	
	// Started running ServiceWorker.
	
	WorkerStart	float64	`json:"workerStart"`
	
	// Finished Starting ServiceWorker.
	
	WorkerReady	float64	`json:"workerReady"`
	
	// Started sending request.
	
	SendStart	float64	`json:"sendStart"`
	
	// Finished sending request.
	
	SendEnd	float64	`json:"sendEnd"`
	
	// Time the server started pushing request.
	
	PushStart	float64	`json:"pushStart"`
	
	// Time the server finished pushing request.
	
	PushEnd	float64	`json:"pushEnd"`
	
	// Finished receiving response headers.
	
	ReceiveHeadersEnd	float64	`json:"receiveHeadersEnd"`
	
}	

// Loading priority of a resource request.
type ResourcePriority string	

// HTTP request data.
type Request struct {
	
	// Request URL (without fragment).
	
	Url	string	`json:"url"`
	
	// Fragment of the requested URL starting with hash, if present.
	
	UrlFragment	string	`json:"urlFragment"`
	
	// HTTP request method.
	
	Method	string	`json:"method"`
	
	// HTTP request headers.
	
	Headers	Headers	`json:"headers"`
	
	// HTTP POST request data.
	
	PostData	string	`json:"postData"`
	
	// True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
	
	HasPostData	bool	`json:"hasPostData,omitempty"`
	
	// The mixed content type of the request.
	
	MixedContentType	security.MixedContentType	`json:"mixedContentType"`
	
	// Priority of the resource request at the time request is sent.
	
	InitialPriority	ResourcePriority	`json:"initialPriority"`
	
	// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	// Possible value: unsafe-url,no-referrer-when-downgrade,no-referrer,origin,origin-when-cross-origin,same-origin,strict-origin,strict-origin-when-cross-origin,
	ReferrerPolicy	string	`json:"referrerPolicy"`
	
	// Whether is loaded via link preload.
	
	IsLinkPreload	bool	`json:"isLinkPreload"`
	
}	

// Details of a signed certificate timestamp (SCT).
type SignedCertificateTimestamp struct {
	
	// Validation status.
	
	Status	string	`json:"status"`
	
	// Origin.
	
	Origin	string	`json:"origin"`
	
	// Log name / description.
	
	LogDescription	string	`json:"logDescription"`
	
	// Log ID.
	
	LogId	string	`json:"logId"`
	
	// Issuance date.
	
	Timestamp	TimeSinceEpoch	`json:"timestamp"`
	
	// Hash algorithm.
	
	HashAlgorithm	string	`json:"hashAlgorithm"`
	
	// Signature algorithm.
	
	SignatureAlgorithm	string	`json:"signatureAlgorithm"`
	
	// Signature data.
	
	SignatureData	string	`json:"signatureData"`
	
}	

// Security details about a request.
type SecurityDetails struct {
	
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	
	Protocol	string	`json:"protocol"`
	
	// Key Exchange used by the connection, or the empty string if not applicable.
	
	KeyExchange	string	`json:"keyExchange"`
	
	// (EC)DH group used by the connection, if applicable.
	
	KeyExchangeGroup	string	`json:"keyExchangeGroup"`
	
	// Cipher name.
	
	Cipher	string	`json:"cipher"`
	
	// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	
	Mac	string	`json:"mac"`
	
	// Certificate ID value.
	
	CertificateId	security.CertificateId	`json:"certificateId"`
	
	// Certificate subject name.
	
	SubjectName	string	`json:"subjectName"`
	
	// Subject Alternative Name (SAN) DNS names and IP addresses.
	
	SanList	[]string	`json:"sanList"`
	
	// Name of the issuing CA.
	
	Issuer	string	`json:"issuer"`
	
	// Certificate valid from date.
	
	ValidFrom	TimeSinceEpoch	`json:"validFrom"`
	
	// Certificate valid to (expiration) date
	
	ValidTo	TimeSinceEpoch	`json:"validTo"`
	
	// List of signed certificate timestamps (SCTs).
	
	SignedCertificateTimestampList	[]SignedCertificateTimestamp	`json:"signedCertificateTimestampList"`
	
	// Whether the request complied with Certificate Transparency policy
	
	CertificateTransparencyCompliance	CertificateTransparencyCompliance	`json:"certificateTransparencyCompliance"`
	
}	

// Whether the request complied with Certificate Transparency policy.
type CertificateTransparencyCompliance string	

// The reason why request was blocked.
type BlockedReason string	

// HTTP response data.
type Response struct {
	
	// Response URL. This URL can be different from CachedResource.url in case of redirect.
	
	Url	string	`json:"url"`
	
	// HTTP response status code.
	
	Status	int	`json:"status"`
	
	// HTTP response status text.
	
	StatusText	string	`json:"statusText"`
	
	// HTTP response headers.
	
	Headers	Headers	`json:"headers"`
	
	// HTTP response headers text.
	
	HeadersText	string	`json:"headersText"`
	
	// Resource mimeType as determined by the browser.
	
	MimeType	string	`json:"mimeType"`
	
	// Refined HTTP request headers that were actually transmitted over the network.
	
	RequestHeaders	Headers	`json:"requestHeaders"`
	
	// HTTP request headers text.
	
	RequestHeadersText	string	`json:"requestHeadersText"`
	
	// Specifies whether physical connection was actually reused for this request.
	
	ConnectionReused	bool	`json:"connectionReused"`
	
	// Physical connection id that was actually used for this request.
	
	ConnectionId	float64	`json:"connectionId"`
	
	// Remote IP address.
	
	RemoteIPAddress	string	`json:"remoteIPAddress"`
	
	// Remote port.
	
	RemotePort	int	`json:"remotePort"`
	
	// Specifies that the request was served from the disk cache.
	
	FromDiskCache	bool	`json:"fromDiskCache"`
	
	// Specifies that the request was served from the ServiceWorker.
	
	FromServiceWorker	bool	`json:"fromServiceWorker"`
	
	// Total number of bytes received for this request so far.
	
	EncodedDataLength	float64	`json:"encodedDataLength"`
	
	// Timing information for the given request.
	
	Timing	ResourceTiming	`json:"timing"`
	
	// Protocol used to fetch this request.
	
	Protocol	string	`json:"protocol"`
	
	// Security state of the request resource.
	
	SecurityState	security.SecurityState	`json:"securityState"`
	
	// Security details for the request.
	
	SecurityDetails	SecurityDetails	`json:"securityDetails"`
	
}	

// WebSocket request data.
type WebSocketRequest struct {
	
	// HTTP request headers.
	
	Headers	Headers	`json:"headers"`
	
}	

// WebSocket response data.
type WebSocketResponse struct {
	
	// HTTP response status code.
	
	Status	int	`json:"status"`
	
	// HTTP response status text.
	
	StatusText	string	`json:"statusText"`
	
	// HTTP response headers.
	
	Headers	Headers	`json:"headers"`
	
	// HTTP response headers text.
	
	HeadersText	string	`json:"headersText"`
	
	// HTTP request headers.
	
	RequestHeaders	Headers	`json:"requestHeaders"`
	
	// HTTP request headers text.
	
	RequestHeadersText	string	`json:"requestHeadersText"`
	
}	

// WebSocket message data. This represents an entire WebSocket message, not just a fragmented frame as the name suggests.
type WebSocketFrame struct {
	
	// WebSocket message opcode.
	
	Opcode	float64	`json:"opcode"`
	
	// WebSocket message mask.
	
	Mask	bool	`json:"mask"`
	
	// WebSocket message payload data.
	// If the opcode is 1, this is a text message and payloadData is a UTF-8 string.
	// If the opcode isn't 1, then payloadData is a base64 encoded string representing binary data.
	
	PayloadData	string	`json:"payloadData"`
	
}	

// Information about the cached resource.
type CachedResource struct {
	
	// Resource URL. This is the url of the original network request.
	
	Url	string	`json:"url"`
	
	// Type of this resource.
	
	Type	ResourceType	`json:"type"`
	
	// Cached response data.
	
	Response	Response	`json:"response"`
	
	// Cached response body size.
	
	BodySize	float64	`json:"bodySize"`
	
}	

// Information about the request initiator.
type Initiator struct {
	
	// Type of this initiator.
	// Possible value: parser,script,preload,SignedExchange,other,
	Type	string	`json:"type"`
	
	// Initiator JavaScript stack trace, set for Script only.
	
	Stack	runtime.StackTrace	`json:"stack"`
	
	// Initiator URL, set for Parser type or for Script type (when script is importing module) or for SignedExchange type.
	
	Url	string	`json:"url"`
	
	// Initiator line number, set for Parser type or for Script type (when script is importing
	// module) (0-based).
	
	LineNumber	float64	`json:"lineNumber"`
	
}	

// Cookie object
type Cookie struct {
	
	// Cookie name.
	
	Name	string	`json:"name"`
	
	// Cookie value.
	
	Value	string	`json:"value"`
	
	// Cookie domain.
	
	Domain	string	`json:"domain"`
	
	// Cookie path.
	
	Path	string	`json:"path"`
	
	// Cookie expiration date as the number of seconds since the UNIX epoch.
	
	Expires	float64	`json:"expires"`
	
	// Cookie size.
	
	Size	int	`json:"size"`
	
	// True if cookie is http-only.
	
	HttpOnly	bool	`json:"httpOnly"`
	
	// True if cookie is secure.
	
	Secure	bool	`json:"secure"`
	
	// True in case of session cookie.
	
	Session	bool	`json:"session"`
	
	// Cookie SameSite type.
	
	SameSite	CookieSameSite	`json:"sameSite"`
	
}	

// Cookie parameter object
type CookieParam struct {
	
	// Cookie name.
	
	Name	string	`json:"name"`
	
	// Cookie value.
	
	Value	string	`json:"value"`
	
	// The request-URI to associate with the setting of the cookie. This value can affect the
	// default domain and path values of the created cookie.
	
	Url	string	`json:"url"`
	
	// Cookie domain.
	
	Domain	string	`json:"domain"`
	
	// Cookie path.
	
	Path	string	`json:"path"`
	
	// True if cookie is secure.
	
	Secure	bool	`json:"secure"`
	
	// True if cookie is http-only.
	
	HttpOnly	bool	`json:"httpOnly"`
	
	// Cookie SameSite type.
	
	SameSite	CookieSameSite	`json:"sameSite"`
	
	// Cookie expiration date, session cookie if not set
	
	Expires	TimeSinceEpoch	`json:"expires"`
	
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

// Stages of the interception to begin intercepting. Request will intercept before the request is
	// sent. Response will intercept after the response is received.
type InterceptionStage string	

// Request pattern for interception.
type RequestPattern struct {
	
	// Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is
	// backslash. Omitting is equivalent to "*".
	
	UrlPattern	string	`json:"urlPattern"`
	
	// If set, only requests for matching resource types will be intercepted.
	
	ResourceType	ResourceType	`json:"resourceType"`
	
	// Stage at wich to begin intercepting requests. Default is Request.
	
	InterceptionStage	InterceptionStage	`json:"interceptionStage"`
	
}	

// Information about a signed exchange signature.
	// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
type SignedExchangeSignature struct {
	
	// Signed exchange signature label.
	
	Label	string	`json:"label"`
	
	// The hex string of signed exchange signature.
	
	Signature	string	`json:"signature"`
	
	// Signed exchange signature integrity.
	
	Integrity	string	`json:"integrity"`
	
	// Signed exchange signature cert Url.
	
	CertUrl	string	`json:"certUrl"`
	
	// The hex string of signed exchange signature cert sha256.
	
	CertSha256	string	`json:"certSha256"`
	
	// Signed exchange signature validity Url.
	
	ValidityUrl	string	`json:"validityUrl"`
	
	// Signed exchange signature date.
	
	Date	int	`json:"date"`
	
	// Signed exchange signature expires.
	
	Expires	int	`json:"expires"`
	
	// The encoded certificates.
	
	Certificates	[]string	`json:"certificates"`
	
}	

// Information about a signed exchange header.
	// https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
type SignedExchangeHeader struct {
	
	// Signed exchange request URL.
	
	RequestUrl	string	`json:"requestUrl"`
	
	// Signed exchange request method.
	
	RequestMethod	string	`json:"requestMethod"`
	
	// Signed exchange response code.
	
	ResponseCode	int	`json:"responseCode"`
	
	// Signed exchange response headers.
	
	ResponseHeaders	Headers	`json:"responseHeaders"`
	
	// Signed exchange response signature.
	
	Signatures	[]SignedExchangeSignature	`json:"signatures"`
	
}	

// Field type for a signed exchange related error.
type SignedExchangeErrorField string	

// Information about a signed exchange response.
type SignedExchangeError struct {
	
	// Error message.
	
	Message	string	`json:"message"`
	
	// The index of the signature which caused the error.
	
	SignatureIndex	int	`json:"signatureIndex"`
	
	// The field which caused the error.
	
	ErrorField	SignedExchangeErrorField	`json:"errorField"`
	
}	

// Information about a signed exchange response.
type SignedExchangeInfo struct {
	
	// The outer response of signed HTTP exchange which was received from network.
	
	OuterResponse	Response	`json:"outerResponse"`
	
	// Information about the signed exchange header.
	
	Header	SignedExchangeHeader	`json:"header"`
	
	// Security details for the signed exchange header.
	
	SecurityDetails	SecurityDetails	`json:"securityDetails"`
	
	// Errors occurred while handling the signed exchagne.
	
	Errors	[]SignedExchangeError	`json:"errors"`
	
}	

