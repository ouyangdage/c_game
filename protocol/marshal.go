package protocol

import "encoding/json"

func MarshalOK(msg *Response) []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err.Error())
	}
	return b
}

func MarshalError(msg *Error) []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err.Error())
	}
	return b
}
