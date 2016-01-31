// Package htmlstrip provides a function to strip script elements from html
package htmlstrip

import (
	"fmt"
	"golang.org/x/net/html"
)

// Strips script elements from a HTML node
func Strip(node *html.Node) {
	var f func(*html.Node, int) bool
	f = func(n *html.Node, d int) bool {
		if n == nil {
			return
		}
		nodeIsScript := false
		if n.Type == html.ElementNode {
			if n.Data == "script" {
				nodeIsScript = true
			}
		}
		scriptNodes := []*html.Node{}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if f(c, d+1) {
				scriptNodes = append(scriptNodes, c)
			}
		}
		for _, c := range scriptNodes {
			n.RemoveChild(c)
		}
		return nodeIsScript
	}
	f(node, -1)
}

// Prints the html node
func PrintParseTree(doc *html.Node) {
	var f func(*html.Node, int)
	f = func(n *html.Node, d int) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode {
			for i := 0; i < d; i++ {
				fmt.Print(" ")
			}
			fmt.Println(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, d+1)
		}
	}
	f(doc, -1)
}
