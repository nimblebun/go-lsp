package lsp

// DiagnosticSeverity denotes the severity of a given problem in a document.
type DiagnosticSeverity int

const (
	// DSError reports an error.
	DSError DiagnosticSeverity = iota + 1

	// DSWarning reports a warning.
	DSWarning

	// DSInformation reports an information.
	DSInformation

	// DSHint reports a hint.
	DSHint
)

// DiagnosticTag denotes the status of a given code snippet in a document.
type DiagnosticTag int

const (
	// DTUnnecessary reports unused or unnecessary code.
	//
	// Clients are allowed to render diagnostics with this tag faded out
	// instead of having an error squiggle.
	DTUnnecessary DiagnosticTag = iota + 1

	// DTDeprecated reports deprecated or obsolete code.
	//
	// Clients are allowed to rendered diagnostics with this tag strike through.
	DTDeprecated
)

// Diagnostic represents a diagnostic, such as a compiler error or warning.
// Diagnostic objects are only valid in the scope of a resource.
type Diagnostic struct {
	// The range at which the message applies.
	Range Range `json:"range"`

	// The diagnostic's severity. Can be omitted. If omitted it is up to the
	// client to interpret diagnostics as error, warning, info or hint.
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	// The diagnostic's code, which might appear in the user interface.
	Code string `json:"code"`

	// A human-readable string describing the source of this diagnostic, e.g.
	// 'typescript' or 'super lint'.
	Source string `json:"omitempty"`

	// The diagnostic's message.
	Message string `json:"message"`

	// Additional metadata about the diagnostic.
	Tags []DiagnosticTag `json:"tags,omitempty"`

	// An array of related diagnostic information, e.g. when symbol-names within
	// a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"diagnosticRelatedInformation,omitempty"`
}

// DiagnosticRelatedInformation represents a related message and source code
// location for a diagnostic. This should be used to point to code locations
// that cause or are related to a diagnostics, for example, when duplicating a
// symbol in a scope.
type DiagnosticRelatedInformation struct {
	// The location of this related diagnostic information.
	Location Location `json:"location"`

	// The message of this related diagnostic information.
	Message string `json:"message"`
}

// PublishDiagnosticsParams contains the parameters sent in a
// `textDocument/publishDiagnostics` notification.
type PublishDiagnosticsParams struct {
	// The URI for which diagnostic information is reported.
	URI DocumentURI `json:"uri"`

	// The version number of the document the diagnostics are published for.
	// Optional.
	Version int `json:"version,omitempty"`

	// An array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}
