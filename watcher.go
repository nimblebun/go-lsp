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
