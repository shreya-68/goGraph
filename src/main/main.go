package main

import (

    "fmt"
    "time"

    )

//Client Node with name, nbr are the first hop neighbours and status is current running status
type Node struct {
    name   string
    nbr    []Node
    status string

}


type Graph struct {

    numNodes int
    nodes    []*Node
    edges    map[*Node]*Node

}



var graph Graph

func initGraph() {

    var nodeA = Node{name:"A", status: "Init"}
    var nodeB = Node{name:"B", status: "Init"}
    nodeA.nbr = []Node{nodeB}
    nodeB.nbr = []Node{nodeA}
    graph = Graph{numNodes: 2, nodes: []*Node{&nodeA, &nodeB}}
    graph.edges = map[*Node]*Node{&nodeA:&nodeB, &nodeB:&nodeA}
    
}

func Client(node *Node) {

    fmt.Printf("Hi my name is %s", node.name)

}

func main() {

    //Create a centralised monitor

    //Initialise graph
    initGraph()

   //Launch Client goroutines
   for i := 0; i < graph.numNodes; i++{

       fmt.Printf(graph.nodes[i].name)
       go Client(graph.nodes[i])

   }
   time.Sleep(2000*time.Millisecond) 
   fmt.Printf("Done!")

}
