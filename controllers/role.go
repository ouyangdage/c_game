package controllers

import (
	"fmt"
	"github.com/fhbzyc/c_game/models"
	"github.com/fhbzyc/c_game/network"
)

func (this Controller) GetPlayer(connect *network.Connect) error {

	if connect.Uid == 0 {
		return ReturnError(connect, lineNum(), fmt.Errorf("请选择服务器"))
	}

	if len(connect.Request.Params) != 1 {
		return ReturnError(connect, lineNum(), fmt.Errorf("Invalid method parameters"))
	}

	areaIdFloat, ok := connect.Request.Params[0].(float64)
	if !ok {
		return ReturnError(connect, lineNum(), fmt.Errorf("Invalid method parameters"))
	}
	areaId := int(areaIdFloat)

	gameArea, err := models.GameArea.FindOne(areaId)
	if err != nil {
		return ReturnError(connect, lineNum(), err)
	}

	if gameArea.AreaId == 0 {
		return ReturnError(connect, lineNum(), fmt.Errorf("Invalid method parameters, 不存在的分区"))
	}

	role, err := models.Role.Insert(connect.Uid, gameArea.AreaId)
	if err != nil {
		return ReturnError(connect, lineNum(), err)
	}

	connect.RoleId = role.RoleId
	connect.AreaId = role.AreaId
	connect.InMap()

	return ReturnSuccess(connect, role)
}
