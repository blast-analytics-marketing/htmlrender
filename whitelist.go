// A whitelist only renders html elements that are included in the passed in
// array of MinimalHtmlNode's
package render_sans

import (
	"io"

	"golang.org/x/net/html"
)

// this is what is not working
func Whitelist(w io.Writer, n *html.Node, filterItems []MinimalHtmlNode) error {

	for _, minNode := range filterItems {
		nodeCopy := html.Node{
			Data: n.Data,
			Attr: n.Attr,
		}
		if TagMatch(nodeCopy, minNode) && AttributeMatch(n.Attr, minNode.Attr) {
			err := html.Render(w, n)

			if err != nil {
				return err
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Whitelist(w, c, filterItems)
	}

	return nil
}
