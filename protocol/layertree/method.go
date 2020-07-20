package layertree

import (
	"github.com/diiyw/cuto/protocol/cdp"
)


// Provides the reasons why the given layer was composited.
const CompositingReasons = "LayerTree.compositingReasons"

type CompositingReasonsParams struct {

	// The id of the layer for which we want to get the reasons it was composited.
	LayerId 	LayerId	`json:"layerId"`
}

type CompositingReasonsResult struct {

	// A list of strings specifying reasons for the given layer to become composited.
	CompositingReasons 	[]string	`json:"compositingReasons"`
}

// Disables compositing tree inspection.
const Disable = "LayerTree.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables compositing tree inspection.
const Enable = "LayerTree.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Returns the snapshot identifier.
const LoadSnapshot = "LayerTree.loadSnapshot"

type LoadSnapshotParams struct {

	// An array of tiles composing the snapshot.
	Tiles 	[]*PictureTile	`json:"tiles"`
}

type LoadSnapshotResult struct {

	// The id of the snapshot.
	SnapshotId 	SnapshotId	`json:"snapshotId"`
}

// Returns the layer snapshot identifier.
const MakeSnapshot = "LayerTree.makeSnapshot"

type MakeSnapshotParams struct {

	// The id of the layer.
	LayerId 	LayerId	`json:"layerId"`
}

type MakeSnapshotResult struct {

	// The id of the layer snapshot.
	SnapshotId 	SnapshotId	`json:"snapshotId"`
}

// 
const ProfileSnapshot = "LayerTree.profileSnapshot"

type ProfileSnapshotParams struct {

	// The id of the layer snapshot.
	SnapshotId 	SnapshotId	`json:"snapshotId"`

	// The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount 	int	`json:"minRepeatCount,omitempty"`

	// The minimum duration (in seconds) to replay the snapshot.
	MinDuration 	float64	`json:"minDuration,omitempty"`

	// The clip rectangle to apply when replaying the snapshot.
	ClipRect 	cdp.Rect	`json:"clipRect,omitempty"`
}

type ProfileSnapshotResult struct {

	// The array of paint profiles, one per run.
	Timings 	[]*PaintProfile	`json:"timings"`
}

// Releases layer snapshot captured by the back-end.
const ReleaseSnapshot = "LayerTree.releaseSnapshot"

type ReleaseSnapshotParams struct {

	// The id of the layer snapshot.
	SnapshotId 	SnapshotId	`json:"snapshotId"`
}

type ReleaseSnapshotResult struct {

}

// Replays the layer snapshot and returns the resulting bitmap.
const ReplaySnapshot = "LayerTree.replaySnapshot"

type ReplaySnapshotParams struct {

	// The id of the layer snapshot.
	SnapshotId 	SnapshotId	`json:"snapshotId"`

	// The first step to replay from (replay from the very start if not specified).
	FromStep 	int	`json:"fromStep,omitempty"`

	// The last step to replay to (replay till the end if not specified).
	ToStep 	int	`json:"toStep,omitempty"`

	// The scale to apply while replaying (defaults to 1).
	Scale 	float64	`json:"scale,omitempty"`
}

type ReplaySnapshotResult struct {

	// A data: URL for resulting image.
	DataURL 	string	`json:"dataURL"`
}

// Replays the layer snapshot and returns canvas log.
const SnapshotCommandLog = "LayerTree.snapshotCommandLog"

type SnapshotCommandLogParams struct {

	// The id of the layer snapshot.
	SnapshotId 	SnapshotId	`json:"snapshotId"`
}

type SnapshotCommandLogResult struct {

	// The array of canvas function calls.
	CommandLog 	[]interface{}	`json:"commandLog"`
}