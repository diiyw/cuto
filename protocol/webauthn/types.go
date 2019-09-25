package webauthn


// 
type AuthenticatorId string	

// 
type AuthenticatorProtocol string	

// 
type AuthenticatorTransport string	

// 
type VirtualAuthenticatorOptions struct {
	
	
	
	Protocol	AuthenticatorProtocol	`json:"protocol"`
	
	
	
	Transport	AuthenticatorTransport	`json:"transport"`
	
	
	
	HasResidentKey	bool	`json:"hasResidentKey"`
	
	
	
	HasUserVerification	bool	`json:"hasUserVerification"`
	
	// If set to true, tests of user presence will succeed immediately.
	// Otherwise, they will not be resolved. Defaults to true.
	
	AutomaticPresenceSimulation	bool	`json:"automaticPresenceSimulation"`
	
}	

// 
type Credential struct {
	
	
	
	CredentialId	string	`json:"credentialId"`
	
	
	
	IsResidentCredential	bool	`json:"isResidentCredential"`
	
	// Relying Party ID the credential is scoped to. Must be set when adding a
	// credential.
	
	RpId	string	`json:"rpId"`
	
	// The ECDSA P-256 private key in PKCS#8 format.
	
	PrivateKey	string	`json:"privateKey"`
	
	// An opaque byte sequence with a maximum size of 64 bytes mapping the
	// credential to a specific user.
	
	UserHandle	string	`json:"userHandle"`
	
	// Signature counter. This is incremented by one for each successful
	// assertion.
	// See https://w3c.github.io/webauthn/#signature-counter
	
	SignCount	int	`json:"signCount"`
	
}	

