package util

import (
	"bufio"
	"bytes"
	"strings"
)

//Substr sub string
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// JoinString  []string join
func JoinString(args ...string) string {
	var buf bytes.Buffer
	for _, v := range args {
		buf.WriteString(v)
	}
	return buf.String()
}

//Cover cover string
func Cover(s, converV string, length int) string {
	currentLen := len(s)
	for i := 0; i < length-currentLen; i++ {
		s = converV + s
	}
	return s
}

// GetLines bytes to []string
func GetLines(bts []byte) []string {
	var plans []string
	var buf bytes.Buffer
	buf.Write(bts)
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		plan := scanner.Text()
		plans = append(plans, plan)
	}
	return plans
}

//TrimPrefixSlash trim prefix /
func TrimPrefixSlash(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	if strings.HasPrefix(path, "/") {
		for {
			path = strings.TrimPrefix(path, "/")
			if !strings.HasPrefix(path, "/") {
				break
			}
		}
	}
	return path
}
