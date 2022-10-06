package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func main() {
	r, err := os.Open("index.html")
	if err != nil {
		fmt.Println("Could not read input file, index.html")
		panic(err)
	}
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println("Could not parse input file, index.html")
		panic(err)
	}

	dom.Children()

	nodeType := []string{"ErrorNode",
		"TextNode",
		"DocumentNode",
		"ElementNode",
		"CommentNode",
		"DoctypeNode",
		"RawNode",
		"scopeMarkerNode"}

	DepthFirstTraversal(l, func(n *html.Node) {
		fmt.Printf("type: %s, data: %s\n", nodeType[n.Type], n.Data)
	})
}

// DepthFirstTraversal on the html node tree. visit is evaluated on leaf nodes and when all children nodes have been evaluated.
func DepthFirstTraversal(n *html.Node, visit func(*html.Node)) {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		DepthFirstTraversal(c, visit)
	}
	visit(n)
}

// // BreadthFirstTraversal on the html node tree. visit is evaluated on a node, then on each subsequent child node.
// func BreadthFirstTraversal(node *html.Node, visit func(*html.Node)) {

// }
