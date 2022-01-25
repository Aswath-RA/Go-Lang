package main

import (
	"fmt"
)

//graph structure
type Graph struct {
	vertices []*Vertex
}
type Vertex struct {
	key       int
	adajacent []*Vertex
}

//Contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

//Adding vertex
func (g *Graph) Addvertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added it is exist already", k)
		fmt.Println(err.Error())
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

//Adding Edge

func (g *Graph) Addedge(From, To int) {
	//get vertex
	fromvertex := g.getVertex(From)
	tovertex := g.getVertex(To)
	//check error
	if fromvertex == nil || tovertex == nil {
		err := fmt.Errorf("invalid edge %v --->%v", From, To)
		fmt.Println(err.Error())
	} else if contains(fromvertex.adajacent, To) {
		err := fmt.Errorf("Existing edge %v --->%v", From, To)
		fmt.Println(err.Error())
	} else { //add edge
		fromvertex.adajacent = append(fromvertex.adajacent, tovertex)
	}

}

//Printing the adjacent list
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adajacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}
func (g Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//Main func
func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.Addvertex(i)
	}

	test.Addedge(2, 1) //Adding edge to a vertex
	test.Addedge(2, 1) //Existing edge
	test.Addvertex(0)  //Adding Existing vertex
	test.Addedge(6, 1) //Adding Invalid Edge

	test.Print()
}
