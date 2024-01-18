package numberstrings

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	/* Hinweis:
	Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
	auch eine Liste mit den Strings für die Ziffern verwenden.
	*/
	return ""
}

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	/* Hinweis:
	Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
	auch eine Liste mit den Strings für die Ziffern verwenden.
	*/
	return ""
}

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	/* Hinweis:
	Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
	auch eine Liste mit den Strings für die Ziffern verwenden.
	*/
	return ""
}

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {
	/* Hinweis:
	Verwenden Sie Modulo- und Divisions-Operationen, um die Ziffern der Zahl zu bestimmen.
	Verwenden Sie die Funktionen DigitString1, DigitString10 und DigitString100,
	um die Ziffern in Strings umzuwandeln.
	*/
	return ""
}

// NumberStringBelow20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	/* Hinweis: Verwenden Sie eine Switch-Anweisung oder eine Liste. */
	return ""
}

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	/* Hinweis:
	Verwenden Sie die Funktionen NumberStringBelow20 und NumberStringGreater20.
	*/
	return ""
}

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {
	/* Hinweis:
	Verwenden Sie Modulo- und Divisions-Operationen, um die Zahl in zwe Dreierblöcke zu zerlegen.
	Verwenden Sie die Funktion NumberString3Digits, um die Ziffern in Strings umzuwandeln.
	*/
	result := ""
	return result
}
