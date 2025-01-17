package steps

import (
	"github.com/pymjer/cayley/graph"
	"github.com/pymjer/cayley/graph/iterator"
	"github.com/pymjer/cayley/query/linkedql"
	"github.com/pymjer/cayley/query/path"
	"github.com/cayleygraph/quad"
	"github.com/cayleygraph/quad/voc"
)

func init() {
	linkedql.Register(&GreaterThanEquals{})
}

var _ linkedql.PathStep = (*GreaterThanEquals)(nil)

// GreaterThanEquals corresponds to gte().
type GreaterThanEquals struct {
	From  linkedql.PathStep `json:"from"`
	Value quad.Value        `json:"value"`
}

// Description implements Step.
func (s *GreaterThanEquals) Description() string {
	return "Greater than equals filters out values that are not greater than or equal given value"
}

// BuildPath implements linkedql.PathStep.
func (s *GreaterThanEquals) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Filter(iterator.CompareGTE, linkedql.AbsoluteValue(s.Value, ns)), nil
}
