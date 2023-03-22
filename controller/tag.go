package controller

import (
	"github.com/awesomexu/go-realworld/models"
	"github.com/gin-gonic/gin"
)

func GetTags(ctx *gin.Context){
	tags, _ := models.FindTags()
	result := []string{}
	for _, tag := range *tags {
		result = append(result, tag.Name)
	}
	ctx.JSON(200, result)
}