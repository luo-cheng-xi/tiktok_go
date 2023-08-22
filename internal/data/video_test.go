package data

import (
	"fmt"
	"testing"
	"tiktok/internal/conf"
	"tiktok/pkg/logging"
)

var videoDao *VideoDao

func init() {
	data, _ := NewData(conf.GetData(), logging.NewLogger())
	videoDao = NewVideoDao(logging.NewLogger(), data)
}
func TestVideoDao_CountVideoByAuthorId(t *testing.T) {
	fmt.Println(videoDao.CountVideoByAuthorId(10))

}
