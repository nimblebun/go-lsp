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
