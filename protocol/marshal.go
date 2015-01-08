package protocol

import (
	"encoding/json"
)

func MarshalOK(msg *Response) []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err.Error())
	}
	return b
}

func MarshalError(requestId, code int, msg string) []byte {

	var Err Error
	Err.Id = requestId
	Err.Error.Code = code
	Err.Error.Message = msg

	b, err := json.Marshal(Err)
	if err != nil {
		panic(err.Error())
	}
	return b
}
