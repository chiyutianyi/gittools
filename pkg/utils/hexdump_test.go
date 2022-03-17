package utils_test

import (
	"testing"

	"github.com/chiyutianyi/gittools/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestHexdump(t *testing.T) {
	testcases := []struct {
		content []byte
		expect  string
	}{
		{
			content: []byte{0x01, 0x02},
			expect: `00000000  01 02                                             |..|
00000010
`,
		},
		{
			content: []byte{
				0x50, 0x41, 0x43, 0x4b, 0x00, 0x00, 0x00, 0x02, 0x0, 0x0, 0x07, 0x8b, 0x91, 0x4c, 0x78, 0x01,
			},
			expect: `00000000  50 41 43 4b 00 00 00 02  00 00 07 8b 91 4c 78 01  |PACK.........Lx.|
00000010
`,
		},
		{
			content: []byte{
				0x50, 0x41, 0x43, 0x4b, 0x00, 0x00, 0x00, 0x02, 0x0, 0x0, 0x07, 0x8b, 0x91, 0x4c, 0x78, 0x01,
				0x8d,
			},
			expect: `00000000  50 41 43 4b 00 00 00 02  00 00 07 8b 91 4c 78 01  |PACK.........Lx.|
00000010  8d                                                |.|
00000020
`,
		},
	}
	for _, testcase := range testcases {
		assert.Equal(t, string(testcase.expect), utils.Hexdump(testcase.content))
	}
}
