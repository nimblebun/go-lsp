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

// WorkspaceEditClientCapabilities contains information about the client's
// workspace edit capabilities.
type WorkspaceEditClientCapabilities struct {
	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges bool `json:"documentChanges,omitempty"`

	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`

	// The failure handling strategy of a client if applying the workspace edit
	// fails.
	FailureHandling []FailureHandlingKind `json:"failureHandling,omitempty"`
}

// WorkspaceSymbolClientCapabilities contains information about the client's
// workspace symbol capabilities.
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

// HoverClientCapabilities contains information about the client's hover
// capabilities.
type HoverClientCapabilities struct {
	// Whether hover supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Client supports the follow content formats for the content property. The
	// order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
}

// SignatureHelpClientCapabilities contains information about the client's
// signature help capabilities.
type SignatureHelpClientCapabilities struct {
	// Whether signature help supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports the following `SignatureInformation` specific
	// properties.
	SignatureInformation struct {
		// Client supports the follow content formats for the documentation
		// property. The order describes the preferred format of the client.
		DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`

		// Client capabilities specific to parameter information.
		ParameterInformation struct {
			// The client supports processing label offsets instead of a simple label
			// string.
			LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
		} `json:"parameterInformation,omitempty"`
	} `json:"signatureInformation,omitempty"`

	// The client supports to send additional context information for a
	// `textDocument/signatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.
	ContextSupport bool `json:"contextSupport,omitempty"`
}

// DeclarationClientCapabilities contains information about the client's
// go-to-declaration capabilities.
type DeclarationClientCapabilities struct {
	// Whether declaration supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of declaration links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// DefinitionClientCapabilities contains information about the client's
// go-to-definition capabilities.
type DefinitionClientCapabilities struct {
	// Whether definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of definition links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// TypeDefinitionClientCapabilities contains information about the client's
// go-to-type-definition capabilities.
type TypeDefinitionClientCapabilities struct {
	// Whether type definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of definition links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// ImplementationClientCapabilities contains information about the client's
// go-to-implementation capabilities.
type ImplementationClientCapabilities struct {
	// Whether implementation supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports additional metadata in the form of definition links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

// ReferenceClientCapabilities contains information about the client's find
// references capabilities.
type ReferenceClientCapabilities struct {
	// Whether references support dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentHighlightClientCapabilities contains information about the client's
// document highlight capabilities.
type DocumentHighlightClientCapabilities struct {
	// Whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentSymbolClientCapabilities contains information about the client's
// document symbol capabilities.
type DocumentSymbolClientCapabilities struct {
	// Whether document symbol supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Specific capabilities for the `SymbolKind` in the
	// `textDocument/documentSymbol` request.
	SymbolKind struct {
		// The symbol kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the symbol kinds from `File` to `Array` as defined in
		// the initial version of the protocol.
		ValueSet []SymbolKind `json:"valueSet,omitempty"`
	} `json:"symbolKind,omitempty"`
}

// DocumentLinkClientCapabilities contains information about the client's
// document link capabilities.
type DocumentLinkClientCapabilities struct {
	// Whether document link supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Whether the client supports the `tooltip` property on `DocumentLink`.
	TooltipSupport bool `json:"tooltipSupport,omitempty"`
}

// DocumentColorClientCapabilities contains information about the client's
// document color capabilities.
type DocumentColorClientCapabilities struct {
	// Whether document color supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentFormattingClientCapabilities contains information about the client's
// document formatting capabilities.
type DocumentFormattingClientCapabilities struct {
	// Whether document formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentRangeFormattingClientCapabilities contains information about the
// client's document formatting capabilities.
type DocumentRangeFormattingClientCapabilities struct {
	// Whether document formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// DocumentOnTypeFormattingClientCapabilities contains information about the
// client's document on-type formatting capabilities.
type DocumentOnTypeFormattingClientCapabilities struct {
	// Whether on-type formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// RenameClientCapabilities contains information about the client's file rename
// capabilities.
type RenameClientCapabilities struct {
	// Whether rename supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Client supports testing for validity of rename operations before execution.
	PrepareSupport bool `json:"prepareSupport,omitempty"`
}

// FoldingRangeClientCapabilities contains information about the client's
// folding range provider capabilities.
type FoldingRangeClientCapabilities struct {
	// Whether the implementation supports dynamic registration for folding range
	// providers.
	// If this is set to `true`, the client supports the new
	// `FoldingRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The maximum number of folding ranges that the client prefers to
	// receive per document.
	// The value serves as a hint, servers are free to follow the limit.
	RangeLimit int `json:"rangeLimit,omitempty"`

	// If set, the client signals that it only supports folding complete lines.
	// If set, the client will ignore specified `startCharacter` and
	// `endCharacter` properties in a FoldingRange.
	LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
}

// SelectionRangeClientCapabilities contains information about the client's
// selection range provider capabilities.
type SelectionRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for selection
	// range providers.
	// If set to `true`, the client supports the new
	// `SelectionRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// TextDocumentClientCapabilities define capabilities the editor / tool provides
// on text documents.
type TextDocumentClientCapabilities struct {
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`

	// Capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`

	// Capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitempty"`

	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitempty"`

	// Capabilities specific to the `textDocument/declaration` request.
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`

	// Capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`

	// Capabilities specific to the `textDocument/typeDefinition` request.
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitempty"`

	// Capabilities specific to the `textDocument/implementation` request.
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`

	// Capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references,omitempty"`

	// Capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitempty"`

	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitempty"`

	// Capabilities specific to the `textDocument/codeAction` request.
	// CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`

	// Capabilities specific to the `textDocument/codeLens` request.
	// CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`

	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitempty"`

	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitempty"`

	// Capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`

	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`

	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`

	// Capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitempty"`

	// Capabilities specific to the `textDocument/publishDiagnostics` notification
	PublishDiagnostics PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`

	// Capabilities specific to the `textDocument/foldingRange` request.
	FoldingRange FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`

	// Capabilities specific to the `textDocument/selectionRange` request.
	SelectionRange SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`
}

// ClientCapabilities defines workspace-specific client capabilities.
type ClientCapabilities struct {
	// Workspace specific client capabilities.
	Workspace *struct {
		// The client supports applying batch edits
		// to the workspace by supporting the request
		// 'workspace/applyEdit'
		ApplyEdit bool `json:"applyEdit,omitempty"`

		// Capabilities specific to `WorkspaceEdit`s
		WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`

		// Capabilities specific to the `workspace/didChangeConfiguration`
		// notification.
		DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`

		// Capabilities specific to the `workspace/didChangeWatchedFiles`
		// notification.
		DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`

		// Capabilities specific to the `workspace/symbol` request.
		Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`

		// Capabilities specific to the `workspace/executeCommand` request.
		ExecuteCommand ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`

		// The client has support for workspace folders.
		WorkspaceFolders bool `json:"workspaceFolders,omitempty"`

		// The client supports `workspace/configuration` requests.
		Configuration bool `json:"configuration,omitempty"`
	} `json:"workspace,omitempty"`

	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	// Window specific client capabilities.
	Window *struct {
		// Whether client supports handling progress notifications.
		// If set, servers are allowed to report in `workDoneProgress` property
		// in the request specific server capabilities.
		WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
	} `json:"window,omitempty"`

	// Experimental client capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
}
