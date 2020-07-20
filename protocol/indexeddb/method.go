package indexeddb

// Clears all entries from an object store.
const ClearObjectStore = "IndexedDB.clearObjectStore"

type ClearObjectStoreParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`

	// Database name.
	DatabaseName 	string	`json:"databaseName"`

	// Object store name.
	ObjectStoreName 	string	`json:"objectStoreName"`
}

type ClearObjectStoreResult struct {

}

// Deletes a database.
const DeleteDatabase = "IndexedDB.deleteDatabase"

type DeleteDatabaseParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`

	// Database name.
	DatabaseName 	string	`json:"databaseName"`
}

type DeleteDatabaseResult struct {

}

// Delete a range of entries from an object store
const DeleteObjectStoreEntries = "IndexedDB.deleteObjectStoreEntries"

type DeleteObjectStoreEntriesParams struct {

	// 
	SecurityOrigin 	string	`json:"securityOrigin"`

	// 
	DatabaseName 	string	`json:"databaseName"`

	// 
	ObjectStoreName 	string	`json:"objectStoreName"`

	// Range of entry keys to delete
	KeyRange 	KeyRange	`json:"keyRange"`
}

type DeleteObjectStoreEntriesResult struct {

}

// Disables events from backend.
const Disable = "IndexedDB.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables events from backend.
const Enable = "IndexedDB.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Requests data from object store or index.
const RequestData = "IndexedDB.requestData"

type RequestDataParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`

	// Database name.
	DatabaseName 	string	`json:"databaseName"`

	// Object store name.
	ObjectStoreName 	string	`json:"objectStoreName"`

	// Index name, empty string for object store data requests.
	IndexName 	string	`json:"indexName"`

	// Number of records to skip.
	SkipCount 	int	`json:"skipCount"`

	// Number of records to fetch.
	PageSize 	int	`json:"pageSize"`

	// Key range.
	KeyRange 	KeyRange	`json:"keyRange,omitempty"`
}

type RequestDataResult struct {

	// Array of object store data entries.
	ObjectStoreDataEntries 	[]*DataEntry	`json:"objectStoreDataEntries"`
	// If true, there are more entries to fetch in the given range.
	HasMore 	bool	`json:"hasMore"`
}

// Gets metadata of an object store
const GetMetadata = "IndexedDB.getMetadata"

type GetMetadataParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`

	// Database name.
	DatabaseName 	string	`json:"databaseName"`

	// Object store name.
	ObjectStoreName 	string	`json:"objectStoreName"`
}

type GetMetadataResult struct {

	// the entries count
	EntriesCount 	float64	`json:"entriesCount"`
	// the current value of key generator, to become the next inserted
	// key into the object store. Valid if objectStore.autoIncrement
	// is true.
	KeyGeneratorValue 	float64	`json:"keyGeneratorValue"`
}

// Requests database with given name in given frame.
const RequestDatabase = "IndexedDB.requestDatabase"

type RequestDatabaseParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`

	// Database name.
	DatabaseName 	string	`json:"databaseName"`
}

type RequestDatabaseResult struct {

	// Database with an array of object stores.
	DatabaseWithObjectStores 	DatabaseWithObjectStores	`json:"databaseWithObjectStores"`
}

// Requests database names for given security origin.
const RequestDatabaseNames = "IndexedDB.requestDatabaseNames"

type RequestDatabaseNamesParams struct {

	// Security origin.
	SecurityOrigin 	string	`json:"securityOrigin"`
}

type RequestDatabaseNamesResult struct {

	// Database names for origin.
	DatabaseNames 	[]string	`json:"databaseNames"`
}