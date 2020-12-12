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

// CodeActionKind defines a set of predefined code action kinds.
type CodeActionKind string

const (
	// CAKEmpty defines an empty kind.
	CAKEmpty CodeActionKind = ""

	// CAKQuickFix is the base kind for quickfix actions: 'quickfix'.
	CAKQuickFix = "quickfix"

	// CAKRefactor is the base kind for refactoring actions: 'refactor'.
	CAKRefactor = "refactor"

	// CAKRefactorExtract is the base kind for refactoring extraction actions:
	// 'refactor.extract'.
	//
	// Example extract actions:
	// - Extract method
	// - Extract function
	// - Extract variable
	// - Extract interface from class
	// - ...
	CAKRefactorExtract = "refactor.extract"

	// CAKRefactorInline is the base kind for refactoring inline actions:
	// 'refactor.inline'.
	//
	// Example inline actions:
	//
	// - Inline function
	// - Inline variable
	// - Inline constant
	// - ...
	CAKRefactorInline = "refactor.inline"

	// CAKRefactorRewrite is the base kind for refactoring rewrite actions:
	// 'refactor.rewrite'.
	//
	// Example rewrite actions:
	//
	// - Convert JavaScript function to class
	// - Add or remove parameter
	// - Encapsulate field
	// - Make method static
	// - Move method to base class
	// - ...
	CAKRefactorRewrite = "refactor.rewrite"

	// CAKSource is the base kind for source actions: `source`.
	//
	// Source code actions apply to the entire file.
	CAKSource = "source"

	// CAKSourceOrganizeImports is the base kind for an organize imports source
	// action `source.organizeImports`.
	CAKSourceOrganizeImports = "source.organizeImports"
)

// CodeActionContext contains additional diagnostic information about the
// context in which a code action is run.
type CodeActionContext struct {
	// An array of diagnostics known on the client side overlapping the range
	// provided to the `textDocument/codeAction` request.
	// They are provided so that the server knows which errors are currently
	// presented to the user for the given range. There is no guarantee that
	// these accurately reflect the error state of the resource.
	// The primary parameter to compute code actions is the provided range.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Requested kind of actions to return.
	//
	// Actions not of this kind are filtered out by the client before
	// being shown, so servers can omit computing them.
	Only []CodeActionKind `json:"only,omitempty"`
}

// CodeAction represents a change that can be performed in code.
// For example, to fix a problem or to refactor code.
//
// A CodeAction must set either `edit` and/or a `command`.
// If both are supplied, the `edit` is applied first, then the `command`
// is executed.
type CodeAction struct {
	// A short, human-readable, title for this code action.
	Title string `json:"title"`

	// The kind of the code action. Used to filter code actions.
	Kind CodeActionKind `json:"kind,omitempty"`

	// The diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`

	// Marks this as a preferred action.
	// Preferred actions are used by the `auto fix` command and can be
	// targeted by keybindings.
	//
	// A quick fix should be marked preferred if it properly addresses the
	// underlying error.
	// A refactoring should be marked preferred if it is the most reasonable
	// choice of actions to take.
	IsPreferred bool `json:"isPreferred,omitempty"`

	// The workspace edit this code action performs.
	Edit WorkspaceEdit `json:"edit,omitempty"`

	// A command this code action executes. If a code action
	// provides an edit and a command, first the edit is
	// executed and then the command.
	Command Command `json:"command,omitempty"`
}

// CodeActionOptions contains the options for the code action handler.
type CodeActionOptions struct {
	WorkDoneProgressOptions

	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`,
	// or the server may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
}

// CodeActionRegistrationOptions contains the options for the code action
// handler registration.
type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

// CodeActionParams contains the fields sent in a `textDocument/codeAction`
// request.
type CodeActionParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range for which the command was invoked.
	Range Range `json:"range"`

	// Context carrying additional information.
	Context CodeActionContext `json:"context"`
}
