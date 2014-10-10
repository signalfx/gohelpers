package workarounds

import "time"

// GolangDoesnotAllowPointerToStringLiteral allows one to take the address of a string literal
func GolangDoesnotAllowPointerToStringLiteral(s string) *string {
	return &s
}

// GolangDoesnotAllowPointerToTimeLiteral allows one to take the address of a time literal
func GolangDoesnotAllowPointerToTimeLiteral(s time.Duration) *time.Duration {
	return &s
}

// GolangDoesnotAllowPointerToUintLiteral allows one to take the address of a uint32 literal
func GolangDoesnotAllowPointerToUintLiteral(s uint32) *uint32 {
	return &s
}
