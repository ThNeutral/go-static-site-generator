package block

import (
	"fmt"
	"strings"

	"github.com/thneutral/static-site-generator/internals/splitter"
	"github.com/thneutral/static-site-generator/internals/textnode"
)

const (
	PARAGRAPH_TYPE = iota
	CODE_TYPE
	QUOTE_TYPE
	UNORDERED_LIST_TYPE
	ORDERED_LIST_TYPE
	HEADING1_TYPE
	HEADING2_TYPE
	HEADING3_TYPE
	HEADING4_TYPE
	HEADING5_TYPE
	HEADING6_TYPE
)

const (
	PARAGRAPH_DELIM       = ""
	CODE_DELIM            = "```"
	QUOTE_DELIM           = ">"
	UNORDERED_LIST_DELIM1 = "*"
	UNORDERED_LIST_DELIM2 = "-"
	ORDERED_LIST_DELIM    = "."

	HEADING1_DELIM = "#"
	HEADING2_DELIM = "##"
	HEADING3_DELIM = "###"
	HEADING4_DELIM = "####"
	HEADING5_DELIM = "#####"
	HEADING6_DELIM = "######"
)

type BlockNode struct {
	NodeType int
	Text     []textnode.TextNode
	Children []BlockNode
}

func (node BlockNode) isHeading() bool {
	return node.NodeType == HEADING1_TYPE || node.NodeType == HEADING2_TYPE || node.NodeType == HEADING3_TYPE || node.NodeType == HEADING4_TYPE || node.NodeType == HEADING5_TYPE || node.NodeType == HEADING6_TYPE
}

func (node BlockNode) ToHTML() string {
	html := ""
	if node.NodeType == UNORDERED_LIST_TYPE || node.NodeType == ORDERED_LIST_TYPE || node.NodeType == QUOTE_TYPE {
		if node.NodeType == UNORDERED_LIST_TYPE {
			html += "<ul>"
		} else if node.NodeType == ORDERED_LIST_TYPE {
			html += "<ol>"
		} else {
			html += "<blockquote>"
		}
		for _, child := range node.Children {
			if child.NodeType == UNORDERED_LIST_TYPE || child.NodeType == ORDERED_LIST_TYPE {
				html += "<li>"
			}
			for _, tn := range child.Text {
				html += tn.ToHTMLNode().ToHTML()
			}
			if child.NodeType == UNORDERED_LIST_TYPE || child.NodeType == ORDERED_LIST_TYPE {
				html += "</li>"
			}
		}
		if node.NodeType == UNORDERED_LIST_TYPE {
			html += "</ul>"
		} else if node.NodeType == ORDERED_LIST_TYPE {
			html += "</ol>"
		} else {
			html += "</blockquote>"
		}
	} else {
		if node.isHeading() {
			html += fmt.Sprintf("<h%v>", headingTypeToCount(node.NodeType))
		} else if node.NodeType == PARAGRAPH_TYPE {
			html += "<p>"
		} else {
			html += "<pre><code>"
		}
		for _, tn := range node.Text {
			html += tn.ToHTMLNode().ToHTML()
		}
		if node.isHeading() {
			html += fmt.Sprintf("</h%v>", headingTypeToCount(node.NodeType))
		} else if node.NodeType == PARAGRAPH_TYPE {
			html += "</p>"
		} else {
			html += "</code></pre>"
		}
	}

	return html
}

func SplitBlocks(input string) []BlockNode {
	splits := strings.Split(input, "\r\n\r\n")
	var blocks []BlockNode
	for _, split := range splits {
		blocks = append(blocks, blockToBlockNode(strings.Trim(split, " ")))
	}
	return blocks
}

func blockToBlockNode(block string) BlockNode {
	var blockNode BlockNode
	if block[0:3] == CODE_DELIM && block[len(block)-3:] == CODE_DELIM {
		blockNode.NodeType = CODE_TYPE
		blockNode.Text = splitter.SplitNodes(block[3:len(block)-3] + "\n")
		return blockNode
	}
	if string(block[0]) == HEADING1_DELIM {
		count := 0
		for _, r := range block[0:6] {
			if string(r) == HEADING1_DELIM {
				count += 1
			} else {
				break
			}

		}
		blockNode.NodeType = countToHeadingType(count)
		blockNode.Text = splitter.SplitNodes(block[count+1:])
		return blockNode
	}
	if string(block[0]) == QUOTE_DELIM {
		blocks := strings.Split(block, "\r\n")
		for _, bl := range blocks {
			var bn BlockNode
			if string(bl[0]) == QUOTE_DELIM {
				bn.NodeType = QUOTE_TYPE
				bn.Text = splitter.SplitNodes(bl[2:])
			} else {
				bn.NodeType = PARAGRAPH_TYPE
				bn.Text = splitter.SplitNodes(bl)
			}
			blockNode.Children = append(blockNode.Children, bn)
		}
		blockNode.NodeType = QUOTE_TYPE
		return blockNode
	}
	if (string(block[0]) == UNORDERED_LIST_DELIM1 || string(block[0]) == UNORDERED_LIST_DELIM2) && string(block[1]) == " " {
		blocks := strings.Split(block, "\r\n")
		for _, bl := range blocks {
			var bn BlockNode
			if string(bl[0]) == UNORDERED_LIST_DELIM1 || string(bl[0]) == UNORDERED_LIST_DELIM2 {
				bn.NodeType = UNORDERED_LIST_TYPE
				bn.Text = splitter.SplitNodes(bl[2:])
			} else {
				bn.NodeType = PARAGRAPH_TYPE
				bn.Text = splitter.SplitNodes(bl)
			}
			blockNode.Children = append(blockNode.Children, bn)
		}
		blockNode.NodeType = UNORDERED_LIST_TYPE
		return blockNode
	}
	if string(block[1]) == ORDERED_LIST_DELIM {
		blocks := strings.Split(block, "\r\n")
		for _, bl := range blocks {
			var bn BlockNode
			if string(bl[1]) == ORDERED_LIST_DELIM {
				bn.NodeType = ORDERED_LIST_TYPE
				bn.Text = splitter.SplitNodes(bl[3:])
			} else {
				bn.NodeType = PARAGRAPH_TYPE
				bn.Text = splitter.SplitNodes(bl)
			}
			blockNode.Children = append(blockNode.Children, bn)
		}
		blockNode.NodeType = ORDERED_LIST_TYPE
		return blockNode
	}
	blockNode.NodeType = PARAGRAPH_TYPE
	blockNode.Text = splitter.SplitNodes(block)
	return blockNode
}

func countToHeadingType(count int) int {
	switch count {
	case 1:
		return HEADING1_TYPE
	case 2:
		return HEADING2_TYPE
	case 3:
		return HEADING3_TYPE
	case 4:
		return HEADING4_TYPE
	case 5:
		return HEADING5_TYPE
	case 6:
		return HEADING6_TYPE
	default:
		return -1
	}
}

func headingTypeToCount(ht int) int {
	switch ht {
	case HEADING1_TYPE:
		return 1
	case HEADING2_TYPE:
		return 2
	case HEADING3_TYPE:
		return 3
	case HEADING4_TYPE:
		return 4
	case HEADING5_TYPE:
		return 5
	case HEADING6_TYPE:
		return 6
	default:
		return -1
	}
}
