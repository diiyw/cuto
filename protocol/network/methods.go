package network

import (

	"github.com/diiyw/goc/protocol/debugger"

	"github.com/diiyw/goc/protocol/io"

)
const (
	
	// Tells whether clearing browser cache is supported.
	CanClearBrowserCache = "Network.canClearBrowserCache"
	
	// Tells whether clearing browser cookies is supported.
	CanClearBrowserCookies = "Network.canClearBrowserCookies"
	
	// Tells whether emulation of network conditions is supported.
	CanEmulateNetworkConditions = "Network.canEmulateNetworkConditions"
	
	// Clears browser cache.
	ClearBrowserCache = "Network.clearBrowserCache"
	
	// Clears browser cookies.
	ClearBrowserCookies = "Network.clearBrowserCookies"
	
	// Response to Network.requestIntercepted which either modifies the request to continue with any
	// modifications, or blocks it, or completes it with the provided response bytes. If a network
	// fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted
	// event will be sent with the same InterceptionId.
	ContinueInterceptedRequest = "Network.continueInterceptedRequest"
	
	// Deletes browser cookies with matching name and url or domain/path pair.
	DeleteCookies = "Network.deleteCookies"
	
	// Disables network tracking, prevents network events from being sent to the client.
	Disable = "Network.disable"
	
	// Activates emulation of network conditions.
	EmulateNetworkConditions = "Network.emulateNetworkConditions"
	
	// Enables network tracking, network events will now be delivered to the client.
	Enable = "Network.enable"
	
	// Returns all browser cookies. Depending on the backend support, will return detailed cookie
	// information in the `cookies` field.
	GetAllCookies = "Network.getAllCookies"
	
	// Returns the DER-encoded certificate.
	GetCertificate = "Network.getCertificate"
	
	// Returns all browser cookies for the current URL. Depending on the backend support, will return
	// detailed cookie information in the `cookies` field.
	GetCookies = "Network.getCookies"
	
	// Returns content served for the given request.
	GetResponseBody = "Network.getResponseBody"
	
	// Returns post data sent with the request. Returns an error when no data was sent with the request.
	GetRequestPostData = "Network.getRequestPostData"
	
	// Returns content served for the given currently intercepted request.
	GetResponseBodyForInterception = "Network.getResponseBodyForInterception"
	
	// Returns a handle to the stream representing the response body. Note that after this command,
	// the intercepted request can't be continued as is -- you either need to cancel it or to provide
	// the response body. The stream only supports sequential read, IO.read will fail if the position
	// is specified.
	TakeResponseBodyForInterceptionAsStream = "Network.takeResponseBodyForInterceptionAsStream"
	
	// This method sends a new XMLHttpRequest which is identical to the original one. The following
	// parameters should be identical: method, url, async, request body, extra headers, withCredentials
	// attribute, user, password.
	ReplayXHR = "Network.replayXHR"
	
	// Searches for given string in response content.
	SearchInResponseBody = "Network.searchInResponseBody"
	
	// Blocks URLs from loading.
	SetBlockedURLs = "Network.setBlockedURLs"
	
	// Toggles ignoring of service worker for each request.
	SetBypassServiceWorker = "Network.setBypassServiceWorker"
	
	// Toggles ignoring cache for each request. If `true`, cache will not be used.
	SetCacheDisabled = "Network.setCacheDisabled"
	
	// Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
	SetCookie = "Network.setCookie"
	
	// Sets given cookies.
	SetCookies = "Network.setCookies"
	
	// For testing.
	SetDataSizeLimitsForTest = "Network.setDataSizeLimitsForTest"
	
	// Specifies whether to always send extra HTTP headers with the requests from this page.
	SetExtraHTTPHeaders = "Network.setExtraHTTPHeaders"
	
	// Sets the requests to intercept that match the provided patterns and optionally resource types.
	SetRequestInterception = "Network.setRequestInterception"
	
	// Allows overriding user agent with the given string.
	SetUserAgentOverride = "Network.setUserAgentOverride"
	
)

// CanClearBrowserCache parameters
type CanClearBrowserCacheParams struct {
	
}

// CanClearBrowserCache returns
type CanClearBrowserCacheReturns struct {
	
	// True if browser cache can be cleared.
	Result	bool	`json:"result"`
	
}

// CanClearBrowserCookies parameters
type CanClearBrowserCookiesParams struct {
	
}

// CanClearBrowserCookies returns
type CanClearBrowserCookiesReturns struct {
	
	// True if browser cookies can be cleared.
	Result	bool	`json:"result"`
	
}

// CanEmulateNetworkConditions parameters
type CanEmulateNetworkConditionsParams struct {
	
}

// CanEmulateNetworkConditions returns
type CanEmulateNetworkConditionsReturns struct {
	
	// True if emulation of network conditions is supported.
	Result	bool	`json:"result"`
	
}

