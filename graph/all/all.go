package all

import (
	// supported backends
	_ "github.com/pymjer/cayley/graph/kv/all"
	_ "github.com/pymjer/cayley/graph/memstore"
	_ "github.com/pymjer/cayley/graph/nosql/all"
	_ "github.com/pymjer/cayley/graph/sql/cockroach"
	_ "github.com/pymjer/cayley/graph/sql/mysql"
	_ "github.com/pymjer/cayley/graph/sql/postgres"
)
