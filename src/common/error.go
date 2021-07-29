package common

type FriendlyError struct {
	Message        string
	DevMessage     string
	InnerException error
	Type           string
}

type ValidationError struct {
	Message string
	Field   string
}
