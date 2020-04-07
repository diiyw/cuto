package log

// Issued when new message was logged.
const EntryAddedEvent = "Log.entryAdded"
type EntryAddedParams struct {

	// The entry.
	Entry 	LogEntry}

