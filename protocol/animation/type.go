package animation
// Animation instance.
type Animation  struct {

	// `Animation`'s id.
	Id	string	`json:"id"`

	// `Animation`'s name.
	Name	string	`json:"name"`

	// `Animation`'s internal paused state.
	PausedState	bool	`json:"pausedState"`

	// `Animation`'s play state.
	PlayState	string	`json:"playState"`

	// `Animation`'s playback rate.
	PlaybackRate	float64	`json:"playbackRate"`

	// `Animation`'s start time.
	StartTime	float64	`json:"startTime"`

	// `Animation`'s current time.
	CurrentTime	float64	`json:"currentTime"`

	// Animation type of `Animation`.
	Type	string	`json:"type"`

	// `Animation`'s source animation node.
	Source	AnimationEffect	`json:"source"`

	// A unique ID for `Animation` representing the sources that triggered this CSS
	// animation/transition.
	CssId	string	`json:"cssId"`
}

// AnimationEffect instance
type AnimationEffect  struct {

	// `AnimationEffect`'s delay.
	Delay	float64	`json:"delay"`

	// `AnimationEffect`'s end delay.
	EndDelay	float64	`json:"endDelay"`

	// `AnimationEffect`'s iteration start.
	IterationStart	float64	`json:"iterationStart"`

	// `AnimationEffect`'s iterations.
	Iterations	float64	`json:"iterations"`

	// `AnimationEffect`'s iteration duration.
	Duration	float64	`json:"duration"`

	// `AnimationEffect`'s playback direction.
	Direction	string	`json:"direction"`

	// `AnimationEffect`'s fill mode.
	Fill	string	`json:"fill"`

	// `AnimationEffect`'s target node.
	BackendNodeId	interface{}	`json:"backendNodeId"`

	// `AnimationEffect`'s keyframes.
	KeyframesRule	KeyframesRule	`json:"keyframesRule"`

	// `AnimationEffect`'s timing function.
	Easing	string	`json:"easing"`
}

// Keyframes Rule
type KeyframesRule  struct {

	// CSS keyframed animation's name.
	Name	string	`json:"name"`

	// List of animation keyframes.
	Keyframes	[]*KeyframeStyle	`json:"keyframes"`
}

// Keyframe Style
type KeyframeStyle  struct {

	// Keyframe's time offset.
	Offset	string	`json:"offset"`

	// `AnimationEffect`'s timing function.
	Easing	string	`json:"easing"`
}
