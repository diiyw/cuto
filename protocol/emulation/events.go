package emulation

const (
	
	// Notification sent after the virtual time has advanced.
	VirtualTimeAdvancedEvent = "Emulation.virtualTimeAdvanced"
	
	// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
	VirtualTimeBudgetExpiredEvent = "Emulation.virtualTimeBudgetExpired"
	
	// Notification sent after the virtual time has paused.
	VirtualTimePausedEvent = "Emulation.virtualTimePaused"
	
)

// Notification sent after the virtual time has advanced.
type VirtualTimeAdvancedParams struct {
	
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first
	// enabled.
	VirtualTimeElapsed	float64	`json:"virtualTimeElapsed"`
	
}

// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
type VirtualTimeBudgetExpiredParams struct {
	
}

// Notification sent after the virtual time has paused.
type VirtualTimePausedParams struct {
	
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first
	// enabled.
	VirtualTimeElapsed	float64	`json:"virtualTimeElapsed"`
	
}

