package audits

import (

	"github.com/diiyw/chr/protocol/network"

)
const (
	
	// Returns the response body and size if it were re-encoded with the specified settings. Only
	// applies to images.
	GetEncodedResponse = "Audits.getEncodedResponse"
	
)

// GetEncodedResponse parameters
type GetEncodedResponseParams struct {
	
	// Identifier of the network request to get content for.
	RequestId	network.RequestId	`json:"requestId"`
	
	// The encoding to use.
	Encoding	string	`json:"encoding"`
	
	// The quality of the encoding (0-1). (defaults to 1)
	Quality	float64	`json:"quality"`
	
	// Whether to only return the size information (defaults to false).
	SizeOnly	bool	`json:"sizeOnly"`
	
}

// GetEncodedResponse returns
type GetEncodedResponseReturns struct {
	
	// The encoded body as a base64 string. Omitted if sizeOnly is true.
	Body	string	`json:"body"`
	
	// Size before re-encoding.
	OriginalSize	int	`json:"originalSize"`
	
	// Size after re-encoding.
	EncodedSize	int	`json:"encodedSize"`
	
}

