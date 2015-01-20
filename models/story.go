package models

import (
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var (
	Story storyModel
)

type storyModel struct {
}

func (this storyModel) FindAll(roleId int) ([]table.StoryTable, error) {
	var result []table.StoryTable
	return result, db.DataBase.Where("role_id = ? ORDER BY story_type ASC , story_id DESC", roleId).Find(&result)
}

func (this storyModel) FindOne(roleId, storyId int, typ table.StoryType) (*table.StoryTable, error) {
	story := new(table.StoryTable)
	_, err := db.DataBase.Where("role_id = ? AND story_id = ? AND story_type = ?", roleId, storyId, typ).Get(story)
	return story, err
}

func (this storyModel) Insert(story *table.StoryTable) error {
	_, err := db.DataBase.Insert(story)
	return err
}

func (this storyModel) Update(story *table.StoryTable) error {
	_, err := db.DataBase.Update(story)
	return err
}
