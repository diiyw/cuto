package security

const (
	
	// There is a certificate error. If overriding certificate errors is enabled, then it should be
	// handled with the `handleCertificateError` command. Note: this event does not fire if the
	// certificate error has been allowed internally. Only one client per target should override
	// certificate errors at the same time.
	CertificateErrorEvent = "Security.certificateError"
	
	// The security state of the page changed.
	SecurityStateChangedEvent = "Security.securityStateChanged"
	
)

// There is a certificate error. If overriding certificate errors is enabled, then it should be
	// handled with the `handleCertificateError` command. Note: this event does not fire if the
	// certificate error has been allowed internally. Only one client per target should override
	// certificate errors at the same time.
type CertificateErrorParams struct {
	
	// The ID of the event.
	EventId	int	`json:"eventId"`
	
	// The type of the error.
	ErrorType	string	`json:"errorType"`
	
	// The url that was requested.
	RequestURL	string	`json:"requestURL"`
	
}

// The security state of the page changed.
type SecurityStateChangedParams struct {
	
	// Security state.
	SecurityState	SecurityState	`json:"securityState"`
	
	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic	bool	`json:"schemeIsCryptographic"`
	
	// List of explanations for the security state. If the overall security state is `insecure` or
	// `warning`, at least one corresponding explanation should be included.
	Explanations	[]SecurityStateExplanation	`json:"explanations"`
	
	// Information about insecure content on the page.
	InsecureContentStatus	InsecureContentStatus	`json:"insecureContentStatus"`
	
	// Overrides user-visible description of the state.
	Summary	string	`json:"summary"`
	
}

