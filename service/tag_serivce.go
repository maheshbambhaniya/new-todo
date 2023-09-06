package service

import (
	"golang-todo-gin/data/request"
	reponse "golang-todo-gin/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Updates(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) reponse.TagsResponse
	FindAll() []reponse.TagsResponse
}
