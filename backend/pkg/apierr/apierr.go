package apierr

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func New(message string, code int, details string) APIError {
	return APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}
