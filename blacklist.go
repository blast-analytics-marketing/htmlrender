// A whitelist only renders html elements that are included in the passed in
// array of MinimalHtmlNode's
package render_sans

import (
	"io"

	"golang.org/x/net/html"
)

func Blacklist(w io.Writer, n *html.Node, filterItems []MinimalHtmlNode) error {

	// returns false if there is ANY  100% match between the node any MinimumHTMLNode
	renderDecisionFunc := func(n html.Node, filterItems []MinimalHtmlNode) bool {
		for _, minNode := range filterItems {
			if TagMatch(n, minNode) && AttributeMatch(n.Attr, minNode.Attr) {
				return false
			}
		}
		return true
	}

	err := RenderSans(w, n, renderDecisionFunc, filterItems)

	if err != nil {
		return err
	}
	return nil
}
