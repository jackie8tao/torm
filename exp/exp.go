package exp

import (
	"fmt"

	"github.com/jackie8tao/torm/sqlx"
)

type Updateable interface {
	// Used internally by update sql
	Set(interface{}) UpdateExpression
}

type (
	Vals []interface{}
	// Parent of all expression types
	Expression interface {
		Clone() Expression
		Expression() Expression
	}
	// An Expression that generates its own sql (e.g Dataset)
	SQLExpression interface {
		Expression
		ToSQL() (string, []interface{}, error)
	}

	AppendableExpression interface {
		Expression
		AppendSQL(b sqlx.SQLBuilder)
		// Returns the alias value as an identifier expression
		GetAs() IdentifierExpression

		// Returns true if this expression returns columns.
		// Used to determine if a Select, Update, Insert, or Delete query returns columns
		ReturnsColumns() bool
	}
	// Expression for Aliased expressions
	//   I("a").As("b") -> "a" AS "b"
	//   SUM("a").As(I("a_sum")) -> SUM("a") AS "a_sum"
	AliasedExpression interface {
		Expression
		// Returns the Epxression being aliased
		Aliased() Expression
		// Returns the alias value as an identiier expression
		GetAs() IdentifierExpression

		// Returns a new IdentifierExpression with the specified schema
		Schema(string) IdentifierExpression
		// Returns a new IdentifierExpression with the specified table
		Table(string) IdentifierExpression
		// Returns a new IdentifierExpression with the specified column
		Col(interface{}) IdentifierExpression
		// Returns a new IdentifierExpression with the column set to *
		//   I("my_table").All() //"my_table".*
		All() IdentifierExpression
	}

	BooleanOperation  int
	BooleanExpression interface {
		Expression
		// Returns the operator for the expression
		Op() BooleanOperation
		// The left hand side of the expression (e.g. I("a")
		LHS() Expression
		// The right hand side of the expression could be a primitive value, dataset, or expression
		RHS() interface{}
	}
	// An Expression that represents another Expression casted to a SQL type
	CastExpression interface {
		Expression
		Aliaseable
		Comparable
		Inable
		Isable
		Likeable
		Orderable
		Distinctable
		Rangeable
		// The exression being casted
		Casted() Expression
		// The the SQL type to cast the expression to
		Type() LiteralExpression
	}
	// A list of columns. Typically used internally by Select, Order, From
	ColumnListExpression interface {
		Expression
		// Returns the list of columns
		Columns() []Expression
		// Returns true if the column list is empty
		IsEmpty() bool
		// Returns a new ColumnListExpression with the columns appended.
		Append(...Expression) ColumnListExpression
	}
	CompoundType       int
	CompoundExpression interface {
		Expression
		Type() CompoundType
		RHS() AppendableExpression
	}
	// An Expression that the ON CONFLICT/ON DUPLICATE KEY portion of an INSERT statement
	ConflictAction     int
	ConflictExpression interface {
		Expression
		Action() ConflictAction
	}
	ConflictUpdateExpression interface {
		ConflictExpression
		TargetColumn() string
		Where(expressions ...Expression) ConflictUpdateExpression
		WhereClause() ExpressionList
		Update() interface{}
	}
	CommonTableExpression interface {
		Expression
		IsRecursive() bool
		// Returns the alias name for the extracted expression
		Name() LiteralExpression
		// Returns the Expression being extracted
		SubQuery() Expression
	}
	ExpressionListType int
	// A list of expressions that should be joined together
	//    And(I("a").Eq(10), I("b").Eq(11)) //(("a" = 10) AND ("b" = 11))
	//    Or(I("a").Eq(10), I("b").Eq(11)) //(("a" = 10) OR ("b" = 11))
	ExpressionList interface {
		Expression
		// Returns type (e.g. OR, AND)
		Type() ExpressionListType
		// Slice of expressions that should be joined together
		Expressions() []Expression
		// Returns a new expression list with the given expressions appended to the current Expressions list
		Append(...Expression) ExpressionList

		IsEmpty() bool
	}
	// An Identifier that can contain schema, table and column identifiers
	IdentifierExpression interface {
		Expression
		Aliaseable
		Comparable
		Inable
		Isable
		Likeable
		Rangeable
		Orderable
		Updateable
		Distinctable
		Castable
		// returns true if this identifier has more more than on part (Schema, Table or Col)
		//	"schema" -> true //cant qualify anymore
		//	"schema.table" -> true
		//	"table" -> false
		// "schema"."table"."col" -> true
		// "table"."col" -> true
		// "col" -> false
		IsQualified() bool
		// Returns a new IdentifierExpression with the specified schema
		Schema(string) IdentifierExpression
		// Returns the current schema
		GetSchema() string
		// Returns a new IdentifierExpression with the specified table
		Table(string) IdentifierExpression
		// Returns the current table
		GetTable() string
		// Returns a new IdentifierExpression with the specified column
		Col(interface{}) IdentifierExpression
		// Returns the current column
		GetCol() interface{}
		// Returns a new IdentifierExpression with the column set to *
		//   I("my_table").All() //"my_table".*
		All() IdentifierExpression

		// Returns true if schema table and identifier are all zero values.
		IsEmpty() bool
	}
	InsertExpression interface {
		Expression
		IsEmpty() bool
		IsInsertFrom() bool
		From() AppendableExpression
		Cols() ColumnListExpression
		SetCols(cols ColumnListExpression) InsertExpression
		Vals() [][]interface{}
		SetVals([][]interface{}) InsertExpression
	}

	JoinType       int
	JoinExpression interface {
		Expression
		JoinType() JoinType
		IsConditioned() bool
		Table() Expression
	}
	// Parent type for join expressions
	ConditionedJoinExpression interface {
		JoinExpression
		Condition() JoinCondition
		IsConditionEmpty() bool
	}
	LateralExpression interface {
		Expression
		Aliaseable
		Table() AppendableExpression
	}

	// Expression for representing "literal" sql.
	//  L("col = 1") -> col = 1)
	//  L("? = ?", I("col"), 1) -> "col" = 1
	LiteralExpression interface {
		Expression
		Aliaseable
		Comparable
		Isable
		Inable
		Likeable
		Rangeable
		Orderable
		// Returns the literal sql
		Literal() string
		// Arguments to be replaced within the sql
		Args() []interface{}
	}

	nullSortType  int
	sortDirection int
	// An expression for specifying sort order and options
	OrderedExpression interface {
		Expression
		// The expression being sorted
		SortExpression() Expression
		// Sort direction (e.g. ASC, DESC)
		IsAsc() bool
		// If the adapter supports it null sort type (e.g. NULLS FIRST, NULLS LAST)
		NullSortType() nullSortType
		// Returns a new OrderedExpression with NullSortType set to NULLS_FIRST
		NullsFirst() OrderedExpression
		// Returns a new OrderedExpression with NullSortType set to NULLS_LAST
		NullsLast() OrderedExpression
	}

	RangeOperation  int
	RangeExpression interface {
		Expression
		// Returns the operator for the expression
		Op() RangeOperation
		// The left hand side of the expression (e.g. I("a")
		LHS() Expression
		// The right hand side of the expression could be a primitive value, dataset, or expression
		RHS() RangeVal
	}
	RangeVal interface {
		Start() interface{}
		End() interface{}
	}

	Windowable interface {
		Over(WindowExpression) SQLWindowFunctionExpression
		OverName(IdentifierExpression) SQLWindowFunctionExpression
	}

	// Expression for representing a SQLFunction(e.g. COUNT, SUM, MIN, MAX...)
	SQLFunctionExpression interface {
		Expression
		Aliaseable
		Rangeable
		Comparable
		Isable
		Inable
		Likeable
		Windowable
		// The function name
		Name() string
		// Arguments to be passed to the function
		Args() []interface{}
	}

	UpdateExpression interface {
		Col() IdentifierExpression
		Val() interface{}
	}

	SQLWindowFunctionExpression interface {
		Expression
		Aliaseable
		Rangeable
		Comparable
		Isable
		Inable
		Likeable
		Func() SQLFunctionExpression

		Window() WindowExpression
		WindowName() IdentifierExpression

		HasWindow() bool
		HasWindowName() bool
	}

	WindowExpression interface {
		Expression

		Name() IdentifierExpression
		HasName() bool

		Parent() IdentifierExpression
		HasParent() bool
		PartitionCols() ColumnListExpression
		HasPartitionBy() bool
		OrderCols() ColumnListExpression
		HasOrder() bool

		Inherit(parent string) WindowExpression
		PartitionBy(cols ...interface{}) WindowExpression
		OrderBy(cols ...interface{}) WindowExpression
	}
	CaseElse interface {
		Result() interface{}
	}
	CaseWhen interface {
		Condition() interface{}
		Result() interface{}
	}
	CaseExpression interface {
		Expression
		Aliaseable
		GetValue() interface{}
		GetWhens() []CaseWhen
		GetElse() CaseElse
		Value(val interface{}) CaseExpression
		When(condition, result interface{}) CaseExpression
		Else(result interface{}) CaseExpression
	}
)

