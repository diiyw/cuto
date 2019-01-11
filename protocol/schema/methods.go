package schema

const (
	
	// Returns supported domains.
	GetDomains = "Schema.getDomains"
	
)

// GetDomains parameters
type GetDomainsParams struct {
	
}

// GetDomains returns
type GetDomainsReturns struct {
	
	// List of supported domains.
	Domains	[]Domain	`json:"domains"`
	
}

