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

// ResourceOperationKind defines a kind of resource operation.
type ResourceOperationKind string

const (
	// ROKCreate refers to a file/folder creation operation.
	ROKCreate ResourceOperationKind = "create"

	// ROKRename refers to a file/folder rename operation.
	ROKRename = "rename"

	// ROKDelete refers to a file/folder deletion operation.
	ROKDelete = "delete"
)

// FailureHandlingKind defines how the client should behave on a workspace edit
// failure.
type FailureHandlingKind string

const (
	// FHKAbort means applying the workspace change is simply aborted if one of
	// the changes provided fails.
	FHKAbort FailureHandlingKind = "abort"

	// FHKTransactional means all operations are executed transactional. That
	// means they either all succeed or no changes at all are applied to the
	// workspace.
	FHKTransactional = "transactional"

	// FHKTextOnlyTransactional means if the workspace edit contains only textual
	// file changes, they are executed transactionally.
	// If resource changes (create, rename or delete file) are part of the
	// change, the failure handling strategy is abort.
	FHKTextOnlyTransactional = "textOnlyTransactional"

	// FHKUndo means the client tries to undo the operations already executed. But
	// there is no guarantee that this is succeeding.
	FHKUndo = "undo"
)

// CreateFileOptions contains options to create a file.
type CreateFileOptions struct {
	// Overwrite existing file. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// Ignore if exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// CreateFile defines a file creation operation.
type CreateFile struct {
	// Should always be "create"
	Kind string `json:"kind"`

	// The resource to create.
	URI DocumentURI `json:"uri"`

	// Additional options.
	Options CreateFileOptions `json:"options,omitempty"`
}

// NewCreateFile instantiates a CreateFile struct.
func NewCreateFile(URI DocumentURI, options CreateFileOptions) *CreateFile {
	return &CreateFile{
		Kind:    "create",
		URI:     URI,
		Options: options,
	}
}

// RenameFileOptions contains options to rename a file.
type RenameFileOptions struct {
	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`.
	Overwrite bool `json:"overwrite,omitempty"`

	// Ignore if target exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

// RenameFile defines a file rename operation.
type RenameFile struct {
	// Should always be "rename"
	Kind string `json:"kind"`

	// The old (existing) location.
	OldURI DocumentURI `json:"oldUri"`

	// The new location.
	NewURI DocumentURI `json:"newUri"`

	// Additional options.
	Options RenameFileOptions `json:"options,omitempty"`
}

// NewRenameFile instantiates a RenameFile struct.
func NewRenameFile(oldURI, newURI DocumentURI, options RenameFileOptions) *RenameFile {
	return &RenameFile{
		Kind:    "rename",
		OldURI:  oldURI,
		NewURI:  newURI,
		Options: options,
	}
}

// DeleteFileOptions contains options to delete a file.
type DeleteFileOptions struct {
	// Delete the content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`

	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

// DeleteFile defines a file deletion operation.
type DeleteFile struct {
	// Should always be "delete".
	Kind string `json:"kind"`

	// The file to delete.
	URI DocumentURI `json:"uri"`

	// Additional options.
	Options DeleteFileOptions `json:"options,omitempty"`
}

// NewDeleteFile instantiates a DeleteFile struct.
func NewDeleteFile(URI DocumentURI, options DeleteFileOptions) *DeleteFile {
	return &DeleteFile{
		Kind:    "delete",
		URI:     URI,
		Options: options,
	}
}

// WorkspaceEditDocumentChange is an interface that can be interpreted as a
// `TextDocumentEdit`, `CreateFile`, `RenameFile`, or `DeleteFile`.
type WorkspaceEditDocumentChange interface{}

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
	DocumentChanges []WorkspaceEditDocumentChange `json:"documentChanges,omitempty"`
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
