package validation

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResult struct {
	Errors []ValidationError `json:"errors"`
}
