package controller

import (
	"errors"
	"net/http"
	"zaxx/backend/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationPostInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ErrorMsg struct {
	Field string `json:"Field"`
	Message string `json:"message"`
}

func getErrorMessage(fe validator.FieldError)string{
	switch fe.Tag(){
	case "required":
		return "This Field is required"
	}
	return "Unknow Error"
}


func FindPost(c *gin.Context){
	var posts []model.Post
	model.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success" : true,
		"message" : "List Data Post",
		"data" : posts,
	})
}

func StorePost(c *gin.Context){
	var input ValidationPostInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMessage(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	post := model.Post{
		Title: input.Title,
		Content: input.Content,
	}

	model.DB.Create(&post)

	c.JSON(201, gin.H{
		"success" : true,
		"message" : "Data Berhasil Ditambah",
		"data" : post,
	})
}

func FindPostById(c *gin.Context){
	var post model.Post
	if err := model.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Data Tidak Ditemukan!"})
		return
	}

	c.JSON(200, gin.H{
		"success" : true,
		"message" : "Data Post Dengan ID : " + c.Param("id"),
		"data" : post,
	})
}

func UpdatePost(c *gin.Context){
	var post model.Post
	if err := model.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Data Tidak Ditemukan"})
		return
	}

	var input ValidationPostInput
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err,&ve){
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMessage(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	model.DB.Model(&post).Updates(input)

	c.JSON(200, gin.H{
		"success" : true,
		"Message" : "Data Post Berhail Diupdate",
		"data" : post,
	})
}

func DeletePost(c *gin.Context){
	var post model.Post
	if err := model.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Data Tidak Ditemukan"})
		return
	}

	model.DB.Delete(&post)

	c.JSON(200, gin.H{
		"success" : true,
		"message" : "Data Berhasil Dihapus",
	})
}