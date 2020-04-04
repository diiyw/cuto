package domsnapshot

// Disables DOM snapshot agent for the given page.
const Disable = "DOMSnapshot.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Enables DOM snapshot agent for the given page.
const Enable = "DOMSnapshot.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Returns a document snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as layout and
// white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
const GetSnapshot = "DOMSnapshot.getSnapshot"

type GetSnapshotParams struct {

	// Whitelist of computed styles to return.
	ComputedStyleWhitelist 	[]string	`json:"computedStyleWhitelist"`

	// Whether or not to retrieve details of DOM listeners (default false).
	IncludeEventListeners 	bool	`json:"includeEventListeners"`

	// Whether to determine and include the paint order index of LayoutTreeNodes (default false).
	IncludePaintOrder 	bool	`json:"includePaintOrder"`

	// Whether to include UA shadow tree in the snapshot (default false).
	IncludeUserAgentShadowTree 	bool	`json:"includeUserAgentShadowTree"`
}

type GetSnapshotResult struct {

	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	DomNodes 	[]*DOMNode	`json:"domNodes"`
	// The nodes in the layout tree.
	LayoutTreeNodes 	[]*LayoutTreeNode	`json:"layoutTreeNodes"`
	// Whitelisted ComputedStyle properties for each node in the layout tree.
	ComputedStyles 	[]*ComputedStyle	`json:"computedStyles"`
}

// Returns a document snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as layout and
// white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
const CaptureSnapshot = "DOMSnapshot.captureSnapshot"

type CaptureSnapshotParams struct {

	// Whitelist of computed styles to return.
	ComputedStyles 	[]string	`json:"computedStyles"`

	// Whether to include layout object paint orders into the snapshot.
	IncludePaintOrder 	bool	`json:"includePaintOrder"`

	// Whether to include DOM rectangles (offsetRects, clientRects, scrollRects) into the snapshot
	IncludeDOMRects 	bool	`json:"includeDOMRects"`
}

type CaptureSnapshotResult struct {

	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Documents 	[]*DocumentSnapshot	`json:"documents"`
	// Shared string table that all string properties refer to with indexes.
	Strings 	[]string	`json:"strings"`
}