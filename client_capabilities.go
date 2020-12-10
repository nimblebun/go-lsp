package lsp

// DidChangeConfigurationClientCapabilities contains information about the
// client's configuration change capabilities.
type DidChangeConfigurationClientCapabilities struct {
	// Did change configuration notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DidChangeWatchedFilesClientCapabilities contains information about the
// client's file watcher change capabilities.
type DidChangeWatchedFilesClientCapabilities struct {
	// Did change watched files notification supports dynamic registration.
	// Please note that the current protocol doesn't support static
	// configuration for file changes from the server side.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
