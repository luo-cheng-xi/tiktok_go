package data

import (
	"fmt"
	"testing"
	"tiktok/internal/conf"
	"tiktok/pkg/logging"
)

func TestRelationDao_GetFollowCount(t *testing.T) {
	data, _ := NewData(conf.GetData(), logging.NewLogger())
	relationDao := NewRelationDao(logging.NewLogger(), data)
	ret := relationDao.GetFollowCount(8)
	fmt.Println(ret)
}

func TestRelationDao_Follow(t *testing.T) {
	data, _ := NewData(conf.GetData(), logging.NewLogger())
	relationDao := NewRelationDao(logging.NewLogger(), data)
	relationDao.Follow(8, 9)
}

func TestRelationDao_GetFollowerCount(t *testing.T) {
	data, _ := NewData(conf.GetData(), logging.NewLogger())
	relationDao := NewRelationDao(logging.NewLogger(), data)
	ret := relationDao.GetFollowerCount(8)
	fmt.Println(ret)
}

func TestRelationDao_IsFollow(t *testing.T) {
	data, _ := NewData(conf.GetData(), logging.NewLogger())
	relationDao := NewRelationDao(logging.NewLogger(), data)
	flag, _ := relationDao.IsFollow(9, 8)
	fmt.Println(flag)
}
