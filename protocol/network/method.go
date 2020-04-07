package network

import (
	"github.com/diiyw/cuto/protocol/debugger"
)


// Tells whether clearing browser cache is supported.
const CanClearBrowserCache = "Network.canClearBrowserCache"

type CanClearBrowserCacheParams struct {
}

type CanClearBrowserCacheResult struct {

	// True if browser cache can be cleared.
	Result 	bool	`json:"result"`
}

// Tells whether clearing browser cookies is supported.
const CanClearBrowserCookies = "Network.canClearBrowserCookies"

type CanClearBrowserCookiesParams struct {
}

type CanClearBrowserCookiesResult struct {

	// True if browser cookies can be cleared.
	Result 	bool	`json:"result"`
}

// Tells whether emulation of network conditions is supported.
const CanEmulateNetworkConditions = "Network.canEmulateNetworkConditions"

type CanEmulateNetworkConditionsParams struct {
}

type CanEmulateNetworkConditionsResult struct {

	// True if emulation of network conditions is supported.
	Result 	bool	`json:"result"`
}

// Clears browser cache.
const ClearBrowserCache = "Network.clearBrowserCache"

type ClearBrowserCacheParams struct {
}

type ClearBrowserCacheResult struct {

}

// Clears browser cookies.
const ClearBrowserCookies = "Network.clearBrowserCookies"

type ClearBrowserCookiesParams struct {
}

type ClearBrowserCookiesResult struct {

}

// Response to Network.requestIntercepted which either modifies the request to continue with any
// modifications, or blocks it, or completes it with the provided response bytes. If a network
// fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted
// event will be sent with the same InterceptionId.
// Deprecated, use Fetch.continueRequest, Fetch.fulfillRequest and Fetch.failRequest instead.
const ContinueInterceptedRequest = "Network.continueInterceptedRequest"

type ContinueInterceptedRequestParams struct {

	// 
	InterceptionId 	InterceptionId	`json:"interceptionId"`

	// If set this causes the request to fail with the given reason. Passing `Aborted` for requests
	// marked with `isNavigationRequest` also cancels the navigation. Must not be set in response
	// to an authChallenge.
	ErrorReason 	ErrorReason	`json:"errorReason"`

	// If set the requests completes using with the provided base64 encoded raw response, including
	// HTTP status line and headers etc... Must not be set in response to an authChallenge.
	RawResponse 	[]byte	`json:"rawResponse"`

	// If set the request url will be modified in a way that's not observable by page. Must not be
	// set in response to an authChallenge.
	Url 	string	`json:"url"`

	// If set this allows the request method to be overridden. Must not be set in response to an
	// authChallenge.
	Method 	string	`json:"method"`

	// If set this allows postData to be set. Must not be set in response to an authChallenge.
	PostData 	string	`json:"postData"`

	// If set this allows the request headers to be changed. Must not be set in response to an
	// authChallenge.
	Headers 	Headers	`json:"headers"`

	// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
	AuthChallengeResponse 	AuthChallengeResponse	`json:"authChallengeResponse"`
}

type ContinueInterceptedRequestResult struct {

}

// Deletes browser cookies with matching name and url or domain/path pair.
const DeleteCookies = "Network.deleteCookies"

type DeleteCookiesParams struct {

	// Name of the cookies to remove.
	Name 	string	`json:"name"`

	// If specified, deletes all the cookies with the given name where domain and path match
	// provided URL.
	Url 	string	`json:"url"`

	// If specified, deletes only cookies with the exact domain.
	Domain 	string	`json:"domain"`

	// If specified, deletes only cookies with the exact path.
	Path 	string	`json:"path"`
}

type DeleteCookiesResult struct {

}

// Disables network tracking, prevents network events from being sent to the client.
const Disable = "Network.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Activates emulation of network conditions.
const EmulateNetworkConditions = "Network.emulateNetworkConditions"

type EmulateNetworkConditionsParams struct {

	// True to emulate internet disconnection.
	Offline 	bool	`json:"offline"`

	// Minimum latency from request sent to response headers received (ms).
	Latency 	float64	`json:"latency"`

	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput 	float64	`json:"downloadThroughput"`

	// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	UploadThroughput 	float64	`json:"uploadThroughput"`

	// Connection type if known.
	ConnectionType 	ConnectionType	`json:"connectionType"`
}

