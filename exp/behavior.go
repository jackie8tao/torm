package exp

// Interface that an expression should implement if it can be aliased.
type Aliaseable interface {
	// Returns an AliasedExpression
	//    I("col").As("other_col") //"col" AS "other_col"
	//    I("col").As(I("other_col")) //"col" AS "other_col"
	As(interface{}) AliasedExpression
}

// Interface that an expression should implement if it can be casted to another SQL type.
type Castable interface {
	// Casts an expression to the specified type
	//   I("a").Cast("numeric")//CAST("a" AS numeric)
	Cast(val string) CastExpression
}

// Interface that an expression should implement if it can be in a list of values.
type Inable interface {
	// Creates a Boolean expression for IN clauses
	//    I("col").In([]string{"a", "b", "c"}) //("col" IN ('a', 'b', 'c'))
	In(...interface{}) BooleanExpression
	// Creates a Boolean expression for NOT IN clauses
	//    I("col").NotIn([]string{"a", "b", "c"}) //("col" NOT IN ('a', 'b', 'c'))
	NotIn(...interface{}) BooleanExpression
}

// Interface that an expression should implement if it can check whether it is null or false.
type Isable interface {
	// Creates an Boolean expression IS clauses
	//   ds.Where(I("a").Is(nil)) //("a" IS NULL)
	//   ds.Where(I("a").Is(true)) //("a" IS TRUE)
	//   ds.Where(I("a").Is(false)) //("a" IS FALSE)
	Is(interface{}) BooleanExpression
	// Creates an Boolean expression IS NOT clauses
	//   ds.Where(I("a").IsNot(nil)) //("a" IS NOT NULL)
	//   ds.Where(I("a").IsNot(true)) //("a" IS NOT TRUE)
	//   ds.Where(I("a").IsNot(false)) //("a" IS NOT FALSE)
	IsNot(interface{}) BooleanExpression
	// Shortcut for Is(nil)
	IsNull() BooleanExpression
	// Shortcut for IsNot(nil)
	IsNotNull() BooleanExpression
	// Shortcut for Is(true)
	IsTrue() BooleanExpression
	// Shortcut for IsNot(true)
	IsNotTrue() BooleanExpression
	// Shortcut for Is(false)
	IsFalse() BooleanExpression
	// Shortcut for IsNot(false)
	IsNotFalse() BooleanExpression
}

// Interface that an expression should implement if it can be used in regexp expression.
type Likeable interface {
	// Creates an Boolean expression for LIKE clauses
	//   ds.Where(I("a").Like("a%")) //("a" LIKE 'a%')
	Like(interface{}) BooleanExpression
	// Creates an Boolean expression for NOT LIKE clauses
	//   ds.Where(I("a").NotLike("a%")) //("a" NOT LIKE 'a%')
	NotLike(interface{}) BooleanExpression
	// Creates an Boolean expression for case insensitive LIKE clauses
	//   ds.Where(I("a").ILike("a%")) //("a" ILIKE 'a%')
	ILike(interface{}) BooleanExpression
	// Creates an Boolean expression for case insensitive NOT LIKE clauses
	//   ds.Where(I("a").NotILike("a%")) //("a" NOT ILIKE 'a%')
	NotILike(interface{}) BooleanExpression

	// Creates an Boolean expression for REGEXP LIKE clauses
	//   ds.Where(I("a").RegexpLike("a%")) //("a" ~ 'a%')
	RegexpLike(interface{}) BooleanExpression
	// Creates an Boolean expression for REGEXP NOT LIKE clauses
	//   ds.Where(I("a").RegexpNotLike("a%")) //("a" !~ 'a%')
	RegexpNotLike(interface{}) BooleanExpression
	// Creates an Boolean expression for case insensitive REGEXP ILIKE clauses
	//   ds.Where(I("a").RegexpILike("a%")) //("a" ~* 'a%')
	RegexpILike(interface{}) BooleanExpression
	// Creates an Boolean expression for case insensitive REGEXP NOT ILIKE clauses
	//   ds.Where(I("a").RegexpNotILike("a%")) //("a" !~* 'a%')
	RegexpNotILike(interface{}) BooleanExpression
}

// Interface that an expression should implement if it can be compared with other values.
type Comparable interface {
	// Creates a Boolean expression comparing equality
	//    I("col").Eq(1) //("col" = 1)
	Eq(interface{}) BooleanExpression
	// Creates a Boolean expression comparing in-equality
	//    I("col").Neq(1) //("col" != 1)
	Neq(interface{}) BooleanExpression
	// Creates a Boolean expression for greater than comparisons
	//    I("col").Gt(1) //("col" > 1)
	Gt(interface{}) BooleanExpression
	// Creates a Boolean expression for greater than or equal to than comparisons
	//    I("col").Gte(1) //("col" >= 1)
	Gte(interface{}) BooleanExpression
	// Creates a Boolean expression for less than comparisons
	//    I("col").Lt(1) //("col" < 1)
	Lt(interface{}) BooleanExpression
	// Creates a Boolean expression for less than or equal to comparisons
	//    I("col").Lte(1) //("col" <= 1)
	Lte(interface{}) BooleanExpression
}

// Interface that an expression should implement if it can be used in a DISTINCT epxression.
type Distinctable interface {
	// Creates a DISTINCT clause
	//   I("a").Distinct() //DISTINCT("a")
	Distinct() SQLFunctionExpression
}

// Interface that an expression should implement if it can be ORDERED.
type Orderable interface {
	// Creates an Ordered Expression for sql ASC order
	//   ds.Order(I("a").Asc()) //ORDER BY "a" ASC
	Asc() OrderedExpression
	// Creates an Ordered Expression for sql DESC order
	//   ds.Order(I("a").Desc()) //ORDER BY "a" DESC
	Desc() OrderedExpression
}

// Interface that an expression should implement if it can be used in a between expression.
type Rangeable interface {
	// Creates a Range expression for between comparisons
	//    I("col").Between(RangeVal{Start:1, End:10}) //("col" BETWEEN 1 AND 10)
	Between(RangeVal) RangeExpression
	// Creates a Range expression for between comparisons
	//    I("col").NotBetween(RangeVal{Start:1, End:10}) //("col" NOT BETWEEN 1 AND 10)
	NotBetween(RangeVal) RangeExpression
}
