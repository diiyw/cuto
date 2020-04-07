package tethering

// Informs that port was successfully bound and got a specified connection id.
const AcceptedEvent = "Tethering.accepted"
type AcceptedParams struct {

	// Port number that was successfully bound.
	Port 	int
	// Connection id to be used.
	ConnectionId 	string}

