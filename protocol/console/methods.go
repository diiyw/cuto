package console

const (
	
	// Does nothing.
	ClearMessages = "Console.clearMessages"
	
	// Disables console domain, prevents further console messages from being reported to the client.
	Disable = "Console.disable"
	
	// Enables console domain, sends the messages collected so far to the client by means of the
	// `messageAdded` notification.
	Enable = "Console.enable"
	
)

// ClearMessages parameters
type ClearMessagesParams struct {
	
}

// ClearMessages returns
type ClearMessagesReturns struct {
	
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

