package main

import (
	
	"golang-todo-gin/config"
	"golang-todo-gin/helper"
	"golang-todo-gin/model"
	"golang-todo-gin/router"
	"golang-todo-gin/service"
	"golang-todo-gin/repository"
	"golang-todo-gin/controller"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Server Strated !")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	
	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		 Addr:  ":3000",
		 Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)


}