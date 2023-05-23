//+build cgo

package all

import (
	// backends requiring cgo
	_ "github.com/pymjer/cayley/graph/sql/sqlite"
)
