package splitter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/thneutral/static-site-generator/internals/textnode"
)

const (
	BOLD_DELIM   = "**"
	ITALIC_DELIM = "*"
	CODE_DELIM   = "`"
	STRIKE_DELIM = "~~"
)

func delimToTextType(delim string) int {
	switch delim {
	case BOLD_DELIM:
		return textnode.BOLD_TYPE
	case ITALIC_DELIM:
		return textnode.ITALIC_TYPE
	case CODE_DELIM:
		return textnode.CODE_TYPE
	case STRIKE_DELIM:
		return textnode.STRIKE_TYPE
	default:
		return -1
	}
}

func SplitNodes(input string) []textnode.TextNode {
	node := textnode.TextNode{
		Text:     input,
		TextType: textnode.TEXT_TYPE,
	}
	nodes := splitNodesTextInternal([]textnode.TextNode{node}, BOLD_DELIM)
	nodes = splitNodesTextInternal(nodes, ITALIC_DELIM)
	nodes = splitNodesTextInternal(nodes, CODE_DELIM)
	nodes = splitNodesTextInternal(nodes, STRIKE_DELIM)
	nodes = splitNodesImages(nodes)
	nodes = splitNodesLinks(nodes)
	return nodes
}

func splitNodesTextInternal(oldNodes []textnode.TextNode, delim string) []textnode.TextNode {
	nodes := []textnode.TextNode{}
	for _, oldNode := range oldNodes {
		splittedTexts := strings.Split(oldNode.Text, delim)
		for index, splittedText := range splittedTexts {
			var typeToAdd int
			if index%2 == 0 {
				typeToAdd = oldNode.TextType
			} else {
				typeToAdd = delimToTextType(delim)
			}
			nodes = append(nodes, textnode.TextNode{
				Text:     splittedText,
				TextType: typeToAdd,
			})
		}
	}
	return nodes
}

func splitNodesImages(oldNodes []textnode.TextNode) []textnode.TextNode {
	nodes := []textnode.TextNode{}
	for _, oldNode := range oldNodes {
		images := extractMarkdownImages(oldNode.Text)
		if len(images) == 0 {
			nodes = append(nodes, oldNode)
			continue
		}
		for alt, url := range images {
			nodes = append(nodes, textnode.TextNode{
				Text:     alt,
				URL:      url,
				TextType: textnode.IMAGE_TYPE,
			})
			splitted := strings.Split(oldNode.Text, fmt.Sprintf("![%v](%v)", alt, url))
			for _, split := range splitted {
				nodes = append(nodes, textnode.TextNode{
					Text:     split,
					TextType: oldNode.TextType,
				})
			}
		}
	}
	return nodes
}

func splitNodesLinks(oldNodes []textnode.TextNode) []textnode.TextNode {
	nodes := []textnode.TextNode{}
	for _, oldNode := range oldNodes {
		links := extractMarkdownLinks(oldNode.Text)
		if len(links) == 0 {
			nodes = append(nodes, oldNode)
			continue
		}
		for alt, url := range links {
			nodes = append(nodes, textnode.TextNode{
				Text:     alt,
				URL:      url,
				TextType: textnode.LINK_TYPE,
			})
			splitted := strings.Split(oldNode.Text, fmt.Sprintf("[%v](%v)", alt, url))
			for _, split := range splitted {
				nodes = append(nodes, textnode.TextNode{
					Text:     split,
					TextType: oldNode.TextType,
				})
			}
		}
	}
	return nodes
}

func extractMarkdownImages(text string) map[string]string {
	r := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	matches := r.FindAllString(text, -1)
	images := make(map[string]string)
	for _, match := range matches {
		parsed := markdownStringToLink(match)
		images[parsed[0]] = parsed[1]
	}
	return images
}

func extractMarkdownLinks(text string) map[string]string {
	r := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	matches := r.FindAllString(text, -1)
	links := make(map[string]string)
	for _, match := range matches {
		parsed := markdownStringToLink(match)
		links[parsed[0]] = parsed[1]
	}
	return links
}

func markdownStringToLink(markdown string) [2]string {
	res := [2]string{"", ""}
	isAlt := true
	for _, r := range markdown {
		char := string(r)
		if char == "!" || char == "[" || char == "(" || char == ")" {
			continue
		}
		if char == "]" {
			isAlt = false
			continue
		}
		if isAlt {
			res[0] += char
		} else {
			res[1] += char
		}
	}
	return res
}
