package lsp

// CallHierarchyOptions specifies the options for settin up call hierarchy
// support for a language server.
type CallHierarchyOptions struct {
	WorkDoneProgressOptions
}

// CallHierarchyRegistrationOptions describes options to be used when
// registering for call hierarchy capabilities.
type CallHierarchyRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CallHierarchyOptions
	StaticRegistrationOptions
}

// CallHierarchyPrepareParams contains the data the client sends through a
// `textDocument/prepareCallHierarchy` request.
type CallHierarchyPrepareParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// CallHierarchyItem represents a single item in the call hierarchy.
type CallHierarchyItem struct {
	// The name of this item.
	Name string `json:"name"`

	// The kind of this item.
	Kind SymbolKind `json:"kind"`

	// Tags for this item.
	Tags []SymbolTag `json:"tags,omitempty"`

	// More detail for this item, e.g. the signature of a function.
	Detail string `json:"detail,omitempty"`

	// The resource identifier of this item.
	URI DocumentURI `json:"uri"`

	// The range enclosing this symbol not including leading/trailing whitespace
	// but everything else, e.g. comments and code.
	Range *Range `json:"range,omitempty"`

	// The range that should be selected and revealed when this symbol is being
	// picked, e.g. the name of a function. Must be contained by the
	// [`range`](#CallHierarchyItem.range).
	SelectionRange *Range `json:"selectionRange,omitempty"`

	// A data entry field that is preserved between a call hierarchy prepare and
	// incoming calls or outgoing calls requests.
	Data interface{} `json:"data,omitempty"`
}

// CallHierarchyIncomingCallsParams contains the data the client sends through a
// `callHierarchy/incomingCalls` request.
type CallHierarchyIncomingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The item that belongs to this call hierarchy incoming call.
	Item *CallHierarchyItem `json:"item"`
}

// CallHierarchyIncomingCall represents a single incoming call.
type CallHierarchyIncomingCall struct {
	// The item that makes the call.
	From *CallHierarchyItem `json:"from"`

	// The ranges at which the calls appear. This is relative to the caller
	// denoted by [`this.from`](#CallHierarchyIncomingCall.from).
	FromRanges []*Range `json:"fromRanges,omitempty"`
}

// CallHierarchyOutgoingCallsParams contains the data the client sends through a
// `callHierarchy/outgoingCalls` request.
type CallHierarchyOutgoingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The item that belongs to this call hierarchy outgoing call.
	Item *CallHierarchyItem `json:"item"`
}

// CallHierarchyOutgoingCall represents a single outgoing call.
type CallHierarchyOutgoingCall struct {
	// The item that is called.
	To *CallHierarchyItem `json:"to"`

	// The range at which this item is called. This is the range relative to
	// the caller, e.g the item passed to `callHierarchy/outgoingCalls` request.
	FromRanges []*Range `json:"fromRanges,omitempty"`
}
