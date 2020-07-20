package target

// Activates (focuses) the target.
const ActivateTarget = "Target.activateTarget"

type ActivateTargetParams struct {

	// 
	TargetId 	TargetID	`json:"targetId"`
}

type ActivateTargetResult struct {

}

// Attaches to the target with given id.
const AttachToTarget = "Target.attachToTarget"

type AttachToTargetParams struct {

	// 
	TargetId 	TargetID	`json:"targetId"`

	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten 	bool	`json:"flatten,omitempty"`
}

type AttachToTargetResult struct {

	// Id assigned to the session.
	SessionId 	SessionID	`json:"sessionId"`
}

// Attaches to the browser target, only uses flat sessionId mode.
const AttachToBrowserTarget = "Target.attachToBrowserTarget"

type AttachToBrowserTargetParams struct {
}

type AttachToBrowserTargetResult struct {

	// Id assigned to the session.
	SessionId 	SessionID	`json:"sessionId"`
}

// Closes the target. If the target is a page that gets closed too.
const CloseTarget = "Target.closeTarget"

type CloseTargetParams struct {

	// 
	TargetId 	TargetID	`json:"targetId"`
}

type CloseTargetResult struct {

	// 
	Success 	bool	`json:"success"`
}

// Inject object to the target's main frame that provides a communication
// channel with browser target.
// 
// Injected object will be available as `window[bindingName]`.
// 
// The object has the follwing API:
// - `binding.send(json)` - a method to send messages over the remote debugging protocol
// - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
const ExposeDevToolsProtocol = "Target.exposeDevToolsProtocol"

type ExposeDevToolsProtocolParams struct {

	// 
	TargetId 	TargetID	`json:"targetId"`

	// Binding name, 'cdp' if not specified.
	BindingName 	string	`json:"bindingName,omitempty"`
}

type ExposeDevToolsProtocolResult struct {

}

// Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than
// one.
const CreateBrowserContext = "Target.createBrowserContext"

type CreateBrowserContextParams struct {
}

type CreateBrowserContextResult struct {

	// The id of the context created.
	BrowserContextId 	BrowserContextID	`json:"browserContextId"`
}

// Returns all browser contexts created with `Target.createBrowserContext` method.
const GetBrowserContexts = "Target.getBrowserContexts"

type GetBrowserContextsParams struct {
}

type GetBrowserContextsResult struct {

	// An array of browser context ids.
	BrowserContextIds 	[]*BrowserContextID	`json:"browserContextIds"`
}

// Creates a new page.
const CreateTarget = "Target.createTarget"

type CreateTargetParams struct {

	// The initial URL the page will be navigated to.
	Url 	string	`json:"url"`

	// Frame width in DIP (headless chrome only).
	Width 	int	`json:"width,omitempty"`

	// Frame height in DIP (headless chrome only).
	Height 	int	`json:"height,omitempty"`

	// The browser context to create the page in.
	BrowserContextId 	BrowserContextID	`json:"browserContextId,omitempty"`

	// Whether BeginFrames for this target will be controlled via DevTools (headless chrome only,
	// not supported on MacOS yet, false by default).
	EnableBeginFrameControl 	bool	`json:"enableBeginFrameControl,omitempty"`

	// Whether to create a new Window or Tab (chrome-only, false by default).
	NewWindow 	bool	`json:"newWindow,omitempty"`

	// Whether to create the target in background or foreground (chrome-only,
	// false by default).
	Background 	bool	`json:"background,omitempty"`
}

type CreateTargetResult struct {

	// The id of the page opened.
	TargetId 	TargetID	`json:"targetId"`
}

// Detaches session with given id.
const DetachFromTarget = "Target.detachFromTarget"

type DetachFromTargetParams struct {

	// Session to detach.
	SessionId 	SessionID	`json:"sessionId,omitempty"`

	// Deprecated.
	TargetId 	TargetID	`json:"targetId,omitempty"`
}

type DetachFromTargetResult struct {

}

// Deletes a BrowserContext. All the belonging pages will be closed without calling their
// beforeunload hooks.
const DisposeBrowserContext = "Target.disposeBrowserContext"

type DisposeBrowserContextParams struct {

	// 
	BrowserContextId 	BrowserContextID	`json:"browserContextId"`
}

type DisposeBrowserContextResult struct {

}

// Returns information about a target.
const GetTargetInfo = "Target.getTargetInfo"

type GetTargetInfoParams struct {

	// 
	TargetId 	TargetID	`json:"targetId,omitempty"`
}

type GetTargetInfoResult struct {

	// 
	TargetInfo 	TargetInfo	`json:"targetInfo"`
}

// Retrieves a list of available targets.
const GetTargets = "Target.getTargets"

type GetTargetsParams struct {
}

type GetTargetsResult struct {

	// The list of targets.
	TargetInfos 	[]*TargetInfo	`json:"targetInfos"`
}

// Sends protocol message over session with given id.
// Consider using flat mode instead; see commands attachToTarget, setAutoAttach,
// and crbug.com/991325.
const SendMessageToTarget = "Target.sendMessageToTarget"

type SendMessageToTargetParams struct {

	// 
	Message 	string	`json:"message"`

	// Identifier of the session.
	SessionId 	SessionID	`json:"sessionId,omitempty"`

	// Deprecated.
	TargetId 	TargetID	`json:"targetId,omitempty"`
}

type SendMessageToTargetResult struct {

}

// Controls whether to automatically attach to new targets which are considered to be related to
// this one. When turned on, attaches to all existing related targets as well. When turned off,
// automatically detaches from all currently attached targets.
const SetAutoAttach = "Target.setAutoAttach"

type SetAutoAttachParams struct {

	// Whether to auto-attach to related targets.
	AutoAttach 	bool	`json:"autoAttach"`

	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart 	bool	`json:"waitForDebuggerOnStart"`

	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten 	bool	`json:"flatten,omitempty"`

	// Auto-attach to the targets created via window.open from current target.
	WindowOpen 	bool	`json:"windowOpen,omitempty"`
}

type SetAutoAttachResult struct {

}

// Controls whether to discover available targets and notify via
// `targetCreated/targetInfoChanged/targetDestroyed` events.
const SetDiscoverTargets = "Target.setDiscoverTargets"

type SetDiscoverTargetsParams struct {

	// Whether to discover available targets.
	Discover 	bool	`json:"discover"`
}

type SetDiscoverTargetsResult struct {

}

// Enables target discovery for the specified locations, when `setDiscoverTargets` was set to
// `true`.
const SetRemoteLocations = "Target.setRemoteLocations"

type SetRemoteLocationsParams struct {

	// List of remote locations.
	Locations 	[]*RemoteLocation	`json:"locations"`
}

type SetRemoteLocationsResult struct {

}