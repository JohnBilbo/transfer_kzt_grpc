package c_errors

var (
	NotFoundErr     = CustomError{StatusCode: 404, ErrCode: "not_found", ErrMessage: "Resource not found"}
	BadRequestErr   = CustomError{StatusCode: 400, ErrCode: "bad_request", ErrMessage: "Bad request"}
	AttemptToRunErr = CustomError{StatusCode: 500, ErrCode: "try_to_run", ErrMessage: "Trying to run"}
)
