package cachestorage

const (
	
	// Deletes a cache.
	DeleteCache = "CacheStorage.deleteCache"
	
	// Deletes a cache entry.
	DeleteEntry = "CacheStorage.deleteEntry"
	
	// Requests cache names.
	RequestCacheNames = "CacheStorage.requestCacheNames"
	
	// Fetches cache entry.
	RequestCachedResponse = "CacheStorage.requestCachedResponse"
	
	// Requests data from cache.
	RequestEntries = "CacheStorage.requestEntries"
	
)

// DeleteCache parameters
type DeleteCacheParams struct {
	
	// Id of cache for deletion.
	CacheId	CacheId	`json:"cacheId"`
	
}

// DeleteCache returns
type DeleteCacheReturns struct {
	
}

// DeleteEntry parameters
type DeleteEntryParams struct {
	
	// Id of cache where the entry will be deleted.
	CacheId	CacheId	`json:"cacheId"`
	
	// URL spec of the request.
	Request	string	`json:"request"`
	
}

// DeleteEntry returns
type DeleteEntryReturns struct {
	
}

// RequestCacheNames parameters
type RequestCacheNamesParams struct {
	
	// Security origin.
	SecurityOrigin	string	`json:"securityOrigin"`
	
}

// RequestCacheNames returns
type RequestCacheNamesReturns struct {
	
	// Caches for the security origin.
	Caches	[]Cache	`json:"caches"`
	
}

// RequestCachedResponse parameters
type RequestCachedResponseParams struct {
	
	// Id of cache that contains the entry.
	CacheId	CacheId	`json:"cacheId"`
	
	// URL spec of the request.
	RequestURL	string	`json:"requestURL"`
	
	// headers of the request.
	RequestHeaders	[]Header	`json:"requestHeaders"`
	
}

// RequestCachedResponse returns
type RequestCachedResponseReturns struct {
	
	// Response read from the cache.
	Response	CachedResponse	`json:"response"`
	
}

// RequestEntries parameters
type RequestEntriesParams struct {
	
	// ID of cache to get entries from.
	CacheId	CacheId	`json:"cacheId"`
	
	// Number of records to skip.
	SkipCount	int	`json:"skipCount"`
	
	// Number of records to fetch.
	PageSize	int	`json:"pageSize"`
	
	// If present, only return the entries containing this substring in the path
	PathFilter	string	`json:"pathFilter"`
	
}

// RequestEntries returns
type RequestEntriesReturns struct {
	
	// Array of object store data entries.
	CacheDataEntries	[]DataEntry	`json:"cacheDataEntries"`
	
	// Count of returned entries from this storage. If pathFilter is empty, it
	// is the count of all entries from this storage.
	ReturnCount	float64	`json:"returnCount"`
	
}

