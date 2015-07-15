package render_sans

import "golang.org/x/net/html"

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

func HasAttribute(array []html.Attribute, attribute html.Attribute) bool {
	for _, a := range array {
		if (a.Key == attribute.Key) && (a.Val == attribute.Val) {
			return true
		}
	}
	return false
}
