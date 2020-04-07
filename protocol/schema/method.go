package schema

// Returns supported domains.
const GetDomains = "Schema.getDomains"

type GetDomainsParams struct {
}

type GetDomainsResult struct {

	// List of supported domains.
	Domains 	[]*Domain	`json:"domains"`
}