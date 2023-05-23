package steps

import (
	"github.com/pymjer/cayley/graph"
	"github.com/pymjer/cayley/query/linkedql"
	"github.com/pymjer/cayley/query/path"
	"github.com/cayleygraph/quad/voc"
)

func init() {
	linkedql.Register(&ReversePropertyNamesAs{})
}

var _ linkedql.PathStep = (*ReversePropertyNamesAs)(nil)

// ReversePropertyNamesAs corresponds to .reversePropertyNamesAs().
type ReversePropertyNamesAs struct {
	From linkedql.PathStep `json:"from"`
	Tag  string            `json:"tag"`
}

// Description implements Step.
func (s *ReversePropertyNamesAs) Description() string {
	return "tags the list of predicates that are pointing in to a node."
}

// BuildPath implements linkedql.PathStep.
func (s *ReversePropertyNamesAs) BuildPath(qs graph.QuadStore, ns *voc.Namespaces) (*path.Path, error) {
	fromPath, err := s.From.BuildPath(qs, ns)
	if err != nil {
		return nil, err
	}
	return fromPath.SavePredicates(true, s.Tag), nil
}
