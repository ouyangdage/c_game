golang-websocket-framework

数据库ORM: [github.com/xorm][1]

Redis: [github.com/garyburd/redigo/redis][2]

配置文件: [github.com/Unknwon/goconfig][3]

日志: [github.com/Sirupsen/logrus][4]

websocket: [github.com/gorilla/websocket][5]

协议仿 jsonrpc

go get -u github.com/fhbzyc/c_game

cd $GOPATH/github.com/fhbzyc/c_game

go build

./c_game

windows 下不可用 因为有用 [github.com/facebookgo/grace/gracehttp][6] 实现优雅重启

或用:

    package main
    
    import ( 
        "github.com/Unknwon/goconfig"
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
    
        if err = Serve(server); err != nil {
            log.Println(err)
        }
    }
    
    func Serve(server *http.Server) error {
        return http.ListenAndServe(server.Addr, server.Handler)
    }

覆盖main.go

  [1]: https://github.com/xorm
  [2]: https://github.com/garyburd/redigo/redis
  [3]: https://github.com/Unknwon/goconfig
  [4]: https://github.com/Sirupsen/logrus
  [5]: https://github.com/gorilla/websocket
  [6]: https://github.com/facebookgo/grace/gracehttp