type EmulateNetworkConditionsResult struct {

}

// Enables network tracking, network events will now be delivered to the client.
const Enable = "Network.enable"

type EnableParams struct {

	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxTotalBufferSize 	int	`json:"maxTotalBufferSize"`

	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxResourceBufferSize 	int	`json:"maxResourceBufferSize"`

	// Longest post body size (in bytes) that would be included in requestWillBeSent notification
	MaxPostDataSize 	int	`json:"maxPostDataSize"`
}

type EnableResult struct {

}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie
// information in the `cookies` field.
const GetAllCookies = "Network.getAllCookies"

type GetAllCookiesParams struct {
}

type GetAllCookiesResult struct {

	// Array of cookie objects.
	Cookies 	[]*Cookie	`json:"cookies"`
}

// Returns the DER-encoded certificate.
const GetCertificate = "Network.getCertificate"

type GetCertificateParams struct {

	// Origin to get certificate for.
	Origin 	string	`json:"origin"`
}

type GetCertificateResult struct {

	// 
	TableNames 	[]string	`json:"tableNames"`
}

// Returns all browser cookies for the current URL. Depending on the backend support, will return
// detailed cookie information in the `cookies` field.
const GetCookies = "Network.getCookies"

type GetCookiesParams struct {

	// The list of URLs for which applicable cookies will be fetched
	Urls 	[]string	`json:"urls"`
}

type GetCookiesResult struct {

	// Array of cookie objects.
	Cookies 	[]*Cookie	`json:"cookies"`
}

// Returns content served for the given request.
const GetResponseBody = "Network.getResponseBody"

type GetResponseBodyParams struct {

	// Identifier of the network request to get content for.
	RequestId 	RequestId	`json:"requestId"`
}

type GetResponseBodyResult struct {

	// Response body.
	Body 	string	`json:"body"`
	// True, if content was sent as base64.
	Base64Encoded 	bool	`json:"base64Encoded"`
}

// Returns post data sent with the request. Returns an error when no data was sent with the request.
const GetRequestPostData = "Network.getRequestPostData"

type GetRequestPostDataParams struct {

	// Identifier of the network request to get content for.
	RequestId 	RequestId	`json:"requestId"`
}

type GetRequestPostDataResult struct {

	// Request body string, omitting files from multipart requests
	PostData 	string	`json:"postData"`
}

// Returns content served for the given currently intercepted request.
const GetResponseBodyForInterception = "Network.getResponseBodyForInterception"

type GetResponseBodyForInterceptionParams struct {

	// Identifier for the intercepted request to get body for.
	InterceptionId 	InterceptionId	`json:"interceptionId"`
}

type GetResponseBodyForInterceptionResult struct {

	// Response body.
	Body 	string	`json:"body"`
	// True, if content was sent as base64.
	Base64Encoded 	bool	`json:"base64Encoded"`
}

// Returns a handle to the stream representing the response body. Note that after this command,
// the intercepted request can't be continued as is -- you either need to cancel it or to provide
// the response body. The stream only supports sequential read, IO.read will fail if the position
// is specified.
const TakeResponseBodyForInterceptionAsStream = "Network.takeResponseBodyForInterceptionAsStream"

type TakeResponseBodyForInterceptionAsStreamParams struct {

	// 
	InterceptionId 	InterceptionId	`json:"interceptionId"`
}

type TakeResponseBodyForInterceptionAsStreamResult struct {

	// 
	Stream 	interface{}	`json:"stream"`
}

// This method sends a new XMLHttpRequest which is identical to the original one. The following
// parameters should be identical: method, url, async, request body, extra headers, withCredentials
// attribute, user, password.
const ReplayXHR = "Network.replayXHR"

type ReplayXHRParams struct {

	// Identifier of XHR to replay.
	RequestId 	RequestId	`json:"requestId"`
}

type ReplayXHRResult struct {

}

// Searches for given string in response content.
const SearchInResponseBody = "Network.searchInResponseBody"

