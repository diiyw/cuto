package tracing

import (

	"github.com/diiyw/chr/protocol/io"

)
const (
	
	
	BufferUsageEvent = "Tracing.bufferUsage"
	
	// Contains an bucket of collected trace events. When tracing is stopped collected events will be
	// send as a sequence of dataCollected events followed by tracingComplete event.
	DataCollectedEvent = "Tracing.dataCollected"
	
	// Signals that tracing is stopped and there is no trace buffers pending flush, all data were
	// delivered via dataCollected events.
	TracingCompleteEvent = "Tracing.tracingComplete"
	
)


type BufferUsageParams struct {
	
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its
	// total size.
	PercentFull	float64	`json:"percentFull"`
	
	// An approximate number of events in the trace log.
	EventCount	float64	`json:"eventCount"`
	
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its
	// total size.
	Value	float64	`json:"value"`
	
}

// Contains an bucket of collected trace events. When tracing is stopped collected events will be
	// send as a sequence of dataCollected events followed by tracingComplete event.
type DataCollectedParams struct {
	
	
	Value	[]interface{}	`json:"value"`
	
}

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were
	// delivered via dataCollected events.
type TracingCompleteParams struct {
	
	// A handle of the stream that holds resulting trace data.
	Stream	io.StreamHandle	`json:"stream"`
	
	// Trace data format of returned stream.
	TraceFormat	StreamFormat	`json:"traceFormat"`
	
	// Compression format of returned stream.
	StreamCompression	StreamCompression	`json:"streamCompression"`
	
}

