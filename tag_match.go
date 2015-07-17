package render_sans

import (
	"strings"

	"golang.org/x/net/html"
)

// According to the HTML spec, tag names are in-case sensative, thus we
// downcase themfor comparison

// Both the html.Node and MinimalHtmlNode structs have a Data attribute, that
// contains the text for the HTML tag it contains be it span, div, etc...

// if there is no tag type within the minHtmlNode, we consider it a match,
// we do this to promote generasism wihthin this package, if we want to exclude
// base purely by attribute IE id or class(es) then not providing a tag should
// be an indication that they want to "skip" the tag portion
func TagMatch(node html.Node, minHtmlNode MinimalHtmlNode) bool {
	nodeTag := strings.ToLower(node.Data)
	minHtmlNodeTag := strings.ToLower(minHtmlNode.Data)

	if minHtmlNodeTag == "" {
		return true
	}

	return (nodeTag == minHtmlNodeTag)
}
