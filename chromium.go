package gotenberg

import (
	"fmt"
	"strconv"
	"time"
)

// Page Properties
const (
	paperWidth        string = "paperWidth"        // Paper width, in inches (default 8.5)
	paperHeight       string = "paperHeight"       // Paper height, in inches (default 11)
	marginTop         string = "marginTop"         // Top margin, in inches (default 0.39)
	marginBottom      string = "marginBottom"      // Bottom margin, in inches (default 0.39)
	marginLeft        string = "marginLeft"        // Left margin, in inches (default 0.39)
	marginRight       string = "marginRight"       // Right margin, in inches (default 0.39)
	preferCssPageSize string = "preferCssPageSize" // Define whether to prefer page size as defined by CSS (default false)
	printBackground   string = "printBackground"   // Print the background graphics (default false)
	omitBackground    string = "omitBackground"    // Hide the default white background and allow generating PDFs with transparency (default false)
	landscape         string = "landscape"         // Set the paper orientation to landscape (default false)
	scale             string = "scale"             // The scale of the page rendering (default 1.0)
	nativePageRanges  string = "nativePageRanges"  // Page ranges to print, e.g., '1-5, 8, 11-13' - empty means all pages
)

// Wait
const (
	waitDelay         string = "waitDelay"         // Duration to wait when loading an HTML document before converting it to PDF
	waitForExpression string = "waitForExpression" // The JavaScript expression to wait before converting an HTML document to PDF until it returns true
)

// HTTP Headers
const (
	userAgent        string = "userAgent"        // Override the default User-Agent header
	extraHttpHeaders string = "extraHttpHeaders" //  HTTP headers to send by Chromium while loading the HTML document (JSON format)
)

// Javascript
const failOnConsoleExceptions string = "failOnConsoleExceptions" // Return a 409 Conflict response if there are exceptions in the Chromium console (default false)

// CSS
const emulatedMediaType string = "emulatedMediaType" // The media type to emulate, either "screen" or "print" - empty means "print"

// PDF
const pdfFormat string = "pdfFormat" // The PDF format of the resulting PDF

// Paper Sizes
var (
	// A0 paper size.
	A0 = [2]float64{33.1, 46.8}
	// A1 paper size.
	A1 = [2]float64{23.4, 33.1}
	// A2 paper size.
	A2 = [2]float64{16.54, 23.4}
	// A3 paper size.
	A3 = [2]float64{11.7, 16.5}
	// A4 paper size.
	A4 = [2]float64{8.27, 11.7}
	// A5 paper size.
	A5 = [2]float64{5.8, 8.3}
	// A6 paper size.
	A6 = [2]float64{4.1, 5.8}
	// Letter paper size.
	Letter = [2]float64{8.5, 11}
	// Legal paper size.
	Legal = [2]float64{8.5, 14}
	// Tabloid paper size.
	Tabloid = [2]float64{11, 17}
	// Ledger paper size.
	Ledger = [2]float64{17, 11}
)

// nolint:gochecknoglobals
var (
	// NoMargins removes margins.
	NoMargins = [4]float64{0, 0, 0, 0}
	// NormalMargins uses 1 inche margins.
	NormalMargins = [4]float64{1, 1, 1, 1}
	// LargeMargins uses 2 inche margins.
	LargeMargins = [4]float64{2, 2, 2, 2}
)

type chromiumRequest struct {
	header Document
	footer Document

	*request
}

func newChromiumRequest() *chromiumRequest {
	return &chromiumRequest{header: nil, footer: nil, request: newRequest()}
}

// Header sets header form file.
func (req *chromiumRequest) Header(header Document) {
	req.header = header
}

// Footer sets footer form file.
func (req *chromiumRequest) Footer(footer Document) {
	req.footer = footer
}

// PaperSize sets paperWidth and paperHeight form fields.
func (req *chromiumRequest) PaperSize(size [2]float64) {
	req.values[paperWidth] = fmt.Sprintf("%f", size[0])
	req.values[paperHeight] = fmt.Sprintf("%f", size[1])
}

// Margins sets marginTop, marginBottom,
// marginLeft and marginRight form fields.
func (req *chromiumRequest) Margins(margins [4]float64) {
	req.values[marginTop] = fmt.Sprintf("%f", margins[0])
	req.values[marginBottom] = fmt.Sprintf("%f", margins[1])
	req.values[marginLeft] = fmt.Sprintf("%f", margins[2])
	req.values[marginRight] = fmt.Sprintf("%f", margins[3])
}

// Landscape sets landscape form field.
func (req *chromiumRequest) Landscape(isLandscape bool) {
	req.values[landscape] = strconv.FormatBool(isLandscape)
}

// NativePageRanges sets pageRanges form field.
func (req *chromiumRequest) NativePageRanges(ranges string) {
	req.values[nativePageRanges] = ranges
}

// Scale sets scale form field
func (req *chromiumRequest) Scale(scaleFactor float64) {
	req.values[scale] = fmt.Sprintf("%f", scaleFactor)
}

// PreferCssPageSize sets preferCssPageSize form field
func (req *chromiumRequest) PreferCssPageSize(isPreferCssPageSize bool) {
	req.values[preferCssPageSize] = strconv.FormatBool(isPreferCssPageSize)
}

// PrintBackground sets printBackground form field
func (req *chromiumRequest) PrintBackground(isPrintBackground bool) {
	req.values[printBackground] = strconv.FormatBool(isPrintBackground)
}

// OmitBackground sets omitBackground form field
func (req *chromiumRequest) OmitBackground(isOmitBackground bool) {
	req.values[omitBackground] = strconv.FormatBool(isOmitBackground)
}

// WaitDelay sets waitDelay form field
func (req *chromiumRequest) WaitDelay(d time.Duration) {
	req.values[waitDelay] = d.String()
}

// WaitForExpression sets waitForExpression form field
func (req *chromiumRequest) WaitForExpression(expression string) {
	req.values[waitForExpression] = expression
}

// UserAgent sets userAgent form field
// e.g.: userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38
//
//	(KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
func (req *chromiumRequest) UserAgent(agent string) {
	req.values[userAgent] = agent
}

// ExtraHttpHeaders sets extraHttpHeaders form field
// e.g.: extraHttpHeaders="{\"MyHeader\": \"MyValue\"}"
func (req *chromiumRequest) ExtraHttpHeaders(headers string) {
	req.values[extraHttpHeaders] = headers
}

// FailOnConsoleExceptions sets failOnConsoleExceptions form field
func (req *chromiumRequest) FailOnConsoleExceptions(isFailOnConsoleExceptions bool) {
	req.values[failOnConsoleExceptions] = strconv.FormatBool(isFailOnConsoleExceptions)
}

// EmulatedMediaType sets emulatedMediaType form field
func (req *chromiumRequest) EmulatedMediaType(mediaType string) {
	req.values[emulatedMediaType] = mediaType
}

// PDFFormat sets pdfFormat form field
func (req *chromiumRequest) PDFFormat(format string) {
	req.values[pdfFormat] = format
}