const (
	UnionCompoundType CompoundType = iota
	UnionAllCompoundType
	IntersectCompoundType
	IntersectAllCompoundType

	DoNothingConflictAction ConflictAction = iota
	DoUpdateConflictAction

	AndType ExpressionListType = iota
	OrType

	InnerJoinType JoinType = iota
	FullOuterJoinType
	RightOuterJoinType
	LeftOuterJoinType
	FullJoinType
	RightJoinType
	LeftJoinType
	NaturalJoinType
	NaturalLeftJoinType
	NaturalRightJoinType
	NaturalFullJoinType
	CrossJoinType

	UsingJoinCondType JoinConditionType = iota
	OnJoinCondType

	// Default null sort type with no null sort order
	NoNullsSortType nullSortType = iota
	// NULLS FIRST
	NullsFirstSortType
	// NULLS LAST
	NullsLastSortType
	// ASC
	ascDir sortDirection = iota
	// DESC
	descSortDir

	// BETWEEN
	BetweenOp RangeOperation = iota
	// NOT BETWEEN
	NotBetweenOp

	// =
	EqOp BooleanOperation = iota
	// != or <>
	NeqOp
	// IS
	IsOp
	// IS NOT
	IsNotOp
	// >
	GtOp
	// >=
	GteOp
	// <
	LtOp
	// <=
	LteOp
	//  IN
	InOp
	//  NOT IN
	NotInOp
	//  LIKE, LIKE BINARY...
	LikeOp
	//  NOT LIKE, NOT LIKE BINARY...
	NotLikeOp
	//  ILIKE, LIKE
	ILikeOp
	//  NOT ILIKE, NOT LIKE
	NotILikeOp
	// ~, REGEXP BINARY
	RegexpLikeOp
	// !~, NOT REGEXP BINARY
	RegexpNotLikeOp
	// ~*, REGEXP
	RegexpILikeOp
	// !~*, NOT REGEXP
	RegexpNotILikeOp

	betweenStr = "between"
)

