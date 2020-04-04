package animation

// Event for when an animation has been cancelled.
const AnimationCanceledEvent = "Animation.animationCanceled"
type AnimationCanceledParams struct {

	// Id of the animation that was cancelled.
	Id 	string}



// Event for each animation that has been created.
const AnimationCreatedEvent = "Animation.animationCreated"
type AnimationCreatedParams struct {

	// Id of the animation that was created.
	Id 	string}



// Event for animation that has been started.
const AnimationStartedEvent = "Animation.animationStarted"
type AnimationStartedParams struct {

	// Animation that was started.
	Animation 	Animation}

