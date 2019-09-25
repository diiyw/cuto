package log

const (
	
	// Clears the log.
	Clear = "Log.clear"
	
	// Disables log domain, prevents further log entries from being reported to the client.
	Disable = "Log.disable"
	
	// Enables log domain, sends the entries collected so far to the client by means of the
	// `entryAdded` notification.
	Enable = "Log.enable"
	
	// start violation reporting.
	StartViolationsReport = "Log.startViolationsReport"
	
	// Stop violation reporting.
	StopViolationsReport = "Log.stopViolationsReport"
	
)

// Clear parameters
type ClearParams struct {
	
}

// Clear returns
type ClearReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// StartViolationsReport parameters
type StartViolationsReportParams struct {
	
	// Configuration for violations.
	Config	[]ViolationSetting	`json:"config"`
	
}

// StartViolationsReport returns
type StartViolationsReportReturns struct {
	
}

// StopViolationsReport parameters
type StopViolationsReportParams struct {
	
}

// StopViolationsReport returns
type StopViolationsReportReturns struct {
	
}

