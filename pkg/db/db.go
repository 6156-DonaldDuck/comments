package db

import (
	"fmt"
	"github.com/6156-DonaldDuck/comments/pkg/config"
	"github.com/6156-DonaldDuck/comments/pkg/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DbConn *gorm.DB

func init() {
	host := config.Configuration.Mysql.Host
	port := config.Configuration.Mysql.Port
	username := config.Configuration.Mysql.Username
	password := config.Configuration.Mysql.Password
	databaseName := config.Configuration.Mysql.DatabaseName
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, databaseName)
	DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("[db.init] error occurred while creating database connection, err=%v\n", err)
	}

	createTables()
}

func createTables() {
	if !DbConn.Migrator().HasTable("comments") {
		log.Infof("[db.createTables] table comments not found, creating new one\n")
		if err := DbConn.Migrator().CreateTable(&model.Comment{}); err != nil {
			log.Errorf("[db.createTables] error occurred while creating table comments, err=%v\n", err)
		}

		// insert test data
		testComment := model.Comment{
			Model: gorm.Model{
				ID: 1,
				CreatedAt: time.Now(),
			},
			ArticleId: 1,
			AuthorId: 1,
			Content: "test",
		}
		result := DbConn.Create(&testComment)
		if result.Error != nil {
			log.Errorf("[db.createTables] error occurred while inserting test data, err=%v\n", result.Error)
		} else {
			log.Infof("[db.createTables] successfully inserted test data, rows affected=%v\n", result.RowsAffected)
		}
	}
}