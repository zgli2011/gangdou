package controller

// ResponseBean : http请求的返回体
type ResponseBean struct {
	Msg     string      `json:"msg"`
	Details interface{} `json:"details"`
}
