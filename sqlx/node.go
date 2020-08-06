package sqlx

import "strings"

// Node sql语法部分节点
type Node interface {
	// Walk 遍历节点，生成sql语句
	Walk(*strings.Builder) error
}

type baseNode struct {
	next Node // 下一个节点
}
