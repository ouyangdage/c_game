package protocol

type Response struct {
	//	Jsonrpc string        `json:"jsonrpc"`
	Id int `json:"id"`
	//Service string        `json:"service"`
	//Method  string        `json:"method"`
	Result interface{} `json:"result"`
}
