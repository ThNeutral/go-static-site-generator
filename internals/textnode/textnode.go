package textnode

import "github.com/thneutral/static-site-generator/internals/htmlnode"

const (
	TEXT_TYPE   = iota
	BOLD_TYPE   = iota
	ITALIC_TYPE = iota
	CODE_TYPE   = iota
	LINK_TYPE   = iota
	IMAGE_TYPE  = iota
	STRIKE_TYPE = iota
)

const (
	TEXT_TAG   = ""
	BOLD_TAG   = "b"
	ITALIC_TAG = "i"
	CODE_TAG   = "code"
	LINK_TAG   = "a"
	IMAGE_TAG  = "img"
	STRIKE_TAG = "s"
)

type TextNode struct {
	Text     string
	TextType int
	URL      string
}

func (tn TextNode) ToString() string {
	return ""
}

func (tn1 TextNode) IsEqual(tn2 TextNode) bool {
	return (tn1.Text == tn2.Text) && (tn1.TextType == tn2.TextType) && (tn1.URL == tn2.URL)
}

func (tn TextNode) ToHTMLNode() htmlnode.HTMLNode {
	var hn htmlnode.HTMLNode
	hn.Props = make(map[string]string)
	switch tn.TextType {
	case TEXT_TYPE:
		{
			hn.Tag = TEXT_TAG
			hn.Value = tn.Text
		}
	case BOLD_TYPE:
		{
			hn.Tag = BOLD_TAG
			hn.Value = tn.Text
		}
	case ITALIC_TYPE:
		{
			hn.Tag = ITALIC_TAG
			hn.Value = tn.Text
		}
	case CODE_TYPE:
		{
			hn.Tag = CODE_TAG
			hn.Value = tn.Text
		}
	case LINK_TYPE:
		{
			hn.Tag = LINK_TAG
			hn.Value = tn.Text
			hn.Props["href"] = tn.URL
		}
	case IMAGE_TYPE:
		{
			hn.Tag = IMAGE_TAG
			hn.Value = ""
			hn.Props["src"] = tn.URL
			hn.Props["alt"] = tn.Text
		}
	case STRIKE_TYPE:
		{
			hn.Tag = STRIKE_TAG
			hn.Value = tn.Text
		}
	}
	return hn
}
