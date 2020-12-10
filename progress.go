package lsp

// ProgressToken represents a token provided by the client or server.
type ProgressToken string

// ProgressParams denotes a token-value pair for a progress.
type ProgressParams struct {
	// The progress token provided by the client or server.
	Token ProgressToken `json:"token"`

	// The progress data.
	Value interface{} `json:"value"`
}

// WorkDoneProgressBegin is the payload that must be sent in a `$/progress`
// notification.
type WorkDoneProgressBegin struct {
	// Should always be "begin".
	Kind string `json:"kind"`

	// Mandatory title of the progress operation. Used to briefly inform about
	// the kind of operation being performed.
	Title string `json:"title"`

	// Controls if a cancel button should show to allow the user to cancel the
	// long running operation.
	// Clients that don't support cancellation can ignore the setting.
	Cancellable bool `json:"cancellable,omitempty"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message"`

	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule.
	Percentage int `json:"percentage"`
}

// WorkDoneProgressReport is the payload that needs to be sent to report
// progress.
type WorkDoneProgressReport struct {
	// Should always be "report".
	Kind string `json:"kind"`

	// Controls enablement state of a cancel button. T
	// This property is only valid if a cancel button is requested in
	// the `WorkDoneProgressBegin` payload.
	//
	// Clients that don't support cancellation or don't support controlling
	// the button's enablement state are allowed to ignore the setting.
	Cancellable bool `json:"cancellable,omitempty"`

	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`

	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule.
	Percentage int `json:"percentage,omitempty"`
}

// WorkDoneProgressEnd is the payload that needs to be sent to signal the end of
// a progress reporting.
type WorkDoneProgressEnd struct {
	// Should always be "end".
	Kind string `json:"kind"`

	// Optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message string `json:"message,omitempty"`
}

// WorkDoneProgressParams contains the parameters of a progress report.
type WorkDoneProgressParams struct {
	// An optional token that a server can use to report work done progress.
	WorkDoneToken ProgressToken `json:"workDoneToken,omitempty"`
}

// WorkDoneProgressOptions is the type definition for the server capability.
type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// WorkDoneProgressCreateParams is the request payload sent in a
// `window/workDoneProgress/create` request.
type WorkDoneProgressCreateParams struct {
	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}

// WorkDoneProgressCancelParams is the request payload sent in a
// `window/workDoneProgress/cancel` request.
type WorkDoneProgressCancelParams struct {
	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}
