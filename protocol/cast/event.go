package cast

// This is fired whenever the list of available sinks changes. A sink is a
// device or a software surface that you can cast to.
const SinksUpdatedEvent = "Cast.sinksUpdated"
type SinksUpdatedParams struct {

	// 
	Sinks 	[]*Sink}



// This is fired whenever the outstanding issue/error message changes.
// |issueMessage| is empty if there is no issue.
const IssueUpdatedEvent = "Cast.issueUpdated"
type IssueUpdatedParams struct {

	// 
	IssueMessage 	string}

