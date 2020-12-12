package lsp

// DeclarationOptions contains the options for the go-to-declaration handler.
type DeclarationOptions struct {
	WorkDoneProgressOptions
}

// DeclarationRegistrationOptions contains the options for the go-to-declaration
// handler registration.
type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// DeclarationParams contains the fields sent in a `textDocument/declaration`
// request.
type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// DefinitionOptions contains the options for the go-to-definition handler.
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

// DefinitionRegistrationOptions contains the options for the go-to-definition
// handler registration.
type DefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

// DefinitionParams contains the fields sent in a `textDocument/definition`
// request.
type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// TypeDefinitionOptions contains the options for the go-to-type-definition
// handler.
type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

// TypeDefinitionRegistrationOptions contains the options for the
// go-to-type-definition handler registration.
type TypeDefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

// TypeDefinitionParams contains the fields sent in a
// `textDocument/typeDefinition` request.
type TypeDefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// ImplementationOptions contains the options for the go-to-implementation
// handler.
type ImplementationOptions struct {
	WorkDoneProgressOptions
}

// ImplementationRegistrationOptions contains the options for the
// go-to-implementation handler registration.
type ImplementationRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

// ImplementationParams contains the fields sent in a
// `textDocument/implementation` request.
type ImplementationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// ReferenceContext defines additional context to a reference parameter.
type ReferenceContext struct {
	// Include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration,omitempty"`
}

// ReferenceOptions contains the options for the find references handler.
type ReferenceOptions struct {
	WorkDoneProgressOptions
}

// ReferenceRegistrationOptions contains the options for the find references
// handler registration.
type ReferenceRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ReferenceOptions
}

// ReferenceParams contains the fields sent in a `textDocument/references`
// request.
type ReferenceParams struct {
	WorkDoneProgressParams
	PartialResultParams

	Context ReferenceContext `json:"context"`
}
