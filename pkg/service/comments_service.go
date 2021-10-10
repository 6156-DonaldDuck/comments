package service

import (
	"github.com/6156-DonaldDuck/comments/pkg/db"
	"github.com/6156-DonaldDuck/comments/pkg/model"
	log "github.com/sirupsen/logrus"
)

func ListAllComments() ([]model.Comment, error) {
	var comments []model.Comment
	result := db.DbConn.Find(&comments)
	if result.Error != nil {
		log.Errorf("[service.ListAllComments] error occurred while listing comments, err=%v\n", result.Error)
	} else {
		log.Infof("[service.ListAllComments] successfully listed comments, rows affected = %v\n", result.RowsAffected)
	}
	return comments, result.Error
}

func GetCommentByCommentId(commentId uint) (model.Comment, error) {
	comment := model.Comment{}
	result := db.DbConn.First(&comment, commentId)
	if result.Error != nil {
		log.Errorf("[service.GetCommentByCommentId] error occurred while getting comment with id %v, err=%v\n", commentId, result.Error)
	} else {
		log.Infof("[service.GetCommentByCommentId] successfully got comment with id %v, rows affected = %v\n", commentId, result.RowsAffected)
	}
	return comment, result.Error
}