package layertree

import (

	"github.com/diiyw/chr/protocol/dom"

)
const (
	
	// Provides the reasons why the given layer was composited.
	CompositingReasons = "LayerTree.compositingReasons"
	
	// Disables compositing tree inspection.
	Disable = "LayerTree.disable"
	
	// Enables compositing tree inspection.
	Enable = "LayerTree.enable"
	
	// Returns the snapshot identifier.
	LoadSnapshot = "LayerTree.loadSnapshot"
	
	// Returns the layer snapshot identifier.
	MakeSnapshot = "LayerTree.makeSnapshot"
	
	
	ProfileSnapshot = "LayerTree.profileSnapshot"
	
	// Releases layer snapshot captured by the back-end.
	ReleaseSnapshot = "LayerTree.releaseSnapshot"
	
	// Replays the layer snapshot and returns the resulting bitmap.
	ReplaySnapshot = "LayerTree.replaySnapshot"
	
	// Replays the layer snapshot and returns canvas log.
	SnapshotCommandLog = "LayerTree.snapshotCommandLog"
	
)

// CompositingReasons parameters
type CompositingReasonsParams struct {
	
	// The id of the layer for which we want to get the reasons it was composited.
	LayerId	LayerId	`json:"layerId"`
	
}

// CompositingReasons returns
type CompositingReasonsReturns struct {
	
	// A list of strings specifying reasons for the given layer to become composited.
	CompositingReasons	[]string	`json:"compositingReasons"`
	
}

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

// LoadSnapshot parameters
type LoadSnapshotParams struct {
	
	// An array of tiles composing the snapshot.
	Tiles	[]PictureTile	`json:"tiles"`
	
}

// LoadSnapshot returns
type LoadSnapshotReturns struct {
	
	// The id of the snapshot.
	SnapshotId	SnapshotId	`json:"snapshotId"`
	
}

// MakeSnapshot parameters
type MakeSnapshotParams struct {
	
	// The id of the layer.
	LayerId	LayerId	`json:"layerId"`
	
}

// MakeSnapshot returns
type MakeSnapshotReturns struct {
	
	// The id of the layer snapshot.
	SnapshotId	SnapshotId	`json:"snapshotId"`
	
}

// ProfileSnapshot parameters
type ProfileSnapshotParams struct {
	
	// The id of the layer snapshot.
	SnapshotId	SnapshotId	`json:"snapshotId"`
	
	// The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount	int	`json:"minRepeatCount"`
	
	// The minimum duration (in seconds) to replay the snapshot.
	MinDuration	float64	`json:"minDuration"`
	
	// The clip rectangle to apply when replaying the snapshot.
	ClipRect	dom.Rect	`json:"clipRect"`
	
}

// ProfileSnapshot returns
type ProfileSnapshotReturns struct {
	
	// The array of paint profiles, one per run.
	Timings	[]PaintProfile	`json:"timings"`
	
}

// ReleaseSnapshot parameters
type ReleaseSnapshotParams struct {
	
	// The id of the layer snapshot.
	SnapshotId	SnapshotId	`json:"snapshotId"`
	
}

// ReleaseSnapshot returns
type ReleaseSnapshotReturns struct {
	
}

// ReplaySnapshot parameters
type ReplaySnapshotParams struct {
	
	// The id of the layer snapshot.
	SnapshotId	SnapshotId	`json:"snapshotId"`
	
	// The first step to replay from (replay from the very start if not specified).
	FromStep	int	`json:"fromStep"`
	
	// The last step to replay to (replay till the end if not specified).
	ToStep	int	`json:"toStep"`
	
	// The scale to apply while replaying (defaults to 1).
	Scale	float64	`json:"scale"`
	
}

// ReplaySnapshot returns
type ReplaySnapshotReturns struct {
	
	// A data: URL for resulting image.
	DataURL	string	`json:"dataURL"`
	
}

// SnapshotCommandLog parameters
type SnapshotCommandLogParams struct {
	
	// The id of the layer snapshot.
	SnapshotId	SnapshotId	`json:"snapshotId"`
	
}

// SnapshotCommandLog returns
type SnapshotCommandLogReturns struct {
	
	// The array of canvas function calls.
	CommandLog	[]interface{}	`json:"commandLog"`
	
}

