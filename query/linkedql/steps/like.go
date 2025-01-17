package steps

import (
	"github.com/pymjer/cayley/graph"
	"github.com/pymjer/cayley/query/linkedql"
	"github.com/pymjer/cayley/query/path"
	"github.com/pymjer/cayley/query/shape"
	"github.com/cayleygraph/quad/voc"
)

func init() {
	linkedql.Register(&Like{})
}

var _ linkedql.PathStep = (*Like)(nil)

// Like corresponds to like().
type Like struct {
	From    linkedql.PathStep `json:"from"`
	Pattern string            `json:"likePattern"`
}

// Description implements Operator.
func (s *Like) Description() string {
	return "Like filters out values that do not match given pattern."
}

// BuildPath implements PathStep.
func (s *Like) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.Filters(shape.Wildcard{Pattern: s.Pattern}), nil
}
