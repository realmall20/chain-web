package handler

type Res struct {
	Code  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func res(code int, desc string, data interface{}) Res {
	var res Res
	res.Code = code
	res.Msg = desc
	res.Data = data==nil
	return res
}
