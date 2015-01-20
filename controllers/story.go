package controllers

import (
	"github.com/fhbzyc/c_game/models"
)

func (this *Controller) StoryList() error {

	roleId := this.Connect.RoleId
	list, err := models.Story.FindAll(roleId)
	if err != nil {
		return this.returnError(lineNum(), err)
	}

	return this.returnSuccess(list)
}
