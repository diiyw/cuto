package inspector

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
const DetachedEvent = "Inspector.detached"
type DetachedParams struct {

	// The reason why connection has been terminated.
	Reason 	string}



// Fired when debugging target has crashed
const TargetCrashedEvent = "Inspector.targetCrashed"
type TargetCrashedParams struct {
}



// Fired when debugging target has reloaded after crash
const TargetReloadedAfterCrashEvent = "Inspector.targetReloadedAfterCrash"
type TargetReloadedAfterCrashParams struct {
}

