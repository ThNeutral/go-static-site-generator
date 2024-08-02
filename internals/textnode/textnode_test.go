package textnode

import "testing"

func TestTextNode(t *testing.T) {
	tn1 := TextNode{Text: "test", TextType: BOLD, URL: ""}
	tn2 := TextNode{Text: "test", TextType: BOLD, URL: ""}
	result := IsEqual(tn1, tn2)
	if !result {
		t.Errorf("IsEqual(%v, %v) = %v; want %v", tn1, tn2, result, true)
	}
}
