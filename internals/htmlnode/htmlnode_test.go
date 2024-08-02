package htmlnode

import "testing"

func TestHTMLNodePropsToHTML(t *testing.T) {
	expect := " href=\"https://www.google.com\" target=\"_blank\" "
	hn := HTMLNode{Props: make(map[string]string)}
	hn.Props["href"] = "https://www.google.com"
	hn.Props["target"] = "_blank"
	if expect != hn.PropsToHTML() {
		t.Errorf("HTMLNode{Props: [\"href\"] = \"https://www.google.com\", [\"target\"] = \"_blank\"}.PropsToHTML() = %v; want %v", hn.PropsToHTML(), expect)
	}
}
