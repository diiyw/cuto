package serviceworker

import (

	"github.com/diiyw/gator/protocol/target"

)

// ServiceWorker registration.
type ServiceWorkerRegistration struct {
	
	
	
	RegistrationId	string	`json:"registrationId"`
	
	
	
	ScopeURL	string	`json:"scopeURL"`
	
	
	
	IsDeleted	bool	`json:"isDeleted"`
	
}	

// 
type ServiceWorkerVersionRunningStatus string	

// 
type ServiceWorkerVersionStatus string	

// ServiceWorker version.
type ServiceWorkerVersion struct {
	
	
	
	VersionId	string	`json:"versionId"`
	
	
	
	RegistrationId	string	`json:"registrationId"`
	
	
	
	ScriptURL	string	`json:"scriptURL"`
	
	
	
	RunningStatus	ServiceWorkerVersionRunningStatus	`json:"runningStatus"`
	
	
	
	Status	ServiceWorkerVersionStatus	`json:"status"`
	
	// The Last-Modified header value of the main script.
	
	ScriptLastModified	float64	`json:"scriptLastModified"`
	
	// The time at which the response headers of the main script were received from the server.
	// For cached script it is the last time the cache entry was validated.
	
	ScriptResponseTime	float64	`json:"scriptResponseTime"`
	
	
	
	ControlledClients	[]target.TargetID	`json:"controlledClients"`
	
	
	
	TargetId	target.TargetID	`json:"targetId"`
	
}	

// ServiceWorker error message.
type ServiceWorkerErrorMessage struct {
	
	
	
	ErrorMessage	string	`json:"errorMessage"`
	
	
	
	RegistrationId	string	`json:"registrationId"`
	
	
	
	VersionId	string	`json:"versionId"`
	
	
	
	SourceURL	string	`json:"sourceURL"`
	
	
	
	LineNumber	int	`json:"lineNumber"`
	
	
	
	ColumnNumber	int	`json:"columnNumber"`
	
}	

