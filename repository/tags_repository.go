package repository

import "golang-todo-gin/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Updates(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
