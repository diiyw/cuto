package headlessexperimental

// Issued when the target starts or stops needing BeginFrames.
// Deprecated. Issue beginFrame unconditionally instead and use result from
// beginFrame to detect whether the frames were suppressed.
const NeedsBeginFramesChangedEvent = "HeadlessExperimental.needsBeginFramesChanged"
type NeedsBeginFramesChangedParams struct {

	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames 	bool}