type SearchInResponseBodyParams struct {

	// Identifier of the network response to search.
	RequestId 	RequestId	`json:"requestId"`

	// String to search for.
	Query 	string	`json:"query"`

	// If true, search is case sensitive.
	CaseSensitive 	bool	`json:"caseSensitive"`

	// If true, treats string parameter as regex.
	IsRegex 	bool	`json:"isRegex"`
}

type SearchInResponseBodyResult struct {

	// List of search matches.
	Result 	[]*debugger.SearchMatch	`json:"result"`
}

// Blocks URLs from loading.
const SetBlockedURLs = "Network.setBlockedURLs"

type SetBlockedURLsParams struct {

	// URL patterns to block. Wildcards ('*') are allowed.
	Urls 	[]string	`json:"urls"`
}

type SetBlockedURLsResult struct {

}

// Toggles ignoring of service worker for each request.
const SetBypassServiceWorker = "Network.setBypassServiceWorker"

type SetBypassServiceWorkerParams struct {

	// Bypass service worker and load from network.
	Bypass 	bool	`json:"bypass"`
}

type SetBypassServiceWorkerResult struct {

}

// Toggles ignoring cache for each request. If `true`, cache will not be used.
const SetCacheDisabled = "Network.setCacheDisabled"

type SetCacheDisabledParams struct {

	// Cache disabled state.
	CacheDisabled 	bool	`json:"cacheDisabled"`
}

type SetCacheDisabledResult struct {

}

// Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
const SetCookie = "Network.setCookie"

type SetCookieParams struct {

	// Cookie name.
	Name 	string	`json:"name"`

	// Cookie value.
	Value 	string	`json:"value"`

	// The request-URI to associate with the setting of the cookie. This value can affect the
	// default domain and path values of the created cookie.
	Url 	string	`json:"url"`

	// Cookie domain.
	Domain 	string	`json:"domain"`

	// Cookie path.
	Path 	string	`json:"path"`

	// True if cookie is secure.
	Secure 	bool	`json:"secure"`

	// True if cookie is http-only.
	HttpOnly 	bool	`json:"httpOnly"`

	// Cookie SameSite type.
	SameSite 	CookieSameSite	`json:"sameSite"`

	// Cookie expiration date, session cookie if not set
	Expires 	TimeSinceEpoch	`json:"expires"`
}

type SetCookieResult struct {

	// True if successfully set cookie.
	Success 	bool	`json:"success"`
}

// Sets given cookies.
const SetCookies = "Network.setCookies"

type SetCookiesParams struct {

	// Cookies to be set.
	Cookies 	[]*CookieParam	`json:"cookies"`
}

type SetCookiesResult struct {

}

// For testing.
const SetDataSizeLimitsForTest = "Network.setDataSizeLimitsForTest"

type SetDataSizeLimitsForTestParams struct {

	// Maximum total buffer size.
	MaxTotalSize 	int	`json:"maxTotalSize"`

	// Maximum per-resource size.
	MaxResourceSize 	int	`json:"maxResourceSize"`
}

type SetDataSizeLimitsForTestResult struct {

}

// Specifies whether to always send extra HTTP headers with the requests from this page.
const SetExtraHTTPHeaders = "Network.setExtraHTTPHeaders"

type SetExtraHTTPHeadersParams struct {

	// Map with extra HTTP headers.
	Headers 	Headers	`json:"headers"`
}

type SetExtraHTTPHeadersResult struct {

}

// Sets the requests to intercept that match the provided patterns and optionally resource types.
// Deprecated, please use Fetch.enable instead.
const SetRequestInterception = "Network.setRequestInterception"

type SetRequestInterceptionParams struct {

	// Requests matching any of these patterns will be forwarded and wait for the corresponding
	// continueInterceptedRequest call.
	Patterns 	[]*RequestPattern	`json:"patterns"`
}

type SetRequestInterceptionResult struct {

}

// Allows overriding user agent with the given string.
const SetUserAgentOverride = "Network.setUserAgentOverride"

type SetUserAgentOverrideParams struct {

	// User agent to use.
	UserAgent 	string	`json:"userAgent"`

	// Browser langugage to emulate.
	AcceptLanguage 	string	`json:"acceptLanguage"`

	// The platform navigator.platform should return.
	Platform 	string	`json:"platform"`
}

type SetUserAgentOverrideResult struct {

}