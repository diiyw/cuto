package cachestorage

// Deletes a cache.
const DeleteCache = "CacheStorage.deleteCache"

type DeleteCacheParams struct {

	// Id of cache for deletion.
	CacheId 	CacheId	`json:"cacheId"`
}

type DeleteCacheResult struct {

}

// Deletes a cache entry.
const DeleteEntry = "CacheStorage.deleteEntry"

type DeleteEntryParams struct {

	// Id of cache where the entry will be deleted.
	CacheId 	CacheId	`json:"cacheId"`

	// URL spec of the request.
	Request 	string	`json:"request"`
}

type DeleteEntryResult struct {

}

// Requests cache names.
const RequestCacheNames = "CacheStorage.requestCacheNames"

type RequestCacheNamesParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`
}

type RequestCacheNamesResult struct {

	// Caches for the security origin.
	Caches 	[]*Cache	`json:"caches"`
}

// Fetches cache entry.
const RequestCachedResponse = "CacheStorage.requestCachedResponse"

type RequestCachedResponseParams struct {

	// Id of cache that contains the entry.
	CacheId 	CacheId	`json:"cacheId"`

	// URL spec of the request.
	RequestURL 	string	`json:"requestURL"`

	// headers of the request.
	RequestHeaders 	[]*Header	`json:"requestHeaders"`
}

type RequestCachedResponseResult struct {

	// Response read from the cache.
	Response 	CachedResponse	`json:"response"`
}

// Requests data from cache.
const RequestEntries = "CacheStorage.requestEntries"

type RequestEntriesParams struct {

	// ID of cache to get entries from.
	CacheId 	CacheId	`json:"cacheId"`

	// Number of records to skip.
	SkipCount 	int	`json:"skipCount"`

	// Number of records to fetch.
	PageSize 	int	`json:"pageSize"`

	// If present, only return the entries containing this substring in the path
	PathFilter 	string	`json:"pathFilter"`
}

type RequestEntriesResult struct {

	// Array of object store data entries.
	CacheDataEntries 	[]*DataEntry	`json:"cacheDataEntries"`
	// Count of returned entries from this storage. If pathFilter is empty, it
	// is the count of all entries from this storage.
	ReturnCount 	float64	`json:"returnCount"`
}