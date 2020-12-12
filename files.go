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
