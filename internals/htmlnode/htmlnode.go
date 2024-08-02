package htmlnode

type HTMLNode struct {
	Tag      string
	Value    string
	Children []HTMLNode
	Props    map[string]string
}

func (hn HTMLNode) ToHTML() string {
	html := ""
	html += "<" + hn.Tag + hn.PropsToHTML() + ">"
	html += "</" + hn.Tag + ">"
	return html
}

func (hn HTMLNode) PropsToHTML() string {
	htmlProps := " "
	for key, value := range hn.Props {
		htmlProps += key + "="
		htmlProps += "\"" + value + "\""
		htmlProps += " "
	}
	return htmlProps
}
