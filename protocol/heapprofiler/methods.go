package heapprofiler

import (

	"github.com/diiyw/gator/protocol/runtime"

)
const (
	
	// Enables console to refer to the node with given id via $x (see Command Line API for more details
	// $x functions).
	AddInspectedHeapObject = "HeapProfiler.addInspectedHeapObject"
	
	
	CollectGarbage = "HeapProfiler.collectGarbage"
	
	
	Disable = "HeapProfiler.disable"
	
	
	Enable = "HeapProfiler.enable"
	
	
	GetHeapObjectId = "HeapProfiler.getHeapObjectId"
	
	
	GetObjectByHeapObjectId = "HeapProfiler.getObjectByHeapObjectId"
	
	
	GetSamplingProfile = "HeapProfiler.getSamplingProfile"
	
	
	StartSampling = "HeapProfiler.startSampling"
	
	
	StartTrackingHeapObjects = "HeapProfiler.startTrackingHeapObjects"
	
	
	StopSampling = "HeapProfiler.stopSampling"
	
	
	StopTrackingHeapObjects = "HeapProfiler.stopTrackingHeapObjects"
	
	
	TakeHeapSnapshot = "HeapProfiler.takeHeapSnapshot"
	
)

// AddInspectedHeapObject parameters
type AddInspectedHeapObjectParams struct {
	
	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectId	HeapSnapshotObjectId	`json:"heapObjectId"`
	
}

// AddInspectedHeapObject returns
type AddInspectedHeapObjectReturns struct {
	
}

// CollectGarbage parameters
type CollectGarbageParams struct {
	
}

// CollectGarbage returns
type CollectGarbageReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// GetHeapObjectId parameters
type GetHeapObjectIdParams struct {
	
	// Identifier of the object to get heap object id for.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
}

// GetHeapObjectId returns
type GetHeapObjectIdReturns struct {
	
	// Id of the heap snapshot object corresponding to the passed remote object id.
	HeapSnapshotObjectId	HeapSnapshotObjectId	`json:"heapSnapshotObjectId"`
	
}

// GetObjectByHeapObjectId parameters
type GetObjectByHeapObjectIdParams struct {
	
	
	ObjectId	HeapSnapshotObjectId	`json:"objectId"`
	
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup	string	`json:"objectGroup"`
	
}

// GetObjectByHeapObjectId returns
type GetObjectByHeapObjectIdReturns struct {
	
	// Evaluation result.
	Result	runtime.RemoteObject	`json:"result"`
	
}

// GetSamplingProfile parameters
type GetSamplingProfileParams struct {
	
}

// GetSamplingProfile returns
type GetSamplingProfileReturns struct {
	
	// Return the sampling profile being collected.
	Profile	SamplingHeapProfile	`json:"profile"`
	
}

// StartSampling parameters
type StartSamplingParams struct {
	
	// Average sample interval in bytes. Poisson distribution is used for the intervals. The
	// default value is 32768 bytes.
	SamplingInterval	float64	`json:"samplingInterval"`
	
}

// StartSampling returns
type StartSamplingReturns struct {
	
}

// StartTrackingHeapObjects parameters
type StartTrackingHeapObjectsParams struct {
	
	
	TrackAllocations	bool	`json:"trackAllocations"`
	
}

// StartTrackingHeapObjects returns
type StartTrackingHeapObjectsReturns struct {
	
}

// StopSampling parameters
type StopSamplingParams struct {
	
}

// StopSampling returns
type StopSamplingReturns struct {
	
	// Recorded sampling heap profile.
	Profile	SamplingHeapProfile	`json:"profile"`
	
}

// StopTrackingHeapObjects parameters
type StopTrackingHeapObjectsParams struct {
	
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken
	// when the tracking is stopped.
	ReportProgress	bool	`json:"reportProgress"`
	
}

// StopTrackingHeapObjects returns
type StopTrackingHeapObjectsReturns struct {
	
}

// TakeHeapSnapshot parameters
type TakeHeapSnapshotParams struct {
	
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress	bool	`json:"reportProgress"`
	
}

// TakeHeapSnapshot returns
type TakeHeapSnapshotReturns struct {
	
}

