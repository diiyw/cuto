package serviceworker

// 
const DeliverPushMessage = "ServiceWorker.deliverPushMessage"

type DeliverPushMessageParams struct {

	// 
	Origin 	string	`json:"origin"`

	// 
	RegistrationId 	RegistrationID	`json:"registrationId"`

	// 
	Data 	string	`json:"data"`
}

type DeliverPushMessageResult struct {

}

// 
const Disable = "ServiceWorker.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// 
const DispatchSyncEvent = "ServiceWorker.dispatchSyncEvent"

type DispatchSyncEventParams struct {

	// 
	Origin 	string	`json:"origin"`

	// 
	RegistrationId 	RegistrationID	`json:"registrationId"`

	// 
	Tag 	string	`json:"tag"`

	// 
	LastChance 	bool	`json:"lastChance"`
}

type DispatchSyncEventResult struct {

}

// 
const DispatchPeriodicSyncEvent = "ServiceWorker.dispatchPeriodicSyncEvent"

type DispatchPeriodicSyncEventParams struct {

	// 
	Origin 	string	`json:"origin"`

	// 
	RegistrationId 	RegistrationID	`json:"registrationId"`

	// 
	Tag 	string	`json:"tag"`
}

type DispatchPeriodicSyncEventResult struct {

}

// 
const Enable = "ServiceWorker.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// 
const InspectWorker = "ServiceWorker.inspectWorker"

type InspectWorkerParams struct {

	// 
	VersionId 	string	`json:"versionId"`
}

type InspectWorkerResult struct {

}

// 
const SetForceUpdateOnPageLoad = "ServiceWorker.setForceUpdateOnPageLoad"

type SetForceUpdateOnPageLoadParams struct {

	// 
	ForceUpdateOnPageLoad 	bool	`json:"forceUpdateOnPageLoad"`
}

type SetForceUpdateOnPageLoadResult struct {

}

// 
const SkipWaiting = "ServiceWorker.skipWaiting"

type SkipWaitingParams struct {

	// 
	ScopeURL 	string	`json:"scopeURL"`
}

type SkipWaitingResult struct {

}

// 
const StartWorker = "ServiceWorker.startWorker"

type StartWorkerParams struct {

	// 
	ScopeURL 	string	`json:"scopeURL"`
}

type StartWorkerResult struct {

}

// 
const StopAllWorkers = "ServiceWorker.stopAllWorkers"

type StopAllWorkersParams struct {
}

type StopAllWorkersResult struct {

}

// 
const StopWorker = "ServiceWorker.stopWorker"

type StopWorkerParams struct {

	// 
	VersionId 	string	`json:"versionId"`
}

type StopWorkerResult struct {

}

// 
const Unregister = "ServiceWorker.unregister"

type UnregisterParams struct {

	// 
	ScopeURL 	string	`json:"scopeURL"`
}

type UnregisterResult struct {

}

// 
const UpdateRegistration = "ServiceWorker.updateRegistration"

type UpdateRegistrationParams struct {

	// 
	ScopeURL 	string	`json:"scopeURL"`
}

type UpdateRegistrationResult struct {

}