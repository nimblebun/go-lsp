package lsp

// ClientInfo contains information about the client.
type ClientInfo struct {
	// The name of the client as defined by the client.
	Name string `json:"name"`

	// The client's version as defined by the client.
	Version string `json:"version,omitempty"`
}

// ServerInfo contains information about the server.
type ServerInfo struct {
	// The name of the server as defined by the server.
	Name string `json:"name"`

	// The server's version as defined by the server.
	Version string `json:"version,omitempty"`
}

// TraceType specifies the type of trace that needs to be done in a
// communication between the client and the server.
type TraceType string

const (
	// TraceOff means tracing is turned off.
	TraceOff TraceType = "off"

	// TraceMessages means only the messages will be traced.
	TraceMessages = "messages"

	// TraceVerbose means tracing will be detailed.
	TraceVerbose = "verbose"
)

// InitializeParams is a struct containing the parameters of an `initialize`
// request.
type InitializeParams struct {
	WorkDoneProgressParams

	// The process ID of the parent process that started the server.
	// Is null if the process has not been started by another process.
	// If the parent process is not alive, then the server should exit
	// (see exit notification) its process.
	ProcessID int `json:"processId,omitempty"`

	// Information about the client.
	ClientInfo ClientInfo `json:"clientInfo,omitempty"`

	// The locale the client is currently showing the user interface in. This must
	// not necessarily be the locale of the operating system.
	//
	// Uses IETF language tags as the value's syntax (See
	// https://en.wikipedia.org/wiki/IETF_language_tag)
	//
	// @since 3.16.0
	Locale string `json:"locale,omitempty"`

	// The rootUri of the workspace.
	RootURI DocumentURI `json:"rootUri,omitempty"`

	// User provided initialization options.
	InitializationOptions interface{} `json:"initializationOptions"`

	// The capabilities provided by the client (editor or tool).
	// Capabilities ClientCapabilities `json:"capabilities"`

	// The initial trace setting. If omitted trace is disabled ('off').
	Trace TraceType `json:"trace,omitempty"`

	// The workspace folders configured in the client when the server starts.
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	WorkspaceFolders []WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

// InitializeResult contains the fields of the `result` parameter in a response
// payload for an `initialize` request.
type InitializeResult struct {
	// The capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`

	// Information about the server.
	ServerInfo ServerInfo `json:"serverInfo,omitempty"`
}
