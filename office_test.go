package gotenberg

import (
	"context"
	"fmt"
	"github.com/commitsmart/gotenberg-go-client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOffice(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.PageRanges("1-1")
	req.Landscape(true)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	//err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficePageRanges(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.PageRanges("1-1")
	resp, err := c.Post(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestOfficeWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.WebhookURL("https://google.com")
	req.WebhookErrorURL("https://google.com")
	req.WaitTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(context.Background(), req)
	assert.Nil(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}
