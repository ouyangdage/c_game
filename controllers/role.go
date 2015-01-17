package controllers

import (
	"fmt"
	"github.com/fhbzyc/c_game/models"
)

func (this *Controller) GetPlayer() error {

	request := this.Request

	if this.Connect.Uid == 0 {
		return this.returnError(lineNum(), fmt.Errorf("请选择服务器"))
	}

	if len(request.Params) != 1 {
		return this.returnError(lineNum(), fmt.Errorf("Invalid method parameters"))
	}

	areaIdFloat, ok := request.Params[0].(float64)
	if !ok {
		return this.returnError(lineNum(), fmt.Errorf("Invalid method parameters"))
	}
	areaId := int(areaIdFloat)

	gameArea, err := models.GameArea.FindOne(areaId)
	if err != nil {
		return this.returnError(lineNum(), err)
	}

	if gameArea.AreaId == 0 {
		return this.returnError(lineNum(), fmt.Errorf("Invalid method parameters, 不存在的分区"))
	}

	role, err := models.Role.FindOneByArea(this.Connect.Uid, gameArea.AreaId)
	if err != nil {
		return this.returnError(lineNum(), err)
	}

	if role.RoleId == 0 {
		role, err = models.Role.Insert(this.Connect.Uid, gameArea.AreaId)
		if err != nil {
			return this.returnError(lineNum(), err)
		}
	}

	this.Connect.RoleId = role.RoleId
	this.Connect.AreaId = role.AreaId
	this.Connect.InMap()

	return this.returnSuccess(role)
}
