package sqlx

// SelectNode select节点
type SelectNode struct {
	baseNode
	columns []string
}

// NewSelectNode 新建select节点
func NewSelectNode() *SelectNode {
	return &SelectNode{
		baseNode: baseNode{},
		columns:  nil,
	}
}

func (n *SelectNode) Walk() error {
	return nil
}
