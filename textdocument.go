package lsp

// TextDocumentIdentifier encapsulates the URI of a given text document.
type TextDocumentIdentifier struct {
	// The text document's URI.
	URI DocumentURI `json:"uri"`
}

// TextDocumentItem is an item to transfer a text document from the client to
// the server.
type TextDocumentItem struct {
	// The text document's URI.
	URI DocumentURI `json:"uri"`

	// The text document's language identifier.
	LanguageID string `json:"languageId"`

	// The version number of this document (it will increase after each change,
	// including undo/redo).
	Version int `json:"version"`

	// The content of the opened text document.
	Text string `json:"text"`
}

// VersionedTextDocumentIdentifier is an identifier to denote a specific version
// of a text document.
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	// The version number of this document.
	// If a versioned text document identifier is sent from the server to the
	// client and the file is not open in the editor (the server has not
	// received an open notification before), the server can send `null` to
	// indicate that the version is known and the content on disk is the
	// master (as specified with document content ownership).
	//
	// The version number of a document will increase after each change,
	// including undo/redo. The number doesn't need to be consecutive.
	Version int `json:"version,omitempty"`
}

// TextDocumentPositionParams is a parameter literal used in requests to pass a
// text document and a position inside that document.
type TextDocumentPositionParams struct {
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The position inside the text document.
	Position Position `json:"position"`
}

// DocumentFilter denotes a document through properties like `language`,
// `scheme` or `pattern`.
type DocumentFilter struct {
	// A language id, like `typescript`.
	Language string `json:"language,omitempty"`

	// A Uri [scheme](#Uri.scheme), like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`

	// A glob pattern, like `*.{ts,js}`.
	Pattern string `json:"pattern,omitempty"`
}

// DocumentSelector is the combination of one or more document filters.
type DocumentSelector []DocumentFilter
