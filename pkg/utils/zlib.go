package utils

import (
	"bytes"
	"compress/zlib"
	"io"
)

// ZlibUncompress do zlib uncompress
func ZlibUncompress(compressSrc []byte) ([]byte, error) {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	io.Copy(&out, r)
	return out.Bytes(), nil
}

//ZlibCompress do zlib compress
func ZlibCompress(src []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err := w.Write(src)
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}
