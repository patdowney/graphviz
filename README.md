# GraphViz Dot Writer [![Build Status](https://travis-ci.org/patdowney/graphviz.svg?branch=master)](https://travis-ci.org/patdowney/graphviz)

## What
Aims to be a simple interface for programatically generating graphviz 
dot files

Still needs some work.

## Examples

```go
g := NewDigraph("G")
g.Properties["fontname"] = "Helvetica"

a := NewAttr("node")
a.Properties["shape"] = "Mrecord"

g.AddAttribute(a)

s := NewClusterSubGraph("GroupA")

n1 := NewNode("node_one")
n2 := NewNode("node_two")

r := n1.AddRelation(n2)

s.AddNode(n1)
s.AddNode(n2)

fmt.Print(g.GraphViz())
```

