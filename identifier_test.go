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

func TestattributeValueMatch(t *testing.T) {
	attributeVal := "primary    secondary      tertiary"
	s := "secondary tertiary primary"

	returnVal := attributeValueMatch(attributeVal, s)
	if returnVal != true {
		t.Fatal("Expcted true, got false")
	}
}

func TestattributeValueMatch_nodeHasAdditionalAttributes(t *testing.T) {
	attributeVal := "primary    secondary      tertiary other"
	s := "secondary tertiary primary"

	returnVal := attributeValueMatch(attributeVal, s)
	if returnVal != true {
		t.Fatal("Expcted true, got false")
	}
}

func TestattributeValueMatch_notAllAttributesPresent(t *testing.T) {
	attributeVal := "primary    secondary      tertiary"
	s := "secondary tertiary primary other"

	returnVal := attributeValueMatch(attributeVal, s)
	if returnVal != false {
		t.Fatal("Expcted false, got true")
	}
}

func TestAttributeMatch(t *testing.T) {
	lookingFor := []html.Attribute{
		{
			Key: "class",
			Val: "primary",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	lookingIn := []html.Attribute{
		{
			Key: "class",
			Val: "primary",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	returnVal := AttributeMatch(lookingIn, lookingFor)
	if returnVal != true {
		t.Fatal("Expcted true, got false")
	}
}

func TestAttributeMatch_LookingInContainsAdditionalAttribute(t *testing.T) {
	lookingFor := []html.Attribute{
		{
			Key: "id",
			Val: "main",
		},
	}

	lookingIn := []html.Attribute{
		{
			Key: "class",
			Val: "primary",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	returnVal := AttributeMatch(lookingIn, lookingFor)
	if returnVal != true {
		t.Fatal("Expcted true, got false")
	}
}

func TestAttributeMatch_LookingInDoesNotContainAllAttributes(t *testing.T) {
	lookingFor := []html.Attribute{
		{
			Key: "class",
			Val: "primary",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	lookingIn := []html.Attribute{
		{
			Key: "class",
			Val: "primary",
		},
	}

	returnVal := AttributeMatch(lookingIn, lookingFor)
	if returnVal != false {
		t.Fatal("Expcted false, got true")
	}
}

func TestAttributeMatch_attributeFormatting(t *testing.T) {
	lookingFor := []html.Attribute{
		{
			Key: "class",
			Val: "primary secondary tertiary",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	lookingIn := []html.Attribute{
		{
			Key: "class",
			Val: "secondary     tertiary    primary    ",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	returnVal := AttributeMatch(lookingIn, lookingFor)
	if returnVal != true {
		t.Fatal("Expcted true, got false")
	}
}

func TestAttributeMatch_attributePositionIndifference(t *testing.T) {
	lookingFor := []html.Attribute{
		{
			Key: "id",
			Val: "main",
		},
		{
			Key: "class",
			Val: "primary secondary tertiary",
		},
	}

	lookingIn := []html.Attribute{
		{
			Key: "class",
			Val: "primary secondary tertiary",
		},
		{
			Key: "id",
			Val: "main",
		},
	}

	returnVal := AttributeMatch(lookingIn, lookingFor)
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
