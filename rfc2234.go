package rfc2234

// IsAlphaByte tests for ALPHA definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsAlphaByte(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

// IsBitByte tests for BIT definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsBitByte(b byte) bool {
	return b == '0' || b == '1'
}

// IsCharByte tests for CHAR definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsCharByte(b byte) bool {
	return b >= '\x01' && b <= '\x7f'
}

// IsCarriageReturnByte tests for CR definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsCarriageReturnByte(b byte) bool {
	return b == '\x0d'
}

// IsControlByte tests for CTL definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsControlByte(b byte) bool {
	return (b >= '\x00' && b <= '\x1f') || b == '\x7f'
}

// IsDigitByte tests for DIGIT definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsDigitByte(b byte) bool {
	return b >= '0' && b <= '9'
}

// IsDoubleQuoteByte tests for DQUOTE definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsDoubleQuoteByte(b byte) bool {
	return b == '\x22'
}

// IsHexDigitByte tests for HEXDIG definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsHexDigitByte(b byte) bool {
	return IsDigitByte(b) || (b >= 'A' && b <= 'F')
}

// IsHorizontalTabByte tests for HTAB definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsHorizontalTabByte(b byte) bool {
	return b == '\x09'
}

// IsLineFeedByte tests for LF definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsLineFeedByte(b byte) bool {
	return b == '\x0A'
}

// IsSpaceByte tests for SP definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsSpaceByte(b byte) bool {
	return b == '\x20'
}

// IsVisibleCharacterByte tests for VCHAR definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsVisibleCharacterByte(b byte) bool {
	return b >= '\x21' && b <= '\x7e'
}

// IsWhitespaceByte tests for WSP definition as defined in https://datatracker.ietf.org/doc/html/rfc2234#section-6.1
func IsWhitespaceByte(b byte) bool {
	return IsSpaceByte(b) || IsHorizontalTabByte(b)
}
