package service

import (
	"golang-todo-gin/data/request"
	reponse "golang-todo-gin/data/response"
	"golang-todo-gin/helper"
	"golang-todo-gin/model"
	"golang-todo-gin/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

// Create implements TagsService
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		ItemName: tags.ItemName,
		Status:   tags.Status,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll() []reponse.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []reponse.TagsResponse
	for _, value := range result {
		tag := reponse.TagsResponse{
			Id:       value.Id,
			ItemName: value.ItemName,
			Status:   value.ItemName,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(tagsId int) reponse.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := reponse.TagsResponse{
		Id:       tagData.Id,
		ItemName: tagData.ItemName,
		Status:   tagData.Status,
	}
	return tagResponse
}

// Update implements TagsService
func (t *TagsServiceImpl) Updates(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.ItemName = tags.ItemName
	tagData.Status = tags.Status
	t.TagsRepository.Updates(tagData)
}
