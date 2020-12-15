package types

// Default return message
type DefaultMessageReturn struct {
    Code    int    `json:"code"`
    Error   string `json:"error,omitempty"`
    Message string `json:"message,omitempty"`
}