package gotenberg

import "strconv"

// Orientation
const (
	landscapeOffice string = "landscape" // Set the paper orientation to landscape (default false)
)

// Page Ranges
const (
	nativePageRangesOffice string = "nativePageRanges" // Page ranges to print, e.g., '1-4' - empty means all pages
)

// PDF Format
const (
	nativePdfFormatOffice string = "nativePdfFormat" // Use unoconv to convert the resulting PDF to the given PDF format
	pdfFormatOffice       string = "pdfFormat"       // The PDF format of the resulting PDF
)

// Merge
const (
	mergeOffice string = "merge" // Merge all PDF files into an individual PDF file
)

type OfficeRequest struct {
	docs []Document

	*request
}

func NewOfficeRequest(docs ...Document) *OfficeRequest {
	return &OfficeRequest{docs: docs, request: newRequest()}
}

// Landscape sets landscape form field.
func (req *OfficeRequest) Landscape(isLandscape bool) {
	req.values[landscapeOffice] = strconv.FormatBool(isLandscape)
}

// PageRanges sets pageRanges form field.
func (req *OfficeRequest) PageRanges(ranges string) {
	req.values[nativePageRangesOffice] = ranges
}

// NativePDFFormat sets nativePdfFormat form field.
func (req *OfficeRequest) NativePDFFormat(pdfFormat string) {
	req.values[nativePdfFormatOffice] = pdfFormat
}

// PDFFormat sets pdfFormat form field.
func (req *OfficeRequest) PDFFormat(pdfFormat string) {
	req.values[pdfFormatOffice] = pdfFormat
}

// Merge sets merge form field.
func (req *OfficeRequest) Merge(merge bool) {
	req.values[mergeOffice] = strconv.FormatBool(merge)
}

func (req *OfficeRequest) postURL() string {
	return "/forms/libreoffice/convert"
}

func (req *OfficeRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, doc := range req.docs {
		files[doc.Filename()] = doc
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(OfficeRequest))
)
