package security

// There is a certificate error. If overriding certificate errors is enabled, then it should be
// handled with the `handleCertificateError` command. Note: this event does not fire if the
// certificate error has been allowed internally. Only one client per target should override
// certificate errors at the same time.
const CertificateErrorEvent = "Security.certificateError"
type CertificateErrorParams struct {

	// The ID of the event.
	EventId 	int
	// The type of the error.
	ErrorType 	string
	// The url that was requested.
	RequestURL 	string}



// The security state of the page changed.
const VisibleSecurityStateChangedEvent = "Security.visibleSecurityStateChanged"
type VisibleSecurityStateChangedParams struct {

	// Security state information about the page.
	VisibleSecurityState 	VisibleSecurityState}



// The security state of the page changed.
const SecurityStateChangedEvent = "Security.securityStateChanged"
type SecurityStateChangedParams struct {

	// Security state.
	SecurityState 	SecurityState
	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic 	bool
	// List of explanations for the security state. If the overall security state is `insecure` or
	// `warning`, at least one corresponding explanation should be included.
	Explanations 	[]*SecurityStateExplanation
	// Information about insecure content on the page.
	InsecureContentStatus 	InsecureContentStatus
	// Overrides user-visible description of the state.
	Summary 	string}

