package render_sans

import "golang.org/x/net/html"

// very much like a node but with limited / minimal information.
// data being the type of tag, ie span, div etc..
// where as the attributes consist of id, class and their respective values
type MinimalHtmlNode struct {
	Data string
	Attr []html.Attribute
}
