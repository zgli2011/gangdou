package controller

// ResponseBean : http请求的返回体
type ResponseBean struct {
	Msg     string      `json:"msg"`
	Status  int64       `json:"status"`
	Details interface{} `json:"details"`
}
