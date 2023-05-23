package linkedql

import (
	"github.com/pymjer/cayley/graph"
	"github.com/pymjer/cayley/query"
	"github.com/pymjer/cayley/query/path"
	"github.com/cayleygraph/quad/voc"
)

// Step is a logical part in the query
type Step interface {
	RegistryItem
}

// IteratorStep is a step that can build an Iterator.
type IteratorStep interface {
	Step
	BuildIterator(qs graph.QuadStore, ns *voc.Namespaces) (query.Iterator, error)
}

// PathStep is a Step that can build a Path.
type PathStep interface {
	Step
	BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error)
}
