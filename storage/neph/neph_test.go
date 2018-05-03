package neph

import (
    "fmt"
    "testing"
)

func TestNew(t *testing.T) {
    s := New(map[string]string{
		"endpoint":        "oss-cn-hongkong.aliyuncs.com",
		"bucketname":      "ctrip-nephele-file-hk",
		"accessKeyId":     "LTAIB2Mh6vqyHL89",
		"accessKeySecret": "0iwE2KRLTWXzAb3URruWKqblB2ZiSK",
	})

	fmt.Println(s.File("[filename]").Bytes())
	fmt.Println(s.Iterator("", "").Next())
}
