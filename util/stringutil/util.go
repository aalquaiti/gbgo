package stringutil

import "strings"

// AsciiToStr Convert Byte Slice to String
func AsciiToStr(src []byte, length int) string {
	sb := strings.Builder{}
	str := make([]byte, length)
	copy(str, src)

	sb.Write(str)

	return sb.String()
}
