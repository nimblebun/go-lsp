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

// WorkspaceSymbolClientCapabilities contains information about the client's
// client's workspace symbol capabilities.
type WorkspaceSymbolClientCapabilities struct {
	// Symbol request supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Specific capabilities for the `SymbolKind` in the `workspace/symbol`
	// request.
	SymbolKind struct {
		ValueSet []SymbolKind `json:"valueSet,omitempty"`
	} `json:"symbolKind,omitempty"`
}

// ExecuteCommandClientCapabilities contains information about the client's
// command execution capabilities.
type ExecuteCommandClientCapabilities struct {
	// Execute command supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentSyncClientCapabilities contains information about the client's
// text document sync capabilities.
type TextDocumentSyncClientCapabilities struct {
	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports sending will save notifications.
	WillSave bool `json:"willSave,omitempty"`

	// The client supports sending a will save request and
	// waits for a response providing text edits which will
	// be applied to the document before it is saved.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// The client supports did save notifications.
	DidSave bool `json:"didSave,omitempty"`
}

// PublishDiagnosticsClientCapabilities contains information about the client's
// diagnostics publishing capabilities.
type PublishDiagnosticsClientCapabilities struct {
	// Whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"relatedInformation,omitempty"`

	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	TagSupport struct {
		ValueSet []DiagnosticTag `json:"valueSet"`
	} `json:"tagSupport,omitempty"`

	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.
	VersionSupport bool `json:"versionSupport,omitempty"`
}

// CompletionClientCapabilities contains information about the client's
// completion capabilities.
type CompletionClientCapabilities struct {
	// Whether completion supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports the following `CompletionItem` specific capabilities.
	CompletionItem struct {
		// Client supports snippets as insert text.
		//
		// A snippet can define tab stops and placeholders with `$1`, `$2`
		// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
		// the end of the snippet.
		// Placeholders with equal identifiers are linked, so that typing in
		// one will update others as well.
		SnippetSupport bool `json:"snippetSupport,omitempty"`

		// Client supports commit characters on a completion item.
		CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`

		// Client supports the follow content formats for the documentation
		// property. The order describes the preferred format of the client.
		DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

		// Client supports the deprecated property on a completion item.
		DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`

		// Client supports the preselect property on a completion item.
		PreselectSupport bool `json:"preselectSupport,omitempty"`

		// Client supports the tag property on a completion item.
		// Clients supporting tags have to handle unknown tags gracefully.
		// Clients especially need to preserve unknown tags when sending
		// a completion item back to the server in a resolve call.
		TagSupport struct {
			// The tags supported by the client.
			ValueSet []CompletionItemTag `json:"valueSet"`
		} `json:"tagSupport,omitempty"`
	} `json:"completionItem,omitempty"`

	// The completion item kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	//
	// If this property is not present the client only supports
	// the completion items kinds from `Text` to `Reference` as defined in
	// the initial version of the protocol.
	CompletionItemKind struct {
		// The completion item kinds supported by the client.
		ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
	} `json:"completionItemKind,omitempty"`

	// The client supports to send additional context information for a
	// `textDocument/completion` request.
	ContextSupport bool `json:"contextSupport,omitempty"`
}
