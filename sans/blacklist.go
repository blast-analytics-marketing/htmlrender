// A Blacklist renders all nodes, EXCEPT that are consitered a match from
// the filterItems argument
package sans

import (
	"io"

	"github.com/blast-analytics-marketing/htmlrender"
	"golang.org/x/net/html"
)

func Blacklist(w io.Writer, n *html.Node, filterItems []htmlrender.MinimalHtmlNode) error {

	// returns false if there is ANY  100% match between the node any MinimumHTMLNode
	renderDecisionFunc := func(n html.Node, filterItems []htmlrender.MinimalHtmlNode) bool {
		for _, minNode := range filterItems {
			if htmlrender.TagMatch(n, minNode) && htmlrender.AttributeMatch(n.Attr, minNode.Attr) {
				return false
			}
		}
		return true
	}

	err := htmlrender.RenderSans(w, n, renderDecisionFunc, filterItems)

	if err != nil {
		return err
	}
	return nil
}
