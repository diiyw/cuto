package console


// Console message.
type ConsoleMessage struct {
	
	// Message source.
	// Possible value: xml,javascript,network,console-api,storage,appcache,rendering,security,other,deprecation,worker,
	Source	string	`json:"source"`
	
	// Message severity.
	// Possible value: log,warning,error,debug,info,
	Level	string	`json:"level"`
	
	// Message text.
	
	Text	string	`json:"text"`
	
	// URL of the message origin.
	
	Url	string	`json:"url"`
	
	// Line number in the resource that generated this message (1-based).
	
	Line	int	`json:"line"`
	
	// Column number in the resource that generated this message (1-based).
	
	Column	int	`json:"column"`
	
}	

