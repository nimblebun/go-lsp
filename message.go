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

// LogMessageParams is the payload that is sent on a `window/logMessage`
// notification.
type LogMessageParams struct {
	// The message type.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}
