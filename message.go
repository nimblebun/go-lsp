package lsp

// MessageType defines the type of a message.
type MessageType int

const (
	// MTError is an error message.
	MTError MessageType = iota + 1

	// MTWarning is a warning message.
	MTWarning

	// MTInfo is an information message.
	MTInfo

	// MTLog is a log message.
	MTLog
)

func (mt MessageType) String() string {
	switch mt {
	case MTError:
		return "error"
	case MTWarning:
		return "warning"
	case MTInfo:
		return "information"
	case MTLog:
		return "log"
	}

	return "<unknown>"
}

// ShowMessageParams is the payload that is sent on a `window/showMessage`
// notification.
type ShowMessageParams struct {
	// The message type.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}

// ShowMessageRequestParams is the payload that is sent on a
// `window/showMessageRequest` request.
type ShowMessageRequestParams struct {
	// The message type.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`

	// The message action items to present.
	Actions []MessageActionItem `json:"actions,omitempty"`
}

// MessageActionItem encapsulates a message action.
type MessageActionItem struct {
	Title string `json:"title"`
}

// Params to show a document.
//
// @since 3.16.0
type ShowDocumentParams struct {
	// The URI of the document to show.
	URI string `json:"uri"`

	// Indicates to show the resource in an external program. To show for example
	// `https://code.visualstudio.com/` in the default WEB browser set `external`
	// to `true`.
	External bool `json:"external,omitempty"`

	// An optional property to indicate whether the editor showing the document
	// should take focus or not. Clients might ignore this property if an external
	// program is started.
	TakeFocus bool `json:"takeFocus,omitempty"`

	// An optional selection range if the document is a text document. Clients
	// might ignore the property if an external program is started or the file is
	// not a text file.
	Selection *Range `json:"selection,omitempty"`
}

// The result of a show document request.
//
// @since 3.16.0
type ShowDocumentResult struct {
	// A boolean indicating if the show was successful.
	Success bool `json:"success"`
}

// LogMessageParams is the payload that is sent on a `window/logMessage`
// notification.
type LogMessageParams struct {
	// The message type.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}
