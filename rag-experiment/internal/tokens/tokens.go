package tokens

import "unicode/utf8"

// Estimate uses the explicit, reproducible approximation of four UTF-8 characters per token.
func Estimate(text string) int { n := utf8.RuneCountInString(text); return (n + 3) / 4 }
