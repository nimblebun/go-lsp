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

// FileSystemWatcher defines a watcher on a given glob pattern.
type FileSystemWatcher struct {
	// The glob pattern to watch.
	GlobPattern string `json:"globPattern"`

	// The kind of events of interest. If omitted it defaults to
	// `WKCreate | WKChange | WKDelete`, which is `7`.
	Kind WatchKind `json:"kind,omitempty"`
}
