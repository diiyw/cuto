package io

// Close the stream, discard any temporary backing storage.
const Close = "IO.close"

type CloseParams struct {

	// Handle of the stream to close.
	Handle 	StreamHandle	`json:"handle"`
}

type CloseResult struct {

}

// Read a chunk of the stream
const Read = "IO.read"

type ReadParams struct {

	// Handle of the stream to read.
	Handle 	StreamHandle	`json:"handle"`

	// Seek to the specified offset before reading (if not specificed, proceed with offset
	// following the last read). Some types of streams may only support sequential reads.
	Offset 	int	`json:"offset"`

	// Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size 	int	`json:"size"`
}

type ReadResult struct {

	// Set if the data is base64-encoded
	Base64Encoded 	bool	`json:"base64Encoded"`
	// Data that were read.
	Data 	string	`json:"data"`
	// Set if the end-of-file condition occured while reading.
	Eof 	bool	`json:"eof"`
}

// Return UUID of Blob object specified by a remote object id.
const ResolveBlob = "IO.resolveBlob"

type ResolveBlobParams struct {

	// Object id of a Blob object wrapper.
	ObjectId 	interface{}	`json:"objectId"`
}

type ResolveBlobResult struct {

	// UUID of the specified Blob.
	Uuid 	string	`json:"uuid"`
}