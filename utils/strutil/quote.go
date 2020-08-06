package strutil

import "fmt"

// Quote 添加字符串引用标志
func Quote(s string) string {
	return fmt.Sprintf("`%s`", s)
}
