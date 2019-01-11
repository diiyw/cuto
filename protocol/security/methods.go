package security

const (
	
	// Disables tracking security state changes.
	Disable = "Security.disable"
	
	// Enables tracking security state changes.
	Enable = "Security.enable"
	
	// Enable/disable whether all certificate errors should be ignored.
	SetIgnoreCertificateErrors = "Security.setIgnoreCertificateErrors"
	
	// Handles a certificate error that fired a certificateError event.
	HandleCertificateError = "Security.handleCertificateError"
	
	// Enable/disable overriding certificate errors. If enabled, all certificate error events need to
	// be handled by the DevTools client and should be answered with `handleCertificateError` commands.
	SetOverrideCertificateErrors = "Security.setOverrideCertificateErrors"
	
)

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// SetIgnoreCertificateErrors parameters
type SetIgnoreCertificateErrorsParams struct {
	
	// If true, all certificate errors will be ignored.
	Ignore	bool	`json:"ignore"`
	
}

// SetIgnoreCertificateErrors returns
type SetIgnoreCertificateErrorsReturns struct {
	
}

// HandleCertificateError parameters
type HandleCertificateErrorParams struct {
	
	// The ID of the event.
	EventId	int	`json:"eventId"`
	
	// The action to take on the certificate error.
	Action	CertificateErrorAction	`json:"action"`
	
}

// HandleCertificateError returns
type HandleCertificateErrorReturns struct {
	
}

// SetOverrideCertificateErrors parameters
type SetOverrideCertificateErrorsParams struct {
	
	// If true, certificate errors will be overridden.
	Override	bool	`json:"override"`
	
}

// SetOverrideCertificateErrors returns
type SetOverrideCertificateErrorsReturns struct {
	
}

