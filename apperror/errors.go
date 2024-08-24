package apperror

var (
	//SUCC   = AppError{0, "success"}
	FAILED = AppError{1, "failed"}

	PARAM    = AppError{11, "parameter error"}
	INTER    = AppError{12, "internal error"}
	TIMEOUT  = AppError{13, "timeout error"}
	EXTERNAL = AppError{14, "external error"}
	RESRC    = AppError{15, "resource does not exist"}
	EXIST    = AppError{16, "resource already exists"}

	AuthFailed  = AppError{100, "authentication failed"}
	InvalidSign = AppError{101, "invalid signature"}

	DAL     = AppError{500, "dal level error"}
	DML     = AppError{501, "dml level error"}
	SERVICE = AppError{502, "service level error"}
	HANDLER = AppError{503, "handler level error"}
)
