package lsp

// SymbolKind specifies the type of a symbol.
type SymbolKind int

const (
	// SKFile denotes a file symbol kind.
	SKFile SymbolKind = iota + 1

	// SKModule denotes a module symbol kind.
	SKModule

	// SKNamespace denotes a namespace symbol kind.
	SKNamespace

	// SKPackage denotes a package symbol kind.
	SKPackage

	// SKClass denotes a class symbol kind.
	SKClass

	// SKMethod denotes a method symbol kind.
	SKMethod

	// SKProperty denotes a property symbol kind.
	SKProperty

	// SKField denotes a field symbol kind.
	SKField

	// SKConstructor denotes a constructor symbol kind.
	SKConstructor

	// SKEnum denotes an enum symbol kind.
	SKEnum

	// SKInterface denotes an interface symbol kind.
	SKInterface

	// SKFunction denotes a function symbol kind.
	SKFunction

	// SKVariable denotes a variable symbol kind.
	SKVariable

	// SKConstant denotes a constant symbol kind.
	SKConstant

	// SKString denotes a string symbol kind.
	SKString

	// SKNumber denotes a number symbol kind.
	SKNumber

	// SKBoolean denotes a boolean symbol kind.
	SKBoolean

	// SKArray denotes an array symbol kind.
	SKArray

	// SKObject denotes an object symbol kind.
	SKObject

	// SKKey denotes a key symbol kind.
	SKKey

	// SKNull denotes a null symbol kind.
	SKNull

	// SKEnumMember denotes an enum member symbol kind.
	SKEnumMember

	// SKStruct denotes a struct symbol kind.
	SKStruct

	// SKEvent denotes an event symbol kind.
	SKEvent

	// SKOperator denotes an operator symbol kind.
	SKOperator

	// SKTypeParameter denotes a type parameter symbol kind.
	SKTypeParameter
)

func (kind SymbolKind) String() string {
	switch kind {
	case SKFile:
		return "file"
	case SKModule:
		return "module"
	case SKNamespace:
		return "namespace"
	case SKPackage:
		return "package"
	case SKClass:
		return "class"
	case SKMethod:
		return "method"
	case SKProperty:
		return "property"
	case SKField:
		return "field"
	case SKConstructor:
		return "constructor"
	case SKEnum:
		return "enum"
	case SKInterface:
		return "interface"
	case SKFunction:
		return "function"
	case SKVariable:
		return "variable"
	case SKConstant:
		return "constant"
	case SKString:
		return "string"
	case SKNumber:
		return "number"
	case SKBoolean:
		return "boolean"
	case SKArray:
		return "array"
	case SKObject:
		return "object"
	case SKKey:
		return "key"
	case SKNull:
		return "null"
	case SKEnumMember:
		return "enum member"
	case SKStruct:
		return "struct"
	case SKEvent:
		return "event"
	case SKOperator:
		return "operator"
	case SKTypeParameter:
		return "type parameter"
	}

	return "<unknown>"
}

// Symbol tags are extra annotations that tweak the rendering of a symbol.
//
// @since 3.16
type SymbolTag int

const (
	// STDeprecated will render a symbol as obsolete, usually using a strike-out.
	STDeprecated SymbolTag = iota + 1
)

func (tag SymbolTag) String() string {
	switch tag {
	case STDeprecated:
		return "deprecated"
	}

	return "<unknown>"
}

// WorkspaceSymbolOptions is a literal for WorkDoneProgressOptions for symbols
// in a workspace.
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
}

// WorkspaceSymbolRegistrationOptions contains the workspace symbol registration
// options.
type WorkspaceSymbolRegistrationOptions struct {
	WorkspaceSymbolOptions
}

// WorkspaceSymbolParams contains the fields sent in a `workspace/symbol`
// request.
type WorkspaceSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// A query string to filter symbols by. Clients may send an empty string here
	// to request all symbols.
	Query string `json:"query"`
}

// DocumentSymbol represents programming constructs like variables, classes,
// interfaces etc. that appear in a document. Document symbols can be
// hierarchical and they have two ranges: one that encloses its definition and
// one that points to its most interesting range, for example, the range of an
// identifier.
type DocumentSymbol struct {
	// The name of this symbol.
	// Will be displayed in the user interface and therefore must not be
	// an empty string or a string only consisting of white spaces.
	Name string `json:"name"`

	// More detail for this symbol, e.g the signature of a function.
	Detail string `json:"detail,omitempty"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this document symbol.
	//
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitempty"`

	// Indicates if this symbol is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// The range enclosing this symbol not including leading/trailing
	// whitespace but everything else like comments.
	// This information is typically used to determine if the client's cursor
	// is inside the symbol to reveal in the symbol in the UI.
	Range *Range `json:"range"`

	// The range that should be selected and revealed when this symbol
	// is being picked, for example, the name of a function.
	// Must be contained by the `range`.
	SelectionRange *Range `json:"selectionRange"`

	// Children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children,omitempty"`
}

// SymbolInformation represents information about programming constructs like
// variables, classes, interfaces etc.
type SymbolInformation struct {
	// The name of this symbol
	Name string `json:"name"`

	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`

	// Tags for this symbol.
	//
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitempty"`

	// Indicates if this symbol is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// The location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in the
	// tool the range's start information is used to position the cursor. So
	// the range usually spans more then the actual symbol's name and does
	// normally include things like visibility modifiers.
	//
	// The range doesn't have to denote a node range in the sense of a abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy of
	// the symbols.
	Location Location `json:"location"`

	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName string `json:"containerName,omitempty"`
}

// DocumentSymbolOptions contains the options for the document symbol handler.
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	// A human-readable string that is shown when multiple outlines trees are
	// shown for the same document.
	//
	// @since 3.16.0
	Label string `json:"label,omitempty"`
}

// DocumentSymbolRegistrationOptions contains the options for the document
// symbol handler registration.
type DocumentSymbolRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

// DocumentSymbolParams contains the fields sent in a
// `textDocument/documentSymbol` request.
type DocumentSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
