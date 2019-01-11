package indexeddb

const (
	
	// Clears all entries from an object store.
	ClearObjectStore = "IndexedDB.clearObjectStore"
	
	// Deletes a database.
	DeleteDatabase = "IndexedDB.deleteDatabase"
	
	// Delete a range of entries from an object store
	DeleteObjectStoreEntries = "IndexedDB.deleteObjectStoreEntries"
	
	// Disables events from backend.
	Disable = "IndexedDB.disable"
	
	// Enables events from backend.
	Enable = "IndexedDB.enable"
	
	// Requests data from object store or index.
	RequestData = "IndexedDB.requestData"
	
	// Requests database with given name in given frame.
	RequestDatabase = "IndexedDB.requestDatabase"
	
	// Requests database names for given security origin.
	RequestDatabaseNames = "IndexedDB.requestDatabaseNames"
	
)

// ClearObjectStore parameters
type ClearObjectStoreParams struct {
	
	// Security origin.
	SecurityOrigin	string	`json:"securityOrigin"`
	
	// Database name.
	DatabaseName	string	`json:"databaseName"`
	
	// Object store name.
	ObjectStoreName	string	`json:"objectStoreName"`
	
}

// ClearObjectStore returns
type ClearObjectStoreReturns struct {
	
}

// DeleteDatabase parameters
type DeleteDatabaseParams struct {
	
	// Security origin.
	SecurityOrigin	string	`json:"securityOrigin"`
	
	// Database name.
	DatabaseName	string	`json:"databaseName"`
	
}

// DeleteDatabase returns
type DeleteDatabaseReturns struct {
	
}

// DeleteObjectStoreEntries parameters
type DeleteObjectStoreEntriesParams struct {
	
	
	SecurityOrigin	string	`json:"securityOrigin"`
	
	
	DatabaseName	string	`json:"databaseName"`
	
	
	ObjectStoreName	string	`json:"objectStoreName"`
	
	// Range of entry keys to delete
	KeyRange	KeyRange	`json:"keyRange"`
	
}

// DeleteObjectStoreEntries returns
type DeleteObjectStoreEntriesReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// RequestData parameters
type RequestDataParams struct {
	
	// Security origin.
	SecurityOrigin	string	`json:"securityOrigin"`
	
	// Database name.
	DatabaseName	string	`json:"databaseName"`
	
	// Object store name.
	ObjectStoreName	string	`json:"objectStoreName"`
	
	// Index name, empty string for object store data requests.
	IndexName	string	`json:"indexName"`
	
	// Number of records to skip.
	SkipCount	int	`json:"skipCount"`
	
	// Number of records to fetch.
	PageSize	int	`json:"pageSize"`
	
	// Key range.
	KeyRange	KeyRange	`json:"keyRange"`
	
}

// RequestData returns
type RequestDataReturns struct {
	
	// Array of object store data entries.
	ObjectStoreDataEntries	[]DataEntry	`json:"objectStoreDataEntries"`
	
	// If true, there are more entries to fetch in the given range.
	HasMore	bool	`json:"hasMore"`
	
}

// RequestDatabase parameters
type RequestDatabaseParams struct {
	
	// Security origin.
	SecurityOrigin	string	`json:"securityOrigin"`
	
	// Database name.
	DatabaseName	string	`json:"databaseName"`
	
}

// RequestDatabase returns
type RequestDatabaseReturns struct {
	
	// Database with an array of object stores.
	DatabaseWithObjectStores	DatabaseWithObjectStores	`json:"databaseWithObjectStores"`
	
}

// RequestDatabaseNames parameters
type RequestDatabaseNamesParams struct {
	
	// Security origin.
	SecurityOrigin	string	`json:"securityOrigin"`
	
}

// RequestDatabaseNames returns
type RequestDatabaseNamesReturns struct {
	
	// Database names for origin.
	DatabaseNames	[]string	`json:"databaseNames"`
	
}

