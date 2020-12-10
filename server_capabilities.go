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
