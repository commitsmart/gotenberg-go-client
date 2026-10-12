package gotenberg

const urlField string = "url" // The URL of the page to convert

type ConvertURLRequest struct {
	*chromiumRequest
}

func NewConvertURLRequest(url string) *ConvertURLRequest {
	req := &ConvertURLRequest{chromiumRequest: newChromiumRequest()}
	req.values[urlField] = url
	return req
}

func (req *ConvertURLRequest) postURL() string {
	return "/forms/chromium/convert/url"
}

func (req *ConvertURLRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}
	return files
}

var _ Request = new(ConvertURLRequest)