// ClearBrowserCache parameters
type ClearBrowserCacheParams struct {
	
}

// ClearBrowserCache returns
type ClearBrowserCacheReturns struct {
	
}

// ClearBrowserCookies parameters
type ClearBrowserCookiesParams struct {
	
}

// ClearBrowserCookies returns
type ClearBrowserCookiesReturns struct {
	
}

// ContinueInterceptedRequest parameters
type ContinueInterceptedRequestParams struct {
	
	
	InterceptionId	InterceptionId	`json:"interceptionId"`
	
	// If set this causes the request to fail with the given reason. Passing `Aborted` for requests
	// marked with `isNavigationRequest` also cancels the navigation. Must not be set in response
	// to an authChallenge.
	ErrorReason	ErrorReason	`json:"errorReason"`
	
	// If set the requests completes using with the provided base64 encoded raw response, including
	// HTTP status line and headers etc... Must not be set in response to an authChallenge.
	RawResponse	string	`json:"rawResponse"`
	
	// If set the request url will be modified in a way that's not observable by page. Must not be
	// set in response to an authChallenge.
	Url	string	`json:"url"`
	
	// If set this allows the request method to be overridden. Must not be set in response to an
	// authChallenge.
	Method	string	`json:"method"`
	
	// If set this allows postData to be set. Must not be set in response to an authChallenge.
	PostData	string	`json:"postData"`
	
	// If set this allows the request headers to be changed. Must not be set in response to an
	// authChallenge.
	Headers	Headers	`json:"headers"`
	
	// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
	AuthChallengeResponse	AuthChallengeResponse	`json:"authChallengeResponse"`
	
}

// ContinueInterceptedRequest returns
type ContinueInterceptedRequestReturns struct {
	
}

// DeleteCookies parameters
type DeleteCookiesParams struct {
	
	// Name of the cookies to remove.
	Name	string	`json:"name"`
	
	// If specified, deletes all the cookies with the given name where domain and path match
	// provided URL.
	Url	string	`json:"url"`
	
	// If specified, deletes only cookies with the exact domain.
	Domain	string	`json:"domain"`
	
	// If specified, deletes only cookies with the exact path.
	Path	string	`json:"path"`
	
}

// DeleteCookies returns
type DeleteCookiesReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// EmulateNetworkConditions parameters
type EmulateNetworkConditionsParams struct {
	
	// True to emulate internet disconnection.
	Offline	bool	`json:"offline"`
	
	// Minimum latency from request sent to response headers received (ms).
	Latency	float64	`json:"latency"`
	
	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput	float64	`json:"downloadThroughput"`
	
	// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	UploadThroughput	float64	`json:"uploadThroughput"`
	
	// Connection type if known.
	ConnectionType	ConnectionType	`json:"connectionType"`
	
}

// EmulateNetworkConditions returns
type EmulateNetworkConditionsReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxTotalBufferSize	int	`json:"maxTotalBufferSize"`
	
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxResourceBufferSize	int	`json:"maxResourceBufferSize"`
	
	// Longest post body size (in bytes) that would be included in requestWillBeSent notification
	MaxPostDataSize	int	`json:"maxPostDataSize"`
	
}

// Enable returns
type EnableReturns struct {
	
}

// GetAllCookies parameters
type GetAllCookiesParams struct {
	
}

// GetAllCookies returns
type GetAllCookiesReturns struct {
	
	// Array of cookie objects.
	Cookies	[]Cookie	`json:"cookies"`
	
}

// GetCertificate parameters
type GetCertificateParams struct {
	
	// Origin to get certificate for.
	Origin	string	`json:"origin"`
	
}

// GetCertificate returns
type GetCertificateReturns struct {
	
	
	TableNames	[]string	`json:"tableNames"`
	
}

// GetCookies parameters
type GetCookiesParams struct {
	
	// The list of URLs for which applicable cookies will be fetched
	Urls	[]string	`json:"urls"`
	
}

// GetCookies returns
type GetCookiesReturns struct {
	
	// Array of cookie objects.
	Cookies	[]Cookie	`json:"cookies"`
	
}

// GetResponseBody parameters
type GetResponseBodyParams struct {
	
	// Identifier of the network request to get content for.
	RequestId	RequestId	`json:"requestId"`
	
}

// GetResponseBody returns
type GetResponseBodyReturns struct {
	
	// Response body.
	Body	string	`json:"body"`
	
	// True, if content was sent as base64.
	Base64Encoded	bool	`json:"base64Encoded"`
	
}

// GetRequestPostData parameters
type GetRequestPostDataParams struct {
	
	// Identifier of the network request to get content for.
	RequestId	RequestId	`json:"requestId"`
	
}

