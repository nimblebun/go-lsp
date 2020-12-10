package lsp

// DidChangeConfigurationParams contains the parameters contained in a
// `workspace/didChangeConfiguration` request.
type DidChangeConfigurationParams struct {
	Settings interface{} `json:"settings"`
}

// ConfigurationItem defines a single configuration item.
type ConfigurationItem struct {
	// The scope to get the configuration section for.
	ScopeURI DocumentURI `json:"scopeUri,omitempty"`

	// The configuration section asked for.
	Section string `json:"section,omitempty"`
}

// ConfigurationParams contains the parameters contained in a
// `workspace/configuration` request.
type ConfigurationParams struct {
	Items []ConfigurationItem `json:"items"`
}
