package storage

const (
	
	// Clears storage for origin.
	ClearDataForOrigin = "Storage.clearDataForOrigin"
	
	// Returns usage and quota in bytes.
	GetUsageAndQuota = "Storage.getUsageAndQuota"
	
	// Registers origin to be notified when an update occurs to its cache storage list.
	TrackCacheStorageForOrigin = "Storage.trackCacheStorageForOrigin"
	
	// Registers origin to be notified when an update occurs to its IndexedDB.
	TrackIndexedDBForOrigin = "Storage.trackIndexedDBForOrigin"
	
	// Unregisters origin from receiving notifications for cache storage.
	UntrackCacheStorageForOrigin = "Storage.untrackCacheStorageForOrigin"
	
	// Unregisters origin from receiving notifications for IndexedDB.
	UntrackIndexedDBForOrigin = "Storage.untrackIndexedDBForOrigin"
	
)

// ClearDataForOrigin parameters
type ClearDataForOriginParams struct {
	
	// Security origin.
	Origin	string	`json:"origin"`
	
	// Comma separated list of StorageType to clear.
	StorageTypes	string	`json:"storageTypes"`
	
}

// ClearDataForOrigin returns
type ClearDataForOriginReturns struct {
	
}

// GetUsageAndQuota parameters
type GetUsageAndQuotaParams struct {
	
	// Security origin.
	Origin	string	`json:"origin"`
	
}

// GetUsageAndQuota returns
type GetUsageAndQuotaReturns struct {
	
	// Storage usage (bytes).
	Usage	float64	`json:"usage"`
	
	// Storage quota (bytes).
	Quota	float64	`json:"quota"`
	
	// Storage usage per type (bytes).
	UsageBreakdown	[]UsageForType	`json:"usageBreakdown"`
	
}

// TrackCacheStorageForOrigin parameters
type TrackCacheStorageForOriginParams struct {
	
	// Security origin.
	Origin	string	`json:"origin"`
	
}

// TrackCacheStorageForOrigin returns
type TrackCacheStorageForOriginReturns struct {
	
}

// TrackIndexedDBForOrigin parameters
type TrackIndexedDBForOriginParams struct {
	
	// Security origin.
	Origin	string	`json:"origin"`
	
}

// TrackIndexedDBForOrigin returns
type TrackIndexedDBForOriginReturns struct {
	
}

// UntrackCacheStorageForOrigin parameters
type UntrackCacheStorageForOriginParams struct {
	
	// Security origin.
	Origin	string	`json:"origin"`
	
}

// UntrackCacheStorageForOrigin returns
type UntrackCacheStorageForOriginReturns struct {
	
}

// UntrackIndexedDBForOrigin parameters
type UntrackIndexedDBForOriginParams struct {
	
	// Security origin.
	Origin	string	`json:"origin"`
	
}

// UntrackIndexedDBForOrigin returns
type UntrackIndexedDBForOriginReturns struct {
	
}

