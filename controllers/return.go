package controllers

import (
	"github.com/fhbzyc/c_game/network"
	"github.com/fhbzyc/c_game/protocol"
	"github.com/fhbzyc/c_game/libs/log"
	"runtime"
)

func lineNum() int {
	_, _, line, ok := runtime.Caller(1)
	if ok {
		return line
	}
	return -1
}

func ReturnSuccess(conn *network.Connect, v interface{}) error {

	response := new(protocol.Response)
	response.Id = conn.Request.Id
	response.Result = v

	conn.Send(protocol.MarshalOK(response))
	return nil
}

func ReturnError(conn *network.Connect, lineNum int, err error) error {

	log.Logger.Errorf("Action Error : line[%d] , %v", lineNum, err)

	message := new(protocol.Error)
	message.Id = conn.Request.Id
	message.Error = protocol.ErrorMessage{}
	message.Error.Code = lineNum
	message.Error.Message = err.Error()

	conn.Send(protocol.MarshalError(message))
	return nil
}
