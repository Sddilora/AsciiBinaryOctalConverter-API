package converters

import (
	"strconv"
	"strings"
)

func AsciiToBinary(input string) string { // ASCII to Binary

	output := ""

	for _, c := range input {

		binary := strconv.FormatInt(int64(c), 2) // ASCII kodunu ikili sayıya çeviriyoruz

		if len(binary) < 8 { // Eğer sayının uzunluğu 8 bit değilse başına sıfır ekliyoruz
			for i := len(binary); i < 8; i++ {
				binary = "0" + binary
			}
		}

		output += binary + " " // Binary çıktısını birleştiriyoruz
	}

	return output
}

func AsciiToOctal(input string) string { // ASCII to Octal

	output := ""

	for _, c := range input {

		octal := strconv.FormatInt(int64(c), 8) // ASCII kodunu sekizli sayıya çeviriyoruz

		for len(octal) < 2 { // Eğer sayının uzunluğu 2 basamak değilse başına sıfır ekliyoruz
			octal = "0" + octal
		}

		output += octal + " " // Sekizli çıktısını birleştiriyoruz
	}

	return output
}

func BinaryToAscii(input string) string { // Binary to ASCII

	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for i := 0; i < len(input); i += 8 { // Her 8 karakter için ASCII karakterlerini birleştiriyoruz

		if i+8 > len(input) { // Eğer son 8 karakterin uzunluğu 8'den az ise başına sıfır ekliyoruz
			input += strings.Repeat("0", i+8-len(input))
		}

		binary := input[i : i+8] // 8 bitlik ikili sayıyı ASCII karakterine dönüştürüyoruz
		ascii, _ := strconv.ParseInt(binary, 2, 64)
		output += string(ascii) + " "
	}

	return output
}

func BinaryToOctal(input string) string { // Binary to Octal
	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for i := 0; i < len(input); i += 8 {

		if i+3 > len(input) {
			input += strings.Repeat("0", i+8-len(input))
		}

		binary := input[i : i+8]
		octal, _ := strconv.ParseInt(binary, 2, 64)
		output += strconv.FormatInt(octal, 8) + " "

	}

	return output
}

func OctalToAscii(input string) string { // Octal to ASCII
	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for i := 0; i < len(input); i += 3 { // Her 3 karakter için ASCII karakterlerini birleştiriyoruz

		if i+3 > len(input) { // Eğer son 3 karakterin uzunluğu 2'den az ise başına sıfır ekliyoruz
			input += strings.Repeat("0", i+3-len(input))
		}

		octal := input[i : i+3] // 3 basamaklı sekizli sayıyı ASCII karakterine dönüştürüyoruz
		ascii, _ := strconv.ParseInt(octal, 8, 64)
		output += string(ascii) + " "
	}

	return output
}

func OctalToBinary(input string) string { // Octal to Binary
	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for i := 0; i < len(input); i += 3 { // Her 3 karakter için ASCII karakterlerini birleştiriyoruz

		if i+3 > len(input) { // Eğer son 3 karakterin uzunluğu 2'den az ise başına sıfır ekliyoruz
			input += strings.Repeat("0", i+3-len(input))
		}

		octal := input[i : i+3] // 3 basamaklı sekizli sayıyı ASCII karakterine dönüştürüyoruz
		ascii, _ := strconv.ParseInt(octal, 8, 64)
		for _, c := range string(ascii) {

			binary := strconv.FormatInt(int64(c), 2) // ASCII kodunu ikili sayıya çeviriyoruz

			if len(binary) < 8 { // Eğer sayının uzunluğu 8 bit değilse başına sıfır ekliyoruz
				for i := len(binary); i < 8; i++ {
					binary = "0" + binary
				}
			}

			output += binary + " " // Binary çıktısını birleştiriyoruz
		}
	}

	return output
}

func main() {
}
