package exp

type (
	JoinConditionType int
	JoinCondition     interface {
		Type() JoinConditionType
		IsEmpty() bool
	}
	JoinOnCondition interface {
		JoinCondition
		On() ExpressionList
	}
	JoinUsingCondition interface {
		JoinCondition
		Using() ColumnListExpression
	}
)
