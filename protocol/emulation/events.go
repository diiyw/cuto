package emulation

const (
	
	// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
	VirtualTimeBudgetExpiredEvent = "Emulation.virtualTimeBudgetExpired"
	
)

// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
type VirtualTimeBudgetExpiredParams struct {
	
}

