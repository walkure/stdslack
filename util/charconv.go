package util

import (
	"bytes"
	"fmt"
	"io"

	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

// ToUTF8 converts to utf-8 charset
func ToUTF8(input io.Reader) (io.Reader, error) {

	bytesInput, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("failure to read input: %w", err)
	}

	detector := chardet.NewTextDetector()
	deetctResult, err := detector.DetectBest(bytesInput)
	if err != nil {
		return nil, fmt.Errorf("failure to detect charset: %w", err)
	}

	reader, err := charset.NewReaderLabel(deetctResult.Charset, bytes.NewReader(bytesInput))
	if err != nil {
		return nil, fmt.Errorf("failure to convert input: %w", err)
	}

	return reader, nil
}
