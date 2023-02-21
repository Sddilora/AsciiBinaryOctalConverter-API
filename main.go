package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
			binaryResult := asciiToBinary(asciiInput)
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
			binaryInput := reqBody.Value // c.Params(reqBody.Value)
			asciiResult := binaryToAscii(binaryInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "octal" {
			binaryInput := reqBody.Value //c.Params(reqBody.Value)
			octalResult := binaryToOctal(binaryInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "ascii" {
			octalInput := reqBody.Value //c.Params(reqBody.Value)
			asciiResult := octalToAscii(octalInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "binary" {
			octalInput := reqBody.Value //c.Params(reqBody.Value)
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

var output string

// ASCII to Binary
func asciiToBinary(input string) string {

	for _, c := range input {

		binary := strconv.FormatInt(int64(c), 2) // ASCII kodunu ikili sayıya çeviriyoruz

		for len(binary) < 7 { // Eğer sayının uzunluğu 7 bit değilse başına sıfır ekliyoruz
			binary = "0" + binary
		}

		output += binary // Binary çıktısını birleştiriyoruz
	}

	return output
}

func asciiToOctal(input string) string { // ASCII to Octal

	output := ""

	for _, c := range input {

		octal := strconv.FormatInt(int64(c), 8) // ASCII kodunu sekizli sayıya çeviriyoruz

		for len(octal) < 2 { // Eğer sayının uzunluğu 2 basamak değilse başına sıfır ekliyoruz
			octal = "0" + octal
		}

		output += octal // Sekizli çıktısını birleştiriyoruz
	}

	return output
}

// Binary to ASCII
func binaryToAscii(input string) string {
	output := ""

	// Her 7 karakter için ASCII karakterlerini birleştiriyoruz
	for i := 0; i < len(input); i += 7 {
		// Eğer son 7 karakterin uzunluğu 7'den az ise başına sıfır ekliyoruz
		if i+7 > len(input) {
			input += strings.Repeat("0", i+7-len(input))
		}

		// 7 bitlik ikili sayıyı ASCII karakterine dönüştürüyoruz
		binary := input[i : i+7]
		ascii, _ := strconv.ParseInt(binary, 2, 64)
		output += string(ascii)
	}

	return output
}

// Binary to Octal
func binaryToOctal(input string) string {
	output := ""

	// Her 3 karakter için sekizli sayıları birleştiriyoruz
	for i := 0; i < len(input); i += 3 {
		// Eğer son 3 karakterin uzunluğu 3'ten az ise başına sıfır ekliyoruz
		if i+3 > len(input) {
			input += strings.Repeat("0", i+3-len(input))
		}

		// 3 bitlik ikili sayıyı sekizli sayıya dönüştürüyoruz
		binary := input[i : i+3]
		octal, _ := strconv.ParseInt(binary, 2, 64)
		output += strconv.FormatInt(octal, 8)
	}

	return output
}

// Octal to ASCII
func octalToAscii(input string) string {
	output := ""

	// Her 2 karakter için ASCII karakterlerini birleştiriyoruz
	for i := 0; i < len(input); i += 2 {
		// Eğer son 2 karakterin uzunluğu 2'den az ise başına sıfır ekliyoruz
		if i+2 > len(input) {
			input += strings.Repeat("0", i+2-len(input))
		}

		// 2 basamaklı sekizli sayıyı ASCII karakterine dönüştürüyoruz
		octal := input[i : i+2]
		ascii, _ := strconv.ParseInt(octal, 8, 64)
		output += string(ascii)
	}

	return output
}

// Octal to Binary
func octalToBinary(input string) string {
	output := ""

	// Her 1 karakter için 3 bitlik ikili sayıları birleştiriyoruz
	for _, c := range input {
		// Sekizli sayıyı 3 bitlik ikili sayıya dönüştürüyoruz
		octal := string(c)
		binary := strconv.FormatInt(int64(octal[0]-'0'), 2)
		for i := 1; i < len(octal); i++ {
			digit := octal[i] - '0'
			padded := fmt.Sprintf("%03d", digit)
			binary += padded
		}

		// 3 bitlik ikili sayıları birleştiriyoruz
		output += binary
	}

	return output
}

/*var octalValues []string
for i := 0; i < len(input); i++ {
	decimalValue := int(input[i])
	octalValue := strconv.FormatInt(int64(decimalValue), 8)
	octalValues = append(octalValues, octalValue)
}
responseJSON := octalValues
return c.JSON(responseJSON)

var binaryValues []string
			for i := 0; i < len(input); i++ {
				decimalValue := int(input[i])
				binaryValue := strconv.FormatInt(int64(decimalValue), 2)
				binaryValues = append(binaryValues, binaryValue)
			}
			responseJSON := binaryValues
			return c.JSON(responseJSON)

*/

//value := c.FormValue("reqBody.value")
//sourceType := c.FormValue("reqBody.SourceType")
//destType := c.FormValue("reqBody.DestType")

// Here, you can access reqBody.DestType and use it in your conversion logic

// return ctx.JSON(fiber.Map{
// 	"message": fmt.Sprintf("%v %s converted to %s", reqBody.Value, reqBody.SourceType, reqBody.DestType),
// })
