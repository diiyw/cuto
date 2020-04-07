package cast
// 
type Sink  struct {

	// 
	Name	string	`json:"name"`

	// 
	Id	string	`json:"id"`

	// Text describing the current session. Present only if there is an active
	// session on the sink.
	Session	string	`json:"session"`
}
