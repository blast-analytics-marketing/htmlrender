// A whitelist only renders html elements that are included in the passed in
// array of MinimalHtmlNode's
package sans

import (
	"io"

	"github.com/blast-analytics-marketing/htmlrender"
	"golang.org/x/net/html"
)

func Whitelist(w io.Writer, n *html.Node, filterItems []htmlrender.MinimalHtmlNode) error {

	for _, minNode := range filterItems {
		nodeCopy := html.Node{
			Data: n.Data,
			Attr: n.Attr,
		}
		if htmlrender.TagMatch(nodeCopy, minNode) && htmlrender.AttributeMatch(n.Attr, minNode.Attr) {
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
