package roman

import "strings"

// NToI erwartet eine Zahl und liefert die entsprechende Anzahl an I als String.
func NToI(n int) string {
	return strings.Repeat("I", n)
}

// NToX erwartet eine Zahl und liefert die entsprechende Anzahl an X als String.
func NToX(n int) string {
	return strings.Repeat("X", n/10)
}

// RomanBelow10 erwartet eine Zahl 1 <= n <= 10 und liefert die römische Schreibweise für n als String.
func RomanBelow10(n int) string {
	if n <= 3 {
		return NToI(n)
	}
	if n <= 5 {
		return NToI(5-n) + "V"
	}
	if n <= 8 {
		return "V" + NToI((n - 5))
	}
	return NToI(10-n) + "X"
}

// RomanBelow100 erwartet eine Zahl 1 <= n <= 100 und liefert die römische Schreibweise für n als String.
func RomanBelow100(n int) string {
	if n <= 39 {
		return NToX(n) + RomanBelow10(n%10)
	}
	if n < 50 {
		return NToX(n-30) + "L" + RomanBelow10(n%10)
	}
	return ""
}
