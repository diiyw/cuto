package animation

const (
	
	// Event for when an animation has been cancelled.
	AnimationCanceledEvent = "Animation.animationCanceled"
	
	// Event for each animation that has been created.
	AnimationCreatedEvent = "Animation.animationCreated"
	
	// Event for animation that has been started.
	AnimationStartedEvent = "Animation.animationStarted"
	
)

// Event for when an animation has been cancelled.
type AnimationCanceledParams struct {
	
	// Id of the animation that was cancelled.
	Id	string	`json:"id"`
	
}

// Event for each animation that has been created.
type AnimationCreatedParams struct {
	
	// Id of the animation that was created.
	Id	string	`json:"id"`
	
}

// Event for animation that has been started.
type AnimationStartedParams struct {
	
	// Animation that was started.
	Animation	Animation	`json:"animation"`
	
}

