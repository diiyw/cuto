package cachestorage
// Unique identifier of the Cache object.
type CacheId string

// type of HTTP response cached
type CachedResponseType string

// Data entry.
type DataEntry  struct {

	// Request URL.
	RequestURL	string	`json:"requestURL"`

	// Request method.
	RequestMethod	string	`json:"requestMethod"`

	// Request headers
	RequestHeaders	[]*Header	`json:"requestHeaders"`

	// Number of seconds since epoch.
	ResponseTime	float64	`json:"responseTime"`

	// HTTP response status code.
	ResponseStatus	int	`json:"responseStatus"`

	// HTTP response status text.
	ResponseStatusText	string	`json:"responseStatusText"`

	// HTTP response type
	ResponseType	CachedResponseType	`json:"responseType"`

	// Response headers
	ResponseHeaders	[]*Header	`json:"responseHeaders"`
}

// Cache identifier.
type Cache  struct {

	// An opaque unique id of the cache.
	CacheId	CacheId	`json:"cacheId"`

	// Security origin of the cache.
	SecurityOrigin	string	`json:"securityOrigin"`

	// The name of the cache.
	CacheName	string	`json:"cacheName"`
}

// 
type Header  struct {

	// 
	Name	string	`json:"name"`

	// 
	Value	string	`json:"value"`
}

// Cached response
type CachedResponse  struct {

	// Entry content, base64-encoded.
	Body	[]byte	`json:"body"`
}
