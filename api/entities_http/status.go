package entities

// The Status type defines a logical error model that is
// suitable for different programming environments, including
// REST APIs and RPC APIs.
type Status struct {

	// The status code, which should be an enum value
	// of [google.rpc.Code][google.rpc.Code].
	Code int32 `json:"code"`

	// A developer-facing error message, which should be in English. Any user-facing error message should be
	// localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `json:"message"`

	// A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Details []StatusDetail `json:"details"`
}

type StatusDetail struct {
	Type string `json:"type"`
}
