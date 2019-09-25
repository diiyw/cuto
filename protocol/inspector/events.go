package inspector

const (
	
	// Fired when remote debugging connection is about to be terminated. Contains detach reason.
	DetachedEvent = "Inspector.detached"
	
	// Fired when debugging target has crashed
	TargetCrashedEvent = "Inspector.targetCrashed"
	
	// Fired when debugging target has reloaded after crash
	TargetReloadedAfterCrashEvent = "Inspector.targetReloadedAfterCrash"
	
)

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
type DetachedParams struct {
	
	// The reason why connection has been terminated.
	Reason	string	`json:"reason"`
	
}

// Fired when debugging target has crashed
type TargetCrashedParams struct {
	
}

// Fired when debugging target has reloaded after crash
type TargetReloadedAfterCrashParams struct {
	
}

