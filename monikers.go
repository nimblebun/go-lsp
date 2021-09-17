package lsp

// MonikerOptions specifies the options for settin up moniker support for a
// language server.
type MonikerOptions struct {
	WorkDoneProgressOptions
}

// MonikerRegistrationOptions describes options to be used when registering for
// moniker capabilities.
type MonikerRegistrationOptions struct {
	TextDocumentRegistrationOptions
	MonikerOptions
}

// MonikerParams contains the data the client sends through a
// `textDocument/moniker` request.
type MonikerParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// UniquenessLevel represents a level of uniqueness for a moniker.
type UniquenessLevel string

const (
	UniquenessLevelDocument UniquenessLevel = "document"
	UniquenessLevelProject  UniquenessLevel = "project"
	UniquenessLevelGroup    UniquenessLevel = "group"
	UniquenessLevelScheme   UniquenessLevel = "scheme"
	UniquenessLevelGlobal   UniquenessLevel = "global"
)

func (level UniquenessLevel) String() string {
	switch level {
	case UniquenessLevelDocument:
		return "document"
	case UniquenessLevelProject:
		return "project"
	case UniquenessLevelGroup:
		return "group"
	case UniquenessLevelScheme:
		return "scheme"
	case UniquenessLevelGlobal:
		return "global"
	}

	return "<unknown>"
}

// MonikerKind represents the kind of moniker.
type MonikerKind string

const (
	MonikerKindImport MonikerKind = "import"
	MonikerKindExport MonikerKind = "export"
	MonikerKindLocal  MonikerKind = "local"
)

func (kind MonikerKind) String() string {
	switch kind {
	case MonikerKindImport:
		return "import"
	case MonikerKindExport:
		return "export"
	case MonikerKindLocal:
		return "local"
	}

	return "<unknown>"
}

// Moniker represents a single moniker.
type Moniker struct {
	// The scheme of the moniker. For example tsc or .Net
	Scheme string `json:"scheme"`

	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`

	// The scope in which the moniker is unique
	Unique UniquenessLevel `json:"unique"`

	// The moniker kind if known.
	Kind MonikerKind `json:"kind,omitempty"`
}
