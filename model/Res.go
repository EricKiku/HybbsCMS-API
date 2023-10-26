package model

type Res struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Token  string
}
