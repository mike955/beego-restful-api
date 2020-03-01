package models

type Response struct {
	Errno int `json:"errno"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}

func NewResponse() *Response{
	return &Response{}
}