package sans

import (
	"bytes"
	"testing"

	"github.com/blast-analytics-marketing/htmlrender"
	"golang.org/x/net/html"
)

func TestBlacklist(t *testing.T) {
	originalHTML := `<html><head></head><body>0&lt;1<p id="A" foo="abc&#34;def">` +
		`2<b empty="">3</b><i backslash="\">&amp;4</i></p>` +
		`5<blockquote></blockquote><br/>6</body></html>`
	originalHTMLAsBuffer := bytes.NewBufferString(originalHTML)

	rootNode, err := html.Parse(originalHTMLAsBuffer)

	if err != nil {
		t.Fatal(err)
	}

	dummyAttributeArray := []htmlrender.MinimalHtmlNode{
		{
			Data: "blockquote",
		},
	}

	w := new(bytes.Buffer)

	want := `<html><head></head><body>0&lt;1<p id="A" foo="abc&#34;def">` +
		`2<b empty="">3</b><i backslash="\">&amp;4</i></p>` +
		`5<br/>6</body></html>`

	if err := Blacklist(w, rootNode, dummyAttributeArray); err != nil {
		t.Fatal(err)
	}
	if got := w.String(); got != want {
		t.Errorf("got vs want:\n%s\n%s\n", got, want)
	}
}
