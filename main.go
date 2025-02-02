package main

import (
	"test/controller"
	"test/model"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	e.POST("/recipes", controller.CreateRecipe)
	e.GET("/recipes", controller.GetRecipes)
	e.GET("/recipes/:id", controller.GetRecipe)
	e.PATCH("/recipes/:id", controller.UpdateRecipe)
	e.DELETE("/recipes/:id", controller.DeleteRecipe)
	e.Logger.Fatal(e.Start(":8080"))
}
