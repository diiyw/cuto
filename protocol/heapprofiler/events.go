package heapprofiler

const (
	
	
	AddHeapSnapshotChunkEvent = "HeapProfiler.addHeapSnapshotChunk"
	
	// If heap objects tracking has been started then backend may send update for one or more fragments
	HeapStatsUpdateEvent = "HeapProfiler.heapStatsUpdate"
	
	// If heap objects tracking has been started then backend regularly sends a current value for last
	// seen object id and corresponding timestamp. If the were changes in the heap since last event
	// then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
	LastSeenObjectIdEvent = "HeapProfiler.lastSeenObjectId"
	
	
	ReportHeapSnapshotProgressEvent = "HeapProfiler.reportHeapSnapshotProgress"
	
	
	ResetProfilesEvent = "HeapProfiler.resetProfiles"
	
)


type AddHeapSnapshotChunkParams struct {
	
	
	Chunk	string	`json:"chunk"`
	
}

// If heap objects tracking has been started then backend may send update for one or more fragments
type HeapStatsUpdateParams struct {
	
	// An array of triplets. Each triplet describes a fragment. The first integer is the fragment
	// index, the second integer is a total count of objects for the fragment, the third integer is
	// a total size of the objects for the fragment.
	StatsUpdate	[]int	`json:"statsUpdate"`
	
}

// If heap objects tracking has been started then backend regularly sends a current value for last
	// seen object id and corresponding timestamp. If the were changes in the heap since last event
	// then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type LastSeenObjectIdParams struct {
	
	
	LastSeenObjectId	int	`json:"lastSeenObjectId"`
	
	
	Timestamp	float64	`json:"timestamp"`
	
}


type ReportHeapSnapshotProgressParams struct {
	
	
	Done	int	`json:"done"`
	
	
	Total	int	`json:"total"`
	
	
	Finished	bool	`json:"finished"`
	
}


type ResetProfilesParams struct {
	
}

