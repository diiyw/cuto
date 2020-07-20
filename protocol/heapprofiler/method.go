package heapprofiler

import (
	"github.com/diiyw/cuto/protocol/runtime"
)


// Enables console to refer to the node with given id via $x (see Command Line API for more details
// $x functions).
const AddInspectedHeapObject = "HeapProfiler.addInspectedHeapObject"

type AddInspectedHeapObjectParams struct {

	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectId 	HeapSnapshotObjectId	`json:"heapObjectId"`
}

type AddInspectedHeapObjectResult struct {

}

// 
const CollectGarbage = "HeapProfiler.collectGarbage"

type CollectGarbageParams struct {
}

type CollectGarbageResult struct {

}

// 
const Disable = "HeapProfiler.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// 
const Enable = "HeapProfiler.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// 
const GetHeapObjectId = "HeapProfiler.getHeapObjectId"

type GetHeapObjectIdParams struct {

	// Identifier of the object to get heap object id for.
	ObjectId 	runtime.RemoteObjectId	`json:"objectId"`
}

type GetHeapObjectIdResult struct {

	// Id of the heap snapshot object corresponding to the passed remote object id.
	HeapSnapshotObjectId 	HeapSnapshotObjectId	`json:"heapSnapshotObjectId"`
}

// 
const GetObjectByHeapObjectId = "HeapProfiler.getObjectByHeapObjectId"

type GetObjectByHeapObjectIdParams struct {

	// 
	ObjectId 	HeapSnapshotObjectId	`json:"objectId"`

	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup 	string	`json:"objectGroup,omitempty"`
}

type GetObjectByHeapObjectIdResult struct {

	// Evaluation result.
	Result 	runtime.RemoteObject	`json:"result"`
}

// 
const GetSamplingProfile = "HeapProfiler.getSamplingProfile"

type GetSamplingProfileParams struct {
}

type GetSamplingProfileResult struct {

	// Return the sampling profile being collected.
	Profile 	SamplingHeapProfile	`json:"profile"`
}

// 
const StartSampling = "HeapProfiler.startSampling"

type StartSamplingParams struct {

	// Average sample interval in bytes. Poisson distribution is used for the intervals. The
	// default value is 32768 bytes.
	SamplingInterval 	float64	`json:"samplingInterval,omitempty"`
}

type StartSamplingResult struct {

}

// 
const StartTrackingHeapObjects = "HeapProfiler.startTrackingHeapObjects"

type StartTrackingHeapObjectsParams struct {

	// 
	TrackAllocations 	bool	`json:"trackAllocations,omitempty"`
}

type StartTrackingHeapObjectsResult struct {

}

// 
const StopSampling = "HeapProfiler.stopSampling"

type StopSamplingParams struct {
}

type StopSamplingResult struct {

	// Recorded sampling heap profile.
	Profile 	SamplingHeapProfile	`json:"profile"`
}

// 
const StopTrackingHeapObjects = "HeapProfiler.stopTrackingHeapObjects"

type StopTrackingHeapObjectsParams struct {

	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken
	// when the tracking is stopped.
	ReportProgress 	bool	`json:"reportProgress,omitempty"`
}

type StopTrackingHeapObjectsResult struct {

}

// 
const TakeHeapSnapshot = "HeapProfiler.takeHeapSnapshot"

type TakeHeapSnapshotParams struct {

	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress 	bool	`json:"reportProgress,omitempty"`
}

type TakeHeapSnapshotResult struct {

}