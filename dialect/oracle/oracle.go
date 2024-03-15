package oracle

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/sqlgen"
)

func DialectOptions() *goqu.SQLDialectOptions {
	opts := goqu.DefaultDialectOptions()

	opts.BooleanDataTypeSupported = false
	opts.UseLiteralIsBools = false

	opts.SupportsReturn = false
	opts.SupportsOrderByOnUpdate = false
	opts.SupportsLimitOnUpdate = false
	opts.SupportsLimitOnDelete = false
	opts.SupportsOrderByOnDelete = true
	opts.SupportsConflictUpdateWhere = false
	opts.SupportsInsertIgnoreSyntax = false
	opts.SupportsConflictTarget = false
	opts.SupportsWithCTE = true
	opts.SupportsWithCTERecursive = false
	opts.SupportsDistinctOn = false
	opts.SupportsWindowFunction = true
	opts.SurroundLimitWithParentheses = true

	opts.PlaceHolderFragment = []byte(":")
	opts.AsFragment = []byte(" ")
	opts.QuoteIdentifiers = true
	opts.UppercaseIdentifiers = true
	opts.IncludePlaceholderNum = true
	opts.DefaultValuesFragment = []byte("")
	opts.True = []byte("1")
	opts.False = []byte("0")

	opts.FetchFragment = []byte(" FETCH FIRST ")

	opts.SelectSQLOrder = []sqlgen.SQLFragmentType{
		sqlgen.CommonTableSQLFragment,
		sqlgen.SelectSQLFragment,
		sqlgen.FromSQLFragment,
		sqlgen.JoinSQLFragment,
		sqlgen.WhereSQLFragment,
		sqlgen.GroupBySQLFragment,
		sqlgen.HavingSQLFragment,
		sqlgen.WindowSQLFragment,
		sqlgen.CompoundsSQLFragment,
		sqlgen.OrderWithOffsetFetchSQLFragment,
		sqlgen.ForSQLFragment,
	}

	return opts
}

func init() {
	goqu.RegisterDialect("oracle", DialectOptions())
	goqu.RegisterDialect("oci8", DialectOptions())
}
