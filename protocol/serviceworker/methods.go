package serviceworker

const (
	
	
	DeliverPushMessage = "ServiceWorker.deliverPushMessage"
	
	
	Disable = "ServiceWorker.disable"
	
	
	DispatchSyncEvent = "ServiceWorker.dispatchSyncEvent"
	
	
	DispatchPeriodicSyncEvent = "ServiceWorker.dispatchPeriodicSyncEvent"
	
	
	Enable = "ServiceWorker.enable"
	
	
	InspectWorker = "ServiceWorker.inspectWorker"
	
	
	SetForceUpdateOnPageLoad = "ServiceWorker.setForceUpdateOnPageLoad"
	
	
	SkipWaiting = "ServiceWorker.skipWaiting"
	
	
	StartWorker = "ServiceWorker.startWorker"
	
	
	StopAllWorkers = "ServiceWorker.stopAllWorkers"
	
	
	StopWorker = "ServiceWorker.stopWorker"
	
	
	Unregister = "ServiceWorker.unregister"
	
	
	UpdateRegistration = "ServiceWorker.updateRegistration"
	
)

// DeliverPushMessage parameters
type DeliverPushMessageParams struct {
	
	
	Origin	string	`json:"origin"`
	
	
	RegistrationId	RegistrationID	`json:"registrationId"`
	
	
	Data	string	`json:"data"`
	
}

// DeliverPushMessage returns
type DeliverPushMessageReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// DispatchSyncEvent parameters
type DispatchSyncEventParams struct {
	
	
	Origin	string	`json:"origin"`
	
	
	RegistrationId	RegistrationID	`json:"registrationId"`
	
	
	Tag	string	`json:"tag"`
	
	
	LastChance	bool	`json:"lastChance"`
	
}

// DispatchSyncEvent returns
type DispatchSyncEventReturns struct {
	
}

// DispatchPeriodicSyncEvent parameters
type DispatchPeriodicSyncEventParams struct {
	
	
	Origin	string	`json:"origin"`
	
	
	RegistrationId	RegistrationID	`json:"registrationId"`
	
	
	Tag	string	`json:"tag"`
	
}

// DispatchPeriodicSyncEvent returns
type DispatchPeriodicSyncEventReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// InspectWorker parameters
type InspectWorkerParams struct {
	
	
	VersionId	string	`json:"versionId"`
	
}

// InspectWorker returns
type InspectWorkerReturns struct {
	
}

// SetForceUpdateOnPageLoad parameters
type SetForceUpdateOnPageLoadParams struct {
	
	
	ForceUpdateOnPageLoad	bool	`json:"forceUpdateOnPageLoad"`
	
}

// SetForceUpdateOnPageLoad returns
type SetForceUpdateOnPageLoadReturns struct {
	
}

// SkipWaiting parameters
type SkipWaitingParams struct {
	
	
	ScopeURL	string	`json:"scopeURL"`
	
}

// SkipWaiting returns
type SkipWaitingReturns struct {
	
}

// StartWorker parameters
type StartWorkerParams struct {
	
	
	ScopeURL	string	`json:"scopeURL"`
	
}

// StartWorker returns
type StartWorkerReturns struct {
	
}

// StopAllWorkers parameters
type StopAllWorkersParams struct {
	
}

// StopAllWorkers returns
type StopAllWorkersReturns struct {
	
}

// StopWorker parameters
type StopWorkerParams struct {
	
	
	VersionId	string	`json:"versionId"`
	
}

// StopWorker returns
type StopWorkerReturns struct {
	
}

// Unregister parameters
type UnregisterParams struct {
	
	
	ScopeURL	string	`json:"scopeURL"`
	
}

// Unregister returns
type UnregisterReturns struct {
	
}

// UpdateRegistration parameters
type UpdateRegistrationParams struct {
	
	
	ScopeURL	string	`json:"scopeURL"`
	
}

// UpdateRegistration returns
type UpdateRegistrationReturns struct {
	
}

