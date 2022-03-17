package git

import (
	"fmt"
	"io/ioutil"

	"github.com/go-git/go-git/v5/plumbing"
)

const (
	// VersionSupported is the packfile version supported by this package
	VersionSupported uint32 = 2

	FirstLengthBits = uint8(4)   // the first byte into object header has 4 bits to store the length
	LengthBits      = uint8(7)   // subsequent bytes has 7 bits to store the length
	MaskFirstLength = 15         // 0000 1111
	MaskContinue    = 0x80       // 1000 0000
	MaskLength      = uint8(127) // 0111 1111
	MaskType        = uint8(112) // 0111 0000
)

func ParseObject(buffer []byte, o plumbing.EncodedObject) ([]byte, error) {
	buffer = append(buffer, []byte("-------------------------------\n")...)
	buffer = append(buffer, []byte(fmt.Sprintf("Hash = %v\n", o.Hash()))...)
	buffer = append(buffer, []byte(fmt.Sprintf("Type = %v\n", o.Type()))...)
	buffer = append(buffer, []byte(fmt.Sprintf("Size = %v\n", o.Size()))...)
	r, err := o.Reader()
	if err != nil {
		return buffer, err
	}
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return buffer, err
	}
	buffer = append(buffer, []byte(fmt.Sprintf("Content bytes length = %v\n", len(bs)))...)
	buffer = append(buffer, []byte(fmt.Sprintf("Content bytes= \n%v\n", bs))...)
	buffer = append(buffer, []byte(fmt.Sprintf("Content = \n%v\n", string(bs)))...)
	return buffer, nil
}

func ParseType(b byte) plumbing.ObjectType {
	return plumbing.ObjectType((b & MaskType) >> FirstLengthBits)
}
