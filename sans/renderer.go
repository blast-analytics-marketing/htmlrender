package sans

import (
	"io"

	"github.com/blast-analytics-marketing/htmlrender"
	"golang.org/x/net/html"
)

type Renderer func(w io.Writer, n *html.Node, filterItems []htmlrender.MinimalHtmlNode) error
