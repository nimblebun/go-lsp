package lsp

// WorkspaceFoldersServerCapabilities contains information about the server's
// workspace folders capabilities.
type WorkspaceFoldersServerCapabilities struct {
	// The server has support for workspace folders.
	Supported bool `json:"supported,omitempty"`

	// Whether the server wants to receive workspace folder
	// change notifications.
	ChangeNotifications string `json:"changeNotifications,omitempty"`
}

// ServerCapabilities defines the capabilities of the language server.
type ServerCapabilities struct {
	// Defines how text documents are synced.
	TextDocumentSync *TextDocumentSyncOptions `json:"textDocumentSync,omitempty"`

	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// The server provides hover support.
	HoverProvider *HoverOptions `json:"hoverProvider,omitempty"`

	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// The server provides go to declaration support.
	DeclarationProvider *DeclarationRegistrationOptions `json:"declarationProvider,omitempty"`

	// The server provides goto definition support.
	DefinitionProvider *DefinitionRegistrationOptions `json:"definitionProvider,omitempty"`

	// The server provides goto type definition support.
	TypeDefinitionProvider *TypeDefinitionRegistrationOptions `json:"typeDefinitionProvider,omitempty"`

	// The server provides goto implementation support.
	ImplementationProvider *ImplementationRegistrationOptions `json:"implementationProvider,omitempty"`

	// The server provides find references support.
	ReferencesProvider *ReferenceRegistrationOptions `json:"referencesProvider,omitempty"`

	// The server provides document highlight support.
	DocumentHighlightProvider *DocumentHighlightRegistrationOptions `json:"documentHighlightProvider,omitempty"`

	// The server provides document symbol support.
	DocumentSymbolProvider *DocumentSymbolRegistrationOptions `json:"documentSymbolProvider,omitempty"`

	// The server provides code actions.
	CodeActionProvider *CodeActionRegistrationOptions `json:"codeActionProvider,omitempty"`

	// The server provides CodeLens.
	CodeLensProvider *CodeLensRegistrationOptions `json:"codeLensProvider,omitempty"`

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkRegistrationOptions `json:"documentLinkProvider,omitempty"`

	// The server provides color provider support.
	ColorProvider *DocumentColorRegistrationOptions `json:"colorProvider,omitempty"`

	// The server provides document formatting.
	DocumentFormattingProvider *DocumentFormattingRegistrationOptions `json:"documentFormattingProvider,omitempty"`

	// The server provides document range formatting.
	DocumentRangeFormattingProvider *DocumentRangeFormattingRegistrationOptions `json:"documentRangeFormattingProvider,omitempty"`

	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingRegistrationOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	// The server provides rename support.
	RenameProvider *RenameRegistrationOptions `json:"renameProvider,omitempty"`

	// The server provides folding provider support.
	FoldingRangeProvider *FoldingRangeRegistrationOptions `json:"foldingRangeProvider,omitempty"`

	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandRegistrationOptions `json:"executeCommandProvider,omitempty"`

	// The server provides selection range support.
	SelectionRangeProvider *SelectionRangeRegistrationOptions `json:"selectionRangeProvider,omitempty"`

	// The server provides linked editing range support.
	//
	// @since 3.16.0
	LinkedEditingRangeProvider *LinkedEditingRangeRegistrationOptions `json:"linkedEditingRangeProvider,omitempty"`

	// The server provides call hierarchy support.
	//
	// @since 3.16.0
	CallHierarchyProvider *CallHierarchyRegistrationOptions `json:"callHierarchyProvider,omitempty"`

	// The server provides semantic tokens support.
	//
	// @since 3.16.0
	SemanticTokensProvider *SemanticTokensRegistrationOptions `json:"semanticTokensProvider,omitempty"`

	// Whether server provides moniker support.
	//
	// @since 3.16.0
	MonikerProvider *MonikerRegistrationOptions `json:"monikerProvider,omitempty"`

	// The server provides workspace symbol support.
	WorkspaceSymbolProvider *WorkspaceSymbolRegistrationOptions `json:"workspaceSymbolProvider,omitempty"`

	// Workspace specific server capabilities.
	Workspace *struct {
		// The server supports workspace folder.
		WorkspaceFolders WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitempty"`

		// The server is interested in file notifications/requests.
		//
		// @since 3.16.0
		FileOperations *struct {
			// The server is interested in receiving didCreateFiles notifications.
			DidCreate *FileOperationRegistrationOptions `json:"didCreate,omitempty"`

			// The server is interested in receiving willCreateFiles requests.
			WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`

			// The server is interested in receiving didRenameFiles notifications.
			DidRename *FileOperationRegistrationOptions `json:"didRename,omitempty"`

			// The server is interested in receiving willRenameFiles requests.
			WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`

			// The server is interested in receiving didDeleteFiles file notifications
			DidDelete *FileOperationRegistrationOptions `json:"didDelete,omitempty"`

			// The server is interested in receiving willDeleteFiles file requests.
			WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
		} `json:"fileOperations,omitempty"`
	} `json:"workspace,omitempty"`

	// Experimental server capabilities.
	Experimental interface{} `json:"experimental,omitempty"`
}
