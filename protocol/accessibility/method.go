package accessibility

// Disables the accessibility domain.
const Disable = "Accessibility.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables the accessibility domain which causes `AXNodeId`s to remain consistent between method calls.
// This turns on accessibility for the page, which can impact performance until accessibility is disabled.
const Enable = "Accessibility.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
const GetPartialAXTree = "Accessibility.getPartialAXTree"

type GetPartialAXTreeParams struct {

	// Identifier of the node to get the partial accessibility tree for.
	NodeId 	interface{}	`json:"nodeId"`

	// Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeId 	interface{}	`json:"backendNodeId"`

	// JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectId 	interface{}	`json:"objectId"`

	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives 	bool	`json:"fetchRelatives"`
}

type GetPartialAXTreeResult struct {

	// The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and
	// children, if requested.
	Nodes 	[]*AXNode	`json:"nodes"`
}

// Fetches the entire accessibility tree
const GetFullAXTree = "Accessibility.getFullAXTree"

type GetFullAXTreeParams struct {
}

type GetFullAXTreeResult struct {

	// 
	Nodes 	[]*AXNode	`json:"nodes"`
}