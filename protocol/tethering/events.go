package tethering

const (
	
	// Informs that port was successfully bound and got a specified connection id.
	AcceptedEvent = "Tethering.accepted"
	
)

// Informs that port was successfully bound and got a specified connection id.
type AcceptedParams struct {
	
	// Port number that was successfully bound.
	Port	int	`json:"port"`
	
	// Connection id to be used.
	ConnectionId	string	`json:"connectionId"`
	
}

