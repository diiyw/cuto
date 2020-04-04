package emulation

// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
const VirtualTimeBudgetExpiredEvent = "Emulation.virtualTimeBudgetExpired"
type VirtualTimeBudgetExpiredParams struct {
}

