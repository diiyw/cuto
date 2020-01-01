package testing

const (
	
	// Generates a report for testing.
	GenerateTestReport = "Testing.generateTestReport"
	
)

// GenerateTestReport parameters
type GenerateTestReportParams struct {
	
	// Message to be displayed in the report.
	Message	string	`json:"message"`
	
	// Specifies the endpoint group to deliver the report to.
	Group	string	`json:"group"`
	
}

// GenerateTestReport returns
type GenerateTestReportReturns struct {
	
}

