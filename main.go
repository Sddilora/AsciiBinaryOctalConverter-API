package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"api/converters"
)

func main() {

	app := Setup()

	log.Fatal(app.Listen(":6027"))

}

func Setup() *fiber.App {

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/convert", func(c *fiber.Ctx) error {

		type Result struct {
			Output string `json:"message"`
		}

		type RequestBody struct {
			Value      string `json:"value"`
			SourceType string `json:"sourceType"`
			DestType   string `json:"destType"`
		}

		var reqBody RequestBody

		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad Request",
			})
		}

		if reqBody.SourceType == "ascii" && reqBody.DestType == "binary" {
			asciiInput := reqBody.Value
			binaryResult := converters.AsciiToBinary(asciiInput)
			result := Result{Output: binaryResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "ascii" && reqBody.DestType == "octal" {
			asciiInput := reqBody.Value
			octalResult := AsciiToOctal(asciiInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "ascii" {
			binaryInput := reqBody.Value
			asciiResult := BinaryToAscii(binaryInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "octal" {
			binaryInput := reqBody.Value
			octalResult := BinaryToOctal(binaryInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "ascii" {
			octalInput := reqBody.Value
			asciiResult := OctalToAscii(octalInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "binary" {
			octalInput := reqBody.Value
			binaryResult := OctalToBinary(octalInput)

			result := Result{Output: binaryResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		}

		return nil
	})
	return app
}

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

		for len(octal) < 3 { // Eğer sayının uzunluğu 2 basamak değilse başına sıfır ekliyoruz
			octal = "0" + octal
		}

		output += octal + " " // Sekizli çıktısını birleştiriyoruz
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

		binary := input[i : i+8]                    // 8 bitlik ikili sayıyı ASCII karakterine dönüştürüyoruz ex.01010101
		ascii, _ := strconv.ParseInt(binary, 2, 64) //(01010101=U)  U=85
		//output = fmt.Sprintf("%v+%s", ascii, " ")
		output += string(rune(ascii)) + " "
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
		output += strconv.FormatInt(octal, 8) + " "

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
		output += string(rune(ascii)) + " "
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

			output += binary + " " // Binary çıktısını birleştiriyoruz
		}
	}

	return output
}
