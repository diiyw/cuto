package headlessexperimental

const (
	
	// Issued when the target starts or stops needing BeginFrames.
	NeedsBeginFramesChangedEvent = "HeadlessExperimental.needsBeginFramesChanged"
	
)

// Issued when the target starts or stops needing BeginFrames.
type NeedsBeginFramesChangedParams struct {
	
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames	bool	`json:"needsBeginFrames"`
	
}

