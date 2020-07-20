package tracing
// Configuration for memory dump. Used only when "memory-infra" category is enabled.
type MemoryDumpConfig interface{}

// 
type TraceConfig  struct {

	// Controls how the trace buffer stores data.
	RecordMode	string	`json:"recordMode,omitempty"`

	// Turns on JavaScript stack sampling.
	EnableSampling	bool	`json:"enableSampling,omitempty"`

	// Turns on system tracing.
	EnableSystrace	bool	`json:"enableSystrace,omitempty"`

	// Turns on argument filter.
	EnableArgumentFilter	bool	`json:"enableArgumentFilter,omitempty"`

	// Included category filters.
	IncludedCategories	[]string	`json:"includedCategories,omitempty"`

	// Excluded category filters.
	ExcludedCategories	[]string	`json:"excludedCategories,omitempty"`

	// Configuration to synthesize the delays in tracing.
	SyntheticDelays	[]string	`json:"syntheticDelays,omitempty"`

	// Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
	MemoryDumpConfig	MemoryDumpConfig	`json:"memoryDumpConfig,omitempty"`
}

// Data format of a trace. Can be either the legacy JSON format or the
	// protocol buffer format. Note that the JSON format will be deprecated soon.
type StreamFormat string

// Compression type to use for traces returned via streams.
type StreamCompression string
