package exp

import "github.com/jackie8tao/torm/sqlx"

// Interface that an expression should implement if it can be aliased.
type Aliaseable interface {
	// Returns an AliasedExpression
	//    I("col").As("other_col") //"col" AS "other_col"
	//    I("col").As(I("other_col")) //"col" AS "other_col"
	As(interface{}) sqlx.Exp
}

// Interface that an expression should implement if it can be casted to another SQL type.
type Castable interface {
	// Casts an expression to the specified type
	//   I("a").Cast("numeric")//CAST("a" AS numeric)
	Cast(val string) sqlx.Exp
}

// Interface that an expression should implement if it can be in a list of values.
type Inable interface {
	// Creates a Boolean expression for IN clauses
	//    I("col").In([]string{"a", "b", "c"}) //("col" IN ('a', 'b', 'c'))
	In(...interface{}) sqlx.Exp
	// Creates a Boolean expression for NOT IN clauses
	//    I("col").NotIn([]string{"a", "b", "c"}) //("col" NOT IN ('a', 'b', 'c'))
	NotIn(...interface{}) sqlx.Exp
}

// Interface that an expression should implement if it can check whether it is null or false.
type Isable interface {
	// Creates an Boolean expression IS clauses
	//   ds.Where(I("a").Is(nil)) //("a" IS NULL)
	//   ds.Where(I("a").Is(true)) //("a" IS TRUE)
	//   ds.Where(I("a").Is(false)) //("a" IS FALSE)
	Is(interface{}) sqlx.Exp
	// Creates an Boolean expression IS NOT clauses
	//   ds.Where(I("a").IsNot(nil)) //("a" IS NOT NULL)
	//   ds.Where(I("a").IsNot(true)) //("a" IS NOT TRUE)
	//   ds.Where(I("a").IsNot(false)) //("a" IS NOT FALSE)
	IsNot(interface{}) sqlx.Exp
	// Shortcut for Is(nil)
	IsNull() sqlx.Exp
	// Shortcut for IsNot(nil)
	IsNotNull() sqlx.Exp
	// Shortcut for Is(true)
	IsTrue() sqlx.Exp
	// Shortcut for IsNot(true)
	IsNotTrue() sqlx.Exp
	// Shortcut for Is(false)
	IsFalse() sqlx.Exp
	// Shortcut for IsNot(false)
	IsNotFalse() sqlx.Exp
}

// Interface that an expression should implement if it can be used in regexp expression.
type Likeable interface {
	// Creates an Boolean expression for LIKE clauses
	//   ds.Where(I("a").Like("a%")) //("a" LIKE 'a%')
	Like(interface{}) sqlx.Exp
	// Creates an Boolean expression for NOT LIKE clauses
	//   ds.Where(I("a").NotLike("a%")) //("a" NOT LIKE 'a%')
	NotLike(interface{}) sqlx.Exp
	// Creates an Boolean expression for case insensitive LIKE clauses
	//   ds.Where(I("a").ILike("a%")) //("a" ILIKE 'a%')
	ILike(interface{}) sqlx.Exp
	// Creates an Boolean expression for case insensitive NOT LIKE clauses
	//   ds.Where(I("a").NotILike("a%")) //("a" NOT ILIKE 'a%')
	NotILike(interface{}) sqlx.Exp

	// Creates an Boolean expression for REGEXP LIKE clauses
	//   ds.Where(I("a").RegexpLike("a%")) //("a" ~ 'a%')
	RegexpLike(interface{}) sqlx.Exp
	// Creates an Boolean expression for REGEXP NOT LIKE clauses
	//   ds.Where(I("a").RegexpNotLike("a%")) //("a" !~ 'a%')
	RegexpNotLike(interface{}) sqlx.Exp
	// Creates an Boolean expression for case insensitive REGEXP ILIKE clauses
	//   ds.Where(I("a").RegexpILike("a%")) //("a" ~* 'a%')
	RegexpILike(interface{}) sqlx.Exp
	// Creates an Boolean expression for case insensitive REGEXP NOT ILIKE clauses
	//   ds.Where(I("a").RegexpNotILike("a%")) //("a" !~* 'a%')
	RegexpNotILike(interface{}) sqlx.Exp
}

// Interface that an expression should implement if it can be compared with other values.
type Comparable interface {
	// Creates a Boolean expression comparing equality
	//    I("col").Eq(1) //("col" = 1)
	Eq(interface{}) sqlx.Exp
	// Creates a Boolean expression comparing in-equality
	//    I("col").Neq(1) //("col" != 1)
	Neq(interface{}) sqlx.Exp
	// Creates a Boolean expression for greater than comparisons
	//    I("col").Gt(1) //("col" > 1)
	Gt(interface{}) sqlx.Exp
	// Creates a Boolean expression for greater than or equal to than comparisons
	//    I("col").Gte(1) //("col" >= 1)
	Gte(interface{}) sqlx.Exp
	// Creates a Boolean expression for less than comparisons
	//    I("col").Lt(1) //("col" < 1)
	Lt(interface{}) sqlx.Exp
	// Creates a Boolean expression for less than or equal to comparisons
	//    I("col").Lte(1) //("col" <= 1)
	Lte(interface{}) sqlx.Exp
}

// Interface that an expression should implement if it can be used in a DISTINCT epxression.
type Distinctable interface {
	// Creates a DISTINCT clause
	//   I("a").Distinct() //DISTINCT("a")
	Distinct() sqlx.Exp
}

// Interface that an expression should implement if it can be ORDERED.
type Orderable interface {
	// Creates an Ordered Expression for sql ASC order
	//   ds.Order(I("a").Asc()) //ORDER BY "a" ASC
	Asc() sqlx.Exp
	// Creates an Ordered Expression for sql DESC order
	//   ds.Order(I("a").Desc()) //ORDER BY "a" DESC
	Desc() sqlx.Exp
}