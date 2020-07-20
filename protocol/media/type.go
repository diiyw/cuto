package media
// Players will get an ID that is unique within the agent context.
type PlayerId string

// 
type Timestamp float64

// Player Property type
type PlayerProperty  struct {

	// 
	Name	string	`json:"name"`

	// 
	Value	string	`json:"value,omitempty"`
}

// Break out events into different types
type PlayerEventType string

// 
type PlayerEvent  struct {

	// 
	Type	PlayerEventType	`json:"type"`

	// Events are timestamped relative to the start of the player creation
	// not relative to the start of playback.
	Timestamp	Timestamp	`json:"timestamp"`

	// 
	Name	string	`json:"name"`

	// 
	Value	string	`json:"value"`
}
