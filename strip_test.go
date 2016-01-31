package htmlstrip

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func hasScript(node *html.Node) bool {
	scripts := false
	var f func(*html.Node, int)
	f = func(n *html.Node, d int) {
		if n.Type == html.ElementNode {
			if n.Data == "script" {
				scripts = true
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, d+1)
		}
	}
	f(node, -1)
	return scripts
}

func TestStripRemovesScript(t *testing.T) {
	htmlString := "<div><script></script></div>"

	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		t.Error(err)
	}

	if !hasScript(doc) {
		t.Error("html should have script")
	}

	Strip(doc)

	if hasScript(doc) {
		t.Error("html should not have script")
	}
}
