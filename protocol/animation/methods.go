package animation

import (

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	// Disables animation domain notifications.
	Disable = "Animation.disable"
	
	// Enables animation domain notifications.
	Enable = "Animation.enable"
	
	// Returns the current time of the an animation.
	GetCurrentTime = "Animation.getCurrentTime"
	
	// Gets the playback rate of the document timeline.
	GetPlaybackRate = "Animation.getPlaybackRate"
	
	// Releases a set of animations to no longer be manipulated.
	ReleaseAnimations = "Animation.releaseAnimations"
	
	// Gets the remote object of the Animation.
	ResolveAnimation = "Animation.resolveAnimation"
	
	// Seek a set of animations to a particular time within each animation.
	SeekAnimations = "Animation.seekAnimations"
	
	// Sets the paused state of a set of animations.
	SetPaused = "Animation.setPaused"
	
	// Sets the playback rate of the document timeline.
	SetPlaybackRate = "Animation.setPlaybackRate"
	
	// Sets the timing of an animation node.
	SetTiming = "Animation.setTiming"
	
)

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

// GetCurrentTime parameters
type GetCurrentTimeParams struct {
	
	// Id of animation.
	Id	string	`json:"id"`
	
}

// GetCurrentTime returns
type GetCurrentTimeReturns struct {
	
	// Current time of the page.
	CurrentTime	float64	`json:"currentTime"`
	
}

// GetPlaybackRate parameters
type GetPlaybackRateParams struct {
	
}

// GetPlaybackRate returns
type GetPlaybackRateReturns struct {
	
	// Playback rate for animations on page.
	PlaybackRate	float64	`json:"playbackRate"`
	
}

// ReleaseAnimations parameters
type ReleaseAnimationsParams struct {
	
	// List of animation ids to seek.
	Animations	[]string	`json:"animations"`
	
}

// ReleaseAnimations returns
type ReleaseAnimationsReturns struct {
	
}

// ResolveAnimation parameters
type ResolveAnimationParams struct {
	
	// Animation id.
	AnimationId	string	`json:"animationId"`
	
}

// ResolveAnimation returns
type ResolveAnimationReturns struct {
	
	// Corresponding remote object.
	RemoteObject	runtime.RemoteObject	`json:"remoteObject"`
	
}

// SeekAnimations parameters
type SeekAnimationsParams struct {
	
	// List of animation ids to seek.
	Animations	[]string	`json:"animations"`
	
	// Set the current time of each animation.
	CurrentTime	float64	`json:"currentTime"`
	
}

// SeekAnimations returns
type SeekAnimationsReturns struct {
	
}

// SetPaused parameters
type SetPausedParams struct {
	
	// Animations to set the pause state of.
	Animations	[]string	`json:"animations"`
	
	// Paused state to set to.
	Paused	bool	`json:"paused"`
	
}

// SetPaused returns
type SetPausedReturns struct {
	
}

// SetPlaybackRate parameters
type SetPlaybackRateParams struct {
	
	// Playback rate for animations on page
	PlaybackRate	float64	`json:"playbackRate"`
	
}

// SetPlaybackRate returns
type SetPlaybackRateReturns struct {
	
}

// SetTiming parameters
type SetTimingParams struct {
	
	// Animation id.
	AnimationId	string	`json:"animationId"`
	
	// Duration of the animation.
	Duration	float64	`json:"duration"`
	
	// Delay of the animation.
	Delay	float64	`json:"delay"`
	
}

// SetTiming returns
type SetTimingReturns struct {
	
}

