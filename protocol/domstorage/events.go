package domstorage

const (
	
	
	DomStorageItemAddedEvent = "DOMStorage.domStorageItemAdded"
	
	
	DomStorageItemRemovedEvent = "DOMStorage.domStorageItemRemoved"
	
	
	DomStorageItemUpdatedEvent = "DOMStorage.domStorageItemUpdated"
	
	
	DomStorageItemsClearedEvent = "DOMStorage.domStorageItemsCleared"
	
)


type DomStorageItemAddedParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
	
	Key	string	`json:"key"`
	
	
	NewValue	string	`json:"newValue"`
	
}


type DomStorageItemRemovedParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
	
	Key	string	`json:"key"`
	
}


type DomStorageItemUpdatedParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
	
	Key	string	`json:"key"`
	
	
	OldValue	string	`json:"oldValue"`
	
	
	NewValue	string	`json:"newValue"`
	
}


type DomStorageItemsClearedParams struct {
	
	
	StorageId	StorageId	`json:"storageId"`
	
}

