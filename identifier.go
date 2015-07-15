package render_sans

import (
	"strings"

	"golang.org/x/net/html"
)

// very much like a node but with limited / minimal information.
type MinimalHtmlNode struct {
	Data string
	Attr []html.Attribute
}

// According to the HTML spec, html attributes are infact case-sensitive
// I have chosen to honor this. Thus class keys or id keys will need to be a
// case-sensitive match.

// we conclude that the HtmlIdentifier and a mode match if the

// make all keys and attributes lowercase ->
// turn all attribute values into arrays and compare said arrays for each key on each node

// the identifier must have a 100% match to the node, meaning the node must
// contain all of the attributes the node contains, while the opposite may not be
// true

// func AttrComparison(node html.Node, identifier Identifier) bool {

// }

// there will be an additonal layer on top of this that iterates over all
// of the attributes
func HasAttribute(array []html.Attribute, attribute html.Attribute) bool {
	for _, a := range array {
		if (a.Key == attribute.Key) && attributeValueMatch(a.Val, attribute.Val) {
			return true
		}
	}
	return false
}

func attributeValueMatch(attributeVal string, lookingForVal string) bool {
	attributeList := strings.Fields(attributeVal)
	lookingForList := strings.Fields(lookingForVal)

	set := make(map[string]bool)

	for _, v := range attributeList {
		set[v] = true
	}

	for _, s := range lookingForList {
		if !hasString(set, s) {
			return false
		}
	}
	return true
}
func hasString(set map[string]bool, s string) bool {
	return set[s]
}
