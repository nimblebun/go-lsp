package lsp

// CodeLens represents a command that should be shown along with
// source text, like the number of references, a way to run tests, etc.
//
// A CodeLens is _unresolved_ when no command is associated to it.
// For performance reasons, the creation of a CodeLens and resolving should
// be done in two stages.
type CodeLens struct {
	// The range in which the CodeLens is valid. Should only span a single line.
	Range Range `json:"range"`

	// The command this CodeLens represents.
	Command Command `json:"command,omitempty"`

	// A data entry field that is preserved on a CodeLens item between
	// a CodeLens and a CodeLens resolve request.
	Data interface{} `json:"data,omitempty"`
}

// CodeLensOptions contains the options for the CodeLens handler.
type CodeLensOptions struct {
	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// CodeLensRegistrationOptions contains the options for the CodeLens handler
// registration.
type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

// CodeLensParams contains the fields sent in a `textDocument/codeLens` request.
type CodeLensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document to request CodeLens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
