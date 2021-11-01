package router

import (
	"errors"
	"github.com/6156-DonaldDuck/comments/pkg/config"
	"github.com/6156-DonaldDuck/comments/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/comments", ListAllComments)
	r.GET("/comments/:commentId", GetCommentByCommentId)
	r.Run(":" + config.Configuration.Port)
}

func ListAllComments(c *gin.Context) {
	offsetStr := c.Param("offset")
	limitStr := c.Param("limit")
	offset, errOffset := strconv.Atoi(offsetStr)
	limit, errLimit := strconv.Atoi(limitStr)
	if errOffset != nil {
		log.Errorf("[router.ListAllComments] failed to parse offset %v, err=%v\n", offsetStr, errOffset)
		c.JSON(http.StatusBadRequest, "invalid offset")
		return
	}
	if errLimit != nil {
		log.Errorf("[router.ListAllComments] failed to parse limit %v, err=%v\n", limitStr, errLimit)
		c.JSON(http.StatusBadRequest, "invalid limit")
		return
	}

	comments, err := service.ListAllComments(offset, limit)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.Error(err)
		}
	} else{
		c.JSON(http.StatusOK, comment)
	}
}