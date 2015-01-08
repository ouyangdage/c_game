package handler

import (
	"fmt"
	"github.com/fhbzyc/c_game/controllers"
	"github.com/fhbzyc/c_game/libs/log"
	"github.com/fhbzyc/c_game/network"
	"github.com/fhbzyc/c_game/protocol"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"reflect"
	"runtime"
)

type Handler struct {
	*controllers.Controller
	funcMap map[string]reflect.Value
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if err != nil {
		log.Logger.Info(err)
		return
	}

	connect := new(network.Connect)
	connect.Conn = conn
	connect.Write()

	h.Read(connect)

	connect.Close()
}

func (h Handler) Read(connect *network.Connect) {

	//	defer func() {
	//		if err := recover(); err != nil {
	//			log.Logger.Critical("Panic occur. %v", err)
	//
	//			_, message := controllers.ReturnError(lineNum(), err)
	//			connect.Send(protocol.MarshalError(message))
	//			h.Read(connect)
	//		}
	//	}()

	log.Logger.Infof("A New Client [address:%s]", connect.Conn.LocalAddr().String())

	if h.Controller == nil {

		h.Controller = new(controllers.Controller)
		h.Connect = connect
		h.Request = new(protocol.Request)
	}

	for {

		_, msg, err := connect.Conn.ReadMessage()
		if err != nil {

			if err.Error() != "websocket: close 1006 "+io.ErrUnexpectedEOF.Error() {
				log.Logger.Info(err)
			} else {
				log.Logger.Info("Client Close Connect")
			}
			return
		}

		h.Request = new(protocol.Request)
		if err = protocol.Unmarshal(msg, h.Request); err != nil {

			controllers.ReturnError(connect, nil, lineNum(), err)
			continue
		}

		h.runMethod()
	}
}

func (h *Handler) runMethod() {

	if h.funcMap == nil {

		h.funcMap = make(map[string]reflect.Value)
		value := reflect.ValueOf(h.Controller)

		numMethod := value.NumMethod()

		for i := 0; i < numMethod; i++ {

			h.funcMap[value.Type().Method(i).Name] = value.Method(i)
		}
	}

	method, ok := h.funcMap[h.Request.Method]
	if !ok {
		controllers.ReturnError(h.Connect, h.Request, lineNum(), fmt.Errorf("Method not found"))
	} else {

		method.Call([]reflect.Value{})
	}
}

func lineNum() int {
	_, _, line, ok := runtime.Caller(1)
	if ok {
		return line
	}
	return -1
}
