package security

// Disables tracking security state changes.
const Disable = "Security.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables tracking security state changes.
const Enable = "Security.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Enable/disable whether all certificate errors should be ignored.
const SetIgnoreCertificateErrors = "Security.setIgnoreCertificateErrors"

type SetIgnoreCertificateErrorsParams struct {

	// If true, all certificate errors will be ignored.
	Ignore 	bool	`json:"ignore"`
}

type SetIgnoreCertificateErrorsResult struct {

}

// Handles a certificate error that fired a certificateError event.
const HandleCertificateError = "Security.handleCertificateError"

type HandleCertificateErrorParams struct {

	// The ID of the event.
	EventId 	int	`json:"eventId"`

	// The action to take on the certificate error.
	Action 	CertificateErrorAction	`json:"action"`
}

type HandleCertificateErrorResult struct {

}

// Enable/disable overriding certificate errors. If enabled, all certificate error events need to
// be handled by the DevTools client and should be answered with `handleCertificateError` commands.
const SetOverrideCertificateErrors = "Security.setOverrideCertificateErrors"

type SetOverrideCertificateErrorsParams struct {

	// If true, certificate errors will be overridden.
	Override 	bool	`json:"override"`
}

type SetOverrideCertificateErrorsResult struct {

}