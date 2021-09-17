package lsp

// LinkedEditingRangeOptions specifies the options for settin up linked editing
// range support for a language server.
type LinkedEditingRangeOptions struct {
	WorkDoneProgressOptions
}

// LinkedEditingRangeRegistrationOptions describes options to be used when
// registering for linked editing range capabilities.
type LinkedEditingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	LinkedEditingRangeOptions
	StaticRegistrationOptions
}

// LinkedEditingRangeParams contains the data the client sends through a
// `textDocument/linkedEditingRange` request.
type LinkedEditingRangeParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// LinkedEditingRanges represents a set of linked editing ranges.
type LinkedEditingRanges struct {
	// A list of ranges that can be renamed together. The ranges must have
	// identical length and contain identical text content. The ranges cannot
	// overlap.
	Ranges []*Range `json:"ranges"`

	// An optional word pattern (regular expression) that describes valid contents
	// for the given ranges. If no pattern is provided, the client configuration's
	// word pattern will be used.
	WordPattern string `json:"wordPattern,omitempty"`
}
