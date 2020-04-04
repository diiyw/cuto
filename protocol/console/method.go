package console

// Does nothing.
const ClearMessages = "Console.clearMessages"

type ClearMessagesParams struct {
}

type ClearMessagesResult struct {

}

// Disables console domain, prevents further console messages from being reported to the client.
const Disable = "Console.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables console domain, sends the messages collected so far to the client by means of the
// `messageAdded` notification.
const Enable = "Console.enable"

type EnableParams struct {
}

type EnableResult struct {

}