// TODO: Clean this up as there is duplication up the wazoo

package htmlrender

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestRenderer(t *testing.T) {
	callback := func(n html.Node, filterItems []MinimalHtmlNode) bool {
		return true
	}

	nodes := [...]*html.Node{
		0: {
			Type: html.ElementNode,
			Data: "html",
		},
		1: {
			Type: html.ElementNode,
			Data: "head",
		},
		2: {
			Type: html.ElementNode,
			Data: "body",
		},
		3: {
			Type: html.TextNode,
			Data: "0<1",
		},
		4: {
			Type: html.ElementNode,
			Data: "p",
			Attr: []html.Attribute{
				{
					Key: "id",
					Val: "A",
				},
				{
					Key: "foo",
					Val: `abc"def`,
				},
			},
		},
		5: {
			Type: html.TextNode,
			Data: "2",
		},
		6: {
			Type: html.ElementNode,
			Data: "b",
			Attr: []html.Attribute{
				{
					Key: "empty",
					Val: "",
				},
			},
		},
		7: {
			Type: html.TextNode,
			Data: "3",
		},
		8: {
			Type: html.ElementNode,
			Data: "i",
			Attr: []html.Attribute{
				{
					Key: "backslash",
					Val: `\`,
				},
			},
		},
		9: {
			Type: html.TextNode,
			Data: "&4",
		},
		10: {
			Type: html.TextNode,
			Data: "5",
		},
		11: {
			Type: html.ElementNode,
			Data: "blockquote",
		},
		12: {
			Type: html.ElementNode,
			Data: "br",
		},
		13: {
			Type: html.TextNode,
			Data: "6",
		},
	}

	// Build a tree out of those nodes, based on a textual representation.
	// Only the ".\t"s are significant. The trailing HTML-like text is
	// just commentary. The "0:" prefixes are for easy cross-reference with
	// the nodes array.
	treeAsText := [...]string{
		0: `<html>`,
		1: `.	<head>`,
		2: `.	<body>`,
		3: `.	.	"0&lt;1"`,
		4: `.	.	<p id="A" foo="abc&#34;def">`,
		5: `.	.	.	"2"`,
		6: `.	.	.	<b empty="">`,
		7: `.	.	.	.	"3"`,
		8: `.	.	.	<i backslash="\">`,
		9: `.	.	.	.	"&amp;4"`,
		10: `.	.	"5"`,
		11: `.	.	<blockquote>`,
		12: `.	.	<br>`,
		13: `.	.	"6"`,
	}
	if len(nodes) != len(treeAsText) {
		t.Fatal("len(nodes) != len(treeAsText)")
	}
	var stack [8]*html.Node
	for i, line := range treeAsText {
		level := 0
		for line[0] == '.' {
			// Strip a leading ".\t".
			line = line[2:]
			level++
		}
		n := nodes[i]
		if level == 0 {
			if stack[0] != nil {
				t.Fatal("multiple root nodes")
			}
			stack[0] = n
		} else {
			stack[level-1].AppendChild(n)
			stack[level] = n
			for i := level + 1; i < len(stack); i++ {
				stack[i] = nil
			}
		}
		// At each stage of tree construction, we check all nodes for consistency.
		for j, m := range nodes {
			if err := checkNodeConsistency(m); err != nil {
				t.Fatalf("i=%d, j=%d: %v", i, j, err)
			}
		}
	}

	want := `<html><head></head><body>0&lt;1<p id="A" foo="abc&#34;def">` +
		`2<b empty="">3</b><i backslash="\">&amp;4</i></p>` +
		`5<blockquote></blockquote><br/>6</body></html>`

	dummyAttributeArray := []MinimalHtmlNode{}
	b := new(bytes.Buffer)

	if err := RenderSans(b, nodes[0], callback, dummyAttributeArray); err != nil {
		t.Fatal(err)
	}
	if got := b.String(); got != want {
		t.Errorf("got vs want:\n%s\n%s\n", got, want)
	}
}

func TestRenderer_withElementExclusion(t *testing.T) {
	callback := func(n html.Node, filterItems []MinimalHtmlNode) bool {
		return false
	}

	nodes := [...]*html.Node{
		0: {
			Type: html.ElementNode,
			Data: "html",
		},
		1: {
			Type: html.ElementNode,
			Data: "head",
		},
		2: {
			Type: html.ElementNode,
			Data: "body",
		},
		3: {
			Type: html.TextNode,
			Data: "0<1",
		},
		4: {
			Type: html.ElementNode,
			Data: "p",
			Attr: []html.Attribute{
				{
					Key: "id",
					Val: "A",
				},
				{
					Key: "foo",
					Val: `abc"def`,
				},
			},
		},
		5: {
			Type: html.TextNode,
			Data: "2",
		},
		6: {
			Type: html.ElementNode,
			Data: "b",
			Attr: []html.Attribute{
				{
					Key: "empty",
					Val: "",
				},
			},
		},
		7: {
			Type: html.TextNode,
			Data: "3",
		},
		8: {
			Type: html.ElementNode,
			Data: "i",
			Attr: []html.Attribute{
				{
					Key: "backslash",
					Val: `\`,
				},
			},
		},
		9: {
			Type: html.TextNode,
			Data: "&4",
		},
		10: {
			Type: html.TextNode,
			Data: "5",
		},
		11: {
			Type: html.ElementNode,
			Data: "blockquote",
		},
		12: {
			Type: html.ElementNode,
			Data: "br",
		},
		13: {
			Type: html.TextNode,
			Data: "6",
		},
	}

	// Build a tree out of those nodes, based on a textual representation.
	// Only the ".\t"s are significant. The trailing HTML-like text is
	// just commentary. The "0:" prefixes are for easy cross-reference with
	// the nodes array.
	treeAsText := [...]string{
		0: `<html>`,
		1: `.	<head>`,
		2: `.	<body>`,
		3: `.	.	"0&lt;1"`,
		4: `.	.	<p id="A" foo="abc&#34;def">`,
		5: `.	.	.	"2"`,
		6: `.	.	.	<b empty="">`,
		7: `.	.	.	.	"3"`,
		8: `.	.	.	<i backslash="\">`,
		9: `.	.	.	.	"&amp;4"`,
		10: `.	.	"5"`,
		11: `.	.	<blockquote>`,
		12: `.	.	<br>`,
		13: `.	.	"6"`,
	}
	if len(nodes) != len(treeAsText) {
		t.Fatal("len(nodes) != len(treeAsText)")
	}
	var stack [8]*html.Node
	for i, line := range treeAsText {
		level := 0
		for line[0] == '.' {
			// Strip a leading ".\t".
			line = line[2:]
			level++
		}
		n := nodes[i]
		if level == 0 {
			if stack[0] != nil {
				t.Fatal("multiple root nodes")
			}
			stack[0] = n
		} else {
			stack[level-1].AppendChild(n)
			stack[level] = n
			for i := level + 1; i < len(stack); i++ {
				stack[i] = nil
			}
		}
		// At each stage of tree construction, we check all nodes for consistency.
		for j, m := range nodes {
			if err := checkNodeConsistency(m); err != nil {
				t.Fatalf("i=%d, j=%d: %v", i, j, err)
			}
		}
	}

	want := ""
	dummyAttributeArray := []MinimalHtmlNode{}
	b := new(bytes.Buffer)

	if err := RenderSans(b, nodes[0], callback, dummyAttributeArray); err != nil {
		t.Fatal(err)
	}
	if got := b.String(); got != want {
		t.Errorf("got vs want:\n%s\n%s\n", got, want)
	}
}
