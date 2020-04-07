package headlessexperimental

// Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a
// screenshot from the resulting frame. Requires that the target was created with enabled
// BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gl/3zHXhB for more background.
const BeginFrame = "HeadlessExperimental.beginFrame"

type BeginFrameParams struct {

	// Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set,
	// the current time will be used.
	FrameTimeTicks 	float64	`json:"frameTimeTicks"`

	// The interval between BeginFrames that is reported to the compositor, in milliseconds.
	// Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval 	float64	`json:"interval"`

	// Whether updates should not be committed and drawn onto the display. False by default. If
	// true, only side effects of the BeginFrame will be run, such as layout and animations, but
	// any visual updates may not be visible on the display or in screenshots.
	NoDisplayUpdates 	bool	`json:"noDisplayUpdates"`

	// If set, a screenshot of the frame will be captured and returned in the response. Otherwise,
	// no screenshot will be captured. Note that capturing a screenshot can fail, for example,
	// during renderer initialization. In such a case, no screenshot data will be returned.
	Screenshot 	ScreenshotParams	`json:"screenshot"`
}

type BeginFrameResult struct {

	// Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the
	// display. Reported for diagnostic uses, may be removed in the future.
	HasDamage 	bool	`json:"hasDamage"`
	// Base64-encoded image data of the screenshot, if one was requested and successfully taken.
	ScreenshotData 	[]byte	`json:"screenshotData"`
}

// Disables headless events for the target.
const Disable = "HeadlessExperimental.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables headless events for the target.
const Enable = "HeadlessExperimental.enable"

type EnableParams struct {
}

type EnableResult struct {

}