package heapprofiler
// Heap snapshot object id.
type HeapSnapshotObjectId string

// Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
type SamplingHeapProfileNode  struct {

	// Function location.
	CallFrame	interface{}	`json:"callFrame"`

	// Allocations size in bytes for the node excluding children.
	SelfSize	float64	`json:"selfSize"`

	// Node id. Ids are unique across all profiles collected between startSampling and stopSampling.
	Id	int	`json:"id"`

	// Child nodes.
	Children	[]*SamplingHeapProfileNode	`json:"children"`
}

// A single sample from a sampling profile.
type SamplingHeapProfileSample  struct {

	// Allocation size in bytes attributed to the sample.
	Size	float64	`json:"size"`

	// Id of the corresponding profile tree node.
	NodeId	int	`json:"nodeId"`

	// Time-ordered sample ordinal number. It is unique across all profiles retrieved
	// between startSampling and stopSampling.
	Ordinal	float64	`json:"ordinal"`
}

// Sampling profile.
type SamplingHeapProfile  struct {

	// 
	Head	SamplingHeapProfileNode	`json:"head"`

	// 
	Samples	[]*SamplingHeapProfileSample	`json:"samples"`
}
