package lsp

// WorkspaceFolder is a structure that defines the reference to a workspace
// folder.
type WorkspaceFolder struct {
	// The associated URI for this workspace folder.
	URI string `json:"uri"`

	// The name of the workspace folder. Used to refer to this workspace folder in
	// the user interface.
	Name string `json:"name"`
}

// WorkspaceFoldersChangeEvent contains information about a workspace folder
// change event.
type WorkspaceFoldersChangeEvent struct {
	// The array of added workspace folders.
	Added []WorkspaceFolder `json:"added"`

	// The array of removed workspace folders.
	Removed []WorkspaceFolder `json:"removed"`
}

// WorkspaceEdit represents changes to many resources managed in the workspace.
type WorkspaceEdit struct {
	// Holds changes to existing resources.
	Changes map[DocumentURI][]TextEdit `json:"changes,omitempty"`

	// The client capability `workspace.workspaceEdit.resourceOperations`
	// determines whether document changes are either an array of
	// `TextDocumentEdit`s to express changes to different text documents,
	// where each text document edit addresses a specific version
	// of a text document, or it can contains the above `TextDocumentEdit`s
	// mixed with create, rename, and delete file / folder operations.
	//
	// Whether a client supports versioned document edits is expressed via
	// `workspace.workspaceEdit.documentChanges` client capability.
	//
	// If a client doesn't support `documentChanges` or
	// `workspace.workspaceEdit.resourceOperations`, then only plain
	// `TextEdit`s using the `changes` property are supported.
	DocumentChanges interface{} `json:"documentChanges,omitempty"`
}

// DidChangeWorkspaceFoldersParams are the parameters contained in a
// `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}

// ApplyWorkspaceEditParams contains the fields sent in a `workspace/applyEdit`
// request.
type ApplyWorkspaceEditParams struct {
	// An optional label of the workspace edit. This label is
	// presented in the user interface for example on an undo
	// stack to undo the workspace edit.
	Label string `json:"label,omitempty"`

	// The edits to apply.
	Edit WorkspaceEdit `json:"edit"`
}

// ApplyWorkspaceEditResponse contains the fields from a `workspace/applyEdit`
// response.
type ApplyWorkspaceEditResponse struct {
	// Indicates whether the edit was applied or not.
	Applied bool `json:"applied"`

	// An optional textual description for why the edit was not applied.
	// This may be used may be used by the server for diagnostic
	// logging or to provide a suitable error for a request that
	// triggered the edit.
	FailureReason string `json:"failureReason,omitempty"`
}
