package gotenberg

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/commitsmart/gotenberg-go-client/test"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestHTMLFromString(t *testing.T) {
	c := &Client{Hostname: "http://gotenberg:3000"}
	index, err := NewDocumentFromString("index.html", "<html>Foo</html>")
	require.Nil(t, err)
	req := NewConvertHTMLRequest(index)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLFromBytes(t *testing.T) {
	c := &Client{Hostname: "http://gotenberg:3000"}
	index, err := NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))
	require.Nil(t, err)
	req := NewConvertHTMLRequest(index)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLComplete(t *testing.T) {
	c := &Client{Hostname: "http://gotenberg:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewConvertHTMLRequest(index)
	header, err := NewDocumentFromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.Nil(t, err)
	req.Header(header)
	footer, err := NewDocumentFromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	req.Footer(footer)
	font, err := NewDocumentFromPath("font.woff", test.HTMLTestFilePath(t, "font.woff"))
	require.Nil(t, err)
	img, err := NewDocumentFromPath("img.gif", test.HTMLTestFilePath(t, "img.gif"))
	require.Nil(t, err)
	style, err := NewDocumentFromPath("style.css", test.HTMLTestFilePath(t, "style.css"))
	require.Nil(t, err)
	req.Assets(font, img, style)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	req.Scale(1.5)
	req.PreferCssPageSize(false)
	req.OmitBackground(false)
	req.PrintBackground(true)
	req.UserAgent("Mozilla")
	req.FailOnConsoleExceptions(true)
	req.EmulatedMediaType("print")
	req.PDFFormat("PDF/A-1a")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLPageRanges(t *testing.T) {
	c := &Client{Hostname: "http://gotenberg:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewConvertHTMLRequest(index)
	req.NativePageRanges("1-1")
	resp, err := c.Post(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
