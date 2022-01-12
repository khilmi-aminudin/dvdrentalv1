package web

type ResponseWeb struct {
	Code   int64       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
