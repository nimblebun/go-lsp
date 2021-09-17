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

func (ctk CompletionTriggerKind) String() string {
	switch ctk {
	case CTKInvoked:
		return "invoked"
	case CTKTriggerCharacter:
		return "trigger character"
	case CTKTriggerForIncompleteCompletions:
		return "trigger for incomplete completions"
	}

	return "<unknown>"
}

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

func (itf InsertTextFormat) String() string {
	switch itf {
	case ITFPlainText:
		return "plain text"
	case ITFSnippet:
		return "snippet"
	}

	return "<unknown>"
}

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
	CIKText CompletionItemKind = iota + 1

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

func (kind CompletionItemKind) String() string {
	switch kind {
	case CIKText:
		return "text"
	case CIKMethod:
		return "method"
	case CIKFunction:
		return "function"
	case CIKConstructor:
		return "constructor"
	case CIKField:
		return "field"
	case CIKVariable:
		return "variable"
	case CIKClass:
		return "class"
	case CIKInterface:
		return "interface"
	case CIKModule:
		return "module"
	case CIKProperty:
		return "property"
	case CIKUnit:
		return "unit"
	case CIKValue:
		return "value"
	case CIKEnum:
		return "enum"
	case CIKKeyword:
		return "keyword"
	case CIKSnippet:
		return "snippet"
	case CIKColor:
		return "color"
	case CIKFile:
		return "file"
	case CIKReference:
		return "reference"
	case CIKFolder:
		return "folder"
	case CIKEnumMember:
		return "enum member"
	case CIKConstant:
		return "constant"
	case CIKStruct:
		return "struct"
	case CIKEvent:
		return "event"
	case CIKOperator:
		return "operator"
	case CIKTypeParameter:
		return "type parameter"
	}

	return "<unknown>"
}

// CompletionItem describes a single completion item.
type CompletionItem struct {
	// The label of this completion item. By default
	// also the text that is inserted when selecting
	// this completion.
	Label string `json:"label,omitempty"`

	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor. The standardized set
	// of available values is defined in `CompletionItemKind`.
	Kind CompletionItemKind `json:"kind,omitempty"`

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

	// How whitespace and indentation is handled during completion item insertion.
	// If not provided the client's default value is used.
	//
	// @since 3.16.0
	InsertTextMode InsertTextMode `json:"insertTextMode,omitempty"`

	// An edit that is applied to a document when selecting this completion.
	// When an edit is provided, the value of `insertText` is ignored.
	//
	// *Note:* The range of the edit must be a single line range and it must
	// contain the position at which completion has been requested.
	//
	// TODO: add support for InsertReplaceEdit
	TextEdit *TextEdit `json:"textEdit,omitempty"`

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

// CompletionList represents a collection of `CompletionItem`s to be presented
// in the editor.
type CompletionList struct {
	// This list it not complete. Further typing should result in recomputing
	// this list.
	IsIncomplete bool `json:"isIncomplete"`

	// The completion items.
	Items []CompletionItem `json:"items"`
}

// CompletionOptions contains the options for the completion handler.
type CompletionOptions struct {
	WorkDoneProgressOptions

	// If code complete should automatically be triggered on characters
	// not being valid inside an identifier (for example `.` in JavaScript),
	// list them in `triggerCharacters`.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// The list of all possible characters that commit a completion.
	// This field can be used if clients don't support individual commit
	// characters per completion item.
	//
	// If a server provides both `allCommitCharacters` and commit characters
	// on an individual completion item, the ones on the completion item win.
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// CompletionRegistrationOptions contains the options for the completion handler
// registration.
type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CompletionOptions
}

// CompletionParams contains the fields sent in a `textDocument/completion`
// request.
type CompletionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	// The completion context.
	Context CompletionContext `json:"context,omitempty"`
}

// A special text edit to provide an insert and a replace operation.
//
// @since 3.16.0
type InsertReplaceEdit struct {
	// The string to be inserted.
	NewText string `json:"newText"`

	// The range if the insert is requested.
	Insert *Range `json:"insert"`

	// The range if the replace is requested.
	Replace *Range `json:"replace"`
}

// How whitespace and indentation is handled during completion item insertion.
//
// @since 3.16.0
type InsertTextMode int

const (
	// InsertTextModeAsIs means the insertion or replace strings is taken as it
	// is. If the value is multi line the lines below the cursor will be inserted
	// using the indentation defined in the string value. The client will not
	// apply any kind of adjustments to the string.
	InsertTextModeAsIs InsertTextMode = iota + 1

	// InsertTextModeAdjustIndentation means The editor adjusts leading whitespace
	// of new lines so that they match the indentation up to the cursor of the
	// line for which the item is accepted.
	//
	// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a multi
	// line completion item is indented using 2 tabs and all following lines
	// inserted will be indented using 2 tabs as well.
	InsertTextModeAdjustIndentation
)
