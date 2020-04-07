package domstorage

// 
const Clear = "DOMStorage.clear"

type ClearParams struct {

	// 
	StorageId 	StorageId	`json:"storageId"`
}

type ClearResult struct {

}

// Disables storage tracking, prevents storage events from being sent to the client.
const Disable = "DOMStorage.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables storage tracking, storage events will now be delivered to the client.
const Enable = "DOMStorage.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// 
const GetDOMStorageItems = "DOMStorage.getDOMStorageItems"

type GetDOMStorageItemsParams struct {

	// 
	StorageId 	StorageId	`json:"storageId"`
}

type GetDOMStorageItemsResult struct {

	// 
	Entries 	[]*Item	`json:"entries"`
}

// 
const RemoveDOMStorageItem = "DOMStorage.removeDOMStorageItem"

type RemoveDOMStorageItemParams struct {

	// 
	StorageId 	StorageId	`json:"storageId"`

	// 
	Key 	string	`json:"key"`
}

type RemoveDOMStorageItemResult struct {

}

// 
const SetDOMStorageItem = "DOMStorage.setDOMStorageItem"

type SetDOMStorageItemParams struct {

	// 
	StorageId 	StorageId	`json:"storageId"`

	// 
	Key 	string	`json:"key"`

	// 
	Value 	string	`json:"value"`
}

type SetDOMStorageItemResult struct {

}