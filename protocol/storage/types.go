package storage


// Enum of possible storage types.
type StorageType string	

// Usage for a storage type.
type UsageForType struct {
	
	// Name of storage type.
	
	StorageType	StorageType	`json:"storageType"`
	
	// Storage usage (bytes).
	
	Usage	float64	`json:"usage"`
	
}	

