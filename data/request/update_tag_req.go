package request

type UpdateTagsRequest struct{
	Id int `validate:"required"`
    ItemName string `validate:"required, min=1,max=200" json="itemname"`
	Status string `validate:"required, min=1,max=200" json="status"`
}