// GetRequestPostData returns
type GetRequestPostDataReturns struct {
	
	// Request body string, omitting files from multipart requests
	PostData	string	`json:"postData"`
	
}

// GetResponseBodyForInterception parameters
type GetResponseBodyForInterceptionParams struct {
	
	// Identifier for the intercepted request to get body for.
	InterceptionId	InterceptionId	`json:"interceptionId"`
	
}

// GetResponseBodyForInterception returns
type GetResponseBodyForInterceptionReturns struct {
	
	// Response body.
	Body	string	`json:"body"`
	
	// True, if content was sent as base64.
	Base64Encoded	bool	`json:"base64Encoded"`
	
}

// TakeResponseBodyForInterceptionAsStream parameters
type TakeResponseBodyForInterceptionAsStreamParams struct {
	
	
	InterceptionId	InterceptionId	`json:"interceptionId"`
	
}

// TakeResponseBodyForInterceptionAsStream returns
type TakeResponseBodyForInterceptionAsStreamReturns struct {
	
	
	Stream	io.StreamHandle	`json:"stream"`
	
}

// ReplayXHR parameters
type ReplayXHRParams struct {
	
	// Identifier of XHR to replay.
	RequestId	RequestId	`json:"requestId"`
	
}

// ReplayXHR returns
type ReplayXHRReturns struct {
	
}

// SearchInResponseBody parameters
type SearchInResponseBodyParams struct {
	
	// Identifier of the network response to search.
	RequestId	RequestId	`json:"requestId"`
	
	// String to search for.
	Query	string	`json:"query"`
	
	// If true, search is case sensitive.
	CaseSensitive	bool	`json:"caseSensitive"`
	
	// If true, treats string parameter as regex.
	IsRegex	bool	`json:"isRegex"`
	
}

// SearchInResponseBody returns
type SearchInResponseBodyReturns struct {
	
	// List of search matches.
	Result	[]debugger.SearchMatch	`json:"result"`
	
}

// SetBlockedURLs parameters
type SetBlockedURLsParams struct {
	
	// URL patterns to block. Wildcards ('*') are allowed.
	Urls	[]string	`json:"urls"`
	
}

// SetBlockedURLs returns
type SetBlockedURLsReturns struct {
	
}

// SetBypassServiceWorker parameters
type SetBypassServiceWorkerParams struct {
	
	// Bypass service worker and load from network.
	Bypass	bool	`json:"bypass"`
	
}

// SetBypassServiceWorker returns
type SetBypassServiceWorkerReturns struct {
	
}

// SetCacheDisabled parameters
type SetCacheDisabledParams struct {
	
	// Cache disabled state.
	CacheDisabled	bool	`json:"cacheDisabled"`
	
}

// SetCacheDisabled returns
type SetCacheDisabledReturns struct {
	
}

// SetCookie parameters
type SetCookieParams struct {
	
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

// SetCookie returns
type SetCookieReturns struct {
	
	// True if successfully set cookie.
	Success	bool	`json:"success"`
	
}

// SetCookies parameters
type SetCookiesParams struct {
	
	// Cookies to be set.
	Cookies	[]CookieParam	`json:"cookies"`
	
}

// SetCookies returns
type SetCookiesReturns struct {
	
}

// SetDataSizeLimitsForTest parameters
type SetDataSizeLimitsForTestParams struct {
	
	// Maximum total buffer size.
	MaxTotalSize	int	`json:"maxTotalSize"`
	
	// Maximum per-resource size.
	MaxResourceSize	int	`json:"maxResourceSize"`
	
}

// SetDataSizeLimitsForTest returns
type SetDataSizeLimitsForTestReturns struct {
	
}

// SetExtraHTTPHeaders parameters
type SetExtraHTTPHeadersParams struct {
	
	// Map with extra HTTP headers.
	Headers	Headers	`json:"headers"`
	
}

// SetExtraHTTPHeaders returns
type SetExtraHTTPHeadersReturns struct {
	
}

// SetRequestInterception parameters
type SetRequestInterceptionParams struct {
	
	// Requests matching any of these patterns will be forwarded and wait for the corresponding
	// continueInterceptedRequest call.
	Patterns	[]RequestPattern	`json:"patterns"`
	
}

// SetRequestInterception returns
type SetRequestInterceptionReturns struct {
	
}

// SetUserAgentOverride parameters
type SetUserAgentOverrideParams struct {
	
	// User agent to use.
	UserAgent	string	`json:"userAgent"`
	
	// Browser langugage to emulate.
	AcceptLanguage	string	`json:"acceptLanguage"`
	
	// The platform navigator.platform should return.
	Platform	string	`json:"platform"`
	
}

// SetUserAgentOverride returns
type SetUserAgentOverrideReturns struct {
	
}

