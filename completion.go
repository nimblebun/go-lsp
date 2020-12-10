package lsp

// CompletionTriggerKind specifies how a completion was triggered.
type CompletionTriggerKind int

const (
	// CTKInvoked means the completion was triggered by typing an identifier
	// (24x7 code complete), manual invocation (e.g Ctrl+Space) or via API.
	CTKInvoked CompletionTriggerKind = iota + 1

	// CTKTriggerCharacter means the completion was triggered by a trigger
	// character specified by the `triggerCharacters` properties of
	// `CompletionRegistrationOptions`.
	CTKTriggerCharacter

	// CTKTriggerForIncompleteCompletions means the completion was re-triggered as
	// the current completion list is incomplete.
	CTKTriggerForIncompleteCompletions
)

// CompletionContext contains additional information about the context in which
// a completion request is triggered.
type CompletionContext struct {
	// How the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`

	// The trigger character (single character) that has trigger code complete.
	// Is undefined if `TriggerKind != CTKTriggerCharacter
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
}

// InsertTextFormat defines whether the insert text in a completion item should
// be interpreted as plain text or a snippet.
type InsertTextFormat int

const (
	// ITFPlainText means the primary text to be inserted is treated as a plain
	// string.
	ITFPlainText InsertTextFormat = iota + 1

	// ITFSnippet means the primary text to be inserted is treated as a snippet.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	ITFSnippet
)

// CompletionItemTag defines extra annotations that tweak the rendering of a
// completion item.
type CompletionItemTag int

// CITDeprecated will render a completion as obsolete, usually using a
// strike-out .
const CITDeprecated CompletionItemTag = 1

// CompletionItemKind specifies the type of a completion.
type CompletionItemKind int

const (
	// CIKText represents a text completion item kind.
	CIKText = iota + 1

	// CIKMethod represents a method completion item kind.
	CIKMethod

	// CIKFunction represents a function completion item kind.
	CIKFunction

	// CIKConstructor represents a constructor completion item kind.
	CIKConstructor

	// CIKField represents a field completion item kind.
	CIKField

	// CIKVariable represents a variable completion item kind.
	CIKVariable

	// CIKClass represents a class completion item kind.
	CIKClass

	// CIKInterface represents an interface completion item kind.
	CIKInterface

	// CIKModule represents a module completion item kind.
	CIKModule

	// CIKProperty represents a property completion item kind.
	CIKProperty

	// CIKUnit represents a unit completion item kind.
	CIKUnit

	// CIKValue represents a value completion item kind.
	CIKValue

	// CIKEnum represents an enum completion item kind.
	CIKEnum

	// CIKKeyword represents a keyword completion item kind.
	CIKKeyword

	// CIKSnippet represents a snippet completion item kind.
	CIKSnippet

	// CIKColor represents a color completion item kind.
	CIKColor

	// CIKFile represents a file completion item kind.
	CIKFile

	// CIKReference represents a reference completion item kind.
	CIKReference

	// CIKFolder represents a folder completion item kind.
	CIKFolder

	// CIKEnumMember represents an enum member completion item kind.
	CIKEnumMember

	// CIKConstant represents a constant completion item kind.
	CIKConstant

	// CIKStruct represents a struct completion item kind.
	CIKStruct

	// CIKEvent represents a event completion item kind.
	CIKEvent

	// CIKOperator represents an operator completion item kind.
	CIKOperator

	// CIKTypeParameter represents a type parameter completion item kind.
	CIKTypeParameter
)

// CompletionItem describes a single completion item.
type CompletionItem struct {
	// The label of this completion item. By default
	// also the text that is inserted when selecting
	// this completion.
	Label string `json:"label,omitempty"`

	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor. The standardized set
	// of available values is defined in `CompletionItemKind`.
	Kind []CompletionItemKind `json:"kind,omitempty"`

	// Tags for this completion item.
	Tags []CompletionItemTag `json:"tags,omitempty"`

	// A human-readable string with additional information
	// about this item, like type or symbol information.
	Detail string `json:"detail,omitempty"`

	// A human-readable string that represents a doc-comment.
	Documentation MarkupContent `json:"documentation,omitempty"`

	// Select this item when showing.
	//
	// *Note* that only one completion item can be selected and that the
	// tool / client decides which item that is. The rule is that the *first*
	// item of those that match best is selected.
	Preselect bool `json:"preselect,omitempty"`

	// A string that should be used when comparing this item
	// with other items. When `falsy` the label is used.
	SortText string `json:"sortText,omitempty"`

	// A string that should be used when filtering a set of
	// completion items. When `falsy` the label is used.
	FilterText string `json:"filterText,omitempty"`

	// A string that should be inserted into a document when selecting
	// this completion. When `falsy` the label is used.
	//
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally.
	// For example, VS Code when code complete is requested in this example
	// `con<cursor position>` and a completion item with an `insertText` of
	// `console` is provided, it will only insert `sole`.
	// Therefore, it is recommended to use `textEdit` instead since it avoids
	// additional client side interpretation.
	InsertText string `json:"insertText,omitempty"`

	// The format of the insert text.
	// The format applies to both the `insertText` property and the `newText`
	// property of a provided `textEdit`.
	// If omitted, defaults to `ITFPlainText`.
	InsertTextFormat InsertTextFormat `json:"insertTextFormat,omitempty"`

	// An edit that is applied to a document when selecting this completion.
	// When an edit is provided, the value of `insertText` is ignored.
	//
	// *Note:* The range of the edit must be a single line range and it must
	// contain the position at which completion has been requested.
	TextEdit TextEdit `json:"textEdit,omitempty"`

	// An optional array of additional text edits that are applied when
	// selecting this completion.
	// Edits must not overlap (including the same insert position) with the
	// main edit nor with themselves.
	//
	// Additional text edits should be used to change text unrelated to the
	// current cursor position (for example adding an import statement at the
	// top of the file if the completion item will insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`

	// An optional set of characters that when pressed, while this completion
	// is active, will accept it first and then type that character.
	// *Note* that all commit characters should have `length=1` and that
	// superfluous characters will be ignored.
	CommitCharacters []string `json:"commitCharacters,omitempty"`

	// An optional command that is executed *after* inserting this completion.
	// *Note* that additional modifications to the current document should be
	// described with the additionalTextEdits-property.
	Command Command `json:"command,omitempty"`

	// A data entry field that is preserved on a completion item between
	// a completion and a completion resolve request.
	Data interface{} `json:"data,omitempty"`
}
