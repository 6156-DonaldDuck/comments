package router

import (
	"fmt"
	docs "github.com/6156-DonaldDuck/comments/docs"
	"github.com/6156-DonaldDuck/comments/pkg/config"
	"github.com/6156-DonaldDuck/comments/pkg/model"
	"github.com/6156-DonaldDuck/comments/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
)

func InitRouter() {
	r := gin.Default()
	r.Use(cors.Default()) // default allows all origin
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/api/v1/comments", ListAllComments)
	r.GET("/api/v1/comments/:commentId", GetCommentByCommentId)
	r.POST("/api/v1/comments", CreateComment)
	r.PUT("/api/v1/comments/:commentId", UpdateCommentById)
	r.DELETE("/api/v1/comments/:commentId", DeleteCommentById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + config.Configuration.Port)
}

// @BasePath /api/v1

// ListAllComments godoc
// @Summary List All Comments
// @Schemes
// @Description List all comments of all articles and all users
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {json} Comments
// @Failure 500 Internal error
// @Router /comments [get]
func ListAllComments(c *gin.Context) {
	articleIdStr := c.DefaultQuery("article_id", "0")

	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		log.Errorf("[router.ListAllComments] failed to parse articleId %v, err=%v\n", articleIdStr, err)
		c.JSON(http.StatusBadRequest, "invalid articleId")
		return
	}

	comments, err := service.ListAllComments(uint(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
	} else {
		c.JSON(http.StatusOK, comments)
	}
}

// @BasePath /api/v1

// GetCommentByCommentId godoc
// @Summary Get Comment By Comment ID
// @Schemes
// @Description Get a specific comment by the given ID
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {json} Comment
// @Failure 400 invalid comment id
// @Router /comments/{commentId} [get]
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

// @BasePath /api/v1

// CreateComment godoc
// @Summary Create Comment
// @Schemes
// @Description Create a new comment
// @Tags Comments
// @Accept json
// @Produce json
// @Success 201 {json} Comment
// @Failure 400 invalid comment
// @Router /comments [post]
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

// @BasePath /api/v1

// UpdateCommentById godoc
// @Summary Update an Existing Comment
// @Schemes
// @Description Update a existing comment
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {json} Comment
// @Failure 400 invalid comment
// @Router /comments/{commentId} [put]
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

// @BasePath /api/v1

// DeleteCommentById godoc
// @Summary Delete an Existing Comment
// @Schemes
// @Description Delete a existing comment
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {json} Comment
// @Failure 400 invalid comment
// @Router /comments/{commentId} [delete]
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
		c.JSON(http.StatusOK, commentId)
	}
}