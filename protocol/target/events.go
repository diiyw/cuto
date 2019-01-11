package target

const (
	
	// Issued when attached to target because of auto-attach or `attachToTarget` command.
	AttachedToTargetEvent = "Target.attachedToTarget"
	
	// Issued when detached from target for any reason (including `detachFromTarget` command). Can be
	// issued multiple times per target if multiple sessions have been attached to it.
	DetachedFromTargetEvent = "Target.detachedFromTarget"
	
	// Notifies about a new protocol message received from the session (as reported in
	// `attachedToTarget` event).
	ReceivedMessageFromTargetEvent = "Target.receivedMessageFromTarget"
	
	// Issued when a possible inspection target is created.
	TargetCreatedEvent = "Target.targetCreated"
	
	// Issued when a target is destroyed.
	TargetDestroyedEvent = "Target.targetDestroyed"
	
	// Issued when a target has crashed.
	TargetCrashedEvent = "Target.targetCrashed"
	
	// Issued when some information about a target has changed. This only happens between
	// `targetCreated` and `targetDestroyed`.
	TargetInfoChangedEvent = "Target.targetInfoChanged"
	
)

// Issued when attached to target because of auto-attach or `attachToTarget` command.
type AttachedToTargetParams struct {
	
	// Identifier assigned to the session used to send/receive messages.
	SessionId	SessionID	`json:"sessionId"`
	
	
	TargetInfo	TargetInfo	`json:"targetInfo"`
	
	
	WaitingForDebugger	bool	`json:"waitingForDebugger"`
	
}

// Issued when detached from target for any reason (including `detachFromTarget` command). Can be
	// issued multiple times per target if multiple sessions have been attached to it.
type DetachedFromTargetParams struct {
	
	// Detached session identifier.
	SessionId	SessionID	`json:"sessionId"`
	
	// Deprecated.
	TargetId	TargetID	`json:"targetId"`
	
}

// Notifies about a new protocol message received from the session (as reported in
	// `attachedToTarget` event).
type ReceivedMessageFromTargetParams struct {
	
	// Identifier of a session which sends a message.
	SessionId	SessionID	`json:"sessionId"`
	
	
	Message	string	`json:"message"`
	
	// Deprecated.
	TargetId	TargetID	`json:"targetId"`
	
}

// Issued when a possible inspection target is created.
type TargetCreatedParams struct {
	
	
	TargetInfo	TargetInfo	`json:"targetInfo"`
	
}

// Issued when a target is destroyed.
type TargetDestroyedParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
}

// Issued when a target has crashed.
type TargetCrashedParams struct {
	
	
	TargetId	TargetID	`json:"targetId"`
	
	// Termination status type.
	Status	string	`json:"status"`
	
	// Termination error code.
	ErrorCode	int	`json:"errorCode"`
	
}

// Issued when some information about a target has changed. This only happens between
	// `targetCreated` and `targetDestroyed`.
type TargetInfoChangedParams struct {
	
	
	TargetInfo	TargetInfo	`json:"targetInfo"`
	
}

