package rfc2234

// IsAlphaByte tests for ALPHA definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsAlphaByte(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

// IsDigitByte tests for DIGIT definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsDigitByte(b byte) bool {
	return b >= '0' && b <= '9'
}
