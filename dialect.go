package hood

type Dialect interface {
	// Marker returns the dialect specific markers for prepared statements,
	// for instance $1, $2, ... and so on. The position starts at 0.
	Marker(pos int) string

	// SqlType returns the SQL type for the provided interface type. The size
	// parameter delcares the data size for the column (e.g. for VARCHARs).
	SqlType(f interface{}, size int) string

	// Insert takes the generated query and modifies it. E.g. Postgres does not
	// return the inserted IDs after executing INSERT, unless a RETURNING
	// statement is appended. If a dialect needs a custom INSERT, it should return
	// implemented == true.
	Insert(hood *Hood, model *Model, query string, args ...interface{}) (id Id, err error, implemented bool)

	// StmtNotNull returns the dialect specific statement for 'NOT NULL'.
	StmtNotNull() string

	// StmtDefault returns the dialect specific statement for 'DEFAULT'.
	StmtDefault(s string) string

	// StmtPrimaryKey returns the dialect specific statement for 'PRIMARY KEY'.
	StmtPrimaryKey() string

	// StmtAutoIncrement returns the dialect specific statement for 'AUTO_INCREMENT'.
	StmtAutoIncrement() string
}
