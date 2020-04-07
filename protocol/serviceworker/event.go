package serviceworker

// 
const WorkerErrorReportedEvent = "ServiceWorker.workerErrorReported"
type WorkerErrorReportedParams struct {

	// 
	ErrorMessage 	ServiceWorkerErrorMessage}



// 
const WorkerRegistrationUpdatedEvent = "ServiceWorker.workerRegistrationUpdated"
type WorkerRegistrationUpdatedParams struct {

	// 
	Registrations 	[]*ServiceWorkerRegistration}



// 
const WorkerVersionUpdatedEvent = "ServiceWorker.workerVersionUpdated"
type WorkerVersionUpdatedParams struct {

	// 
	Versions 	[]*ServiceWorkerVersion}

