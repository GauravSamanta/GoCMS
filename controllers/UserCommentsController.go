package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Processed"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	postId := c.Param("post_id")

	userID, _ := c.Get("userID")

	commentbody := c.Request.PostFormValue("comment")

	comment := models.Comments{
		Comment: commentbody,
	}

	result := connections.DB.Create(&comment)
	if result.Error != nil {
		render.Render(c, http.StatusInternalServerError, views.Failure("Error while creating the comment", "please try again"))
		return
	}

	link := models.LinkUserPostComment{
		UserID:    userID.(uint),
		PostID:    StringToUint(postId),
		CommentID: comment.ID,
	}
	linkresult := connections.DB.Create(&link)

	if linkresult.Error != nil {
		render.Render(c, http.StatusInternalServerError, views.Failure("Error while creating the comment", "please try again"))
		return
	}

}
