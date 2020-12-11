package lsp

// ParameterInformation represents a parameter of a callable-signature. A
// parameter can have a label and a doc-comment.
type ParameterInformation struct {
	// The label of this parameter information.
	Label string `json:"label,omitempty"`

	// The human-readable doc-comment of this parameter. Will be shown in the UI
	// but can be omitted.
	Documentation MarkupContent `json:"documentation,omitempty"`
}

// SignatureInformation represents the signature of something callable. A
// signature can have a label, like a function-name, a doc-comment, and a set of
// parameters.
type SignatureInformation struct {
	// The label of this signature. Will be shown in the UI.
	Label string `json:"label,omitempty"`

	// The human-readable doc-comment of this signature. Will be shown in the UI
	// but can be omitted.
	Documentation MarkupContent `json:"documentation,omitempty"`

	// The parameters of this signature.
	Parameters []ParameterInformation `json:"parameters"`
}

// SignatureHelp represents the signature of something callable. There can be
// multiple signatures but only one active and only one active parameter.
type SignatureHelp struct {
	// One or more signatures. If no signatures are available the signature help
	// request should return `null`.
	Signatures []SignatureInformation `json:"signatures"`

	// The active signature. If omitted or the value lies outside the
	// range of `signatures` the value defaults to zero or is ignore if
	// the `SignatureHelp` as no signatures.
	//
	// Whenever possible implementors should make an active decision about
	// the active signature and shouldn't rely on a default value.
	//
	// In future version of the protocol this property might become
	// mandatory to better express this.
	ActiveSignature int `json:"activeSignature"`

	// The active parameter of the active signature. If omitted or the value
	// lies outside the range of `signatures[activeSignature].parameters`
	// defaults to 0 if the active signature has parameters. If
	// the active signature has no parameters it is ignored.
	// In future version of the protocol this property might become
	// mandatory to better express the active parameter if the
	// active signature does have any.
	ActiveParameter int `json:"activeParameter"`
}

// SignatureHelpTriggerKind defines how a signature help was triggered.
type SignatureHelpTriggerKind int

const (
	// SHTriggerKindInvoked means the signature help was invoked manually by the
	// user or by a command.
	SHTriggerKindInvoked SignatureHelpTriggerKind = iota + 1

	// SHTriggerKindTriggerCharacter means the signature help was triggered by a
	// trigger character.
	SHTriggerKindTriggerCharacter

	// SHTriggerKindContentChange means the signature help was triggered by the
	// cursor moving or by the document content changing.
	SHTriggerKindContentChange
)

// SignatureHelpContext contains additional information about the context in
// which a signature help request was triggered.
type SignatureHelpContext struct {
	// Action that caused signature help to be triggered.
	TriggerKind SignatureHelpTriggerKind `json:"triggerKind"`

	// Character that caused signature help to be triggered.
	// This is undefined when `TriggerKind != SHTriggerKindTriggerCharacter`.
	TriggerCharacter string `json:"triggerCharacter,omitempty"`

	// `true` if signature help was already showing when it was triggered.
	//
	// Retriggers occur when the signature help is already active and can be
	// caused by actions such as typing a trigger character, a cursor move,
	// or document content changes.
	IsRetrigger bool `json:"isRetrigger"`

	// The currently active `SignatureHelp`.
	//
	// The `ActiveSignatureHelp` has its `SignatureHelp.ActiveSignature` field
	// updated based on the user navigating through available signatures.
	ActiveSignatureHelp SignatureHelp `json:"activeSignatureHelp,omitempty"`
}

// SignatureHelpOptions contains the options for the signature help handler.
type SignatureHelpOptions struct {
	WorkDoneProgressOptions

	// The characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// List of characters that re-trigger signature help.
	//
	// These trigger characters are only active when signature help is already
	// showing.
	// All trigger characters are also counted as re-trigger characters.
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// SignatureHelpRegistrationOptions contains the options for the signature help
// handler registration.
type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

// SignatureHelpParams contains the fields sent in a
// `textDocument/signatureHelp` request.
type SignatureHelpParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams

	// The signature help context.
	// This is only available if the client specifies to send this using the
	// client capability `TextDocument.SignatureHelp.ContextSupport == true`.
	Context SignatureHelpContext `json:"context,omitempty"`
}
