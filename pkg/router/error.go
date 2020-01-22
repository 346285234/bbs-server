package router

type AppError struct {
	error   error
	Message string
	Code    int
}

func (ae *AppError) Error() string {
	return ae.Message + string(ae.Code)
}

func NewAppError(e error) *AppError {
	return &AppError{e, "", 500}
}
