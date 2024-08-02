package htmlnode

import (
	"testing"
)

func TestParentNodePropsToHTML(t *testing.T) {
	expect := "  href=\"https://www.google.com\" target=\"_blank\" "
	hn := HTMLNode{Props: make(map[string]string)}
	hn.Props["href"] = "https://www.google.com"
	hn.Props["target"] = "_blank"
	if expect != hn.PropsToHTML() {
		t.Errorf("HTMLNode{Props: [\"href\"] = \"https://www.google.com\", [\"target\"] = \"_blank\"}.PropsToHTML() = %v; want %v;", hn.PropsToHTML(), expect)
	}
}

func TestSingleHTMLNodeToHTML(t *testing.T) {
	ln := HTMLNode{
		Tag:   "a",
		Value: "Click me!",
		Props: make(map[string]string),
	}
	ln.Props["href"] = "https://www.google.com"
	expect := "<a href=\"https://www.google.com\" >Click me!</a>"
	if expect != ln.ToHTML() {
		t.Errorf("HTMLNode{Props: [\"href\"] = \"https://www.google.com\", Tag: \"a\", Value: \"Click me!\"}.ToHTML() = %v; want %v;", ln.ToHTML(), expect)
	}
}

func TestMultipleHTMLNodeToHTML(t *testing.T) {
	ln := HTMLNode{
		Tag: "p",
		Children: []HTMLNode{
			{
				Tag:   "b",
				Value: "Bold text",
			},
			{
				Tag:   "",
				Value: "Normal text",
			},
			{
				Tag:   "i",
				Value: "Italic text",
			},
			{
				Tag:   "",
				Value: "Normal text",
			},
		},
	}
	expect := "<p><b>Bold text</b>Normal text<i>Italic text</i>Normal text</p>"
	if ln.ToHTML() != expect {
		t.Errorf("HTMLNode.ToHTML() = %v; want %v;", ln.ToHTML(), expect)
	}
}
