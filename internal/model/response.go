package model

type ResponseData struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"msg"     dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}
