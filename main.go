package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/fhbzyc/c_game/handler"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// websocket
	http.Handle("/ws", handler.Handler{})
}

func main() {

	c, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	port, err := c.GetValue("Server", "port")
	if err != nil {
		fmt.Println(err)
		return
	}

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = http.Serve(listen, nil); err != nil {
		fmt.Println(err)
		return
	}
}
