package applicationcache

import (

	"github.com/diiyw/goc/protocol/frame"

)
const (
	
	
	ApplicationCacheStatusUpdatedEvent = "ApplicationCache.applicationCacheStatusUpdated"
	
	
	NetworkStateUpdatedEvent = "ApplicationCache.networkStateUpdated"
	
)


type ApplicationCacheStatusUpdatedParams struct {
	
	// Identifier of the frame containing document whose application cache updated status.
	FrameId	frame.FrameId	`json:"frameId"`
	
	// Manifest URL.
	ManifestURL	string	`json:"manifestURL"`
	
	// Updated application cache status.
	Status	int	`json:"status"`
	
}


type NetworkStateUpdatedParams struct {
	
	
	IsNowOnline	bool	`json:"isNowOnline"`
	
}

