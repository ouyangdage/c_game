package main

import (
	"github.com/Unknwon/goconfig"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/fhbzyc/c_game/handler"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		log.Println(err)

	}

	port, err := c.GetValue("Server", "port")
	if err != nil {
		log.Println(err)

	}

	// websocket
	http.Handle("/ws", handler.Handler{})

	server := new(http.Server)
	server.Addr = ":" + port

	if err = gracehttp.Serve(server); err != nil {
		log.Println(err)
	}
}
