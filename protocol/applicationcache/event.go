package applicationcache

import (
	"github.com/diiyw/cuto/protocol/cdp"
)


// 
const ApplicationCacheStatusUpdatedEvent = "ApplicationCache.applicationCacheStatusUpdated"
type ApplicationCacheStatusUpdatedParams struct {

	// Identifier of the frame containing document whose application cache updated status.
	FrameId 	cdp.FrameId
	// Manifest URL.
	ManifestURL 	string
	// Updated application cache status.
	Status 	int}



// 
const NetworkStateUpdatedEvent = "ApplicationCache.networkStateUpdated"
type NetworkStateUpdatedParams struct {

	// 
	IsNowOnline 	bool}

