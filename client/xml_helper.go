package client

import (
	"encoding/xml"
	"io"

	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
)

var SelectedCharsetReader func(string, io.Reader) (io.Reader, error) = nil

func GetXmlReader(input io.Reader, strict bool) *xml.Decoder {
	decoder := xml.NewDecoder(input)
	// if SelectedCharsetReader != nil {
	// 	decoder.CharsetReader = SelectedCharsetReader
	// }
	decoder.CharsetReader = charset.NewReader
	decoder.Strict = strict
	return decoder
}
