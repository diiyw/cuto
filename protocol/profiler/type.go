package profiler

import (
	"github.com/diiyw/cuto/protocol/runtime"
)

// Profile node. Holds callsite information, execution statistics and child nodes.
type ProfileNode  struct {

	// Unique id of the node.
	Id	int	`json:"id"`

	// Function location.
	CallFrame	runtime.CallFrame	`json:"callFrame"`

	// Number of samples where this node was on top of the call stack.
	HitCount	int	`json:"hitCount,omitempty"`

	// Child node ids.
	Children	[]int	`json:"children,omitempty"`

	// The reason of being not optimized. The function may be deoptimized or marked as don't
	// optimize.
	DeoptReason	string	`json:"deoptReason,omitempty"`

	// An array of source position ticks.
	PositionTicks	[]*PositionTickInfo	`json:"positionTicks,omitempty"`
}

// Profile.
type Profile  struct {

	// The list of profile nodes. First item is the root node.
	Nodes	[]*ProfileNode	`json:"nodes"`

	// Profiling start timestamp in microseconds.
	StartTime	float64	`json:"startTime"`

	// Profiling end timestamp in microseconds.
	EndTime	float64	`json:"endTime"`

	// Ids of samples top nodes.
	Samples	[]int	`json:"samples,omitempty"`

	// Time intervals between adjacent samples in microseconds. The first delta is relative to the
	// profile startTime.
	TimeDeltas	[]int	`json:"timeDeltas,omitempty"`
}

// Specifies a number of samples attributed to a certain source position.
type PositionTickInfo  struct {

	// Source line number (1-based).
	Line	int	`json:"line"`

	// Number of samples attributed to the source line.
	Ticks	int	`json:"ticks"`
}

// Coverage data for a source range.
type CoverageRange  struct {

	// JavaScript script source offset for the range start.
	StartOffset	int	`json:"startOffset"`

	// JavaScript script source offset for the range end.
	EndOffset	int	`json:"endOffset"`

	// Collected execution count of the source range.
	Count	int	`json:"count"`
}

// Coverage data for a JavaScript function.
type FunctionCoverage  struct {

	// JavaScript function name.
	FunctionName	string	`json:"functionName"`

	// Source ranges inside the function with coverage data.
	Ranges	[]*CoverageRange	`json:"ranges"`

	// Whether coverage data for this function has block granularity.
	IsBlockCoverage	bool	`json:"isBlockCoverage"`
}

// Coverage data for a JavaScript script.
type ScriptCoverage  struct {

	// JavaScript script id.
	ScriptId	runtime.ScriptId	`json:"scriptId"`

	// JavaScript script name or url.
	Url	string	`json:"url"`

	// Functions contained in the script that has coverage data.
	Functions	[]*FunctionCoverage	`json:"functions"`
}

// Describes a type collected during runtime.
type TypeObject  struct {

	// Name of a type collected with type profiling.
	Name	string	`json:"name"`
}

// Source offset and types for a parameter or return value.
type TypeProfileEntry  struct {

	// Source offset of the parameter or end of function for return values.
	Offset	int	`json:"offset"`

	// The types for this parameter or return value.
	Types	[]*TypeObject	`json:"types"`
}

// Type profile data collected during runtime for a JavaScript script.
type ScriptTypeProfile  struct {

	// JavaScript script id.
	ScriptId	runtime.ScriptId	`json:"scriptId"`

	// JavaScript script name or url.
	Url	string	`json:"url"`

	// Type profile entries for parameters and return values of the functions in the script.
	Entries	[]*TypeProfileEntry	`json:"entries"`
}
