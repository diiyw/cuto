package security


// An internal certificate ID value.
type CertificateId int	

// A description of mixed content (HTTP resources on HTTPS pages), as defined by
	// https://www.w3.org/TR/mixed-content/#categories
type MixedContentType string	

// The security level of a page or resource.
type SecurityState string	

// An explanation of an factor contributing to the security state.
type SecurityStateExplanation struct {
	
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
type InsecureContentStatus struct {
	
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

