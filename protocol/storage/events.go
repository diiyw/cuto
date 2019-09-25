package storage

const (
	
	// A cache's contents have been modified.
	CacheStorageContentUpdatedEvent = "Storage.cacheStorageContentUpdated"
	
	// A cache has been added/deleted.
	CacheStorageListUpdatedEvent = "Storage.cacheStorageListUpdated"
	
	// The origin's IndexedDB object store has been modified.
	IndexedDBContentUpdatedEvent = "Storage.indexedDBContentUpdated"
	
	// The origin's IndexedDB database list has been modified.
	IndexedDBListUpdatedEvent = "Storage.indexedDBListUpdated"
	
)

// A cache's contents have been modified.
type CacheStorageContentUpdatedParams struct {
	
	// Origin to update.
	Origin	string	`json:"origin"`
	
	// Name of cache in origin.
	CacheName	string	`json:"cacheName"`
	
}

// A cache has been added/deleted.
type CacheStorageListUpdatedParams struct {
	
	// Origin to update.
	Origin	string	`json:"origin"`
	
}

// The origin's IndexedDB object store has been modified.
type IndexedDBContentUpdatedParams struct {
	
	// Origin to update.
	Origin	string	`json:"origin"`
	
	// Database to update.
	DatabaseName	string	`json:"databaseName"`
	
	// ObjectStore to update.
	ObjectStoreName	string	`json:"objectStoreName"`
	
}

// The origin's IndexedDB database list has been modified.
type IndexedDBListUpdatedParams struct {
	
	// Origin to update.
	Origin	string	`json:"origin"`
	
}

