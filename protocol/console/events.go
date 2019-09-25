package console

const (
	
	// Issued when new console message is added.
	MessageAddedEvent = "Console.messageAdded"
	
)

// Issued when new console message is added.
type MessageAddedParams struct {
	
	// Console message that has been added.
	Message	ConsoleMessage	`json:"message"`
	
}

