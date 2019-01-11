package applicationcache

import (

	"github.com/diiyw/gator/protocol/frame"

)
const (
	
	// Enables application cache domain notifications.
	Enable = "ApplicationCache.enable"
	
	// Returns relevant application cache data for the document in given frame.
	GetApplicationCacheForFrame = "ApplicationCache.getApplicationCacheForFrame"
	
	// Returns array of frame identifiers with manifest urls for each frame containing a document
	// associated with some application cache.
	GetFramesWithManifests = "ApplicationCache.getFramesWithManifests"
	
	// Returns manifest URL for document in the given frame.
	GetManifestForFrame = "ApplicationCache.getManifestForFrame"
	
)

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// GetApplicationCacheForFrame parameters
type GetApplicationCacheForFrameParams struct {
	
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameId	frame.FrameId	`json:"frameId"`
	
}

// GetApplicationCacheForFrame returns
type GetApplicationCacheForFrameReturns struct {
	
	// Relevant application cache data for the document in given frame.
	ApplicationCache	ApplicationCache	`json:"applicationCache"`
	
}

// GetFramesWithManifests parameters
type GetFramesWithManifestsParams struct {
	
}

// GetFramesWithManifests returns
type GetFramesWithManifestsReturns struct {
	
	// Array of frame identifiers with manifest urls for each frame containing a document
	// associated with some application cache.
	FrameIds	[]FrameWithManifest	`json:"frameIds"`
	
}

// GetManifestForFrame parameters
type GetManifestForFrameParams struct {
	
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameId	frame.FrameId	`json:"frameId"`
	
}

// GetManifestForFrame returns
type GetManifestForFrameReturns struct {
	
	// Manifest URL for document in the given frame.
	ManifestURL	string	`json:"manifestURL"`
	
}

