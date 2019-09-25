package memory

const (
	
	
	GetDOMCounters = "Memory.getDOMCounters"
	
	
	PrepareForLeakDetection = "Memory.prepareForLeakDetection"
	
	// Simulate OomIntervention by purging V8 memory.
	ForciblyPurgeJavaScriptMemory = "Memory.forciblyPurgeJavaScriptMemory"
	
	// Enable/disable suppressing memory pressure notifications in all processes.
	SetPressureNotificationsSuppressed = "Memory.setPressureNotificationsSuppressed"
	
	// Simulate a memory pressure notification in all processes.
	SimulatePressureNotification = "Memory.simulatePressureNotification"
	
	// Start collecting native memory profile.
	StartSampling = "Memory.startSampling"
	
	// Stop collecting native memory profile.
	StopSampling = "Memory.stopSampling"
	
	// Retrieve native memory allocations profile
	// collected since renderer process startup.
	GetAllTimeSamplingProfile = "Memory.getAllTimeSamplingProfile"
	
	// Retrieve native memory allocations profile
	// collected since browser process startup.
	GetBrowserSamplingProfile = "Memory.getBrowserSamplingProfile"
	
	// Retrieve native memory allocations profile collected since last
	// `startSampling` call.
	GetSamplingProfile = "Memory.getSamplingProfile"
	
)

// GetDOMCounters parameters
type GetDOMCountersParams struct {
	
}

// GetDOMCounters returns
type GetDOMCountersReturns struct {
	
	
	Documents	int	`json:"documents"`
	
	
	Nodes	int	`json:"nodes"`
	
	
	JsEventListeners	int	`json:"jsEventListeners"`
	
}

// PrepareForLeakDetection parameters
type PrepareForLeakDetectionParams struct {
	
}

// PrepareForLeakDetection returns
type PrepareForLeakDetectionReturns struct {
	
}

// ForciblyPurgeJavaScriptMemory parameters
type ForciblyPurgeJavaScriptMemoryParams struct {
	
}

// ForciblyPurgeJavaScriptMemory returns
type ForciblyPurgeJavaScriptMemoryReturns struct {
	
}

// SetPressureNotificationsSuppressed parameters
type SetPressureNotificationsSuppressedParams struct {
	
	// If true, memory pressure notifications will be suppressed.
	Suppressed	bool	`json:"suppressed"`
	
}

// SetPressureNotificationsSuppressed returns
type SetPressureNotificationsSuppressedReturns struct {
	
}

// SimulatePressureNotification parameters
type SimulatePressureNotificationParams struct {
	
	// Memory pressure level of the notification.
	Level	PressureLevel	`json:"level"`
	
}

// SimulatePressureNotification returns
type SimulatePressureNotificationReturns struct {
	
}

// StartSampling parameters
type StartSamplingParams struct {
	
	// Average number of bytes between samples.
	SamplingInterval	int	`json:"samplingInterval"`
	
	// Do not randomize intervals between samples.
	SuppressRandomness	bool	`json:"suppressRandomness"`
	
}

// StartSampling returns
type StartSamplingReturns struct {
	
}

// StopSampling parameters
type StopSamplingParams struct {
	
}

// StopSampling returns
type StopSamplingReturns struct {
	
}

// GetAllTimeSamplingProfile parameters
type GetAllTimeSamplingProfileParams struct {
	
}

// GetAllTimeSamplingProfile returns
type GetAllTimeSamplingProfileReturns struct {
	
	
	Profile	SamplingProfile	`json:"profile"`
	
}

// GetBrowserSamplingProfile parameters
type GetBrowserSamplingProfileParams struct {
	
}

// GetBrowserSamplingProfile returns
type GetBrowserSamplingProfileReturns struct {
	
	
	Profile	SamplingProfile	`json:"profile"`
	
}

// GetSamplingProfile parameters
type GetSamplingProfileParams struct {
	
}

// GetSamplingProfile returns
type GetSamplingProfileReturns struct {
	
	
	Profile	SamplingProfile	`json:"profile"`
	
}

