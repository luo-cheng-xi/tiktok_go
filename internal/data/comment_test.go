package data

import (
	"testing"
	"tiktok/internal/conf"
	"tiktok/internal/model"
	"tiktok/pkg/logging"
)

var d, _ = NewData(conf.GetData(), logging.NewLogger())
var commentDao = NewCommentDao(logging.NewLogger(), d)

func TestCommentDao_SaveComment(t *testing.T) {
	c := model.Comment{
		VideoId: 11,
		UserId:  8,
		Content: "123456",
	}
	commentDao.SaveComment(c)
}

func TestCommentDao_DeleteComment(t *testing.T) {
	commentDao.DeleteComment(4)
}
