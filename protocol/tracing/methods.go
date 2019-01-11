package tracing

const (
	
	// Stop trace events collection.
	End = "Tracing.end"
	
	// Gets supported tracing categories.
	GetCategories = "Tracing.getCategories"
	
	// Record a clock sync marker in the trace.
	RecordClockSyncMarker = "Tracing.recordClockSyncMarker"
	
	// Request a global memory dump.
	RequestMemoryDump = "Tracing.requestMemoryDump"
	
	// Start trace events collection.
	Start = "Tracing.start"
	
)

// End parameters
type EndParams struct {
	
}

// End returns
type EndReturns struct {
	
}

// GetCategories parameters
type GetCategoriesParams struct {
	
}

// GetCategories returns
type GetCategoriesReturns struct {
	
	// A list of supported tracing categories.
	Categories	[]string	`json:"categories"`
	
}

// RecordClockSyncMarker parameters
type RecordClockSyncMarkerParams struct {
	
	// The ID of this clock sync marker
	SyncId	string	`json:"syncId"`
	
}

// RecordClockSyncMarker returns
type RecordClockSyncMarkerReturns struct {
	
}

// RequestMemoryDump parameters
type RequestMemoryDumpParams struct {
	
}

// RequestMemoryDump returns
type RequestMemoryDumpReturns struct {
	
	// GUID of the resulting global memory dump.
	DumpGuid	string	`json:"dumpGuid"`
	
	// True iff the global memory dump succeeded.
	Success	bool	`json:"success"`
	
}

// Start parameters
type StartParams struct {
	
	// Category/tag filter
	Categories	string	`json:"categories"`
	
	// Tracing options
	Options	string	`json:"options"`
	
	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval	float64	`json:"bufferUsageReportingInterval"`
	
	// Whether to report trace events as series of dataCollected events or to save trace to a
	// stream (defaults to `ReportEvents`).
	TransferMode	string	`json:"transferMode"`
	
	// Compression format to use. This only applies when using `ReturnAsStream`
	// transfer mode (defaults to `none`)
	StreamCompression	StreamCompression	`json:"streamCompression"`
	
	
	TraceConfig	TraceConfig	`json:"traceConfig"`
	
}

// Start returns
type StartReturns struct {
	
}

