package tracing

// Stop trace events collection.
const End = "Tracing.end"

type EndParams struct {
}

type EndResult struct {

}

// Gets supported tracing categories.
const GetCategories = "Tracing.getCategories"

type GetCategoriesParams struct {
}

type GetCategoriesResult struct {

	// A list of supported tracing categories.
	Categories 	[]string	`json:"categories"`
}

// Record a clock sync marker in the trace.
const RecordClockSyncMarker = "Tracing.recordClockSyncMarker"

type RecordClockSyncMarkerParams struct {

	// The ID of this clock sync marker
	SyncId 	string	`json:"syncId"`
}

type RecordClockSyncMarkerResult struct {

}

// Request a global memory dump.
const RequestMemoryDump = "Tracing.requestMemoryDump"

type RequestMemoryDumpParams struct {

	// Enables more deterministic results by forcing garbage collection
	Deterministic 	bool	`json:"deterministic,omitempty"`
}

type RequestMemoryDumpResult struct {

	// GUID of the resulting global memory dump.
	DumpGuid 	string	`json:"dumpGuid"`
	// True iff the global memory dump succeeded.
	Success 	bool	`json:"success"`
}

// Start trace events collection.
const Start = "Tracing.start"

type StartParams struct {

	// Category/tag filter
	Categories 	string	`json:"categories,omitempty"`

	// Tracing options
	Options 	string	`json:"options,omitempty"`

	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval 	float64	`json:"bufferUsageReportingInterval,omitempty"`

	// Whether to report trace events as series of dataCollected events or to save trace to a
	// stream (defaults to `ReportEvents`).
	TransferMode 	string	`json:"transferMode,omitempty"`

	// Trace data format to use. This only applies when using `ReturnAsStream`
	// transfer mode (defaults to `json`).
	StreamFormat 	StreamFormat	`json:"streamFormat,omitempty"`

	// Compression format to use. This only applies when using `ReturnAsStream`
	// transfer mode (defaults to `none`)
	StreamCompression 	StreamCompression	`json:"streamCompression,omitempty"`

	// 
	TraceConfig 	TraceConfig	`json:"traceConfig,omitempty"`
}

type StartResult struct {

}