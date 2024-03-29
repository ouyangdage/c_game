package protocol

import "encoding/json"

type Request struct {
	//	Jsonrpc string        `json:"jsonrpc"`
	Id     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

func Unmarshal(s []byte, msg *Request) error {
	return json.Unmarshal(s, msg)
}
