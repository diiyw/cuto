package target


// 
type TargetID string	

// Unique identifier of attached debugging session.
type SessionID string	

// 
type BrowserContextID string	

// 
type TargetInfo struct {
	
	
	
	TargetId	TargetID	`json:"targetId"`
	
	
	
	Type	string	`json:"type"`
	
	
	
	Title	string	`json:"title"`
	
	
	
	Url	string	`json:"url"`
	
	// Whether the target has an attached client.
	
	Attached	bool	`json:"attached"`
	
	// Opener target Id
	
	OpenerId	TargetID	`json:"openerId"`
	
	
	
	BrowserContextId	BrowserContextID	`json:"browserContextId"`
	
}	

// 
type RemoteLocation struct {
	
	
	
	Host	string	`json:"host"`
	
	
	
	Port	int	`json:"port"`
	
}	

