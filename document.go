package gotenberg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Document interface {
	Filename() string
	Reader() (io.ReadCloser, error)
}

type document struct {
	filename string
}

func (doc *document) Filename() string {
	return doc.filename
}

type documentFromPath struct {
	fpath string

	*document
}

// NewDocumentFromPath creates a Document from
// a file path.
func NewDocumentFromPath(filename, fpath string) (Document, error) {
	if !fileExists(fpath) {
		return nil, fmt.Errorf("%s: file %s does not exist", fpath, filename)
	}
	return &documentFromPath{
		fpath,
		&document{filename},
	}, nil
}

func (doc *documentFromPath) Reader() (io.ReadCloser, error) {
	in, err := os.Open(doc.fpath)
	if err != nil {
		return nil, fmt.Errorf("%s: opening file: %v", doc.Filename(), err)
	}
	return in, nil
}

type documentFromString struct {
	data string

	*document
}

// NewDocumentFromString creates a Document from
// a string.
func NewDocumentFromString(filename, data string) (Document, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%s: string is empty", filename)
	}
	return &documentFromString{
		data,
		&document{filename},
	}, nil
}

func (doc *documentFromString) Reader() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(doc.data)), nil
}

type documentFromBytes struct {
	data []byte

	*document
}

// NewDocumentFromBytes creates a Document from
// bytes.
func NewDocumentFromBytes(filename string, data []byte) (Document, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%s: bytes are empty", filename)
	}
	return &documentFromBytes{
		data,
		&document{filename},
	}, nil
}

func (doc *documentFromBytes) Reader() (io.ReadCloser, error) {
	return io.NopCloser(bytes.NewReader(doc.data)), nil
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ Document = new(documentFromPath)
	_ Document = new(documentFromString)
	_ Document = new(documentFromBytes)
)
