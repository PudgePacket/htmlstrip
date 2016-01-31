# htmlstrip

Strip script elements from a html node.

Example code

```go
// html with a script inside
htmlString := "<div><script></script></div>"

doc, err := html.Parse(strings.NewReader(htmlString))
if err != nil {
	t.Error(err)
}

htmlstrip.Strip(doc)

// doc is now equivalent to "<div></div>"
```
