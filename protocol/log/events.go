package log

const (
	
	// Issued when new message was logged.
	EntryAddedEvent = "Log.entryAdded"
	
)

// Issued when new message was logged.
type EntryAddedParams struct {
	
	// The entry.
	Entry	LogEntry	`json:"entry"`
	
}

