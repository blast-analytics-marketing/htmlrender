package htmlrender

import (
	"strings"

	"golang.org/x/net/html"
)

// According to the HTML spec, html attributes are infact case-sensitive
// I have chosen to honor this. Thus class keys or id keys will need to be a
// case-sensitive match.

// we conclude that the HtmlIdentifier and a mode match if the

// make all keys and attributes lowercase ->
// turn all attribute values into arrays and compare said arrays for each key on each node

// the identifier must have a 100% match to the node, meaning the node must
// contain all of the attributes the node contains, while the opposite may not be
// true

// ensures all attributes in lookingFor are present witin the looking in arg.
// attribute values are case sensative according to the HTML spec

// in the event that the lookingFor arugment does not contain any attributes,
// we return true, as this would be an indication that the user looking for some
// type of element that does not have any attributes.

func AttributeMatch(lookingIn, lookingFor []html.Attribute) bool {
	for _, prospectAttr := range lookingFor {
		if !HasAttribute(lookingIn, prospectAttr) {
			return false
		}
	}
	return true
}

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
