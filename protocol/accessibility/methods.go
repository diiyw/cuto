package accessibility

import (

	"github.com/diiyw/goc/protocol/dom"

	"github.com/diiyw/goc/protocol/runtime"

)
const (
	
	// Disables the accessibility domain.
	Disable = "Accessibility.disable"
	
	// Enables the accessibility domain which causes `AXNodeId`s to remain consistent between method calls.
	// This turns on accessibility for the page, which can impact performance until accessibility is disabled.
	Enable = "Accessibility.enable"
	
	// Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
	GetPartialAXTree = "Accessibility.getPartialAXTree"
	
	// Fetches the entire accessibility tree
	GetFullAXTree = "Accessibility.getFullAXTree"
	
)

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

// GetPartialAXTree parameters
type GetPartialAXTreeParams struct {
	
	// Identifier of the node to get the partial accessibility tree for.
	NodeId	dom.NodeId	`json:"nodeId"`
	
	// Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeId	dom.BackendNodeId	`json:"backendNodeId"`
	
	// JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectId	runtime.RemoteObjectId	`json:"objectId"`
	
	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives	bool	`json:"fetchRelatives"`
	
}

// GetPartialAXTree returns
type GetPartialAXTreeReturns struct {
	
	// The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and
	// children, if requested.
	Nodes	[]AXNode	`json:"nodes"`
	
}

// GetFullAXTree parameters
type GetFullAXTreeParams struct {
	
}

// GetFullAXTree returns
type GetFullAXTreeReturns struct {
	
	
	Nodes	[]AXNode	`json:"nodes"`
	
}

