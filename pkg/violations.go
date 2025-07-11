package pkg

import (
	"encoding/json"
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

type Text struct {
	Log_id         int    `json:"log_id"`
	Error_code     int    `json:"error_code"`
	Error_msg      string `json:"error_msg"`
	Conclusion     string `json:"conclusion"`
	ConclusionType int    `json:"conclusionType"`
}

func Violation(str string) Text {
	client := censor.NewClient("KrJUmun9Y5zEmrdtmnhgtx1O", "O800Fug86yASdOfGIjt8tCbdSnTnuFxS")
	//如果是百度云ak sk,使用下面的客户端
	res := client.TextCensor(str)
	var context Text
	json.Unmarshal([]byte(res), &context)
	return context
}
