package converters

import (
	"strconv"
	"strings"
)

func AsciiToBinary(input string) string { // ASCII to Binary

	output := ""

	for i := 0; i < len([]rune(input)); i++ {

		binary := strconv.FormatInt(int64(input[i]), 2) // ASCII kodunu ikili sayıya çeviriyoruz

		if len(binary) < 8 { // Eğer sayının uzunluğu 8 bit değilse başına sıfır ekliyoruz
			for i := len(binary); i < 8; i++ {
				binary = "0" + binary
			}
		}

		if i < len([]rune(input))-1 {
			output += binary + " "
		} else {
			output += binary
		}
	}

	return output
}

func AsciiToOctal(input string) string { // ASCII to Octal

	output := ""

	for i := 0; i < len([]rune(input)); i++ {

		octal := strconv.FormatInt(int64(input[i]), 8) // ASCII kodunu sekizli sayıya çeviriyoruz

		for len(octal) < 3 { // Eğer sayının uzunluğu 2 basamak değilse başına sıfır ekliyoruz
			octal = "0" + octal
		}

		if i < len([]rune(input))-1 {
			output += octal + " "
		} else {
			output += octal
		}
	}

	return output
}

func BinaryToAscii(input string) string { // Binary to ASCII

	output := ""

	unwanted := " "
	input = strings.Map(func(r rune) rune {
		if strings.ContainsRune(unwanted, r) {
			return -1
		}
		return r
	}, input)

	for i := 0; i < len(input); i += 8 { // Her 8 karakter için ASCII karakterlerini birleştiriyoruz

		if i+8 > len(input) {
			x := strings.Repeat("0", i+8-len(input))
			input = x + input
		}

		binary := input[i : i+8]                    // 8 bitlik ikili sayıyı ASCII karakterine dönüştürüyoruz ex.U -> 01010101
		ascii, _ := strconv.ParseInt(binary, 2, 64) //(01010101=U)  (U=85 in 10)

		if i+8 == (len([]rune(input))) {
			output += string(rune(ascii))
		} else {
			output += string(rune(ascii)) + " "
		}

	}

	return output
}

func BinaryToOctal(input string) string { // Binary to Octal
	output := ""

	unwanted := " "
	input = strings.Map(func(r rune) rune {
		if strings.ContainsRune(unwanted, r) {
			return -1
		}
		return r
	}, input)

	for i := 0; i < len(input); i += 8 {

		if i+8 > len(input) {
			x := strings.Repeat("0", i+8-len(input))
			input = x + input
		}

		binary := input[i : i+8]
		octal, _ := strconv.ParseInt(binary, 2, 64)

		if i+8 == (len([]rune(input))) {
			output += strconv.FormatInt(octal, 8)
		} else {
			output += strconv.FormatInt(octal, 8) + " "
		}

	}

	return output
}

func OctalToAscii(input string) string { // Octal to ASCII
	output := ""

	unwanted := " "
	input = strings.Map(func(r rune) rune {
		if strings.ContainsRune(unwanted, r) {
			return -1
		}
		return r
	}, input)

	for i := 0; i < len(input); i += 3 { // Her 3 karakter için ASCII karakterlerini birleştiriyoruz

		if i+3 > len(input) { // Eğer son 3 karakterin uzunluğu 2'den az ise başına sıfır ekliyoruz
			x := strings.Repeat("0", i+3-len(input))
			input = x + input
		}

		octal := input[i : i+3] // 3 basamaklı sekizli sayıyı ASCII karakterine dönüştürüyoruz
		ascii, _ := strconv.ParseInt(octal, 8, 64)

		if i+3 == (len([]rune(input))) {
			output += string(rune(ascii))
		} else {
			output += string(rune(ascii)) + " "
		}

	}

	return output
}

func OctalToBinary(input string) string { // Octal to Binary
	output := ""

	unwanted := " "
	input = strings.Map(func(r rune) rune {
		if strings.ContainsRune(unwanted, r) {
			return -1
		}
		return r
	}, input)

	for i := 0; i < len(input); i += 3 { // Her 3 karakter için ASCII karakterlerini birleştiriyoruz

		if i+3 > len(input) { // Eğer son 3 karakterin uzunluğu 2'den az ise başına sıfır ekliyoruz
			x := strings.Repeat("0", i+3-len(input))
			input = x + input
		}

		octal := input[i : i+3] // 3 basamaklı sekizli sayıyı ASCII karakterine dönüştürüyoruz
		ascii, _ := strconv.ParseInt(octal, 8, 64)
		for _, c := range string(rune(ascii)) {

			binary := strconv.FormatInt(int64(c), 2) // ASCII kodunu ikili sayıya çeviriyoruz

			if len(binary) < 8 { // Eğer sayının uzunluğu 8 bit değilse başına sıfır ekliyoruz
				for i := len(binary); i < 8; i++ {
					binary = "0" + binary
				}
			}

			if i+3 == (len([]rune(input))) {
				output += binary // Binary çıktısını birleştiriyoruz
			} else {
				output += binary + " " // Binary çıktısını birleştiriyoruz
			}

		}
	}

	return output
}
