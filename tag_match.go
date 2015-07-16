package render_sans

import (
	"strings"

	"golang.org/x/net/html"
)

// According to the HTML spec, tag names are in-case sensative, thus we
// downcase themfor comparison

// Both the html.Node and MinimalHtmlNode structs have a Data attribute, that
// contains the text for the HTML tag it contains be it span, div, etc...
func TagMatch(node html.Node, minHtmlNode MinimalHtmlNode) bool {
	nodeTag := strings.ToLower(node.Data)
	minHtmlNodeTag := strings.ToLower(minHtmlNode.Data)

	return (nodeTag == minHtmlNodeTag)
}
