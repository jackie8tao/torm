package torm

import "strings"

type NodeType int

const (
	NTSelect NodeType = iota + 1
	NTFrom
	NTWhere
)

// Node sql语法部分节点
type Node interface {
	// Next 返回下一个节点，如果为链表结束，则返回错误
	Next() (Node, error)

	// Walk 遍历节点，生成sql语句
	Walk(strings.Builder) error
}
