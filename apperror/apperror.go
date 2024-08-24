package apperror

type AppError struct {
	Code int
	Msg  string
}

func (e AppError) Error() string {
	return e.String()
}

func (e AppError) String() string {
	return e.Msg
}

func New(code int, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

func Is(target error) bool {
	_, ok := target.(AppError)
	return ok
}
