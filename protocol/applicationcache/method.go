package applicationcache

// Enables application cache domain notifications.
const Enable = "ApplicationCache.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Returns relevant application cache data for the document in given frame.
const GetApplicationCacheForFrame = "ApplicationCache.getApplicationCacheForFrame"

type GetApplicationCacheForFrameParams struct {

	// Identifier of the frame containing document whose application cache is retrieved.
	FrameId 	interface{}	`json:"frameId"`
}

type GetApplicationCacheForFrameResult struct {

	// Relevant application cache data for the document in given frame.
	ApplicationCache 	ApplicationCache	`json:"applicationCache"`
}

// Returns array of frame identifiers with manifest urls for each frame containing a document
// associated with some application cache.
const GetFramesWithManifests = "ApplicationCache.getFramesWithManifests"

type GetFramesWithManifestsParams struct {
}

type GetFramesWithManifestsResult struct {

	// Array of frame identifiers with manifest urls for each frame containing a document
	// associated with some application cache.
	FrameIds 	[]*FrameWithManifest	`json:"frameIds"`
}

// Returns manifest URL for document in the given frame.
const GetManifestForFrame = "ApplicationCache.getManifestForFrame"

type GetManifestForFrameParams struct {

	// Identifier of the frame containing document whose manifest is retrieved.
	FrameId 	interface{}	`json:"frameId"`
}

type GetManifestForFrameResult struct {

	// Manifest URL for document in the given frame.
	ManifestURL 	string	`json:"manifestURL"`
}