package console

// Issued when new console message is added.
const MessageAddedEvent = "Console.messageAdded"
type MessageAddedParams struct {

	// Console message that has been added.
	Message 	ConsoleMessage}

