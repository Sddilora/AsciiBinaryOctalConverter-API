package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"api/converters"
)

func main() {

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
			octalResult := asciiToOctal(asciiInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "ascii" {
			binaryInput := reqBody.Value
			asciiResult := binaryToAscii(binaryInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "octal" {
			binaryInput := reqBody.Value
			octalResult := binaryToOctal(binaryInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "ascii" {
			octalInput := reqBody.Value
			asciiResult := octalToAscii(octalInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "binary" {
			octalInput := reqBody.Value
			binaryResult := octalToBinary(octalInput)

			result := Result{Output: binaryResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		}

		return nil
	})

	log.Fatal(app.Listen(":6027"))

}

func asciiToOctal(input string) string { // ASCII to Octal

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

func binaryToAscii(input string) string { // Binary to ASCII

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

func binaryToOctal(input string) string { // Binary to Octal
	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for i := 0; i < len(input); i += 3 { // Her 3 karakter için sekizli sayıları birleştiriyoruz

		if i+3 > len(input) { // Eğer son 3 karakterin uzunluğu 3'ten az ise başına sıfır ekliyoruz
			input += strings.Repeat("0", i+3-len(input))
		}

		binary := string(input[i : i+3]) // 3 bitlik ikili sayıyı sekizli sayıya dönüştürüyoruz
		octal, _ := strconv.ParseInt(binary, 2, 64)
		output += strconv.FormatInt(octal, 8) + " "
	}

	return output
}

func octalToAscii(input string) string { // Octal to ASCII
	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for i := 0; i < len(input); i += 2 { // Her 2 karakter için ASCII karakterlerini birleştiriyoruz

		if i+2 > len(input) { // Eğer son 2 karakterin uzunluğu 2'den az ise başına sıfır ekliyoruz
			input += strings.Repeat("0", i+2-len(input))
		}

		octal := input[i : i+2] // 2 basamaklı sekizli sayıyı ASCII karakterine dönüştürüyoruz
		ascii, _ := strconv.ParseInt(octal, 8, 64)
		output += string(ascii) + " "
	}

	return output
}

func octalToBinary(input string) string { // Octal to Binary
	output := ""

	var inputholder []byte

	for i := 0; i < len(input); i++ {

		if string(input[i]) != " " {
			inputholder = append(inputholder, input[i])
		}
	}

	input = string(inputholder)

	for _, c := range input { // Her 1 karakter için 3 bitlik ikili sayıları birleştiriyoruz

		octal := string(c) // Sekizli sayıyı 3 bitlik ikili sayıya dönüştürüyoruz
		binary := strconv.FormatInt(int64(octal[0]-'0'), 2)
		for i := 1; i < len(octal); i++ {
			digit := octal[i] - '0'
			padded := fmt.Sprintf("%03d", digit)
			binary += padded
		}

		output += binary + " " // 3 bitlik ikili sayıları birleştiriyoruz
	}

	return output
}
