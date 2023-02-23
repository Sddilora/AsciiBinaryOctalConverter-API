package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertEndpoint(t *testing.T) {
	// Create a new fiber app
	app := Setup()
	defer app.Shutdown()

	///////////////////////////////////////////////// ASCİİTOBİNARY
	// Create a test request
	req := httptest.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(`{
        "value": "dilara",
        "sourceType": "ascii",
        "destType": "binary"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	expected := `{"message":"01100100 01101001 01101100 01100001 01110010 01100001 "}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	//////////////////////////////////////////////////ASCİİTOOCTAL
	// Create a test request
	req = httptest.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(`{
        "value": "dilara0",
        "sourceType": "ascii",
        "destType": "octal"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err = app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	expected = `{"message":"144 151 154 141 162 141 060 "}`
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	/////////////////////////////////////////////////////////BİNARYTOASCİİ
	// Create a test request
	req = httptest.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(`{
        "value": "1010101",
        "sourceType": "binary",
        "destType": "ascii"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err = app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	expected = `{"message":"U "}`
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	///////////////////////////////////////////////////////BİNARYTOOCTAL
	// Create a test request
	req = httptest.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(`{
        "value": "1010101",
        "sourceType": "binary",
        "destType": "octal"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err = app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	expected = `{"message":"125 "}`
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	///////////////////////////////////////////////////////////////////////OCTALTOASCİİ
	// Create a test request
	req = httptest.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(`{
        "value": "60",
        "sourceType": "octal",
        "destType": "ascii"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err = app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	expected = `{"message":"0 "}`
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

	/////////////////////////////////////////////////////////////OCTALTOBİNARY
	req = httptest.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(`{
        "value": "60",
        "sourceType": "octal",
        "destType": "binary"
    }`)))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err = app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the response body
	expected = `{"message":"00110000 "}`
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("Expected response body '%s' but got '%s'", expected, string(body))
	}

}
