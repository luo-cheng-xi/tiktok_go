package util

import (
	"fmt"
	"testing"
	"tiktok/internal/conf"
)

//		func TestIDGenerator_GetNextID(t *testing.T) {
//			gen, err := NewIDGenerator(1, 1)
//			if err != nil {
//				return
//			}
//			for i := 0; i < 100; i++ {
//				fmt.Println(gen.GetNextID())
//			}
//	}
//
// https://lcx-tiktok.oss-cn-beijing.aliyuncs.com/d771e213-0a87-4877-b42b-5e48fd1b740e.txt?Expires=1692591236&OSSAccessKeyId=TMP.3KhWfqPpEuFCQ1KVXzvLwBPGXzEdb2EcwWScsGHYNMbEkhHQo9VkzJ8UsaQUz9HkhikpGJFVbrFq1TYrt7eekd2p4CcBNL&Signature=Apvmt5AUbewwJT8te4XWXItrbcw%3D
func TestOssUtil_GetUrl(t *testing.T) {
	fmt.Println(GetOssUtil(conf.GetOSSConf()).GetUrl("d771e213-0a87-4877-b42b-5e48fd1b740e.txt"))
}
