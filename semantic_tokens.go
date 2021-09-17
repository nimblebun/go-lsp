package lsp

// SemanticTokenType represents the type of a semantic token.
type SemanticTokenType string

const (
	STTNamespace SemanticTokenType = "namespace"

	// STTType represents a generic type. Acts as a fallback for types which can't
	// be mapped to a specific type like class or enum.
	STTType SemanticTokenType = "type"

	STTClass         SemanticTokenType = "class"
	STTEnum          SemanticTokenType = "enum"
	STTInterface     SemanticTokenType = "interface"
	STTStruct        SemanticTokenType = "struct"
	STTTypeParameter SemanticTokenType = "typeParameter"
	STTParameter     SemanticTokenType = "parameter"
	STTVariable      SemanticTokenType = "variable"
	STTProperty      SemanticTokenType = "property"
	STTEnumMember    SemanticTokenType = "enumMember"
	STTEvent         SemanticTokenType = "event"
	STTFunction      SemanticTokenType = "function"
	STTMethod        SemanticTokenType = "method"
	STTMacro         SemanticTokenType = "macro"
	STTKeyword       SemanticTokenType = "keyword"
	STTModifier      SemanticTokenType = "modifier"
	STTComment       SemanticTokenType = "comment"
	STTString        SemanticTokenType = "string"
	STTNumber        SemanticTokenType = "number"
	STTRegExp        SemanticTokenType = "regexp"
	STTOperator      SemanticTokenType = "operator"
)

func (stt SemanticTokenType) String() string {
	switch stt {
	case STTNamespace:
		return "namespace"
	case STTType:
		return "type"
	case STTClass:
		return "class"
	case STTEnum:
		return "enum"
	case STTInterface:
		return "interface"
	case STTStruct:
		return "struct"
	case STTTypeParameter:
		return "typeParameter"
	case STTParameter:
		return "parameter"
	case STTVariable:
		return "variable"
	case STTProperty:
		return "property"
	case STTEnumMember:
		return "enumMember"
	case STTEvent:
		return "event"
	case STTFunction:
		return "function"
	case STTMethod:
		return "method"
	case STTMacro:
		return "macro"
	case STTKeyword:
		return "keyword"
	case STTModifier:
		return "modifier"
	case STTComment:
		return "comment"
	case STTString:
		return "string"
	case STTNumber:
		return "number"
	case STTRegExp:
		return "regexp"
	case STTOperator:
		return "operator"
	}

	return "<unknown>"
}

// SemanticTokenModifiers represents a modifier for a semantic token.
type SemanticTokenModifiers string

const (
	STMDeclaration    SemanticTokenModifiers = "declaration"
	STMDefinition     SemanticTokenModifiers = "definition"
	STMReadOnly       SemanticTokenModifiers = "readonly"
	STMStatic         SemanticTokenModifiers = "static"
	STMDeprecated     SemanticTokenModifiers = "deprecated"
	STMAbstract       SemanticTokenModifiers = "abstract"
	STMAsync          SemanticTokenModifiers = "async"
	STMModification   SemanticTokenModifiers = "modification"
	STMDocumentation  SemanticTokenModifiers = "documentation"
	STMDefaultLibrary SemanticTokenModifiers = "defaultLibrary"
)

func (stm SemanticTokenModifiers) String() string {
	switch stm {
	case STMDeclaration:
		return "declaration"
	case STMDefinition:
		return "definition"
	case STMReadOnly:
		return "readonly"
	case STMStatic:
		return "static"
	case STMDeprecated:
		return "deprecated"
	case STMAbstract:
		return "abstract"
	case STMAsync:
		return "async"
	case STMModification:
		return "modification"
	case STMDocumentation:
		return "documentation"
	case STMDefaultLibrary:
		return "defaultLibrary"
	}

	return "<unknown>"
}

// TokenFormat represents the format of a semantic token.
type TokenFormat string

const (
	TokenFormatRelative TokenFormat = "relative"
)

func (tf TokenFormat) String() string {
	switch tf {
	case TokenFormatRelative:
		return "relative"
	}

	return "<unknown>"
}

// SemanticTokensLegend represents a mapping of semantic token types and
// modifiers.
type SemanticTokensLegend struct {
	// The token types a server uses.
	TokenTypes []SemanticTokenType `json:"tokenTypes"`

	// The token modifiers a server uses.
	TokenModifiers []SemanticTokenModifiers `json:"tokenModifiers"`
}

// SemanticTokensOptions specifies the options for settin up semantic token
// support for a language server.
type SemanticTokensOptions struct {
	WorkDoneProgressOptions

	// The legend used by the server
	Legend *SemanticTokensLegend `json:"legend,omitempty"`

	// Server supports providing semantic tokens for a specific range of a
	// document.
	Range bool `json:"range,omitempty"`

	// Server supports providing semantic tokens for a full document.
	Full struct {
		// The server supports deltas for full documents.
		Delta bool `json:"delta,omitempty"`
	} `json:"full,omitempty"`
}

// SemanticTokensRegistrationOptions describes options to be used when
// registering for semantic token capabilities.
type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	StaticRegistrationOptions
}

// SemanticTokenParams contains the data the client sends through a
// `textDocument/semanticTokens/full` request.
type SemanticTokensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	TextDocument *TextDocumentIdentifier `json:"textDocument"`
}

// SemanticTokens represents a set of semantic tokens for a document.
type SemanticTokens struct {
	// An optional result id. If provided and clients support delta updating
	// the client will include the result id in the next semantic token request.
	// A server can then instead of computing all semantic tokens again simply
	// send a delta.
	ResultID string `json:"resultId,omitempty"`

	// The actual tokens.
	Data []uint `json:"data"`
}

// SemanticTokensPartialResult is similar to SemanticTokens but only contains
// the tokens (without the result ID).
//
// Note: this type is probably redundant, but the spec defines it.
type SemanticTokensPartialResult struct {
	Data []uint `json:"data"`
}

// SemanticTokensDeltaParams contains the data the client sends through a
// `textDocument/semanticTokens/full/delta` request.
type SemanticTokensDeltaParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The result id of a previous response. The result Id can either point to
	// a full response or a delta response depending on what was received last.
	PreviousResultID string `json:"previousResultId"`
}

// SemanticTokensEdit represents data about a single semantic tokens edit within
// the document.
type SemanticTokensEdit struct {
	// The start offset of the edit.
	Start uint `json:"start"`

	// The count of elements to remove.
	DeleteCount uint `json:"deleteCount"`

	// The elements to insert.
	Data []uint `json:"data,omitempty"`
}

// SemanticTokensDelta represents a set of semantic tokens edits for a document.
type SemanticTokensDelta struct {
	ResultID string `json:"resultId,omitempty"`

	// The semantic token edits to transform a previous result into a new result.
	Edits []SemanticTokensEdit `json:"edits"`
}

// SemanticTokensDeltaPartialResult is similar to SemanticTokensDelta but only
// contains the edits (without the result ID).
type SemanticTokensDeltaPartialResult struct {
	Edits []SemanticTokensEdit `json:"edits"`
}

// SemanticTokensRangeParams contains the data the client sends through a
// `textDocument/semanticTokens/range` request.
type SemanticTokensRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document
	TextDocument *TextDocumentIdentifier `json:"textDocument"`

	// The range of the document to get semantic tokens for.
	Range *Range `json:"range"`
}
