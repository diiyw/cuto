package webauthn

const (
	
	// Enable the WebAuthn domain and start intercepting credential storage and
	// retrieval with a virtual authenticator.
	Enable = "WebAuthn.enable"
	
	// Disable the WebAuthn domain.
	Disable = "WebAuthn.disable"
	
	// Creates and adds a virtual authenticator.
	AddVirtualAuthenticator = "WebAuthn.addVirtualAuthenticator"
	
	// Removes the given authenticator.
	RemoveVirtualAuthenticator = "WebAuthn.removeVirtualAuthenticator"
	
	// Adds the credential to the specified authenticator.
	AddCredential = "WebAuthn.addCredential"
	
	// Returns a single credential stored in the given virtual authenticator that
	// matches the credential ID.
	GetCredential = "WebAuthn.getCredential"
	
	// Returns all the credentials stored in the given virtual authenticator.
	GetCredentials = "WebAuthn.getCredentials"
	
	// Removes a credential from the authenticator.
	RemoveCredential = "WebAuthn.removeCredential"
	
	// Clears all the credentials from the specified device.
	ClearCredentials = "WebAuthn.clearCredentials"
	
	// Sets whether User Verification succeeds or fails for an authenticator.
	// The default is true.
	SetUserVerified = "WebAuthn.setUserVerified"
	
)

// Enable parameters
type EnableParams struct {
	
}

// Enable returns
type EnableReturns struct {
	
}

// Disable parameters
type DisableParams struct {
	
}

// Disable returns
type DisableReturns struct {
	
}

// AddVirtualAuthenticator parameters
type AddVirtualAuthenticatorParams struct {
	
	
	Options	VirtualAuthenticatorOptions	`json:"options"`
	
}

// AddVirtualAuthenticator returns
type AddVirtualAuthenticatorReturns struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
}

// RemoveVirtualAuthenticator parameters
type RemoveVirtualAuthenticatorParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
}

// RemoveVirtualAuthenticator returns
type RemoveVirtualAuthenticatorReturns struct {
	
}

// AddCredential parameters
type AddCredentialParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
	
	Credential	Credential	`json:"credential"`
	
}

// AddCredential returns
type AddCredentialReturns struct {
	
}

// GetCredential parameters
type GetCredentialParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
	
	CredentialId	string	`json:"credentialId"`
	
}

// GetCredential returns
type GetCredentialReturns struct {
	
	
	Credential	Credential	`json:"credential"`
	
}

// GetCredentials parameters
type GetCredentialsParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
}

// GetCredentials returns
type GetCredentialsReturns struct {
	
	
	Credentials	[]Credential	`json:"credentials"`
	
}

// RemoveCredential parameters
type RemoveCredentialParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
	
	CredentialId	string	`json:"credentialId"`
	
}

// RemoveCredential returns
type RemoveCredentialReturns struct {
	
}

// ClearCredentials parameters
type ClearCredentialsParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
}

// ClearCredentials returns
type ClearCredentialsReturns struct {
	
}

// SetUserVerified parameters
type SetUserVerifiedParams struct {
	
	
	AuthenticatorId	AuthenticatorId	`json:"authenticatorId"`
	
	
	IsUserVerified	bool	`json:"isUserVerified"`
	
}

// SetUserVerified returns
type SetUserVerifiedReturns struct {
	
}

