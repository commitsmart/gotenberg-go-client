package gotenberg

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recordedForm struct {
	path   string
	values map[string]string
	files  []string
}

func recordingGotenberg(t *testing.T) (*Client, *recordedForm) {
	t.Helper()
	got := &recordedForm{values: map[string]string{}}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.path = r.URL.Path
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for name, vals := range r.MultipartForm.Value {
			got.values[name] = vals[0]
		}
		for _, fh := range r.MultipartForm.File["files"] {
			got.files = append(got.files, fh.Filename)
		}
		w.Header().Set("Content-Type", "application/pdf")
		_, _ = w.Write([]byte("%PDF-1.4"))
	}))
	t.Cleanup(srv.Close)
	return NewClient(srv.URL, srv.Client()), got
}

func postURLRequest(t *testing.T, c *Client, req *ConvertURLRequest) []byte {
	t.Helper()
	resp, err := c.Post(context.Background(), req)
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	return body
}

func TestURLPostsToTheChromiumURLRoute(t *testing.T) {
	c, got := recordingGotenberg(t)
	req := NewConvertURLRequest("https://example.com/offer/AB12CD34?preview=1&print=1&lang=hu")

	body := postURLRequest(t, c, req)

	assert.Equal(t, "%PDF-1.4", string(body))
	assert.Equal(t, "/forms/chromium/convert/url", got.path)
	assert.Equal(t, "https://example.com/offer/AB12CD34?preview=1&print=1&lang=hu", got.values["url"])
	assert.Empty(t, got.files, "a URL conversion sends no index.html")
}

func TestURLForwardsChromiumOptions(t *testing.T) {
	c, got := recordingGotenberg(t)
	req := NewConvertURLRequest("https://example.com")
	req.PaperSize(A4)
	req.Margins(NoMargins)
	req.PrintBackground(true)
	req.EmulatedMediaType("print")
	req.WaitForExpression("window.__ready === true")
	require.NoError(t, req.MetaData(MetaData{Title: "Offer AB12CD34"}))

	postURLRequest(t, c, req)

	assert.Equal(t, "8.270000", got.values["paperWidth"])
	assert.Equal(t, "11.700000", got.values["paperHeight"])
	for _, m := range []string{"marginTop", "marginBottom", "marginLeft", "marginRight"} {
		assert.Equal(t, "0.000000", got.values[m], m)
	}
	assert.Equal(t, "true", got.values["printBackground"])
	assert.Equal(t, "print", got.values["emulatedMediaType"])
	assert.Equal(t, "window.__ready === true", got.values["waitForExpression"])
	assert.JSONEq(t, `{"Title":"Offer AB12CD34"}`, got.values["metadata"])
}

func TestURLSendsHeaderAndFooterOnlyWhenSet(t *testing.T) {
	c, got := recordingGotenberg(t)
	req := NewConvertURLRequest("https://example.com")
	header, err := NewDocumentFromString("header.html", "<p>head</p>")
	require.NoError(t, err)
	req.Header(header)

	postURLRequest(t, c, req)

	assert.Equal(t, []string{"header.html"}, got.files)
}
