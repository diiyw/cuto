package indexeddb

import (
	"github.com/diiyw/cuto/protocol/runtime"
)

// Database with an array of object stores.
type DatabaseWithObjectStores  struct {

	// Database name.
	Name	string	`json:"name"`

	// Database version (type is not 'integer', as the standard
	// requires the version number to be 'unsigned long long')
	Version	float64	`json:"version"`

	// Object stores in this database.
	ObjectStores	[]*ObjectStore	`json:"objectStores"`
}

// Object store.
type ObjectStore  struct {

	// Object store name.
	Name	string	`json:"name"`

	// Object store key path.
	KeyPath	KeyPath	`json:"keyPath"`

	// If true, object store has auto increment flag set.
	AutoIncrement	bool	`json:"autoIncrement"`

	// Indexes in this object store.
	Indexes	[]*ObjectStoreIndex	`json:"indexes"`
}

// Object store index.
type ObjectStoreIndex  struct {

	// Index name.
	Name	string	`json:"name"`

	// Index key path.
	KeyPath	KeyPath	`json:"keyPath"`

	// If true, index is unique.
	Unique	bool	`json:"unique"`

	// If true, index allows multiple entries for a key.
	MultiEntry	bool	`json:"multiEntry"`
}

// Key.
type Key  struct {

	// Key type.
	Type	string	`json:"type"`

	// Number value.
	Number	float64	`json:"number,omitempty"`

	// String value.
	String	string	`json:"string,omitempty"`

	// Date value.
	Date	float64	`json:"date,omitempty"`

	// Array value.
	Array	[]*Key	`json:"array,omitempty"`
}

// Key range.
type KeyRange  struct {

	// Lower bound.
	Lower	Key	`json:"lower,omitempty"`

	// Upper bound.
	Upper	Key	`json:"upper,omitempty"`

	// If true lower bound is open.
	LowerOpen	bool	`json:"lowerOpen"`

	// If true upper bound is open.
	UpperOpen	bool	`json:"upperOpen"`
}

// Data entry.
type DataEntry  struct {

	// Key object.
	Key	runtime.RemoteObject	`json:"key"`

	// Primary key object.
	PrimaryKey	runtime.RemoteObject	`json:"primaryKey"`

	// Value object.
	Value	runtime.RemoteObject	`json:"value"`
}

// Key path.
type KeyPath  struct {

	// Key path type.
	Type	string	`json:"type"`

	// String value.
	String	string	`json:"string,omitempty"`

	// Array value.
	Array	[]string	`json:"array,omitempty"`
}
