package router

import (
	"fmt"
	"github.com/6156-DonaldDuck/comments/pkg/config"
	"github.com/6156-DonaldDuck/comments/pkg/model"
	"github.com/6156-DonaldDuck/comments/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/api/v1/comments", ListAllComments)
	r.GET("/api/v1/comments/:commentId", GetCommentByCommentId)
	r.POST("/api/v1/comments", CreateComment)
	r.PUT("/api/v1/comments/:commentId", UpdateCommentById)
	r.DELETE("/api/v1/comments/:commentId", DeleteCommentById)
	r.Run(":" + config.Configuration.Port)
}

func ListAllComments(c *gin.Context) {
	comments, err := service.ListAllComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
	} else {
		c.JSON(http.StatusOK, comments)
	}
}

func GetCommentByCommentId(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		log.Errorf("[router.GetCommentByCommentId] failed to parse comment id %v, err=%v\n", commentIdStr, err)
		c.JSON(http.StatusBadRequest, "invalid comment id")
		return
	}
	comment, err := service.GetCommentByCommentId(uint(commentId))
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

func CreateComment(c *gin.Context) {
	comment := model.Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("invalid comment, err=%v", err))
		return
	}
	if id, err := service.CreateComment(comment); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error occurred while creating comment, err=%v", err))
	} else {
		c.JSON(http.StatusCreated, id)
	}
}

func UpdateCommentById(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		log.Errorf("[router.UpdateCommentById] failed to parse comment id %v, err=%v\n", commentIdStr, err)
		c.JSON(http.StatusBadRequest, "invalid comment id")
		return
	}
	comment := model.Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("invalid comment, err=%v", err))
		return
	}
	if comment.ID == 0 {
		comment.ID = uint(commentId)
	}
	if id, err := service.UpdateComment(comment); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error occurred while updating comment, err=%v", err))
	} else {
		c.JSON(http.StatusOK, id)
	}
}

func DeleteCommentById(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		log.Errorf("[router.DeleteCommentById] failed to parse comment id %v, err=%v\n", commentIdStr, err)
		c.JSON(http.StatusBadRequest, "invalid comment id")
		return
	}
	if err := service.DeleteCommentById(uint(commentId)); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error occurred while deleting comment, err=%v", err))
	} else {
		c.JSON(http.StatusOK, "success")
	}
}