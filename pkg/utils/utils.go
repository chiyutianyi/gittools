package utils

func Isprint(x byte) bool {
	return x >= 0x20 && (x) <= 0x7e
}
