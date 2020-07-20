package profiler

// 
const Disable = "Profiler.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// 
const Enable = "Profiler.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Collect coverage data for the current isolate. The coverage data may be incomplete due to
// garbage collection.
const GetBestEffortCoverage = "Profiler.getBestEffortCoverage"

type GetBestEffortCoverageParams struct {
}

type GetBestEffortCoverageResult struct {

	// Coverage data for the current isolate.
	Result 	[]*ScriptCoverage	`json:"result"`
}

// Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
const SetSamplingInterval = "Profiler.setSamplingInterval"

type SetSamplingIntervalParams struct {

	// New sampling interval in microseconds.
	Interval 	int	`json:"interval"`
}

type SetSamplingIntervalResult struct {

}

// 
const Start = "Profiler.start"

type StartParams struct {
}

type StartResult struct {

}

// Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code
// coverage may be incomplete. Enabling prevents running optimized code and resets execution
// counters.
const StartPreciseCoverage = "Profiler.startPreciseCoverage"

type StartPreciseCoverageParams struct {

	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount 	bool	`json:"callCount,omitempty"`

	// Collect block-based coverage.
	Detailed 	bool	`json:"detailed,omitempty"`
}

type StartPreciseCoverageResult struct {

}

// Enable type profile.
const StartTypeProfile = "Profiler.startTypeProfile"

type StartTypeProfileParams struct {
}

type StartTypeProfileResult struct {

}

// 
const Stop = "Profiler.stop"

type StopParams struct {
}

type StopResult struct {

	// Recorded profile.
	Profile 	Profile	`json:"profile"`
}

// Disable precise code coverage. Disabling releases unnecessary execution count records and allows
// executing optimized code.
const StopPreciseCoverage = "Profiler.stopPreciseCoverage"

type StopPreciseCoverageParams struct {
}

type StopPreciseCoverageResult struct {

}

// Disable type profile. Disabling releases type profile data collected so far.
const StopTypeProfile = "Profiler.stopTypeProfile"

type StopTypeProfileParams struct {
}

type StopTypeProfileResult struct {

}

// Collect coverage data for the current isolate, and resets execution counters. Precise code
// coverage needs to have started.
const TakePreciseCoverage = "Profiler.takePreciseCoverage"

type TakePreciseCoverageParams struct {
}

type TakePreciseCoverageResult struct {

	// Coverage data for the current isolate.
	Result 	[]*ScriptCoverage	`json:"result"`
}

// Collect type profile.
const TakeTypeProfile = "Profiler.takeTypeProfile"

type TakeTypeProfileParams struct {
}

type TakeTypeProfileResult struct {

	// Type profile for all scripts since startTypeProfile() was turned on.
	Result 	[]*ScriptTypeProfile	`json:"result"`
}