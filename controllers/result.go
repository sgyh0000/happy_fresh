package controllers

type Result struct {
	Code    int
	Success bool
	Msg     string
	Err     error
	Data    interface{}
}

func Success(data interface{}) Result {
	return Result{
		Code:    200,
		Success: true,
		Msg:     "Success",
		Err:     nil,
		Data:    data,
	}
}

func FailedCompletely(code int, msg string, err error) Result {
	return Result{
		Code:    code,
		Success: false,
		Msg:     msg,
		Err:     err,
		Data:    nil,
	}
}

func FailedWithCode(code int, err error) Result {
	return FailedCompletely(code, err.Error(), err)
}

func Failed(err error) Result {
	return FailedWithCode(500, err)
}
