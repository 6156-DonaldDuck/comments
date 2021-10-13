package service

import (
	"fmt"
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
		log.Infof("[service.ListAllComments] successfully listed comments, rows affected=%v\n", result.RowsAffected)
	}
	return comments, result.Error
}

func GetCommentByCommentId(commentId uint) (model.Comment, error) {
	comment := model.Comment{}
	result := db.DbConn.First(&comment, commentId)
	if result.Error != nil {
		log.Errorf("[service.ListAllComments] error occurred while getting comment with id %v, err=%v\n", commentId, result.Error)
	} else {
		log.Infof("[service.ListAllComments] successfully got comment with id %v, rows affected=%v\n", commentId, result.RowsAffected)
	}
	return comment, result.Error
}

func CreateComment(comment model.Comment) (uint, error) {
	result := db.DbConn.Create(&comment)
	if result.Error != nil {
		log.Errorf("[service.CreateComment] error occurred while creating comment, err=%v\n", result.Error)
		return 0, result.Error
	}
	log.Infof("[service.CreateComment] successfully created the comment with id=%v\n", comment.ID)
	return comment.ID, nil
}

func UpdateComment(comment model.Comment) (uint, error) {
	if comment.ID == 0 {
		err := fmt.Errorf("[service.UpdateComment] comment id cannot be 0")
		return 0, err
	}
	result := db.DbConn.Model(&comment).Updates(comment)
	if result.Error != nil {
		log.Errorf("[service.UpdateComment] error occurred while updating comment %v, err=%v\n", comment.ID, result.Error)
		return 0, result.Error
	}
	log.Infof("[service.UpdateComment] successfully updated the comment with id=%v\n", comment.ID)
	return comment.ID, nil
}

func DeleteCommentById(commentId uint) error {
	result := db.DbConn.Delete(&model.Comment{}, commentId)
	if result.Error != nil {
		log.Errorf("[service.DeleteCommentById] error occurred while deleting comment %v, err=%v\n", commentId, result.Error)
		return result.Error
	}
	log.Infof("[service.DeleteCommentById] successfully deleted comment with id %v, rows affected=%v\n", commentId, result.RowsAffected)
	return nil
}