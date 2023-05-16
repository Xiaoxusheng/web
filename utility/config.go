package utility

type ResponseError struct {
	Code int    `json:"code"` //错误码
	Msg  string `json:"msg"`  //错误消息
}
