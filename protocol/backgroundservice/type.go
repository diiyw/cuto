package backgroundservice

import (
	"github.com/diiyw/cuto/protocol/cdp"
	"github.com/diiyw/cuto/protocol/serviceworker"
)

// The Background Service that will be associated with the commands/events.
	// Every Background Service operates independently, but they share the same
	// API.
type ServiceName string

// A key-value pair for additional event information to pass along.
type EventMetadata  struct {

	// 
	Key	string	`json:"key"`

	// 
	Value	string	`json:"value"`
}

// 
type BackgroundServiceEvent  struct {

	// Timestamp of the event (in seconds).
	Timestamp	cdp.TimeSinceEpoch	`json:"timestamp"`

	// The origin this event belongs to.
	Origin	string	`json:"origin"`

	// The Service Worker ID that initiated the event.
	ServiceWorkerRegistrationId	serviceworker.RegistrationID	`json:"serviceWorkerRegistrationId"`

	// The Background Service this event belongs to.
	Service	ServiceName	`json:"service"`

	// A description of the event.
	EventName	string	`json:"eventName"`

	// An identifier that groups related events together.
	InstanceId	string	`json:"instanceId"`

	// A list of event-specific information.
	EventMetadata	[]*EventMetadata	`json:"eventMetadata"`
}
