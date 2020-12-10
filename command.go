package lsp

// Command represents a reference to a command. Provides a title which will be
// used to represent a command in the UI. Commands are identified by a string
// identifier. The recommended way to handle commands is to implement their
// execution on the server side if the client and server provides the
// corresponding capabilities. Alternatively the tool extension code could
// handle the command. The protocol currently doesn’t specify a set of
// well-known commands.
type Command struct {
	// Title of the command, like `save`.
	Title string `json:"title"`

	// The identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command handler should be invoked with.
	Arguments []interface{} `json:"arguments"`
}

// ExecuteCommandOptions contains the options for the command execution handler.
type ExecuteCommandOptions struct {
	WorkDoneProgressOptions

	// The commands to be executed on the server.
	Commands []string `json:"commands"`
}

// ExecuteCommandRegistrationOptions contains the command execution registration
// options.
type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

// ExecuteCommandParams contains the fields sent in a `workspace/executeCommand`
// request.
type ExecuteCommandParams struct {
	WorkDoneProgressParams

	// The identifier of the actual command handler.
	Command string `json:"command"`

	// Arguments that the command should be invoked with.
	Arguments interface{} `json:"arguments,omitempty"`
}
