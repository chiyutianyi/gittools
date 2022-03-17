package utils

import "fmt"

func Hexdump(content []byte) string {
	var (
		i          = 0
		b          byte
		buffer     []byte
		asc_buffer []byte
		out        []byte
	)

	for i, b = range content {
		if i&15 == 0 && i != 0 {
			dumpLine(&out, &buffer, &asc_buffer, i-16)
			buffer = nil
			asc_buffer = nil
		}
		if i&7 == 0 && i&15 != 0 {
			buffer = append(buffer, ' ')
		}
		buffer = append(buffer, fmt.Sprintf(" %02x", b)...)
		if Isprint(b) {
			asc_buffer = append(asc_buffer, b)
		} else {
			asc_buffer = append(asc_buffer, '.')
		}
	}
	if len(buffer) > 0 {
		dumpLine(&out, &buffer, &asc_buffer, i)
	}
	out = append(out, []byte(fmt.Sprintf("%08x\n", i>>4<<4+16))...)
	return string(out)
}

func dumpLine(out, buffer, asc_buffer *[]byte, i int) {
	*out = append(*out, []byte(fmt.Sprintf("%08x %-49s  |", i>>4<<4, string(*buffer)))...)
	*out = append(*out, *asc_buffer...)
	*out = append(*out, '|', '\n')
}
