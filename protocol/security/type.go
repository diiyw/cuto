package security
// An internal certificate ID value.
type CertificateId int

// A description of mixed content (HTTP resources on HTTPS pages), as defined by
	// https://www.w3.org/TR/mixed-content/#categories
type MixedContentType string

// The security level of a page or resource.
type SecurityState string

// Details about the security state of the page certificate.
type CertificateSecurityState  struct {

	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol	string	`json:"protocol"`

	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange	string	`json:"keyExchange"`

	// (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup	string	`json:"keyExchangeGroup"`

	// Cipher name.
	Cipher	string	`json:"cipher"`

	// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac	string	`json:"mac"`

	// Page certificate.
	Certificate	[]string	`json:"certificate"`

	// Certificate subject name.
	SubjectName	string	`json:"subjectName"`

	// Name of the issuing CA.
	Issuer	string	`json:"issuer"`

	// Certificate valid from date.
	ValidFrom	interface{}	`json:"validFrom"`

	// Certificate valid to (expiration) date
	ValidTo	interface{}	`json:"validTo"`

	// True if the certificate uses a weak signature aglorithm.
	CertifcateHasWeakSignature	bool	`json:"certifcateHasWeakSignature"`

	// True if modern SSL
	ModernSSL	bool	`json:"modernSSL"`

	// True if the connection is using an obsolete SSL protocol.
	ObsoleteSslProtocol	bool	`json:"obsoleteSslProtocol"`

	// True if the connection is using an obsolete SSL key exchange.
	ObsoleteSslKeyExchange	bool	`json:"obsoleteSslKeyExchange"`

	// True if the connection is using an obsolete SSL cipher.
	ObsoleteSslCipher	bool	`json:"obsoleteSslCipher"`

	// True if the connection is using an obsolete SSL signature.
	ObsoleteSslSignature	bool	`json:"obsoleteSslSignature"`
}

// Security state information about the page.
type VisibleSecurityState  struct {

	// The security level of the page.
	SecurityState	SecurityState	`json:"securityState"`

	// Security state details about the page certificate.
	CertificateSecurityState	CertificateSecurityState	`json:"certificateSecurityState"`

	// Array of security state issues ids.
	SecurityStateIssueIds	[]string	`json:"securityStateIssueIds"`
}

// An explanation of an factor contributing to the security state.
type SecurityStateExplanation  struct {

	// Security state representing the severity of the factor being explained.
	SecurityState	SecurityState	`json:"securityState"`

	// Title describing the type of factor.
	Title	string	`json:"title"`

	// Short phrase describing the type of factor.
	Summary	string	`json:"summary"`

	// Full text explanation of the factor.
	Description	string	`json:"description"`

	// The type of mixed content described by the explanation.
	MixedContentType	MixedContentType	`json:"mixedContentType"`

	// Page certificate.
	Certificate	[]string	`json:"certificate"`

	// Recommendations to fix any issues.
	Recommendations	[]string	`json:"recommendations"`
}

// Information about insecure content on the page.
type InsecureContentStatus  struct {

	// Always false.
	RanMixedContent	bool	`json:"ranMixedContent"`

	// Always false.
	DisplayedMixedContent	bool	`json:"displayedMixedContent"`

	// Always false.
	ContainedMixedForm	bool	`json:"containedMixedForm"`

	// Always false.
	RanContentWithCertErrors	bool	`json:"ranContentWithCertErrors"`

	// Always false.
	DisplayedContentWithCertErrors	bool	`json:"displayedContentWithCertErrors"`

	// Always set to unknown.
	RanInsecureContentStyle	SecurityState	`json:"ranInsecureContentStyle"`

	// Always set to unknown.
	DisplayedInsecureContentStyle	SecurityState	`json:"displayedInsecureContentStyle"`
}

// The action to take when a certificate error occurs. continue will continue processing the
	// request and cancel will cancel the request.
type CertificateErrorAction string
