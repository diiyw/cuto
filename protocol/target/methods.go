package target

const (
	
	// Activates (focuses) the target.
	ActivateTarget = "Target.activateTarget"
	
	// Attaches to the target with given id.
	AttachToTarget = "Target.attachToTarget"
	
	// Attaches to the browser target, only uses flat sessionId mode.
	AttachToBrowserTarget = "Target.attachToBrowserTarget"
	
	// Closes the target. If the target is a page that gets closed too.
	CloseTarget = "Target.closeTarget"
	
	// Inject object to the target's main frame that provides a communication
	// channel with browser target.
	// 
	// Injected object will be available as `window[bindingName]`.
	// 
	// The object has the follwing API:
	// - `binding.send(json)` - a method to send messages over the remote debugging protocol
	// - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
	ExposeDevToolsProtocol = "Target.exposeDevToolsProtocol"
	
	// Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than
	// one.
	CreateBrowserContext = "Target.createBrowserContext"
	
	// Returns all browser contexts created with `Target.createBrowserContext` method.
	GetBrowserContexts = "Target.getBrowserContexts"
	
	// Creates a new page.
	CreateTarget = "Target.createTarget"
	
	// Detaches session with given id.
	DetachFromTarget = "Target.detachFromTarget"
	
	// Deletes a BrowserContext. All the belonging pages will be closed without calling their
	// beforeunload hooks.
	DisposeBrowserContext = "Target.disposeBrowserContext"
	
	// Returns information about a target.
	GetTargetInfo = "Target.getTargetInfo"
	
	// Retrieves a list of available targets.
	GetTargets = "Target.getTargets"
	
	// Sends protocol message over session with given id.
	SendMessageToTarget = "Target.sendMessageToTarget"
	
	// Controls whether to automatically attach to new targets which are considered to be related to
	// this one. When turned on, attaches to all existing related targets as well. When turned off,
	// automatically detaches from all currently attached targets.
	SetAutoAttach = "Target.setAutoAttach"
	
	// Controls whether to discover available targets and notify via
	// `targetCreated/targetInfoChanged/targetDestroyed` events.
	SetDiscoverTargets = "Target.setDiscoverTargets"
	
	// Enables target discovery for the specified locations, when `setDiscoverTargets` was set to
	// `true`.
	SetRemoteLocations = "Target.setRemoteLocations"
	
)

// ActivateTarget parameters
type ActivateTargetParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
}

// ActivateTarget returns
type ActivateTargetReturns struct {
	
}

// AttachToTarget parameters
type AttachToTargetParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	Flatten	bool	`json:"flatten"`
	
}

// AttachToTarget returns
type AttachToTargetReturns struct {
	
	// Id assigned to the session.
	SessionId	SessionID	`json:"sessionId"`
	
}

// AttachToBrowserTarget parameters
type AttachToBrowserTargetParams struct {
	
}

// AttachToBrowserTarget returns
type AttachToBrowserTargetReturns struct {
	
	// Id assigned to the session.
	SessionId	SessionID	`json:"sessionId"`
	
}

// CloseTarget parameters
type CloseTargetParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
}

// CloseTarget returns
type CloseTargetReturns struct {
	
	
	Success	bool	`json:"success"`
	
}

// ExposeDevToolsProtocol parameters
type ExposeDevToolsProtocolParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
	// Binding name, 'cdp' if not specified.
	BindingName	string	`json:"bindingName"`
	
}

// ExposeDevToolsProtocol returns
type ExposeDevToolsProtocolReturns struct {
	
}

// CreateBrowserContext parameters
type CreateBrowserContextParams struct {
	
}

// CreateBrowserContext returns
type CreateBrowserContextReturns struct {
	
	// The id of the context created.
	BrowserContextId	BrowserContextID	`json:"browserContextId"`
	
}

// GetBrowserContexts parameters
type GetBrowserContextsParams struct {
	
}

// GetBrowserContexts returns
type GetBrowserContextsReturns struct {
	
	// An array of browser context ids.
	BrowserContextIds	[]BrowserContextID	`json:"browserContextIds"`
	
}

// CreateTarget parameters
type CreateTargetParams struct {
	
	// The initial URL the page will be navigated to.
	Url	string	`json:"url"`
	
	// Frame width in DIP (headless chrome only).
	Width	int	`json:"width"`
	
	// Frame height in DIP (headless chrome only).
	Height	int	`json:"height"`
	
	// The browser context to create the page in.
	BrowserContextId	BrowserContextID	`json:"browserContextId"`
	
	// Whether BeginFrames for this target will be controlled via DevTools (headless chrome only,
	// not supported on MacOS yet, false by default).
	EnableBeginFrameControl	bool	`json:"enableBeginFrameControl"`
	
}

// CreateTarget returns
type CreateTargetReturns struct {
	
	// The id of the page opened.
	TargetId	TargetID	`json:"targetId"`
	
}

// DetachFromTarget parameters
type DetachFromTargetParams struct {
	
	// Session to detach.
	SessionId	SessionID	`json:"sessionId"`
	
	// Deprecated.
	TargetId	TargetID	`json:"targetId"`
	
}

// DetachFromTarget returns
type DetachFromTargetReturns struct {
	
}

// DisposeBrowserContext parameters
type DisposeBrowserContextParams struct {
	
	
	BrowserContextId	BrowserContextID	`json:"browserContextId"`
	
}

// DisposeBrowserContext returns
type DisposeBrowserContextReturns struct {
	
}

// GetTargetInfo parameters
type GetTargetInfoParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
}

// GetTargetInfo returns
type GetTargetInfoReturns struct {
	
	
	TargetInfo	TargetInfo	`json:"targetInfo"`
	
}

// GetTargets parameters
type GetTargetsParams struct {
	
}

// GetTargets returns
type GetTargetsReturns struct {
	
	// The list of targets.
	TargetInfos	[]TargetInfo	`json:"targetInfos"`
	
}

// SendMessageToTarget parameters
type SendMessageToTargetParams struct {
	
	
	Message	string	`json:"message"`
	
	// Identifier of the session.
	SessionId	SessionID	`json:"sessionId"`
	
	// Deprecated.
	TargetId	TargetID	`json:"targetId"`
	
}

// SendMessageToTarget returns
type SendMessageToTargetReturns struct {
	
}

// SetAutoAttach parameters
type SetAutoAttachParams struct {
	
	// Whether to auto-attach to related targets.
	AutoAttach	bool	`json:"autoAttach"`
	
	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart	bool	`json:"waitForDebuggerOnStart"`
	
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	Flatten	bool	`json:"flatten"`
	
}

// SetAutoAttach returns
type SetAutoAttachReturns struct {
	
}

// SetDiscoverTargets parameters
type SetDiscoverTargetsParams struct {
	
	// Whether to discover available targets.
	Discover	bool	`json:"discover"`
	
}

// SetDiscoverTargets returns
type SetDiscoverTargetsReturns struct {
	
}

// SetRemoteLocations parameters
type SetRemoteLocationsParams struct {
	
	// List of remote locations.
	Locations	[]RemoteLocation	`json:"locations"`
	
}

// SetRemoteLocations returns
type SetRemoteLocationsReturns struct {
	
}

