package gotenberg

type ConvertHTMLRequest struct {
	index  Document
	assets []Document

	*chromiumRequest
}

func NewConvertHTMLRequest(index Document) *ConvertHTMLRequest {
	return &ConvertHTMLRequest{index: index, assets: []Document{}, chromiumRequest: newChromiumRequest()}
}

func (req *ConvertHTMLRequest) postURL() string {
	return "/forms/chromium/convert/html"
}

// Assets sets assets form files.
func (req *ConvertHTMLRequest) Assets(assets ...Document) {
	req.assets = assets
}

func (req *ConvertHTMLRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	files["index.html"] = req.index
	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}
	for _, asset := range req.assets {
		files[asset.Filename()] = asset
	}
	return files
}

var _ Request = new(ConvertHTMLRequest)
