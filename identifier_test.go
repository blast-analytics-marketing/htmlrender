package render_sans

import (
	"testing"

	"golang.org/x/net/html"
)

func TestHasAttribute(t *testing.T) {
	array := []html.Attribute{
		{
			Key: "class",
			Val: "primary",
		},
	}

	singleItem := html.Attribute{
		Key: "class",
		Val: "primary",
	}

	returnVal := HasAttribute(array, singleItem)
	if returnVal != true {
		t.Fatal("Expcted true, got false")
	}
}

// func TestComparison(t *testing.T) {
// 	min := MinimalHtmlNode{
// 		Data: "div",
// 		Attr: []html.Attribute{
// 			{
// 				Key: "class",
// 				Val: "primary",
// 			},
// 		},
// 	}

// 	node := html.Node{
// 		Data: "div",
// 		Attr: []html.Attribute{
// 			{
// 				Key: "class",
// 				Val: "primary",
// 			},
// 		},
// 	}

// 	returnVal := AttrComparison(min, node)
// 	if returnVal != true {
// 		t.Fatal("Expcted true, got false")
// 	}
// }

// func TestComparison_attributeNotPresent(t *testing.T) {
// 	min := MinimalHtmlNode{
// 		Data: "div",
// 		Attr: []html.Attribute{
// 			{
// 				Key: "class",
// 				Val: "primary",
// 			},
// 		},
// 	}

// 	node := html.Node{
// 		Data: "div",
// 		Attr: []html.Attribute{
// 			{
// 				Key: "id",
// 				Val: "secondary",
// 			},
// 		},
// 	}

// 	returnVal := AttrComparison(min, node)
// 	if returnVal != true {
// 		t.Fatal("Expcted true, got false")
// 	}
// }
