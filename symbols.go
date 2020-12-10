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
