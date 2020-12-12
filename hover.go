package lsp

// Hover is the result of a hover request.
type Hover struct {
	// The hover's content.
	Contents MarkupContent `json:"contents"`

	// An optional range is a range inside a text document that is used to
	// visualize a hover, e.g. by changing the background color.
	Range *Range `json:"range,omitempty"`
}

// HoverOptions contains the options for the hover handler.
type HoverOptions struct {
	WorkDoneProgressOptions
}

// HoverRegistrationOptions contains the options for the hover handler
// registration.
type HoverRegistrationOptions struct {
	TextDocumentRegistrationOptions
	HoverOptions
}

// HoverParams contains the fields sent in a `textDocument/hover` request.
type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}
