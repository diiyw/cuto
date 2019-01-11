package domstorage

const (
	
	
	Clear = "DOMStorage.clear"
	
	// Disables storage tracking, prevents storage events from being sent to the client.
	Disable = "DOMStorage.disable"
	
	// Enables storage tracking, storage events will now be delivered to the client.
	Enable = "DOMStorage.enable"
	
	
	GetDOMStorageItems = "DOMStorage.getDOMStorageItems"
	
	
	RemoveDOMStorageItem = "DOMStorage.removeDOMStorageItem"
	
	
	SetDOMStorageItem = "DOMStorage.setDOMStorageItem"
	
)

// Clear parameters
type ClearParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
}

// Clear returns
type ClearReturns struct {
	
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

// GetDOMStorageItems parameters
type GetDOMStorageItemsParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
}

// GetDOMStorageItems returns
type GetDOMStorageItemsReturns struct {
	
	
	Entries	[]Item	`json:"entries"`
	
}

// RemoveDOMStorageItem parameters
type RemoveDOMStorageItemParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
	
	Key	string	`json:"key"`
	
}

// RemoveDOMStorageItem returns
type RemoveDOMStorageItemReturns struct {
	
}

// SetDOMStorageItem parameters
type SetDOMStorageItemParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
	
	Key	string	`json:"key"`
	
	
	Value	string	`json:"value"`
	
}

// SetDOMStorageItem returns
type SetDOMStorageItemReturns struct {
	
}

