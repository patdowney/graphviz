package graphviz

type Relation struct {
	LeftID     string
	RightID    string
	Properties Properties
}

func (r *Relation) GraphViz() string {
	const graphTemplate = `{{.LeftID}} -> {{.RightID}} [{{ range $n, $v :=     .Properties}} {{$n}}="{{$v}}"{{end}} ];`

	return RenderTemplate(graphTemplate, r)
}

func NewRelation(left *Node, right *Node) *Relation {
	r := Relation{
		LeftID:     left.ID,
		RightID:    right.ID,
		Properties: make(Properties)}

	return &r
}
