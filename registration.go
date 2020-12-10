package lsp

// Registration contains general parameters to register for a capability.
type Registration struct {
	// The id used to register the request. The id can be used to deregister
	// the request again.
	ID string `json:"id"`

	// The method/capability to register for.
	Method string `json:"method"`

	// Options necessary for the registration.
	RegisterOptions interface{} `json:"registerOptions,omitempty"`
}

// RegistrationParams is the req payload sent in a `client/registerCapability`
// request.
type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

// Unregistration contains general parameters to unregister a capability.
type Unregistration struct {
	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	ID string `json:"id"`

	// The method/capability to unregister for.
	Method string `json:"method"`
}

// UnregistrationParams is the payload sent in a `client/unregisterCapability`
// request.
type UnregistrationParams struct {
	Unregistrations []Unregistration `json:"unregisterations"`
}
