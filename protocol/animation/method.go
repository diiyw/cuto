package animation

// Disables animation domain notifications.
const Disable = "Animation.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables animation domain notifications.
const Enable = "Animation.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Returns the current time of the an animation.
const GetCurrentTime = "Animation.getCurrentTime"

type GetCurrentTimeParams struct {

	// Id of animation.
	Id 	string	`json:"id"`
}

type GetCurrentTimeResult struct {

	// Current time of the page.
	CurrentTime 	float64	`json:"currentTime"`
}

// Gets the playback rate of the document timeline.
const GetPlaybackRate = "Animation.getPlaybackRate"

type GetPlaybackRateParams struct {
}

type GetPlaybackRateResult struct {

	// Playback rate for animations on page.
	PlaybackRate 	float64	`json:"playbackRate"`
}

// Releases a set of animations to no longer be manipulated.
const ReleaseAnimations = "Animation.releaseAnimations"

type ReleaseAnimationsParams struct {

	// List of animation ids to seek.
	Animations 	[]string	`json:"animations"`
}

type ReleaseAnimationsResult struct {

}

// Gets the remote object of the Animation.
const ResolveAnimation = "Animation.resolveAnimation"

type ResolveAnimationParams struct {

	// Animation id.
	AnimationId 	string	`json:"animationId"`
}

type ResolveAnimationResult struct {

	// Corresponding remote object.
	RemoteObject 	interface{}	`json:"remoteObject"`
}

// Seek a set of animations to a particular time within each animation.
const SeekAnimations = "Animation.seekAnimations"

type SeekAnimationsParams struct {

	// List of animation ids to seek.
	Animations 	[]string	`json:"animations"`

	// Set the current time of each animation.
	CurrentTime 	float64	`json:"currentTime"`
}

type SeekAnimationsResult struct {

}

// Sets the paused state of a set of animations.
const SetPaused = "Animation.setPaused"

type SetPausedParams struct {

	// Animations to set the pause state of.
	Animations 	[]string	`json:"animations"`

	// Paused state to set to.
	Paused 	bool	`json:"paused"`
}

type SetPausedResult struct {

}

// Sets the playback rate of the document timeline.
const SetPlaybackRate = "Animation.setPlaybackRate"

type SetPlaybackRateParams struct {

	// Playback rate for animations on page
	PlaybackRate 	float64	`json:"playbackRate"`
}

type SetPlaybackRateResult struct {

}

// Sets the timing of an animation node.
const SetTiming = "Animation.setTiming"

type SetTimingParams struct {

	// Animation id.
	AnimationId 	string	`json:"animationId"`

	// Duration of the animation.
	Duration 	float64	`json:"duration"`

	// Delay of the animation.
	Delay 	float64	`json:"delay"`
}

type SetTimingResult struct {

}