package audits

import (
	"github.com/diiyw/cuto/protocol/network"
)


// Returns the response body and size if it were re-encoded with the specified settings. Only
// applies to images.
const GetEncodedResponse = "Audits.getEncodedResponse"

type GetEncodedResponseParams struct {

	// Identifier of the network request to get content for.
	RequestId 	network.RequestId	`json:"requestId"`

	// The encoding to use.
	Encoding 	string	`json:"encoding"`

	// The quality of the encoding (0-1). (defaults to 1)
	Quality 	float64	`json:"quality,omitempty"`

	// Whether to only return the size information (defaults to false).
	SizeOnly 	bool	`json:"sizeOnly,omitempty"`
}

type GetEncodedResponseResult struct {

	// The encoded body as a base64 string. Omitted if sizeOnly is true.
	Body 	[]byte	`json:"body"`
	// Size before re-encoding.
	OriginalSize 	int	`json:"originalSize"`
	// Size after re-encoding.
	EncodedSize 	int	`json:"encodedSize"`
}