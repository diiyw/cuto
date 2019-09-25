package media

const (
	
	// This can be called multiple times, and can be used to set / override /
	// remove player properties. A null propValue indicates removal.
	PlayerPropertiesChangedEvent = "Media.playerPropertiesChanged"
	
	// Send events as a list, allowing them to be batched on the browser for less
	// congestion. If batched, events must ALWAYS be in chronological order.
	PlayerEventsAddedEvent = "Media.playerEventsAdded"
	
	// Called whenever a player is created, or when a new agent joins and recieves
	// a list of active players. If an agent is restored, it will recieve the full
	// list of player ids and all events again.
	PlayersCreatedEvent = "Media.playersCreated"
	
)

// This can be called multiple times, and can be used to set / override /
	// remove player properties. A null propValue indicates removal.
type PlayerPropertiesChangedParams struct {
	
	
	PlayerId	PlayerId	`json:"playerId"`
	
	
	Properties	[]PlayerProperty	`json:"properties"`
	
}

// Send events as a list, allowing them to be batched on the browser for less
	// congestion. If batched, events must ALWAYS be in chronological order.
type PlayerEventsAddedParams struct {
	
	
	PlayerId	PlayerId	`json:"playerId"`
	
	
	Events	[]PlayerEvent	`json:"events"`
	
}

// Called whenever a player is created, or when a new agent joins and recieves
	// a list of active players. If an agent is restored, it will recieve the full
	// list of player ids and all events again.
type PlayersCreatedParams struct {
	
	
	Players	[]PlayerId	`json:"players"`
	
}

