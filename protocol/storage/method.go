package storage

// Clears storage for origin.
const ClearDataForOrigin = "Storage.clearDataForOrigin"

type ClearDataForOriginParams struct {

	// Security origin.
	Origin 	string	`json:"origin"`

	// Comma separated list of StorageType to clear.
	StorageTypes 	string	`json:"storageTypes"`
}

type ClearDataForOriginResult struct {

}

// Returns usage and quota in bytes.
const GetUsageAndQuota = "Storage.getUsageAndQuota"

type GetUsageAndQuotaParams struct {

	// Security origin.
	Origin 	string	`json:"origin"`
}

type GetUsageAndQuotaResult struct {

	// Storage usage (bytes).
	Usage 	float64	`json:"usage"`
	// Storage quota (bytes).
	Quota 	float64	`json:"quota"`
	// Storage usage per type (bytes).
	UsageBreakdown 	[]*UsageForType	`json:"usageBreakdown"`
}

// Registers origin to be notified when an update occurs to its cache storage list.
const TrackCacheStorageForOrigin = "Storage.trackCacheStorageForOrigin"

type TrackCacheStorageForOriginParams struct {

	// Security origin.
	Origin 	string	`json:"origin"`
}

type TrackCacheStorageForOriginResult struct {

}

// Registers origin to be notified when an update occurs to its IndexedDB.
const TrackIndexedDBForOrigin = "Storage.trackIndexedDBForOrigin"

type TrackIndexedDBForOriginParams struct {

	// Security origin.
	Origin 	string	`json:"origin"`
}

type TrackIndexedDBForOriginResult struct {

}

// Unregisters origin from receiving notifications for cache storage.
const UntrackCacheStorageForOrigin = "Storage.untrackCacheStorageForOrigin"

type UntrackCacheStorageForOriginParams struct {

	// Security origin.
	Origin 	string	`json:"origin"`
}

type UntrackCacheStorageForOriginResult struct {

}

// Unregisters origin from receiving notifications for IndexedDB.
const UntrackIndexedDBForOrigin = "Storage.untrackIndexedDBForOrigin"

type UntrackIndexedDBForOriginParams struct {

	// Security origin.
	Origin 	string	`json:"origin"`
}

type UntrackIndexedDBForOriginResult struct {

}