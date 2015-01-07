package protocol

type Error struct {
	Id    string       `json:"id"`
	Error ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
