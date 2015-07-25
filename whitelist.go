// A whitelist only renders html elements that are included in the passed in
// array of MinimalHtmlNode's
package render_sans

import (
	"io"

	"golang.org/x/net/html"
)

// this is what is not working
func Whitelist(w io.Writer, n *html.Node, filterItems []MinimalHtmlNode) error {

	renderDecisionFunc := func(n html.Node, filterItems []MinimalHtmlNode) bool {
		for _, minNode := range filterItems {
			if TagMatch(n, minNode) && AttributeMatch(n.Attr, minNode.Attr) {
				return true
			}
		}
		return false
	}

	err := RenderSans(w, n, renderDecisionFunc, filterItems)

	if err != nil {
		return err
	}
	return nil
}
