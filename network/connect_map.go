package network

import (
	"sync"
)

var playerMap *PlayerMap = new(PlayerMap)

type PlayerMap struct {
	Lock    *sync.RWMutex
	AreaMap map[int]map[int]*Connect //   map[区服id]map[RoleId]*Connect
}

func init() {
	playerMap.Lock = new(sync.RWMutex)
	playerMap.AreaMap = make(map[int]map[int]*Connect)
}

func (p *PlayerMap) Get(areaId, roleId int) *Connect {
	p.Lock.RLock()
	defer p.Lock.RUnlock()
	if area, ok := playerMap.AreaMap[areaId]; ok {
		if c, ok := area[roleId]; ok {
			return c
		}
	}
	return nil
}

func (p *PlayerMap) Set(areaId, roleId int, connect *Connect) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if _, ok := p.AreaMap[areaId]; !ok {
		p.AreaMap[areaId] = make(map[int]*Connect)
	}
	p.AreaMap[areaId][roleId] = connect
}

func (p *PlayerMap) Delete(areaId, roleId int) {
	p.Lock.Lock()
	defer p.Lock.Unlock()
	if _, ok := p.AreaMap[areaId]; !ok {
		return
	}
	delete(p.AreaMap[areaId], roleId)
}

func SendMessage(areaId, roleId int, s []byte) {
	playerMap.Lock.RLock()
	defer playerMap.Lock.RUnlock()
	if area, ok := playerMap.AreaMap[areaId]; ok {
		if c, ok := area[roleId]; ok {
			c.Chan <- s
		}
	}
}
