package heapprofiler

// 
const AddHeapSnapshotChunkEvent = "HeapProfiler.addHeapSnapshotChunk"
type AddHeapSnapshotChunkParams struct {

	// 
	Chunk 	string}



// If heap objects tracking has been started then backend may send update for one or more fragments
const HeapStatsUpdateEvent = "HeapProfiler.heapStatsUpdate"
type HeapStatsUpdateParams struct {

	// An array of triplets. Each triplet describes a fragment. The first integer is the fragment
	// index, the second integer is a total count of objects for the fragment, the third integer is
	// a total size of the objects for the fragment.
	StatsUpdate 	[]int}



// If heap objects tracking has been started then backend regularly sends a current value for last
// seen object id and corresponding timestamp. If the were changes in the heap since last event
// then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
const LastSeenObjectIdEvent = "HeapProfiler.lastSeenObjectId"
type LastSeenObjectIdParams struct {

	// 
	LastSeenObjectId 	int
	// 
	Timestamp 	float64}



// 
const ReportHeapSnapshotProgressEvent = "HeapProfiler.reportHeapSnapshotProgress"
type ReportHeapSnapshotProgressParams struct {

	// 
	Done 	int
	// 
	Total 	int
	// 
	Finished 	bool}



// 
const ResetProfilesEvent = "HeapProfiler.resetProfiles"
type ResetProfilesParams struct {
}

