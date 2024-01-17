package numberstrings

import "fmt"

func ExampleDigitString1() {
	fmt.Println(DigitString1(0))
	fmt.Println(DigitString1(1))
	fmt.Println(DigitString1(2))
	fmt.Println(DigitString1(3))
	fmt.Println(DigitString1(4))
	fmt.Println(DigitString1(5))
	fmt.Println(DigitString1(6))
	fmt.Println(DigitString1(7))
	fmt.Println(DigitString1(8))
	fmt.Println(DigitString1(9))

	// Output:
	//
	// einund
	// zweiund
	// dreiund
	// vierund
	// fünfund
	// sechsund
	// siebenund
	// achtund
	// neunund
}

func ExampleDigitString10() {
	fmt.Println(DigitString10(0))
	fmt.Println(DigitString10(1))
	fmt.Println(DigitString10(2))
	fmt.Println(DigitString10(3))
	fmt.Println(DigitString10(4))
	fmt.Println(DigitString10(5))
	fmt.Println(DigitString10(6))
	fmt.Println(DigitString10(7))
	fmt.Println(DigitString10(8))
	fmt.Println(DigitString10(9))

	// Output:
	//
	//
	// zwanzig
	// dreißig
	// vierzig
	// fünfzig
	// sechzig
	// siebzig
	// achtzig
	// neunzig
}

func ExampleDigitString100() {
	fmt.Println(DigitString100(0))
	fmt.Println(DigitString100(1))
	fmt.Println(DigitString100(2))
	fmt.Println(DigitString100(3))
	fmt.Println(DigitString100(4))
	fmt.Println(DigitString100(5))
	fmt.Println(DigitString100(6))
	fmt.Println(DigitString100(7))
	fmt.Println(DigitString100(8))
	fmt.Println(DigitString100(9))

	// Output:
	//
	// einhundert
	// zweihundert
	// dreihundert
	// vierhundert
	// fünfhundert
	// sechshundert
	// siebenhundert
	// achthundert
	// neunhundert
}

func ExampleNumberStringGreater20() {
	fmt.Println(NumberStringGreater20(20))
	fmt.Println(NumberStringGreater20(21))
	fmt.Println(NumberStringGreater20(22))
	fmt.Println(NumberStringGreater20(30))
	fmt.Println(NumberStringGreater20(31))
	fmt.Println(NumberStringGreater20(32))
	fmt.Println(NumberStringGreater20(41))
	fmt.Println(NumberStringGreater20(42))
	fmt.Println(NumberStringGreater20(173))
	fmt.Println(NumberStringGreater20(852))
	fmt.Println(NumberStringGreater20(999))
	fmt.Println(NumberStringGreater20(103))
	fmt.Println(NumberStringGreater20(113))

	// Output:
	// zwanzig
	// einundzwanzig
	// zweiundzwanzig
	// dreißig
	// einunddreißig
	// zweiunddreißig
	// einundvierzig
	// zweiundvierzig
	// einhundertdreiundsiebzig
	// achthundertzweiundfünfzig
	// neunhundertneunundneunzig
	// einhundertdrei
	// einhundertdreizehn
}

func ExampleNumberStringBelow20() {
	for i := 0; i < 20; i++ {
		fmt.Println(NumberStringBelow20(i))
	}

	// Output:
	// null
	// eins
	// zwei
	// drei
	// vier
	// fünf
	// sechs
	// sieben
	// acht
	// neun
	// zehn
	// elf
	// zwölf
	// dreizehn
	// vierzehn
	// fünfzehn
	// sechzehn
	// siebzehn
	// achtzehn
	// neunzehn
}

func ExampleNumberString3Digits() {
	fmt.Println(NumberString3Digits(0))
	fmt.Println(NumberString3Digits(1))
	fmt.Println(NumberString3Digits(2))
	fmt.Println(NumberString3Digits(3))
	fmt.Println(NumberString3Digits(4))
	fmt.Println(NumberString3Digits(5))
	fmt.Println(NumberString3Digits(6))
	fmt.Println(NumberString3Digits(7))
	fmt.Println(NumberString3Digits(8))
	fmt.Println(NumberString3Digits(9))
	fmt.Println(NumberString3Digits(10))
	fmt.Println(NumberString3Digits(11))
	fmt.Println(NumberString3Digits(12))
	fmt.Println(NumberString3Digits(13))
	fmt.Println(NumberString3Digits(14))
	fmt.Println(NumberString3Digits(15))
	fmt.Println(NumberString3Digits(16))
	fmt.Println(NumberString3Digits(17))
	fmt.Println(NumberString3Digits(18))
	fmt.Println(NumberString3Digits(19))
	fmt.Println(NumberString3Digits(20))
	fmt.Println(NumberString3Digits(21))
	fmt.Println(NumberString3Digits(22))
	fmt.Println(NumberString3Digits(30))
	fmt.Println(NumberString3Digits(31))
	fmt.Println(NumberString3Digits(32))
	fmt.Println(NumberString3Digits(41))
	fmt.Println(NumberString3Digits(42))
	fmt.Println(NumberString3Digits(173))
	fmt.Println(NumberString3Digits(852))
	fmt.Println(NumberString3Digits(999))

	// Output:
	// null
	// eins
	// zwei
	// drei
	// vier
	// fünf
	// sechs
	// sieben
	// acht
	// neun
	// zehn
	// elf
	// zwölf
	// dreizehn
	// vierzehn
	// fünfzehn
	// sechzehn
	// siebzehn
	// achtzehn
	// neunzehn
	// zwanzig
	// einundzwanzig
	// zweiundzwanzig
	// dreißig
	// einunddreißig
	// zweiunddreißig
	// einundvierzig
	// zweiundvierzig
	// einhundertdreiundsiebzig
	// achthundertzweiundfünfzig
	// neunhundertneunundneunzig
}

func ExampleNumberString6Digits() {
	fmt.Println(NumberString6Digits(113642))
	fmt.Println(NumberString6Digits(100000))
	fmt.Println(NumberString6Digits(100013))
	fmt.Println(NumberString6Digits(786))
	fmt.Println(NumberString6Digits(16384))
	fmt.Println(NumberString6Digits(0))
	fmt.Println(NumberString6Digits(103))

	// Output:
	// einhundertdreizehntausendsechshundertzweiundvierzig
	// einhunderttausend
	// einhunderttausenddreizehn
	// siebenhundertsechsundachtzig
	// sechzehntausenddreihundertvierundachtzig
	// null
	// einhundertdrei
}
