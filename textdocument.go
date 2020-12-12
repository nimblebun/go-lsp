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

// TextEdit represents a textual edit applicable to a text document.
type TextEdit struct {
	// The range of the text document to be manipulated. To insert
	// text into a document create a range where start === end.
	Range Range `json:"range"`

	// The string to be inserted. For delete operations use an
	// empty string.
	NewText string `json:"newText"`
}

// TextDocumentEdit describes textual changes on a single text document.
type TextDocumentEdit struct {
	// The text document to change.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// The edits to be applied.
	Edits []TextEdit `json:"edits"`
}

// TextDocumentSyncKind defines how the host (editor) should sync document
// changes to the language server.
type TextDocumentSyncKind int

const (
	// TDSyncKindNone means documents should not be synced at all.
	TDSyncKindNone TextDocumentSyncKind = iota

	// TDSyncKindFull means cocuments are synced by always sending the full
	// content of the document.
	TDSyncKindFull

	// TDSyncKindIncremental means documents are synced by sending the full
	// content on open. After that only incremental updates to the document are
	// sent.
	TDSyncKindIncremental
)

func (kind TextDocumentSyncKind) String() string {
	switch kind {
	case TDSyncKindNone:
		return "none"
	case TDSyncKindFull:
		return "full"
	case TDSyncKindIncremental:
		return "incremental"
	}

	return "<unknown>"
}

// TextDocumentSyncOptions specifies the options for setting up text document
// sync.
type TextDocumentSyncOptions struct {
	// Open and close notifications are sent to the server.
	// If omitted open close notification should not be sent.
	OpenClose bool `json:"openClose,omitempty"`

	// Change notifications are sent to the server.
	// If omitted, it defaults to TDSyncKindNone.
	Change TextDocumentSyncKind `json:"change,omitempty"`

	// If present will save notifications are sent to the server.
	// If omitted, the notification should not be sent.
	WillSave bool `json:"willSave,omitempty"`

	// If present will save wait until requests are sent to the server.
	// If omitted, the request should not be sent.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`

	// If present save notifications are sent to the server.
	// If omitted, the notification should not be sent.
	Save *SaveOptions `json:"save,omitempty"`
}

// DidOpenTextDocumentParams contains the data the client sends through a
// `textDocument/didOpen` request.
type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

// TextDocumentChangeRegistrationOptions describes options to be used when
// registering for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// How documents are synced to the server.
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}

// TextDocumentContentChangeEvent is an event describing a change to a text
// document.
type TextDocumentContentChangeEvent struct {
	// The range of the document that changed.
	Range *Range `json:"range,omitempty"`

	// The new text for the provided range or the whole document.
	Text string `json:"text"`
}

// DidChangeTextDocumentParams contains the data the client sends through a
// `textDocument/didChange` notification.
type DidChangeTextDocumentParams struct {
	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// The actual content changes.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentSaveReason represents reasons why a text document is saved.
type TextDocumentSaveReason int

const (
	// TDSaveReasonManual represents a document save that was manually triggered,
	// for example, by the user pressing save, by starting debugging, or by an API
	// call.
	TDSaveReasonManual TextDocumentSaveReason = iota + 1

	// TDSaveReasonAfterDelay represents a document save that was automatic after
	// a delay.
	TDSaveReasonAfterDelay

	// TDSaveReasonFocusOut represents a document save that happened when the
	// editor lost focus.
	TDSaveReasonFocusOut
)

func (reason TextDocumentSaveReason) String() string {
	switch reason {
	case TDSaveReasonManual:
		return "manual"
	case TDSaveReasonAfterDelay:
		return "after delay"
	case TDSaveReasonFocusOut:
		return "focus out"
	}

	return "<unknown>"
}

// WillSaveTextDocumentParams contains the parameters sent in a will save text
// document notification.
type WillSaveTextDocumentParams struct {
	// The document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason"`
}

// SaveOptions contains the options that need to be taken into consideration
// when saving.
type SaveOptions struct {
	// The client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}

// TextDocumentSaveRegistrationOptions contains the registration options of
// textDocument/save.
type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// The client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}

// DidSaveTextDocumentParams contains the parameters sent in a
// `textDocument/didSave` notification.
type DidSaveTextDocumentParams struct {
	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`
}

// DidCloseTextDocumentParams contains the parameters sent in a
// `textDocument/didClose` notification.
type DidCloseTextDocumentParams struct {
	// The document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