var (
	ConditionedJoinTypes = map[JoinType]bool{
		InnerJoinType:      true,
		FullOuterJoinType:  true,
		RightOuterJoinType: true,
		LeftOuterJoinType:  true,
		FullJoinType:       true,
		RightJoinType:      true,
		LeftJoinType:       true,
	}
	// used internally for inverting operators
	operatorInversions = map[BooleanOperation]BooleanOperation{
		IsOp:             IsNotOp,
		EqOp:             NeqOp,
		GtOp:             LteOp,
		GteOp:            LtOp,
		LtOp:             GteOp,
		LteOp:            GtOp,
		InOp:             NotInOp,
		LikeOp:           NotLikeOp,
		ILikeOp:          NotILikeOp,
		RegexpLikeOp:     RegexpNotLikeOp,
		RegexpILikeOp:    RegexpNotILikeOp,
		IsNotOp:          IsOp,
		NeqOp:            EqOp,
		NotInOp:          InOp,
		NotLikeOp:        LikeOp,
		NotILikeOp:       ILikeOp,
		RegexpNotLikeOp:  RegexpLikeOp,
		RegexpNotILikeOp: RegexpILikeOp,
	}
)

func (bo BooleanOperation) String() string {
	switch bo {
	case EqOp:
		return "eq"
	case NeqOp:
		return "neq"
	case IsOp:
		return "is"
	case IsNotOp:
		return "isnot"
	case GtOp:
		return "gt"
	case GteOp:
		return "gte"
	case LtOp:
		return "lt"
	case LteOp:
		return "lte"
	case InOp:
		return "in"
	case NotInOp:
		return "notin"
	case LikeOp:
		return "like"
	case NotLikeOp:
		return "notlike"
	case ILikeOp:
		return "ilike"
	case NotILikeOp:
		return "notilike"
	case RegexpLikeOp:
		return "regexplike"
	case RegexpNotLikeOp:
		return "regexpnotlike"
	case RegexpILikeOp:
		return "regexpilike"
	case RegexpNotILikeOp:
		return "regexpnotilike"
	}
	return fmt.Sprintf("%d", bo)
}

func (ro RangeOperation) String() string {
	switch ro {
	case BetweenOp:
		return betweenStr
	case NotBetweenOp:
		return "not between"
	}
	return fmt.Sprintf("%d", ro)
}

func (jt JoinType) String() string {
	switch jt {
	case InnerJoinType:
		return "InnerJoinType"
	case FullOuterJoinType:
		return "FullOuterJoinType"
	case RightOuterJoinType:
		return "RightOuterJoinType"
	case LeftOuterJoinType:
		return "LeftOuterJoinType"
	case FullJoinType:
		return "FullJoinType"
	case RightJoinType:
		return "RightJoinType"
	case LeftJoinType:
		return "LeftJoinType"
	case NaturalJoinType:
		return "NaturalJoinType"
	case NaturalLeftJoinType:
		return "NaturalLeftJoinType"
	case NaturalRightJoinType:
		return "NaturalRightJoinType"
	case NaturalFullJoinType:
		return "NaturalFullJoinType"
	case CrossJoinType:
		return "CrossJoinType"
	}
	return fmt.Sprintf("%d", jt)
}
