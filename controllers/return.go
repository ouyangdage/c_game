package controllers

import (
	"github.com/fhbzyc/c_game/libs/log"
	"github.com/fhbzyc/c_game/network"
	"github.com/fhbzyc/c_game/protocol"
	"runtime"
)

func (this *Controller) returnSuccess(request *protocol.Request, v interface{}) error {
	return ReturnSuccess(this.Connect, request, v)
}

func (this *Controller) returnError(request *protocol.Request, lineNum int, err error) error {
	return ReturnError(this.Connect, request, lineNum, err)
}

func ReturnSuccess(conn *network.Connect, request *protocol.Request, v interface{}) error {

	response := new(protocol.Response)
	response.Id = request.Id
	response.Result = v

	conn.Send(protocol.MarshalOK(response))
	return nil
}

func ReturnError(conn *network.Connect, request *protocol.Request, lineNum int, err error) error {

	log.Logger.Errorf("Line[%d] , %s", lineNum, err.Error())

	message := new(protocol.Error)

	if request != nil {
		message.Id = request.Id
	}

	message.Error = protocol.ErrorMessage{}
	message.Error.Code = lineNum
	message.Error.Message = err.Error()

	conn.Send(protocol.MarshalError(message))
	return nil
}

func lineNum() int {
	_, _, line, ok := runtime.Caller(1)
	if ok {
		return line
	}
	return -1
}
