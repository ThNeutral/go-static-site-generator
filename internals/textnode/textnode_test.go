package textnode

import "testing"

func TestTextNode(t *testing.T) {
	tn1 := TextNode{Text: "test", TextType: BOLD_TYPE, URL: ""}
	tn2 := TextNode{Text: "test", TextType: BOLD_TYPE, URL: ""}
	result := tn1.IsEqual(tn2)
	if !result {
		t.Errorf("%v.IsEqual(%v) = %v; want %v;", tn1, tn2, result, true)
	}
}

func TestToHTMLNode(t *testing.T) {
	tn := TextNode{Text: "alt", TextType: IMAGE_TYPE, URL: "test-url"}
	html := tn.ToHTMLNode().ToHTML()
	expect := "<img alt=\"alt\" src=\"test-url\" ></img>"
	if html != expect {
		t.Errorf("%v.ToHTMLNode().ToHTML() = %v; want %v;", tn, html, expect)
	}
}
