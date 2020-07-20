package tracing

import (
	"github.com/diiyw/cuto/protocol/io"
)


// 
const BufferUsageEvent = "Tracing.bufferUsage"
type BufferUsageParams struct {

	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its
	// total size.
	PercentFull 	float64
	// An approximate number of events in the trace log.
	EventCount 	float64
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its
	// total size.
	Value 	float64}



// Contains an bucket of collected trace events. When tracing is stopped collected events will be
// send as a sequence of dataCollected events followed by tracingComplete event.
const DataCollectedEvent = "Tracing.dataCollected"
type DataCollectedParams struct {

	// 
	Value 	[]interface{}}



// Signals that tracing is stopped and there is no trace buffers pending flush, all data were
// delivered via dataCollected events.
const TracingCompleteEvent = "Tracing.tracingComplete"
type TracingCompleteParams struct {

	// Indicates whether some trace data is known to have been lost, e.g. because the trace ring
	// buffer wrapped around.
	DataLossOccurred 	bool
	// A handle of the stream that holds resulting trace data.
	Stream 	io.StreamHandle
	// Trace data format of returned stream.
	TraceFormat 	StreamFormat
	// Compression format of returned stream.
	StreamCompression 	StreamCompression}

