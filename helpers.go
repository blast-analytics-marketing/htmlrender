package render_sans

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// wihtin the x/html package, escape function wich the render functions depend on
// is not exported, thus I have a coppiednversion of it here, wiht the corresponding constant
const escapedChars = "&'<>\"\r"

func escape(w writer, s string) error {
	i := strings.IndexAny(s, escapedChars)
	for i != -1 {
		if _, err := w.WriteString(s[:i]); err != nil {
			return err
		}
		var esc string
		switch s[i] {
		case '&':
			esc = "&amp;"
		case '\'':
			// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
			esc = "&#39;"
		case '<':
			esc = "&lt;"
		case '>':
			esc = "&gt;"
		case '"':
			// "&#34;" is shorter than "&quot;".
			esc = "&#34;"
		case '\r':
			esc = "&#13;"
		default:
			panic("unrecognized escape character")
		}
		s = s[i+1:]
		if _, err := w.WriteString(esc); err != nil {
			return err
		}
		i = strings.IndexAny(s, escapedChars)
	}
	_, err := w.WriteString(s)
	return err
}

// checkNodeConsistency checks that a node's parent/child/sibling relationships
// are consistent.
func checkNodeConsistency(n *html.Node) error {
	if n == nil {
		return nil
	}

	nParent := 0
	for p := n.Parent; p != nil; p = p.Parent {
		nParent++
		if nParent == 1e4 {
			return fmt.Errorf("html: parent list looks like an infinite loop")
		}
	}

	nForward := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nForward++
		if nForward == 1e6 {
			return fmt.Errorf("html: forward list of children looks like an infinite loop")
		}
		if c.Parent != n {
			return fmt.Errorf("html: inconsistent child/parent relationship")
		}
	}

	nBackward := 0
	for c := n.LastChild; c != nil; c = c.PrevSibling {
		nBackward++
		if nBackward == 1e6 {
			return fmt.Errorf("html: backward list of children looks like an infinite loop")
		}
		if c.Parent != n {
			return fmt.Errorf("html: inconsistent child/parent relationship")
		}
	}

	if n.Parent != nil {
		if n.Parent == n {
			return fmt.Errorf("html: inconsistent parent relationship")
		}
		if n.Parent == n.FirstChild {
			return fmt.Errorf("html: inconsistent parent/first relationship")
		}
		if n.Parent == n.LastChild {
			return fmt.Errorf("html: inconsistent parent/last relationship")
		}
		if n.Parent == n.PrevSibling {
			return fmt.Errorf("html: inconsistent parent/prev relationship")
		}
		if n.Parent == n.NextSibling {
			return fmt.Errorf("html: inconsistent parent/next relationship")
		}

		parentHasNAsAChild := false
		for c := n.Parent.FirstChild; c != nil; c = c.NextSibling {
			if c == n {
				parentHasNAsAChild = true
				break
			}
		}
		if !parentHasNAsAChild {
			return fmt.Errorf("html: inconsistent parent/child relationship")
		}
	}

	if n.PrevSibling != nil && n.PrevSibling.NextSibling != n {
		return fmt.Errorf("html: inconsistent prev/next relationship")
	}
	if n.NextSibling != nil && n.NextSibling.PrevSibling != n {
		return fmt.Errorf("html: inconsistent next/prev relationship")
	}

	if (n.FirstChild == nil) != (n.LastChild == nil) {
		return fmt.Errorf("html: inconsistent first/last relationship")
	}
	if n.FirstChild != nil && n.FirstChild == n.LastChild {
		// We have a sole child.
		if n.FirstChild.PrevSibling != nil || n.FirstChild.NextSibling != nil {
			return fmt.Errorf("html: inconsistent sole child's sibling relationship")
		}
	}

	seen := map[*html.Node]bool{}

	var last *html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if seen[c] {
			return fmt.Errorf("html: inconsistent repeated child")
		}
		seen[c] = true
		last = c
	}
	if last != n.LastChild {
		return fmt.Errorf("html: inconsistent last relationship")
	}

	var first *html.Node
	for c := n.LastChild; c != nil; c = c.PrevSibling {
		if !seen[c] {
			return fmt.Errorf("html: inconsistent missing child")
		}
		delete(seen, c)
		first = c
	}
	if first != n.FirstChild {
		return fmt.Errorf("html: inconsistent first relationship")
	}

	if len(seen) != 0 {
		return fmt.Errorf("html: inconsistent forwards/backwards child list")
	}

	return nil
}
