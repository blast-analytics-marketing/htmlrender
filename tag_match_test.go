package htmlrender

import (
	"testing"

	"golang.org/x/net/html"
)

func TestTagMatch(t *testing.T) {
	node := html.Node{
		Data: "span",
	}

	minHtmlNode := MinimalHtmlNode{
		Data: "span",
	}

	returnVal := TagMatch(node, minHtmlNode)

	if returnVal != true {
		t.Errorf("expected true, got false")
	}
}

func TestTagMatch_caseInsensitive(t *testing.T) {
	node := html.Node{
		Data: "SpAn",
	}

	minHtmlNode := MinimalHtmlNode{
		Data: "sPaN",
	}

	returnVal := TagMatch(node, minHtmlNode)

	if returnVal != true {
		t.Errorf("expected true, got false")
	}
}

func TestTagMatch_blankMinNodeData(t *testing.T) {
	node := html.Node{
		Data: "SpAn",
	}

	minHtmlNode := MinimalHtmlNode{
		Data: "",
	}

	returnVal := TagMatch(node, minHtmlNode)

	if returnVal != true {
		t.Errorf("expected true, got false")
	}
}

func TestTagMatch_failure(t *testing.T) {
	node := html.Node{
		Data: "div",
	}

	minHtmlNode := MinimalHtmlNode{
		Data: "span",
	}

	returnVal := TagMatch(node, minHtmlNode)

	if returnVal != false {
		t.Errorf("expected false, got true")
	}
}
