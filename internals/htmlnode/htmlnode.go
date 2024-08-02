package htmlnode

type HTMLNode struct {
	Props    map[string]string
	Tag      string
	Value    string
	Children []HTMLNode
}

func (hn HTMLNode) ToHTML() string {
	html := ""

	if hn.Tag != "" {
		html += "<" + hn.Tag + hn.PropsToHTML() + ">"
	}

	if hn.IsLeaf() {
		html += hn.Value
	} else {
		for _, child := range hn.Children {
			html += child.ToHTML()
		}
	}

	if hn.Tag != "" {
		html += "</" + hn.Tag + ">"
	}

	return html
}

func (hn HTMLNode) PropsToHTML() string {
	htmlProps := ""
	if len(hn.Props) > 0 {
		htmlProps = " "
	}
	for key, value := range hn.Props {
		htmlProps += key + "="
		htmlProps += "\"" + value + "\""
		htmlProps += " "
	}
	return htmlProps
}

func (hn HTMLNode) IsLeaf() bool {
	return len(hn.Children) == 0
}
