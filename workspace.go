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

// DidChangeWorkspaceFoldersParams are the parameters contained in a
// `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}
