package storage

// A cache's contents have been modified.
const CacheStorageContentUpdatedEvent = "Storage.cacheStorageContentUpdated"
type CacheStorageContentUpdatedParams struct {

	// Origin to update.
	Origin 	string
	// Name of cache in origin.
	CacheName 	string}



// A cache has been added/deleted.
const CacheStorageListUpdatedEvent = "Storage.cacheStorageListUpdated"
type CacheStorageListUpdatedParams struct {

	// Origin to update.
	Origin 	string}



// The origin's IndexedDB object store has been modified.
const IndexedDBContentUpdatedEvent = "Storage.indexedDBContentUpdated"
type IndexedDBContentUpdatedParams struct {

	// Origin to update.
	Origin 	string
	// Database to update.
	DatabaseName 	string
	// ObjectStore to update.
	ObjectStoreName 	string}



// The origin's IndexedDB database list has been modified.
const IndexedDBListUpdatedEvent = "Storage.indexedDBListUpdated"
type IndexedDBListUpdatedParams struct {

	// Origin to update.
	Origin 	string}

