package GORM

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"GoLib/GORM/dal"
	"GoLib/GORM/dto"
	"GoLib/GORM/model"
)

var db *gorm.DB
var log zerolog.Logger

func TestMain(t *testing.M) {
	dsn := "root:root@tcp(127.0.0.1:3306)/kratos?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Stack()
		return
	}
	log = zerolog.New(os.Stdout).With().Timestamp().Logger()

	os.Exit(t.Run())
}

func TestUserListWithGen(t *testing.T) {
	ctx := context.Background()
	query := dal.Use(db).User

	users, err := query.WithContext(ctx).Select(query.UserID, query.Username).Find()
	if err != nil {
		log.Error().Err(err).Stack()
		return
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestUserList(t *testing.T) {
	query := db.Model(&model.User{})
	var users []dto.UserInfo
	if err := query.Select("user_id,username,email").Find(&users).Error; err != nil {
		log.Error().Err(err).Stack()
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
