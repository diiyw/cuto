package webauthn

// Enable the WebAuthn domain and start intercepting credential storage and
// retrieval with a virtual authenticator.
const Enable = "WebAuthn.enable"

type EnableParams struct {
}

type EnableResult struct {

}

// Disable the WebAuthn domain.
const Disable = "WebAuthn.disable"

type DisableParams struct {
}

type DisableResult struct {

}

// Creates and adds a virtual authenticator.
const AddVirtualAuthenticator = "WebAuthn.addVirtualAuthenticator"

type AddVirtualAuthenticatorParams struct {

	// 
	Options 	VirtualAuthenticatorOptions	`json:"options"`
}

type AddVirtualAuthenticatorResult struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`
}

// Removes the given authenticator.
const RemoveVirtualAuthenticator = "WebAuthn.removeVirtualAuthenticator"

type RemoveVirtualAuthenticatorParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`
}

type RemoveVirtualAuthenticatorResult struct {

}

// Adds the credential to the specified authenticator.
const AddCredential = "WebAuthn.addCredential"

type AddCredentialParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`

	// 
	Credential 	Credential	`json:"credential"`
}

type AddCredentialResult struct {

}

// Returns a single credential stored in the given virtual authenticator that
// matches the credential ID.
const GetCredential = "WebAuthn.getCredential"

type GetCredentialParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`

	// 
	CredentialId 	[]byte	`json:"credentialId"`
}

type GetCredentialResult struct {

	// 
	Credential 	Credential	`json:"credential"`
}

// Returns all the credentials stored in the given virtual authenticator.
const GetCredentials = "WebAuthn.getCredentials"

type GetCredentialsParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`
}

type GetCredentialsResult struct {

	// 
	Credentials 	[]*Credential	`json:"credentials"`
}

// Removes a credential from the authenticator.
const RemoveCredential = "WebAuthn.removeCredential"

type RemoveCredentialParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`

	// 
	CredentialId 	[]byte	`json:"credentialId"`
}

type RemoveCredentialResult struct {

}

// Clears all the credentials from the specified device.
const ClearCredentials = "WebAuthn.clearCredentials"

type ClearCredentialsParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`
}

type ClearCredentialsResult struct {

}

// Sets whether User Verification succeeds or fails for an authenticator.
// The default is true.
const SetUserVerified = "WebAuthn.setUserVerified"

type SetUserVerifiedParams struct {

	// 
	AuthenticatorId 	AuthenticatorId	`json:"authenticatorId"`

	// 
	IsUserVerified 	bool	`json:"isUserVerified"`
}

type SetUserVerifiedResult struct {

}