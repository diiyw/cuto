package cast

const (
	
	// This is fired whenever the list of available sinks changes. A sink is a
	// device or a software surface that you can cast to.
	SinksUpdatedEvent = "Cast.sinksUpdated"
	
	// This is fired whenever the outstanding issue/error message changes.
	// |issueMessage| is empty if there is no issue.
	IssueUpdatedEvent = "Cast.issueUpdated"
	
)

// This is fired whenever the list of available sinks changes. A sink is a
	// device or a software surface that you can cast to.
type SinksUpdatedParams struct {
	
	
	SinkNames	[]string	`json:"sinkNames"`
	
}

// This is fired whenever the outstanding issue/error message changes.
	// |issueMessage| is empty if there is no issue.
type IssueUpdatedParams struct {
	
	
	IssueMessage	string	`json:"issueMessage"`
	
}

