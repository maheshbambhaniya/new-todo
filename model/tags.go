package model

type Tags struct{
	Id int `gorm:"type:int;primary_key"`
	ItemName string `gorm:"type:varchar(225)"`
	Status string `gorm:"type:varchar(225)"`
}