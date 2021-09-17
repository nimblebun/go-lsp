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

	// Whether the client normalizes line endings to the client specific setting.
	// If set to `true` the client will normalize line ending characters in a
	// workspace edit to the client specific new line character(s).
	//
	// @since 3.16.0
	NormalizesLineEndings bool `json:"normalizesLineEndings,omitempty"`

	// Whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	//
	// @since 3.16.0
	ChangeAnnotationSupport struct {
		// Whether the client groups edits with equal labels into tree nodes, for
		// instance all edits labelled with "Changes in Strings" would be a tree
		// node.
		GroupsOnLabel bool `json:"groupsOnLabel,omitempty"`
	} `json:"changeAnnotationSupport,omitempty"`
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

	// The client supports tags on `SymbolInformation`.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport struct {
		ValueSet []SymbolTag `json:"valueSet,omitempty"`
	} `json:"tagSupport,omitempty"`
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

	// Client supports a codeDescription property
	//
	// @since 3.16.0
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`

	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	//
	// @since 3.16.0
	DataSupport bool `json:"dataSupport,omitempty"`
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

		// Client supports insert replace edit to control different behavior if a
		// completion item is inserted in the text or should replace text.
		//
		// @since 3.16.0
		InsertReplaceSupport bool `json:"insertReplaceSupport,omitempty"`

		// Indicates which properties a client can resolve lazily on a completion
		// item. Before version 3.16.0 only the predefined properties
		// `documentation` and `detail` could be resolved lazily.
		//
		// @since 3.16.0
		ResolveSupport struct {
			// The properties that a client can resolve lazily.
			Properties []string `json:"properties"`
		} `json:"resolveSupport,omitempty"`

		// The client supports the `insertTextMode` property on a completion item to
		// override the whitespace handling mode as defined by the client.
		//
		// @since 3.16.0
		InsertTextModeSupport struct {
			ValueSet []InsertTextMode `json:"valueSet"`
		} `json:"insertTextModeSupport,omitempty"`
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

		// The client supports the `activeParameter` property on
		// `SignatureInformation` literal.
		//
		// @since 3.16.0
		ActiveParameterSupport bool `json:"activeParameterSupport,omitempty"`
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

	// The client supports hierarchical document symbols.
	//
	// @since 3.16.0
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`

	// The client supports tags on `SymbolInformation`. Tags are supported on
	// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport struct {
		// The tags supported by the client.
		ValueSet []SymbolTag
	} `json:"tagSupport,omitempty"`

	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.
	//
	// @since 3.16.0
	LabelSupport bool `json:"labelSupport,omitempty"`
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

// PrepareSupportDefaultBehavior indicates the default behavior for the client's
// prepare support.
//
// @since 3.16.0
type PrepareSupportDefaultBehavior int

const (
	// PSDefaultBehaviorIdentifier means the client's default behavior is to
	// select the identifier according the to language's syntax rule.
	PSDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = iota + 1
)

// RenameClientCapabilities contains information about the client's file rename
// capabilities.
type RenameClientCapabilities struct {
	// Whether rename supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Client supports testing for validity of rename operations before execution.
	PrepareSupport bool `json:"prepareSupport,omitempty"`

	// Client supports the default behavior result
	// (`{ defaultBehavior: boolean }`).
	//
	// The value indicates the default behavior used by the client.
	//
	// @since 3.16.0
	PrepareSupportDefaultBehavior PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitempty"`

	// Whether the client honors the change annotations in text edits and resource
	// operations returned via the rename request's workspace edit by for example
	// presenting the workspace edit in the user interface and asking for
	// confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
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

// CodeActionClientCapabilities contains information about the client's code
// action capabilities.
type CodeActionClientCapabilities struct {
	// Whether code action supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// The client supports code action literals as a valid
	// response of the `textDocument/codeAction` request.
	CodeActionLiteralSupport struct {
		// The code action kind is supported with the following value set.
		CodeActionKind struct {
			// The code action kind values the client supports. When this
			// property exists the client also guarantees that it will
			// handle values outside its set gracefully and falls back
			// to a default value when unknown.
			ValueSet []CodeActionKind `json:"valueSet"`
		} `json:"codeActionKind,omitempty"`
	} `json:"codeActionLiteralSupport,omitempty"`

	// Whether code action supports the `isPreferred` property.
	IsPreferredSupport bool `json:"isPreferredSupport,omitempty"`

	// Whether code action supports the `disabled` property.
	//
	// @since 3.16.0
	DisabledSupport bool `json:"disabledSupport,omitempty"`

	// Whether code action supports the `data` property which is preserved between
	// a `textDocument/codeAction` and a `codeAction/resolve` request.
	//
	// @since 3.16.0
	DataSupport bool `json:"dataSupport,omitempty"`

	// Whether the client supports resolving additional code action properties via
	// a separate `codeAction/resolve` request.
	//
	// @since 3.16.0
	ResolveSupport struct {
		// The properties that a client can resolve lazily.
		Properties []string `json:"properties"`
	} `json:"resolveSupport,omitempty"`

	// Whether the client honors the change annotations in text edits and resource
	// operations returned via the `CodeAction#edit` property by for example
	// presenting the workspace edit in the user interface and asking for
	// confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
}

// CodeLensClientCapabilities contains information about the client's CodeLens
// capabilities.
type CodeLensClientCapabilities struct {
	// Whether CodeLens supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// Client capabilities specific to the used markdown parser.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// The name of the parser.
	Parser string `json:"parser"`

	// The version of the parser.
	Version string `json:"version,omitempty"`
}

// Client capabilities specific to regular expressions.
//
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
	// The name of the engine.
	Engine string `json:"engine"`

	// The version of the engine.
	Version string `json:"version,omitempty"`
}

