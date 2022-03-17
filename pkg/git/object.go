package git

import (
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

func ParseType(b byte) plumbing.ObjectType {
	return plumbing.ObjectType((b & MaskType) >> FirstLengthBits)
}
