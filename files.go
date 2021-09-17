package lsp

// WatchKind defines the kind of the file system watcher.
type WatchKind int

const (
	// WKCreate means the watcher is interested in create events.
	WKCreate WatchKind = 1 << iota

	// WKChange means the watcher is interested in change events.
	WKChange

	// WKDelete means the watcher is interested in delete events.
	WKDelete
)

func (kind WatchKind) String() string {
	switch kind {
	case WKCreate:
		return "create"
	case WKChange:
		return "change"
	case WKDelete:
		return "delete"
	case WKCreate | WKChange:
		return "create/change"
	case WKCreate | WKDelete:
		return "create/delete"
	case WKChange | WKDelete:
		return "change/delete"
	case WKCreate | WKChange | WKDelete:
		return "create/change/delete"
	}

	return "<unknown>"
}

// FileChangeType defines the type of a file event.
type FileChangeType int

const (
	// FCTCreated means the file got created.
	FCTCreated FileChangeType = 1 << iota

	// FCTChanged means the file got changed.
	FCTChanged

	// FCTDeleted means the file got deleted.
	FCTDeleted
)

func (fct FileChangeType) String() string {
	switch fct {
	case FCTCreated:
		return "created"
	case FCTChanged:
		return "changed"
	case FCTDeleted:
		return "deleted"
	case FCTCreated | FCTChanged:
		return "created/changed"
	case FCTCreated | FCTDeleted:
		return "created/deleted"
	case FCTChanged | FCTDeleted:
		return "changed/deleted"
	case FCTCreated | FCTChanged | FCTDeleted:
		return "created/changed/deleted"
	}

	return "<unknown>"
}

// FileSystemWatcher defines a watcher on a given glob pattern.
type FileSystemWatcher struct {
	// The glob pattern to watch.
	GlobPattern string `json:"globPattern"`

	// The kind of events of interest. If omitted it defaults to
	// `WKCreate | WKChange | WKDelete`, which is `7`.
	Kind WatchKind `json:"kind,omitempty"`
}

// FileEvent is an event describing a file change.
type FileEvent struct {
	// The file's URI.
	URI DocumentURI `json:"uri"`

	// The change type.
	Type FileChangeType `json:"type"`
}

// DidChangeWatchedFilesRegistrationOptions describes options to be used when
// registering for file system change events.
type DidChangeWatchedFilesRegistrationOptions struct {
	// The watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}

// DidChangeWatchedFilesParams contains the parameters of a request of method
// `workspace/didChangeWatchedFiles`.
type DidChangeWatchedFilesParams struct {
	Changes []FileEvent `json:"changes"`
}

// RenameOptions contains the options for the rename handler.
type RenameOptions struct {
	WorkDoneProgressOptions

	// Renames should be checked and tested before being executed.
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

// RenameRegistrationOptions contains the options for the rename handler
// registration.
type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions
	RenameOptions
}

// RenameParams contains the fields sent in a `textDocument/rename` request.
type RenameParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams

	// The new name of the symbol. If the given name is not valid the
	// request must return a [ResponseError](#ResponseError) with an
	// appropriate message set.
	NewName string `json:"newName"`
}

// PrepareRenameParams contains the fields sent in a `textDocument/prepareRename`
// request.
type PrepareRenameParams struct {
	TextDocumentPositionParams
}

// A pattern kind describing if a glob pattern matches a file a folder or both.
//
// @since 3.16.0
type FileOperationPatternKind string

const (
	// FOPKindFile means the pattern matches a file only.
	FOPKindFile FileOperationPatternKind = "file"

	// FOPKindFolder means the pattern matches a folder only.
	FOPKindFolder FileOperationPatternKind = "folder"
)

// Matching options for the file operation pattern.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
	// The pattern should be matched ignoring casing.
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// A pattern to describe in which file operation requests or notifications the
// server is interested in.
//
// @since 3.16.0
type FileOperationPattern struct {
	// The glob pattern to match. Glob patterns can have the following syntax:
	// - `*` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `**` to match any number of path segments, including none
	// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}`
	//   matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment
	//   (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment
	//   (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but
	//   not `example.0`)
	Glob string `json:"glob"`

	// Whether to match files or folders with this pattern.
	//
	// Matches both if undefined.
	Matches FileOperationPatternKind `json:"matches,omitempty"`

	// Additional options used during matching.
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}

// A filter to describe in which file operation requests or notifications the
// server is interested in.
//
// @since 3.16.0
type FileOperationFilter struct {
	// A Uri like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`

	// The actual file operation pattern.
	Pattern *FileOperationPattern `json:"pattern"`
}

// The options to register for file operations.
//
// @since 3.16.0
type FileOperationRegistrationOptions struct {
	Filters []FileOperationFilter `json:"filters"`
}

// Represents information on a file/folder rename.
//
// @since 3.16.0
type FileRename struct {
	// A file:// URI for the original location of the file/folder being renamed.
	OldURI string `json:"oldUri"`

	// A file:// URI for the new location of the file/folder being renamed.
	NewURI string `json:"newUri"`
}

// The parameters sent in notifications/requests for user-initiated renames of
// files.
//
// @since 3.16.0
type RenameFilesParams struct {
	// An array of all files/folders renamed in this operation. When a folder is
	// renamed, only the folder will be included, and not its children.
	Files []FileRename `json:"files"`
}

// Represents information on a file/folder delete.
//
// @since 3.16.0
type FileDelete struct {
	URI string `json:"uri"`
}

// The parameters sent in notifications/requests for user-initiated deletes of
// files.
//
// @since 3.16.0
type DeleteFilesParams struct {
	// A file:// URI for the location of the file/folder being deleted.
	Files []FileDelete `json:"files"`
}