// CodeLensWorkspaceCapabilities contains information about the client's
// workspace CodeLens capabilities.
//
// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// code lenses currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// LinkedEditingRangeClientCapabilities contains information about the client's
// linked editing range capabilities.
//
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration.
	// If this is set to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// CallHierarchyClientCapabilities contains information about the client's
// call hierarchy capabilities.
//
// @since 3.16.0
type CallHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// SemanticTokensClientCapabilities contains information about the client's
// semantic tokens capabilities.
//
// @since 3.16.0
type SemanticTokensClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

	// Which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might not
	// show semantic tokens or degrade some of the user experience if a range
	// or full request is advertised by the client but not provided by the
	// server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or might
	// even decide to not show any semantic tokens at all.
	Requests struct {
		// The client will send the `textDocument/semanticTokens/range` request
		// if the server provides a corresponding handler.
		Range bool `json:"range,omitempty"`

		// The client will send the `textDocument/semanticTokens/full` request
		// if the server provides a corresponding handler.
		Full *struct {
			// The client will send the `textDocument/semanticTokens/full/delta`
			// request if the server provides a corresponding handler.
			Delta bool `json:"delta,omitempty"`
		} `json:"full,omitempty"`
	} `json:"requests"`

	// The token types that the client supports
	TokenTypes []SemanticTokenType `json:"tokenTypes"`

	// The token modifiers that the client supports.
	TokenModifiers []SemanticTokenModifiers `json:"tokenModifiers"`

	// The formats the clients supports.
	Formats []TokenFormat `json:"formats"`

	// Whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport bool `json:"overlappingTokenSupport,omitempty"`

	// Whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport bool `json:"multilineTokenSupport,omitempty"`
}

// SemanticTokensWorkspaceClientCapabilities contains information about the
// client's semantic tokens workspace capabilities.
//
// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// semantic tokens currently shown. It should be used with absolute care
	// and is useful for situation where a server for example detect a project
	// wide change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// MonikerClientCapabilities contains information about the client's moniker
// capabilities.
//
// @since 3.16.0
type MonikerClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to
	// `true` the client supports the new `(TextDocumentRegistrationOptions &
	// StaticRegistrationOptions)` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

// ShowMessageRequestClientCapabilities contains information about the client's
// show message request capabilities.
//
// @since 3.16.0
type ShowMessageRequestClientCapabilities struct {
	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *struct {
		// Whether the client supports additional attributes which are preserved and
		// sent back to the server in the request's response.
		AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
	} `json:"messageActionItem"`
}

// ShowDocumentClientCapabilities contains information about the client's
// show document capabilities.
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// The client has support for the show document request.
	Support bool `json:"support"`
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
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`

	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`

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
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`

	// Capabilities specific to the `textDocument/foldingRange` request.
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`

	// Capabilities specific to the `textDocument/selectionRange` request.
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`

	// Capabilities specific to the `textDocument/linkedEditingRange` request.
	//
	// @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitempty"`

	// Capabilities specific to the various call hierarchy requests.
	//
	// @since 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitempty"`

	// Capabilities specific to the various semantic token requests.
	//
	// @since 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`

	// Capabilities specific to the `textDocument/moniker` request.
	//
	// @since 3.16.0
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`
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
		ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`

		// The client has support for workspace folders.
		WorkspaceFolders bool `json:"workspaceFolders,omitempty"`

		// The client supports `workspace/configuration` requests.
		Configuration bool `json:"configuration,omitempty"`

		// Capabilities specific to the semantic token requests scoped to the
		// workspace.
		//
		// @since 3.16.0
		SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`

		// Capabilities specific to the code lens requests scoped to the
		// workspace.
		//
		// @since 3.16.0
		CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`

		// The client has support for file requests/notifications.
		//
		// @since 3.16.0
		FileOperations *struct {
			// Whether the client supports dynamic registration for file requests or
			// notifications.
			DynamicRegistration bool `json:"dynamicRegistration,omitempty"`

			// The client has support for sending didCreateFiles notifications.
			DidCreate bool `json:"didCreate,omitempty"`

			// The client has support for sending willCreateFiles requests.
			WillCreate bool `json:"willCreate,omitempty"`

			// The client has support for sending didRenameFiles notifications.
			DidRename bool `json:"didRename,omitempty"`

			// The client has support for sending willRenameFiles requests.
			WillRename bool `json:"willRename,omitempty"`

			// The client has support for sending didDeleteFiles notifications.
			DidDelete bool `json:"didDelete,omitempty"`

			// The client has support for sending willDeleteFiles requests.
			WillDelete bool `json:"willDelete,omitempty"`
		} `json:"fileOperations,omitempty"`
	} `json:"workspace,omitempty"`

	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	// Window specific client capabilities.
	Window *struct {
		// Whether client supports handling progress notifications.
		// If set, servers are allowed to report in `workDoneProgress` property
		// in the request specific server capabilities.
		WorkDoneProgress bool `json:"workDoneProgress,omitempty"`

		// Capabilities specific to the showMessage request
		//
		// @since 3.16.0
		ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`

		// Client capabilities for the show document request.
		//
		// @since 3.16.0
		ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
	} `json:"window,omitempty"`

	// General client capabilities.
	//
	// @since 3.16.0
	General *struct {
		// Client capabilities specific to regular expressions.
		//
		// @since 3.16.0
		RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`

		// Client capabilities specific to the client's markdown parser.
		//
		// @since 3.16.0
		Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`
	} `json:"general,omitempty"`

	// Experimental client capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
}
