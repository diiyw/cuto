package target

// Issued when attached to target because of auto-attach or `attachToTarget` command.
const AttachedToTargetEvent = "Target.attachedToTarget"
type AttachedToTargetParams struct {

	// Identifier assigned to the session used to send/receive messages.
	SessionId 	SessionID
	// 
	TargetInfo 	TargetInfo
	// 
	WaitingForDebugger 	bool}



// Issued when detached from target for any reason (including `detachFromTarget` command). Can be
// issued multiple times per target if multiple sessions have been attached to it.
const DetachedFromTargetEvent = "Target.detachedFromTarget"
type DetachedFromTargetParams struct {

	// Detached session identifier.
	SessionId 	SessionID
	// Deprecated.
	TargetId 	TargetID}



// Notifies about a new protocol message received from the session (as reported in
// `attachedToTarget` event).
const ReceivedMessageFromTargetEvent = "Target.receivedMessageFromTarget"
type ReceivedMessageFromTargetParams struct {

	// Identifier of a session which sends a message.
	SessionId 	SessionID
	// 
	Message 	string
	// Deprecated.
	TargetId 	TargetID}



// Issued when a possible inspection target is created.
const TargetCreatedEvent = "Target.targetCreated"
type TargetCreatedParams struct {

	// 
	TargetInfo 	TargetInfo}



// Issued when a target is destroyed.
const TargetDestroyedEvent = "Target.targetDestroyed"
type TargetDestroyedParams struct {

	// 
	TargetId 	TargetID}



// Issued when a target has crashed.
const TargetCrashedEvent = "Target.targetCrashed"
type TargetCrashedParams struct {

	// 
	TargetId 	TargetID
	// Termination status type.
	Status 	string
	// Termination error code.
	ErrorCode 	int}



// Issued when some information about a target has changed. This only happens between
// `targetCreated` and `targetDestroyed`.
const TargetInfoChangedEvent = "Target.targetInfoChanged"
type TargetInfoChangedParams struct {

	// 
	TargetInfo 	TargetInfo}

