package _errors

var (
	UserNotFound        = &AppError{Code: "USER_NOT_FOUND", Message: "User not found"}
	UserAlreadyExists   = &AppError{Code: "USER_ALREADY_EXISTS", Message: "User already exists"}
	ErrInvalidToken     = &AppError{Code: "INVALID_TOKEN", Message: "Invalid token"}
	InvalidCredentials  = &AppError{Code: "INVALID_CREDENTIALS", Message: "Invalid email or password"}
	InternalServerError = &AppError{Code: "INTERNAL_ERROR", Message: "Internal server error"}

	ValidationError = &AppError{Code: "VALIDATION_ERROR", Message: "Validation failed"}
)

func NewValidationError(details interface{}) *AppError {
	return &AppError{
		Code:    ValidationError.Code,
		Message: ValidationError.Message,
		Details: details,
	}
}
