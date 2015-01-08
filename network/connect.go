package network

import (
	"github.com/fhbzyc/c_game/libs/log"
	"github.com/gorilla/websocket"
)

const ()

type Connect struct {
	Uid    int
	AreaId int
	RoleId int
	Conn   *websocket.Conn
	Chan   chan []byte
}

func (this *Connect) Send(s []byte) error {
	this.Chan <- s
	return nil
}

func (this *Connect) Write() {
	this.Chan = make(chan []byte, 10)
	go func() {
		for s := range this.Chan {
			if err := this.Conn.WriteMessage(websocket.TextMessage, s); err != nil {
				log.Logger.Warn("Send Message Error: ", err)
			} else {
				log.Logger.Info("Send Success")
			}
		}
	}()
}

func (this *Connect) Close() {

	playerMap.Delete(this.AreaId, this.RoleId)
	this.Conn.Close()
	close(this.Chan)
}

func (this *Connect) InMap() {
	playerMap.Set(this.AreaId, this.RoleId, this)
}
