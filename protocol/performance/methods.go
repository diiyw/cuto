package performance

const (
	
	// Disable collecting and reporting metrics.
	Disable = "Performance.disable"
	
	// Enable collecting and reporting metrics.
	Enable = "Performance.enable"
	
	// Sets time domain to use for collecting and reporting duration metrics.
	// Note that this must be called before enabling metrics collection. Calling
	// this method while metrics collection is enabled returns an error.
	SetTimeDomain = "Performance.setTimeDomain"
	
	// Retrieve current values of run-time metrics.
	GetMetrics = "Performance.getMetrics"
	
)

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

// SetTimeDomain parameters
type SetTimeDomainParams struct {
	
	// Time domain
	TimeDomain	string	`json:"timeDomain"`
	
}

// SetTimeDomain returns
type SetTimeDomainReturns struct {
	
}

// GetMetrics parameters
type GetMetricsParams struct {
	
}

// GetMetrics returns
type GetMetricsReturns struct {
	
	// Current values for run-time metrics.
	Metrics	[]Metric	`json:"metrics"`
	
}

