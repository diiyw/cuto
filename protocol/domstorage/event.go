package domstorage

// 
const DomStorageItemAddedEvent = "DOMStorage.domStorageItemAdded"
type DomStorageItemAddedParams struct {

	// 
	StorageId 	StorageId
	// 
	Key 	string
	// 
	NewValue 	string}



// 
const DomStorageItemRemovedEvent = "DOMStorage.domStorageItemRemoved"
type DomStorageItemRemovedParams struct {

	// 
	StorageId 	StorageId
	// 
	Key 	string}



// 
const DomStorageItemUpdatedEvent = "DOMStorage.domStorageItemUpdated"
type DomStorageItemUpdatedParams struct {

	// 
	StorageId 	StorageId
	// 
	Key 	string
	// 
	OldValue 	string
	// 
	NewValue 	string}



// 
const DomStorageItemsClearedEvent = "DOMStorage.domStorageItemsCleared"
type DomStorageItemsClearedParams struct {

	// 
	StorageId 	StorageId}

