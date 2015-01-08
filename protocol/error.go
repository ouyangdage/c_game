package protocol

type Error struct {
	Id    int          `json:"id"`
	Error ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	_ = iota
	ERROR_TOKEN
)
