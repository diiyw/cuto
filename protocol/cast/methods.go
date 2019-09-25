package cast

const (
	
	// Starts observing for sinks that can be used for tab mirroring, and if set,
	// sinks compatible with |presentationUrl| as well. When sinks are found, a
	// |sinksUpdated| event is fired.
	// Also starts observing for issue messages. When an issue is added or removed,
	// an |issueUpdated| event is fired.
	Enable = "Cast.enable"
	
	// Stops observing for sinks and issues.
	Disable = "Cast.disable"
	
	// Sets a sink to be used when the web page requests the browser to choose a
	// sink via Presentation API, Remote Playback API, or Cast SDK.
	SetSinkToUse = "Cast.setSinkToUse"
	
	// Starts mirroring the tab to the sink.
	StartTabMirroring = "Cast.startTabMirroring"
	
	// Stops the active Cast session on the sink.
	StopCasting = "Cast.stopCasting"
	
)

// Enable parameters
type EnableParams struct {
	
	
	PresentationUrl	string	`json:"presentationUrl"`
	
}

// Enable returns
type EnableReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// SetSinkToUse parameters
type SetSinkToUseParams struct {
	
	
	SinkName	string	`json:"sinkName"`
	
}

// SetSinkToUse returns
type SetSinkToUseReturns struct {
	
}

// StartTabMirroring parameters
type StartTabMirroringParams struct {
	
	
	SinkName	string	`json:"sinkName"`
	
}

// StartTabMirroring returns
type StartTabMirroringReturns struct {
	
}

// StopCasting parameters
type StopCastingParams struct {
	
	
	SinkName	string	`json:"sinkName"`
	
}

// StopCasting returns
type StopCastingReturns struct {
	
}

