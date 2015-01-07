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

type Handler struct{}

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

	log.Logger.Info("A new client")

	val := reflect.ValueOf(connect)
	params := make([]reflect.Value, 1, 1)
	params[0] = val
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

		request := new(protocol.Request)
		if err = protocol.Unmarshal(msg, request); err != nil {

			controllers.ReturnError(connect, lineNum(), err)
			continue
		}

		method, ok := getMethod(request)
		if !ok {
			controllers.ReturnError(connect, lineNum(), fmt.Errorf("Method not found"))
			continue
		} else {
			connect.Request = request
			method.Call(params)
		}
	}
}

func getMethod(request *protocol.Request) (reflect.Value, bool) {

	v, ok := controllers.FuncMap[request.Method]
	return v, ok
}

func lineNum() int {
	_, _, line, ok := runtime.Caller(1)
	if ok {
		return line
	}
	return -1
}
