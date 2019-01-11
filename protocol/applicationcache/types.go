package applicationcache

import (

	"github.com/diiyw/gator/protocol/frame"

)

// Detailed application cache resource information.
type ApplicationCacheResource struct {
	
	// Resource url.
	
	Url	string	`json:"url"`
	
	// Resource size.
	
	Size	int	`json:"size"`
	
	// Resource type.
	
	Type	string	`json:"type"`
	
}	

// Detailed application cache information.
type ApplicationCache struct {
	
	// Manifest URL.
	
	ManifestURL	string	`json:"manifestURL"`
	
	// Application cache size.
	
	Size	float64	`json:"size"`
	
	// Application cache creation time.
	
	CreationTime	float64	`json:"creationTime"`
	
	// Application cache update time.
	
	UpdateTime	float64	`json:"updateTime"`
	
	// Application cache resources.
	
	Resources	[]ApplicationCacheResource	`json:"resources"`
	
}	

// Frame identifier - manifest URL pair.
type FrameWithManifest struct {
	
	// Frame identifier.
	
	FrameId	frame.FrameId	`json:"frameId"`
	
	// Manifest URL.
	
	ManifestURL	string	`json:"manifestURL"`
	
	// Application cache status.
	
	Status	int	`json:"status"`
	
}	

