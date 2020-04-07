package domstorage
// DOM Storage identifier.
type StorageId  struct {

	// Security origin for the storage.
	SecurityOrigin	string	`json:"securityOrigin"`

	// Whether the storage is local storage (not session storage).
	IsLocalStorage	bool	`json:"isLocalStorage"`
}

// DOM Storage item.
type Item 	[]string
