package memory

// 
const GetDOMCounters = "Memory.getDOMCounters"

type GetDOMCountersParams struct {
}

type GetDOMCountersResult struct {

	// 
	Documents 	int	`json:"documents"`
	// 
	Nodes 	int	`json:"nodes"`
	// 
	JsEventListeners 	int	`json:"jsEventListeners"`
}

// 
const PrepareForLeakDetection = "Memory.prepareForLeakDetection"

type PrepareForLeakDetectionParams struct {
}

type PrepareForLeakDetectionResult struct {

}

// Simulate OomIntervention by purging V8 memory.
const ForciblyPurgeJavaScriptMemory = "Memory.forciblyPurgeJavaScriptMemory"

type ForciblyPurgeJavaScriptMemoryParams struct {
}

type ForciblyPurgeJavaScriptMemoryResult struct {

}

// Enable/disable suppressing memory pressure notifications in all processes.
const SetPressureNotificationsSuppressed = "Memory.setPressureNotificationsSuppressed"

type SetPressureNotificationsSuppressedParams struct {

	// If true, memory pressure notifications will be suppressed.
	Suppressed 	bool	`json:"suppressed"`
}

type SetPressureNotificationsSuppressedResult struct {

}

// Simulate a memory pressure notification in all processes.
const SimulatePressureNotification = "Memory.simulatePressureNotification"

type SimulatePressureNotificationParams struct {

	// Memory pressure level of the notification.
	Level 	PressureLevel	`json:"level"`
}

type SimulatePressureNotificationResult struct {

}

// Start collecting native memory profile.
const StartSampling = "Memory.startSampling"

type StartSamplingParams struct {

	// Average number of bytes between samples.
	SamplingInterval 	int	`json:"samplingInterval,omitempty"`

	// Do not randomize intervals between samples.
	SuppressRandomness 	bool	`json:"suppressRandomness,omitempty"`
}

type StartSamplingResult struct {

}

// Stop collecting native memory profile.
const StopSampling = "Memory.stopSampling"

type StopSamplingParams struct {
}

type StopSamplingResult struct {

}

// Retrieve native memory allocations profile
// collected since renderer process startup.
const GetAllTimeSamplingProfile = "Memory.getAllTimeSamplingProfile"

type GetAllTimeSamplingProfileParams struct {
}

type GetAllTimeSamplingProfileResult struct {

	// 
	Profile 	SamplingProfile	`json:"profile"`
}

// Retrieve native memory allocations profile
// collected since browser process startup.
const GetBrowserSamplingProfile = "Memory.getBrowserSamplingProfile"

type GetBrowserSamplingProfileParams struct {
}

type GetBrowserSamplingProfileResult struct {

	// 
	Profile 	SamplingProfile	`json:"profile"`
}

// Retrieve native memory allocations profile collected since last
// `startSampling` call.
const GetSamplingProfile = "Memory.getSamplingProfile"

type GetSamplingProfileParams struct {
}

type GetSamplingProfileResult struct {

	// 
	Profile 	SamplingProfile	`json:"profile"`
}