package lsp

// DocumentURI is a document's unique resource identifier.
type DocumentURI string

// Position in a text document expressed as zero-based line and zero-based
// character offset. A position is between two characters like an ‘insert’
// cursor in a editor. Special values like for example `-1` to denote the end of
// a line are not supported.
type Position struct {
	// Line position in a document (zero-based).
	Line int `json:"line"`

	// Character offset on a line in a document (zero-based). Assuming that the
	// line is represented as a string, the `character` value represents the gap
	// between the `character` and `character + 1`.
	//
	// If the character value is greater than the line length it defaults back
	// to the line length.
	Character int `json:"character"`
}

// Range is a range in a text document expressed as (zero-based) start and end
// positions. A range is comparable to a selection in an editor. Therefore the
// end position is exclusive. If you want to specify a range that contains a
// line including the line ending character(s) then use an end position denoting
// the start of the next line.
type Range struct {
	// The range's start position.
	Start Position `json:"start"`

	// The range's end position.
	End Position `json:"end"`
}

// Location represents a location inside a resource, such as a line inside a
// text file.
type Location struct {
	URI   DocumentURI `json:"uri"`
	Range Range       `json:"range"`
}

// LocationLink represents a link between a source and a target location.
type LocationLink struct {
	// Span of the origin of this link.
	//
	// Used as the underlined span for mouse interaction.
	// Defaults to the word range at the mouse position.
	OriginSelectionRange Range `json:"originSelectionRange,omitempty"`

	// The target resource identifier of this link.
	TargetURI DocumentURI `json:"targetUri"`

	// The full target range of this link.
	// For example, if the target is a symbol, then target range is the range
	// enclosing this symbol not including leading/trailing whitespace but
	// everything else like comments.
	// This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"targetRange"`

	// The range that should be selected and revealed when this link is being
	// followed, for example, the name of a function.
	// Must be contained by the the `targetRange`.
	TargetSelectionRange Range `json:"targetSelectionRange"`
}

// DocumentHighlightKind defines a document highlight kind.
type DocumentHighlightKind int

const (
	// DocumentHighlightKindText denotes a textual occurrence.
	DocumentHighlightKindText DocumentHighlightKind = iota + 1

	// DocumentHighlightKindRead denotes read-access of a symbol, like reading a
	// variable.
	DocumentHighlightKindRead

	// DocumentHighlightKindWrite denotes write-access of a symbol, like writing
	// to a variable.
	DocumentHighlightKindWrite
)

func (hk DocumentHighlightKind) String() string {
	switch hk {
	case DocumentHighlightKindText:
		return "text"
	case DocumentHighlightKindRead:
		return "string"
	case DocumentHighlightKindWrite:
		return "write"
	}

	return "<unknown>"
}

// DocumentHighlight is a range inside a text document which deserves special
// attention. Usually a document highlight is visualized by changing the
// background color of its range.
type DocumentHighlight struct {
	// The range this highlight applies to.
	Range Range `json:"range"`

	// The highlight kind, default is DocumentHighlightKindText.
	Kind DocumentHighlightKind `json:"kind"`
}

// DocumentHighlightOptions contains the options for the document highlight
// handler.
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

// DocumentHighlightRegistrationOptions contains the options for the document
// highlight handler registration.
type DocumentHighlightRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

// DocumentHighlightParams contains the fields sent in a
// `textDocument/documentHighlight` request.
type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// DocumentLink is a range in a text document that links to an internal or
// external resource, like another text document or a web site.
type DocumentLink struct {
	// The range this link applies to.
	Range Range `json:"range"`

	// The uri this link points to. If missing a resolve request is sent later.
	Target DocumentURI `json:"target,omitempty"`

	// The tooltip text when you hover over this link.
	//
	// If a tooltip is provided, it will be displayed in a string that
	// includes instructions on how to trigger the link,
	// such as `{0} (ctrl + click)`.
	// The specific instructions vary depending on OS, user settings,
	// and localization.
	Tooltip string `json:"tooltip,omitempty"`

	// A data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data interface{} `json:"data,omitempty"`
}

// DocumentLinkOptions contains the options for the document link handler.
type DocumentLinkOptions struct {
	WorkDoneProgressOptions

	// Document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentLinkRegistrationOptions contains the options for the document link
// handler registration.
type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

// DocumentLinkParams contains the fields sent in a
// `textDocument/documentLink` request.
type DocumentLinkParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// Color represents a color in RGBA space.
type Color struct {
	// The red component of this color in the range [0-1].
	Red float64 `json:"red"`

	// The green component of this color in the range [0-1].
	Green float64 `json:"green"`

	// The blue component of this color in the range [0-1].
	Blue float64 `json:"blue"`

	// The alpha component of this color in the range [0-1].
	Alpha float64 `json:"alpha"`
}

// ColorInformation contains the color information for a given range.
type ColorInformation struct {
	// The range in the document where this color appears.
	Range Range `json:"range"`

	// The actual color value for this color range.
	Color Color `json:"color"`
}

// DocumentColorOptions contains the options for the document color handler.
type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

// DocumentColorRegistrationOptions contains the options for the document color
// handler registration.
type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
	DocumentColorOptions
}

