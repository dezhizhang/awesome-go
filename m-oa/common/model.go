package common

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Result) Fail(code int, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}
