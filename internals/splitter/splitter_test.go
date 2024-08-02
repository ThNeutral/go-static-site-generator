package splitter

import (
	"testing"

	"github.com/thneutral/static-site-generator/internals/htmlnode"
)

func TestSplitter(t *testing.T) {
	input := "`code` normal text **bald** *italian* [to boot dev](https://www.boot.dev) ![rick roll](https://i.imgur.com/aKaOqIh.gif)"
	nodes := SplitNodes(input)
	var hn htmlnode.HTMLNode
	for _, node := range nodes {
		hn.Children = append(hn.Children, node.ToHTMLNode())
	}
	html := hn.ToHTML()
	expect := "<code>code</code> normal text <b>bald</b> <i>italian</i><img src=\"https://i.imgur.com/aKaOqIh.gif\" alt=\"rick  roll\" ></img><a href=\"https://www.boot.dev\" >to boot dev</a>  "
	if html != expect {
		t.Errorf("SplitNodes(%v) = %v; want %v;", input, html, expect)
	}
}

func TestMarkdownLinkExtractor(t *testing.T) {
	text := "This is text with a link [to boot dev](https://www.boot.dev) and [to youtube](https://www.youtube.com/@bootdotdev)"
	tested := extractMarkdownLinks(text)
	expect := make(map[string]string)
	expect["to boot dev"] = "https://www.boot.dev"
	expect["to youtube"] = "https://www.youtube.com/@bootdotdev"
	if (expect["to boot dev"] != tested["to boot dev"]) || (expect["to youtube"] != tested["to youtube"]) {
		t.Errorf("extractMarkdownLinks(%v) = %v; want %v;", text, tested, expect)
	}
}

func TestMarkdownImageExtractor(t *testing.T) {
	text := "This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)"
	tested := extractMarkdownImages(text)
	expect := make(map[string]string)
	expect["rick roll"] = "https://i.imgur.com/aKaOqIh.gif"
	expect["obi wan"] = "https://i.imgur.com/fJRm4Vk.jpeg"
	if (expect["rick roll"] != tested["rick roll"]) || (expect["obi wan"] != tested["obi wan"]) {
		t.Errorf("extractMarkdownLinks(%v) = %v; want %v;", text, tested, expect)
	}
}
