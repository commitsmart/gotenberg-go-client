**⚠️ Working with Gotenberg >= 7 ⚠️**

# Gotenberg Go client

A simple Go client for interacting with a Gotenberg API.

Inspired by [thecodingmachine/gotenberg-go-client](https://github.com/thecodingmachine/gotenberg-go-client)

## Install

```bash
$ go get -u github.com/commitsmart/gotenberg-go-client
```

## Usage

```golang
ctx := context.Background()
httpClient := &http.Client{
    Timeout: time.Duration(5) * time.Second,
}
client := gotenberg.NewClient("localhost:3000", httpClient)

// from a path.
index, err := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")
check(err)
// ... or from a string.
// index, err := gotenberg.NewDocumentFromString("index.html", "<html>Foo</html>")
// ... or from bytes.
// index, err := gotenberg.NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))

header, err := gotenberg.NewDocumentFromPath("header.html", "/path/to/file")
check(err)
footer, err := gotenberg.NewDocumentFromPath("footer.html", "/path/to/file")
check(err)
style, err := gotenberg.NewDocumentFromPath("style.css", "/path/to/file")
check(err)
img, err := gotenberg.NewDocumentFromPath("img.png", "/path/to/file")
check(err)

req := gotenberg.NewConvertHTMLRequest(index)
req.Header(header)
req.Footer(footer)
req.Assets(style, img)
req.PaperSize(gotenberg.A4)
req.Margins(gotenberg.NoMargins)
req.Scale(0.75)
req.PreferCssPageSize(false)
req.OmitBackground(false)
req.PrintBackground(true)
//req.UserAgent("")
req.FailOnConsoleExceptions(true)
req.EmulatedMediaType("print")
req.PDFFormat("PDF/A-1a")
req.ExtraHttpHeaders("{\"MyHeader\": \"MyValue\"}")

// store method allows you to... store the resulting PDF in a particular destination.
client.Store(ctx, req, "path/you/want/the/pdf/to/be/stored.pdf")
// if you wish to redirect the response directly to the browser, you may also use:
resp, err := client.Post(ctx, req)
check(err)

bb, err := io.ReadAll(resp.Body)
check(err)
```

For more complete guides read the [documentation](https://gotenberg.dev/docs/about).