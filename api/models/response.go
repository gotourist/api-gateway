package models

// InternalServerError ...
type InternalServerError struct {
	NonFieldErrors []string `json:"non_field_errors"`
}

// ValidationError ...
type ValidationError struct {
	NonFieldErrors []string `json:"non_field_errors"`
}

// ObjectNotFoundError ...
type ObjectNotFoundError struct {
	NonFieldErrors []string `json:"non_field_errors"`
}
