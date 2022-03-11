package rfc2234

// IsAlphaByte test for ALPHA definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsAlphaByte(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

// IsDigitByte test for DIGIT definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsDigitByte(b byte) bool {
	return b >= '0' && b <= '9'
}
