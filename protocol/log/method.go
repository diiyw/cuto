package log

// Clears the log.
const Clear = "Log.clear"

type ClearParams struct {
}

type ClearResult struct {

}

// Disables log domain, prevents further log entries from being reported to the client.
const Disable = "Log.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables log domain, sends the entries collected so far to the client by means of the
// `entryAdded` notification.
const Enable = "Log.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// start violation reporting.
const StartViolationsReport = "Log.startViolationsReport"

type StartViolationsReportParams struct {

	// Configuration for violations.
	Config 	[]*ViolationSetting	`json:"config"`
}

type StartViolationsReportResult struct {

}

// Stop violation reporting.
const StopViolationsReport = "Log.stopViolationsReport"

type StopViolationsReportParams struct {
}

type StopViolationsReportResult struct {

}