package controller

import (
	"test/model"

	"github.com/labstack/echo/v4"
)

func CreateRecipe(c echo.Context) error {
	recipe := model.Recipe{}
	if err := c.Bind(&recipe); err != nil {
		return c.JSON(200, map[string]string{
			"message":  "Recipe creation failed!",
			"required": "title, making_time, serves, ingredients, cost"})
	}
	if recipe.Title == "" || recipe.MakingTime == "" || recipe.Serves == "" || recipe.Ingredients == "" || recipe.Cost == 0 {
		return c.JSON(200, map[string]string{
			"message":  "Recipe creation failed!",
			"required": "title, making_time, serves, ingredients, cost"})
	}

	model.DB.Create(&recipe)
	return c.JSON(200, map[string]interface{}{
		"message": "Recipe successfully created!",
		"recipe":  []model.Recipe{recipe},
	})
}

func GetRecipes(c echo.Context) error {
	var recipes []model.Recipe
	model.DB.Find(&recipes)
	return c.JSON(200, map[string]interface{}{
		"recipes": recipes,
	})
}

func GetRecipe(c echo.Context) error {
	id := c.Param("id")
	var recipe model.Recipe
	if err := model.DB.First(&recipe, id).Error; err != nil {
		return c.JSON(200, map[string]string{"message": "No Recipe found"})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "Recipe details by id",
		"recipe":  []model.Recipe{recipe},
	})
}

func UpdateRecipe(c echo.Context) error {
	id := c.Param("id")
	recipe := model.Recipe{}
	if err := c.Bind(&recipe); err != nil {
		return c.JSON(200, map[string]string{"message": "Recipe update failed!"})
	}
	model.DB.Model(&model.Recipe{}).Where("id = ?", id).Updates(&recipe)
	return c.JSON(200, map[string]interface{}{
		"message": "Recipe successfully updated!",
		"recipe":  []model.Recipe{recipe},
	})
}

func DeleteRecipe(c echo.Context) error {
	id := c.Param("id")
	var recipe model.Recipe
	if err := model.DB.First(&recipe, id).Error; err != nil {
		return c.JSON(200, map[string]string{"message": "No Recipe found"})
	}
	model.DB.Delete(&recipe)
	return c.JSON(200, map[string]interface{}{
		"message": "Recipe successfully removed!",
	})
}
