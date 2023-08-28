package utils

type CustomError struct {
	Code        int
	Description string
}

func (ce *CustomError) Error() string {
	return ce.Description
}
