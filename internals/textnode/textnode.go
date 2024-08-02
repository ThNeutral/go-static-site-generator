package textnode

const (
	BOLD = iota
)

type TextNode struct {
	Text     string
	TextType int
	URL      string
}

func (tn TextNode) ToString() string {
	return ""
}

func IsEqual(tn1 TextNode, tn2 TextNode) bool {
	return (tn1.Text == tn2.Text) && (tn1.TextType == tn2.TextType) && (tn1.URL == tn2.URL)
}
