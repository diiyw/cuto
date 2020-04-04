package cast

// Starts observing for sinks that can be used for tab mirroring, and if set,
// sinks compatible with |presentationUrl| as well. When sinks are found, a
// |sinksUpdated| event is fired.
// Also starts observing for issue messages. When an issue is added or removed,
// an |issueUpdated| event is fired.
const Enable = "Cast.enable"

type EnableParams struct {

	// 
	PresentationUrl 	string	`json:"presentationUrl"`
}

type EnableResult struct {

}

// Stops observing for sinks and issues.
const Disable = "Cast.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Sets a sink to be used when the web page requests the browser to choose a
// sink via Presentation API, Remote Playback API, or Cast SDK.
const SetSinkToUse = "Cast.setSinkToUse"

type SetSinkToUseParams struct {

	// 
	SinkName 	string	`json:"sinkName"`
}

type SetSinkToUseResult struct {

}

// Starts mirroring the tab to the sink.
const StartTabMirroring = "Cast.startTabMirroring"

type StartTabMirroringParams struct {

	// 
	SinkName 	string	`json:"sinkName"`
}

type StartTabMirroringResult struct {

}

// Stops the active Cast session on the sink.
const StopCasting = "Cast.stopCasting"

type StopCastingParams struct {

	// 
	SinkName 	string	`json:"sinkName"`
}

type StopCastingResult struct {

}