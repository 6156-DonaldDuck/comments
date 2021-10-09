package router

import (
	"github.com/6156-DonaldDuck/comments/pkg/config"
	"github.com/6156-DonaldDuck/comments/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/comments", ListAllComments)
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