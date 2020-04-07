package media

// This can be called multiple times, and can be used to set / override /
// remove player properties. A null propValue indicates removal.
const PlayerPropertiesChangedEvent = "Media.playerPropertiesChanged"
type PlayerPropertiesChangedParams struct {

	// 
	PlayerId 	PlayerId
	// 
	Properties 	[]*PlayerProperty}



// Send events as a list, allowing them to be batched on the browser for less
// congestion. If batched, events must ALWAYS be in chronological order.
const PlayerEventsAddedEvent = "Media.playerEventsAdded"
type PlayerEventsAddedParams struct {

	// 
	PlayerId 	PlayerId
	// 
	Events 	[]*PlayerEvent}



// Called whenever a player is created, or when a new agent joins and recieves
// a list of active players. If an agent is restored, it will recieve the full
// list of player ids and all events again.
const PlayersCreatedEvent = "Media.playersCreated"
type PlayersCreatedParams struct {

	// 
	Players 	[]*PlayerId}

