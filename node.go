package torm

import (
	"fmt"
	"strings"
)

type NodeType int

// 节点类型标志
const (
	NTSelect NodeType = iota + 1
	NTFrom
	NTWhere
)

// Node sql语法部分节点
type Node interface {
	// Walk 遍历节点，生成sql语句
	Walk(*strings.Builder) error
}

// BaseNode 基础节点
type BaseNode struct {
	// Type 节点类型
	Type NodeType

	// Next 下一个节点
	Next Node
}

// SelectNode select节点
type SelectNode struct {
	BaseNode
	columns []string
}

func (s *SelectNode) Walk(sql *strings.Builder) (err error) {
	_, err = sql.WriteString("SELECT ")
	if err != nil {
		return
	}

	if len(s.columns) <= 0 {
		sql.WriteString("* ")
	} else {
		for _, v := range s.columns {
			sql.WriteString(fmt.Sprintf("%s, ", v))
		}
	}

	err = s.Next.Walk(sql)
	return
}

// FromNode from节点
type FromNode struct {
	BaseNode
	table string
}

func (f *FromNode) Walk(sql *strings.Builder) (err error) {
	_, err = sql.WriteString(fmt.Sprintf("FROM %s ", f.table))
	if err != nil {
		return
	}

	err = f.Next.Walk(sql)
	return
}

// WhereNode where节点
type WhereNode struct {
	BaseNode
	cond string
}

func (w *WhereNode) Walk(sql *strings.Builder) (err error) {
	_, err = sql.WriteString(fmt.Sprintf("WHERE %s ", w.cond))
	if err != nil {
		return
	}

	err = w.Next.Walk(sql)
	return
}
