package c_errors

type CError interface {
	Error() string
	GetStatusCode() int
}

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	ErrCode    string `json:"errCode"`
	ErrMessage string `json:"errMessage"`
}

func (e CustomError) Error() string {
	return e.ErrMessage
}

func (e CustomError) GetStatusCode() int {
	return e.StatusCode
}
