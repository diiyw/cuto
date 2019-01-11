package profiler

const (
	
	
	Disable = "Profiler.disable"
	
	
	Enable = "Profiler.enable"
	
	// Collect coverage data for the current isolate. The coverage data may be incomplete due to
	// garbage collection.
	GetBestEffortCoverage = "Profiler.getBestEffortCoverage"
	
	// Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
	SetSamplingInterval = "Profiler.setSamplingInterval"
	
	
	Start = "Profiler.start"
	
	// Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code
	// coverage may be incomplete. Enabling prevents running optimized code and resets execution
	// counters.
	StartPreciseCoverage = "Profiler.startPreciseCoverage"
	
	// Enable type profile.
	StartTypeProfile = "Profiler.startTypeProfile"
	
	
	Stop = "Profiler.stop"
	
	// Disable precise code coverage. Disabling releases unnecessary execution count records and allows
	// executing optimized code.
	StopPreciseCoverage = "Profiler.stopPreciseCoverage"
	
	// Disable type profile. Disabling releases type profile data collected so far.
	StopTypeProfile = "Profiler.stopTypeProfile"
	
	// Collect coverage data for the current isolate, and resets execution counters. Precise code
	// coverage needs to have started.
	TakePreciseCoverage = "Profiler.takePreciseCoverage"
	
	// Collect type profile.
	TakeTypeProfile = "Profiler.takeTypeProfile"
	
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

// GetBestEffortCoverage parameters
type GetBestEffortCoverageParams struct {
	
}

// GetBestEffortCoverage returns
type GetBestEffortCoverageReturns struct {
	
	// Coverage data for the current isolate.
	Result	[]ScriptCoverage	`json:"result"`
	
}

// SetSamplingInterval parameters
type SetSamplingIntervalParams struct {
	
	// New sampling interval in microseconds.
	Interval	int	`json:"interval"`
	
}

// SetSamplingInterval returns
type SetSamplingIntervalReturns struct {
	
}

// Start parameters
type StartParams struct {
	
}

// Start returns
type StartReturns struct {
	
}

// StartPreciseCoverage parameters
type StartPreciseCoverageParams struct {
	
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount	bool	`json:"callCount"`
	
	// Collect block-based coverage.
	Detailed	bool	`json:"detailed"`
	
}

// StartPreciseCoverage returns
type StartPreciseCoverageReturns struct {
	
}

// StartTypeProfile parameters
type StartTypeProfileParams struct {
	
}

// StartTypeProfile returns
type StartTypeProfileReturns struct {
	
}

// Stop parameters
type StopParams struct {
	
}

// Stop returns
type StopReturns struct {
	
	// Recorded profile.
	Profile	Profile	`json:"profile"`
	
}

// StopPreciseCoverage parameters
type StopPreciseCoverageParams struct {
	
}

// StopPreciseCoverage returns
type StopPreciseCoverageReturns struct {
	
}

// StopTypeProfile parameters
type StopTypeProfileParams struct {
	
}

// StopTypeProfile returns
type StopTypeProfileReturns struct {
	
}

// TakePreciseCoverage parameters
type TakePreciseCoverageParams struct {
	
}

// TakePreciseCoverage returns
type TakePreciseCoverageReturns struct {
	
	// Coverage data for the current isolate.
	Result	[]ScriptCoverage	`json:"result"`
	
}

// TakeTypeProfile parameters
type TakeTypeProfileParams struct {
	
}

// TakeTypeProfile returns
type TakeTypeProfileReturns struct {
	
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result	[]ScriptTypeProfile	`json:"result"`
	
}

