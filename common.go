package lsp

// StaticRegistrationOptions can be used to register a feature in the initialize
// result with a given server control ID to be able to un-register the feature
// later on.
type StaticRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	ID string `json:"id,omitempty"`
}

// TextDocumentRegistrationOptions denotes options to dynamically register for
// requests for a set of text documents
type TextDocumentRegistrationOptions struct {
	// A document selector to identify the scope of the registration.
	// If not provided, the document selector provided on the client side
	// will be used.
	DocumentSelector DocumentSelector `json:"documentSelector,omitempty"`
}

// MarkupKind describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
type MarkupKind string

const (
	// MKPlainText denotes that plaintext is supported as a content format.
	MKPlainText MarkupKind = "plaintext"

	// MKMarkdown denotes that markdown is supported as a content format.
	MKMarkdown = "markdown"
)

// MarkupContent represents a string value, which content is interpreted based
// on its kind flag.
type MarkupContent struct {
	// The type of the Markup.
	Kind MarkupKind `json:"kind"`

	// The content itself.
	Value string `json:"value"`
}

// PartialResultParams is a parameter literal used to pass a partial result
// token.
type PartialResultParams struct {
	// An optional token that a server can use to report partial results
	// (for example, streaming) to the client.
	PartialResultToken ProgressToken `json:"partialResultToken,omitempty"`
}
