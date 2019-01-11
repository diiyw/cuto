package serviceworker

const (
	
	
	WorkerErrorReportedEvent = "ServiceWorker.workerErrorReported"
	
	
	WorkerRegistrationUpdatedEvent = "ServiceWorker.workerRegistrationUpdated"
	
	
	WorkerVersionUpdatedEvent = "ServiceWorker.workerVersionUpdated"
	
)


type WorkerErrorReportedParams struct {
	
	
	ErrorMessage	ServiceWorkerErrorMessage	`json:"errorMessage"`
	
}


type WorkerRegistrationUpdatedParams struct {
	
	
	Registrations	[]ServiceWorkerRegistration	`json:"registrations"`
	
}


type WorkerVersionUpdatedParams struct {
	
	
	Versions	[]ServiceWorkerVersion	`json:"versions"`
	
}