// DocumentColorParams contains the fields sent in a
// `textDocument/documentColor` request.
type DocumentColorParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// ColorPresentation contains the presentation data of a color value.
type ColorPresentation struct {
	// The label of this color presentation. It will be shown on the color
	// picker header.
	// By default, this is also the text that is inserted when selecting
	// this color presentation.
	Label string `json:"label"`

	// A `TextEdit` which is applied to a document when selecting
	// this presentation for the color.
	// When `falsy`, the `ColorPresentation.Label` is used.
	TextEdit TextEdit `json:"textEdit,omitempty"`

	// An optional array of additional `TextEdit`s that are
	// applied when selecting this color presentation.
	// Edits must not overlap with the main `ColorPresentation.TextEdit`
	// nor with themselves.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
}

// ColorPresentationParams contains the fields sent in a
// `textDocument/colorPresentation` request.
type ColorPresentationParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The color information to request presentations for.
	Color Color `json:"color"`

	// The range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
}

// FormattingOptions is a value-object describing what options formatting should
// use.
type FormattingOptions struct {
	// Size of a tab in spaces.
	TabSize int `json:"tabSize"`

	// Prefer spaces over tabs.
	InsertSpaces bool `json:"insertSpaces"`

	// Trim trailing whitespace on a line.
	TrimTrailingWhitespace bool `json:"trimTrailingWhitespace,omitempty"`

	// Insert a newline character at the end of the file if one does not exist.
	InsertFinalNewline bool `json:"insertFinalNewline,omitempty"`

	// Trim all newlines after the final newline at the end of the file.
	TrimFinalNewlines bool `json:"trimFinalNewlines,omitempty"`
}

// DocumentFormattingOptions contains the options for the document formatting
// handler.
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

// DocumentFormattingRegistrationOptions contains the options for the document
// formatting handler registration.
type DocumentFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentFormattingOptions
}

// DocumentFormattingParams contains the fields sent in a
// `textDocument/formatting` request.
type DocumentFormattingParams struct {
	WorkDoneProgressParams

	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The format options.
	Options FormattingOptions `json:"options"`
}

// DocumentRangeFormattingOptions contains the options for the document range
// formatting handler.
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

// DocumentRangeFormattingRegistrationOptions contains the options for the
// document range formatting handler registration.
type DocumentRangeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentRangeFormattingOptions
}

// DocumentRangeFormattingParams contains the fields sent in a
// `textDocument/rangeFormatting` request.
type DocumentRangeFormattingParams struct {
	WorkDoneProgressParams

	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The range to format.
	Range Range `json:"range"`

	// The format options.
	Options FormattingOptions `json:"options"`
}

// DocumentOnTypeFormattingOptions contains the options for the document on-type
// formatting handler.
type DocumentOnTypeFormattingOptions struct {
	// A character on which formatting should be triggered, like `}`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

// DocumentOnTypeFormattingRegistrationOptions contains the options for the
// document on-type formatting handler registration.
type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

// DocumentOnTypeFormattingParams contains the fields sent in a
// `textDocument/onTypeFormatting` request.
type DocumentOnTypeFormattingParams struct {
	TextDocumentPositionParams

	// The character that has been typed.
	Character string `json:"ch"`

	// The format options.
	Options FormattingOptions `json:"options"`
}

// FoldingRangeKind contains the known range kinds.
type FoldingRangeKind string

const (
	// FoldingRangeKindComment denotes that the folding range is for a comment.
	FoldingRangeKindComment FoldingRangeKind = "comment"

	// FoldingRangeKindImports denotes that the folding range is for imports or
	// includes.
	FoldingRangeKindImports = "imports"

	// FoldingRangeKindRegion denotes that the folding range is for a region
	// (e.g. `#region`).
	FoldingRangeKindRegion = "region"
)

// FoldingRange represents a folding range for the client.
type FoldingRange struct {
	// The zero-based line number from where the folded range starts.
	StartLine int `json:"startLine"`

	// The zero-based character offset from where the folded range starts.
	// If not defined, defaults to the length of the start line.
	StartCharacter int `json:"startCharacter,omitempty"`

	// The zero-based line number where the folded range ends.
	EndLine int `json:"endLine"`

	// The zero-based character offset before the folded range ends.
	// If not defined, defaults to the length of the end line.
	EndCharacter int `json:"endCharacter,omitempty"`

	// Describes the kind of the folding range such as `comment` or `region`.
	// The kind is used to categorize folding ranges and used by commands
	// like 'Fold all comments'.
	// See FoldingRangeKind for an enumeration of standardized kinds.
	Kind FoldingRangeKind `json:"kind,omitempty"`
}

// FoldingRangeOptions contains the options for the folding range provider
// handler.
type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}

// FoldingRangeRegistrationOptions contains the options for the folding range
// provider handler registration.
type FoldingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	StaticRegistrationOptions
}

// FoldingRangeParams contains the fields sent in a `textDocument/foldingRange`
// request.
type FoldingRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// SelectionRangeOptions contains the options for the selection range provider
// handler.
type SelectionRangeOptions struct {
	WorkDoneProgressOptions
}

// SelectionRangeRegistrationOptions contains the options for the selection
// range provider handler registration.
type SelectionRangeRegistrationOptions struct {
	SelectionRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// SelectionRangeParams contains the fields sent in a
// `textDocument/selectionRange` request.
type SelectionRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The positions inside the text document.
	Positions []Position `json:"positions"`
}
