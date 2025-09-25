package utils


type AppError struct {
    Message    string            `json:"message"`
    StatusCode int               `json:"status_code"`
    // Errors     map[string]string `json:"errors,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Helper function to quickly create new errors
func NewAppError(message string, statusCode int) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

type ValidationError struct {
    Message    string            `json:"message"`
    StatusCode int               `json:"status_code"`
    Errors     map[string]string `json:"errors"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func ValidationAppError(message string, statusCode int,errors map[string]string) *ValidationError {
	return &ValidationError{
		Message:    message,
		StatusCode: statusCode,
		Errors:errors,
	}
}