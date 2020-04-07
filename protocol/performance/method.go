package performance

// Disable collecting and reporting metrics.
const Disable = "Performance.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enable collecting and reporting metrics.
const Enable = "Performance.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Sets time domain to use for collecting and reporting duration metrics.
// Note that this must be called before enabling metrics collection. Calling
// this method while metrics collection is enabled returns an error.
const SetTimeDomain = "Performance.setTimeDomain"

type SetTimeDomainParams struct {

	// Time domain
	TimeDomain 	string	`json:"timeDomain"`
}

type SetTimeDomainResult struct {

}

// Retrieve current values of run-time metrics.
const GetMetrics = "Performance.getMetrics"

type GetMetricsParams struct {
}

type GetMetricsResult struct {

	// Current values for run-time metrics.
	Metrics 	[]*Metric	`json:"metrics"`
}